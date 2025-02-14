// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	net "net"

	mock "github.com/stretchr/testify/mock"

	netlink "github.com/vishvananda/netlink"
)

// NetLinkOps is an autogenerated mock type for the NetLinkOps type
type NetLinkOps struct {
	mock.Mock
}

// AddrAdd provides a mock function with given fields: link, addr
func (_m *NetLinkOps) AddrAdd(link netlink.Link, addr *netlink.Addr) error {
	ret := _m.Called(link, addr)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, *netlink.Addr) error); ok {
		r0 = rf(link, addr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddrDel provides a mock function with given fields: link, addr
func (_m *NetLinkOps) AddrDel(link netlink.Link, addr *netlink.Addr) error {
	ret := _m.Called(link, addr)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, *netlink.Addr) error); ok {
		r0 = rf(link, addr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddrList provides a mock function with given fields: link, family
func (_m *NetLinkOps) AddrList(link netlink.Link, family int) ([]netlink.Addr, error) {
	ret := _m.Called(link, family)

	var r0 []netlink.Addr
	if rf, ok := ret.Get(0).(func(netlink.Link, int) []netlink.Addr); ok {
		r0 = rf(link, family)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]netlink.Addr)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(netlink.Link, int) error); ok {
		r1 = rf(link, family)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConntrackDeleteFilter provides a mock function with given fields: table, family, filter
func (_m *NetLinkOps) ConntrackDeleteFilter(table netlink.ConntrackTableType, family netlink.InetFamily, filter netlink.CustomConntrackFilter) (uint, error) {
	ret := _m.Called(table, family, filter)

	var r0 uint
	if rf, ok := ret.Get(0).(func(netlink.ConntrackTableType, netlink.InetFamily, netlink.CustomConntrackFilter) uint); ok {
		r0 = rf(table, family, filter)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(netlink.ConntrackTableType, netlink.InetFamily, netlink.CustomConntrackFilter) error); ok {
		r1 = rf(table, family, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkByName provides a mock function with given fields: ifaceName
func (_m *NetLinkOps) LinkByName(ifaceName string) (netlink.Link, error) {
	ret := _m.Called(ifaceName)

	var r0 netlink.Link
	if rf, ok := ret.Get(0).(func(string) netlink.Link); ok {
		r0 = rf(ifaceName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(netlink.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ifaceName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkDelete provides a mock function with given fields: link
func (_m *NetLinkOps) LinkDelete(link netlink.Link) error {
	ret := _m.Called(link)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkList provides a mock function with given fields:
func (_m *NetLinkOps) LinkList() ([]netlink.Link, error) {
	ret := _m.Called()

	var r0 []netlink.Link
	if rf, ok := ret.Get(0).(func() []netlink.Link); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]netlink.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkSetDown provides a mock function with given fields: link
func (_m *NetLinkOps) LinkSetDown(link netlink.Link) error {
	ret := _m.Called(link)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetHardwareAddr provides a mock function with given fields: link, hwaddr
func (_m *NetLinkOps) LinkSetHardwareAddr(link netlink.Link, hwaddr net.HardwareAddr) error {
	ret := _m.Called(link, hwaddr)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, net.HardwareAddr) error); ok {
		r0 = rf(link, hwaddr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetMTU provides a mock function with given fields: link, mtu
func (_m *NetLinkOps) LinkSetMTU(link netlink.Link, mtu int) error {
	ret := _m.Called(link, mtu)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int) error); ok {
		r0 = rf(link, mtu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetName provides a mock function with given fields: link, newName
func (_m *NetLinkOps) LinkSetName(link netlink.Link, newName string) error {
	ret := _m.Called(link, newName)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, string) error); ok {
		r0 = rf(link, newName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetNsFd provides a mock function with given fields: link, fd
func (_m *NetLinkOps) LinkSetNsFd(link netlink.Link, fd int) error {
	ret := _m.Called(link, fd)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int) error); ok {
		r0 = rf(link, fd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetTxQLen provides a mock function with given fields: link, qlen
func (_m *NetLinkOps) LinkSetTxQLen(link netlink.Link, qlen int) error {
	ret := _m.Called(link, qlen)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link, int) error); ok {
		r0 = rf(link, qlen)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LinkSetUp provides a mock function with given fields: link
func (_m *NetLinkOps) LinkSetUp(link netlink.Link) error {
	ret := _m.Called(link)

	var r0 error
	if rf, ok := ret.Get(0).(func(netlink.Link) error); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NeighAdd provides a mock function with given fields: neigh
func (_m *NetLinkOps) NeighAdd(neigh *netlink.Neigh) error {
	ret := _m.Called(neigh)

	var r0 error
	if rf, ok := ret.Get(0).(func(*netlink.Neigh) error); ok {
		r0 = rf(neigh)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NeighList provides a mock function with given fields: linkIndex, family
func (_m *NetLinkOps) NeighList(linkIndex int, family int) ([]netlink.Neigh, error) {
	ret := _m.Called(linkIndex, family)

	var r0 []netlink.Neigh
	if rf, ok := ret.Get(0).(func(int, int) []netlink.Neigh); ok {
		r0 = rf(linkIndex, family)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]netlink.Neigh)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(linkIndex, family)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RouteAdd provides a mock function with given fields: route
func (_m *NetLinkOps) RouteAdd(route *netlink.Route) error {
	ret := _m.Called(route)

	var r0 error
	if rf, ok := ret.Get(0).(func(*netlink.Route) error); ok {
		r0 = rf(route)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RouteDel provides a mock function with given fields: route
func (_m *NetLinkOps) RouteDel(route *netlink.Route) error {
	ret := _m.Called(route)

	var r0 error
	if rf, ok := ret.Get(0).(func(*netlink.Route) error); ok {
		r0 = rf(route)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RouteList provides a mock function with given fields: link, family
func (_m *NetLinkOps) RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	ret := _m.Called(link, family)

	var r0 []netlink.Route
	if rf, ok := ret.Get(0).(func(netlink.Link, int) []netlink.Route); ok {
		r0 = rf(link, family)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]netlink.Route)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(netlink.Link, int) error); ok {
		r1 = rf(link, family)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RouteListFiltered provides a mock function with given fields: family, filter, filterMask
func (_m *NetLinkOps) RouteListFiltered(family int, filter *netlink.Route, filterMask uint64) ([]netlink.Route, error) {
	ret := _m.Called(family, filter, filterMask)

	var r0 []netlink.Route
	if rf, ok := ret.Get(0).(func(int, *netlink.Route, uint64) []netlink.Route); ok {
		r0 = rf(family, filter, filterMask)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]netlink.Route)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *netlink.Route, uint64) error); ok {
		r1 = rf(family, filter, filterMask)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
