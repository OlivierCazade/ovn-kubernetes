package node

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/config"
	kubeMocks "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/kube/mocks"
	ovntest "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/testing"
	mocks "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/testing/mocks/github.com/vishvananda/netlink"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/types"
	"github.com/ovn-org/ovn-kubernetes/go-controller/pkg/util"
	utilMocks "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/util/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/vishvananda/netlink"
	"net"
)

func genOVSAddMgmtPortCmd(nodeName string) string {
	return fmt.Sprintf("ovs-vsctl --timeout=15 -- --may-exist add-port br-int %s -- set interface %s external-ids:iface-id=%s",
		types.K8sMgmtIntfName, types.K8sMgmtIntfName, types.K8sPrefix+nodeName)
}

var _ = Describe("Mananagement port smart-nic tests", func() {
	origNetlinkOps := util.GetNetLinkOps()
	var netlinkOpsMock *utilMocks.NetLinkOps
	var execMock *ovntest.FakeExec
	var nodeAnnotatorMock *kubeMocks.Annotator
	var waiter *startupWaiter

	config.PrepareTestConfig()
	util.ResetRunner()

	BeforeEach(func() {
		netlinkOpsMock = &utilMocks.NetLinkOps{}
		nodeAnnotatorMock = &kubeMocks.Annotator{}
		execMock = ovntest.NewFakeExec()
		err := util.SetExec(execMock)
		Expect(err).NotTo(HaveOccurred())
		waiter = newStartupWaiter()
		util.SetNetLinkOpMockInst(netlinkOpsMock)
	})

	AfterEach(func() {
		config.PrepareTestConfig()
		util.SetNetLinkOpMockInst(origNetlinkOps)
		util.ResetRunner()
	})

	Context("Create Management port Smart-NIC", func() {
		It("Fails if representor and ovn-k8s-mp0 netdev is not found", func() {
			mgmtPortSnic := managementPortSmartNIC{
				vfRepName: "non-existent-netdev",
			}
			netlinkOpsMock.On("LinkByName", "non-existent-netdev").Return(
				nil, fmt.Errorf("failed to get interface"))
			netlinkOpsMock.On("LinkByName", types.K8sMgmtIntfName).Return(
				nil, fmt.Errorf("failed to get interface"))

			_, err := mgmtPortSnic.Create(nodeAnnotatorMock, waiter)
			Expect(err).To(HaveOccurred())
		})

		It("Fails if set Name to ovn-k8s-mp0 fails", func() {
			mgmtPortSnic := managementPortSmartNIC{
				vfRepName: "enp3s0f0v0",
			}
			linkMock := &mocks.Link{}
			linkMock.On("Attrs").Return(&netlink.LinkAttrs{Name: "enp3s0f0v0", MTU: 1400})

			netlinkOpsMock.On("LinkByName", "enp3s0f0v0").Return(
				linkMock, nil)
			netlinkOpsMock.On("LinkSetDown", linkMock).Return(nil)
			netlinkOpsMock.On("LinkSetName", linkMock, types.K8sMgmtIntfName).Return(fmt.Errorf("failed to set name"))

			_, err := mgmtPortSnic.Create(nodeAnnotatorMock, waiter)
			Expect(err).To(HaveOccurred())
		})

		It("Configures VF representor and connects it to OVS bridge", func() {
			_, ipnet, err := net.ParseCIDR("192.168.0.1/24")
			Expect(err).ToNot(HaveOccurred())
			expectedMgmtPortMac := util.IPAddrToHWAddr(util.GetNodeManagementIfAddr(ipnet).IP)
			config.Default.MTU = 1400
			mgmtPortSnic := managementPortSmartNIC{
				nodeName:    "k8s-worker0",
				hostSubnets: []*net.IPNet{ipnet},
				vfRepName:   "enp3s0f0v0",
			}
			nodeAnnotatorMock.On("Set", mock.Anything, expectedMgmtPortMac.String()).Return(nil)
			linkMock := &mocks.Link{}
			linkMock.On("Attrs").Return(&netlink.LinkAttrs{Name: "enp3s0f0v0", MTU: 1500})

			netlinkOpsMock.On("LinkByName", "enp3s0f0v0").Return(
				linkMock, nil)
			netlinkOpsMock.On("LinkSetDown", linkMock).Return(nil)
			netlinkOpsMock.On("LinkSetName", linkMock, types.K8sMgmtIntfName).Return(nil)
			netlinkOpsMock.On("LinkSetMTU", linkMock, config.Default.MTU).Return(nil)
			netlinkOpsMock.On("LinkSetUp", linkMock).Return(nil)
			execMock.AddFakeCmd(&ovntest.ExpectedCmd{
				Cmd: genOVSAddMgmtPortCmd(mgmtPortSnic.nodeName),
			})

			mpcfg, err := mgmtPortSnic.Create(nodeAnnotatorMock, waiter)
			Expect(execMock.CalledMatchesExpected()).To(BeTrue(), execMock.ErrorDesc)
			Expect(err).ToNot(HaveOccurred())
			Expect(mpcfg.ifName).To(Equal(types.K8sMgmtIntfName))
			Expect(mpcfg.link).To(Equal(linkMock))
		})

		It("Brings interface up and attemps to add ovn-k8s-mp0 to OVS if interface already configured", func() {
			_, ipnet, err := net.ParseCIDR("192.168.0.1/24")
			Expect(err).ToNot(HaveOccurred())
			expectedMgmtPortMac := util.IPAddrToHWAddr(util.GetNodeManagementIfAddr(ipnet).IP)
			config.Default.MTU = 1400
			mgmtPortSnic := managementPortSmartNIC{
				nodeName:    "k8s-worker0",
				hostSubnets: []*net.IPNet{ipnet},
				vfRepName:   "enp3s0f0v0",
			}
			nodeAnnotatorMock.On("Set", mock.Anything, expectedMgmtPortMac.String()).Return(nil)
			linkMock := &mocks.Link{}
			linkMock.On("Attrs").Return(&netlink.LinkAttrs{Name: "ovn-k8s-mp0", MTU: config.Default.MTU})

			netlinkOpsMock.On("LinkByName", "enp3s0f0v0").Return(
				nil, fmt.Errorf("failed to get link device"))
			netlinkOpsMock.On("LinkByName", "ovn-k8s-mp0").Return(
				linkMock, nil)
			netlinkOpsMock.On("LinkSetUp", linkMock).Return(nil)
			execMock.AddFakeCmd(&ovntest.ExpectedCmd{
				Cmd: genOVSAddMgmtPortCmd(mgmtPortSnic.nodeName),
			})

			mpcfg, err := mgmtPortSnic.Create(nodeAnnotatorMock, waiter)
			Expect(execMock.CalledMatchesExpected()).To(BeTrue(), execMock.ErrorDesc)
			Expect(err).ToNot(HaveOccurred())
			Expect(mpcfg.ifName).To(Equal(types.K8sMgmtIntfName))
			Expect(mpcfg.link).To(Equal(linkMock))
		})
	})

	Context("Create Management port Smart-NIC host", func() {
		It("Fails if netdev does not exist", func() {
			mgmtPortSnicHost := managementPortSmartNICHost{
				netdevName: "non-existent-netdev",
			}
			netlinkOpsMock.On("LinkByName", "non-existent-netdev").Return(
				nil, fmt.Errorf("failed to get interface"))
			netlinkOpsMock.On("LinkByName", types.K8sMgmtIntfName).Return(
				nil, fmt.Errorf("failed to get interface"))

			_, err := mgmtPortSnicHost.Create(nil, waiter)
			Expect(err).To(HaveOccurred())
		})

		It("Configures VF and calls createPlatformManagementPort", func() {
			_, ipnet, err := net.ParseCIDR("192.168.0.1/24")
			Expect(err).ToNot(HaveOccurred())
			expectedMgmtPortMac := util.IPAddrToHWAddr(util.GetNodeManagementIfAddr(ipnet).IP)
			currentMgmtPortMac, err := net.ParseMAC("00:bb:cc:dd:ee:11")
			Expect(err).ToNot(HaveOccurred())
			config.Default.MTU = 1400
			mgmtPortSnicHost := managementPortSmartNICHost{
				hostSubnets: []*net.IPNet{ipnet},
				netdevName:  "enp3s0f0v0",
			}
			linkMock := &mocks.Link{}
			linkMock.On("Attrs").Return(&netlink.LinkAttrs{
				Name: "enp3s0f0v0", MTU: 1500, HardwareAddr: currentMgmtPortMac})

			netlinkOpsMock.On("LinkByName", "enp3s0f0v0").Return(
				linkMock, nil)
			netlinkOpsMock.On("LinkSetDown", linkMock).Return(nil)
			netlinkOpsMock.On("LinkSetHardwareAddr", linkMock, expectedMgmtPortMac).Return(nil)
			netlinkOpsMock.On("LinkSetName", linkMock, types.K8sMgmtIntfName).Return(nil)
			netlinkOpsMock.On("LinkSetMTU", linkMock, config.Default.MTU).Return(nil)
			netlinkOpsMock.On("LinkSetUp", linkMock).Return(nil, nil)

			// mock createPlatformManagementPort, we fail it as it covers what we want to test without the
			// need to mock the entire flow down to routes and iptable rules.
			netlinkOpsMock.On("LinkByName", mock.Anything).Return(nil, fmt.Errorf(
				"createPlatformManagementPort error"))

			_, err = mgmtPortSnicHost.Create(nil, nil)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("createPlatformManagementPort error"))
		})

		It("Does not configure VF if already configured", func() {
			_, ipnet, err := net.ParseCIDR("192.168.0.1/24")
			Expect(err).ToNot(HaveOccurred())
			_, clusterCidr, err := net.ParseCIDR("192.168.0.0/16")
			Expect(err).ToNot(HaveOccurred())
			expectedMgmtPortMac := util.IPAddrToHWAddr(util.GetNodeManagementIfAddr(ipnet).IP)
			config.Default.MTU = 1400
			config.Default.ClusterSubnets = []config.CIDRNetworkEntry{{CIDR: clusterCidr, HostSubnetLength: 8}}
			mgmtPortSnicHost := managementPortSmartNICHost{
				hostSubnets: []*net.IPNet{ipnet},
				netdevName:  "enp3s0f0v0",
			}
			linkMock := &mocks.Link{}
			linkMock.On("Attrs").Return(&netlink.LinkAttrs{
				Name: "ovn-k8s-mp0", MTU: 1400, HardwareAddr: expectedMgmtPortMac})

			netlinkOpsMock.On("LinkByName", "enp3s0f0v0").Return(
				nil, fmt.Errorf("failed to get link")).Once()
			netlinkOpsMock.On("LinkByName", "ovn-k8s-mp0").Return(
				linkMock, nil).Once()
			netlinkOpsMock.On("LinkSetUp", linkMock).Return(nil, nil).Once()

			// mock createPlatformManagementPort, we fail it as it covers what we want to test without the
			// need to mock the entire flow down to routes and iptable rules.
			netlinkOpsMock.On("LinkByName", mock.Anything).Return(nil, fmt.Errorf(
				"createPlatformManagementPort error")).Once()

			_, err = mgmtPortSnicHost.Create(nil, nil)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(
				"createPlatformManagementPort error"))
		})
	})
})
