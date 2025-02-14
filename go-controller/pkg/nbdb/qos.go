// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

type (
	QoSAction    = string
	QoSBandwidth = string
	QoSDirection = string
)

var (
	QoSActionDSCP         QoSAction    = "dscp"
	QoSBandwidthBurst     QoSBandwidth = "burst"
	QoSBandwidthRate      QoSBandwidth = "rate"
	QoSDirectionFromLport QoSDirection = "from-lport"
	QoSDirectionToLport   QoSDirection = "to-lport"
)

// QoS defines an object in QoS table
type QoS struct {
	UUID        string            `ovsdb:"_uuid"`
	Action      map[string]int    `ovsdb:"action"`
	Bandwidth   map[string]int    `ovsdb:"bandwidth"`
	Direction   QoSDirection      `ovsdb:"direction"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	Match       string            `ovsdb:"match"`
	Priority    int               `ovsdb:"priority"`
}
