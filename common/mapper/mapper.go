package mapper

type Mapper struct {
	//Name: specifies the mapper name.
	Type string `json:"type"`

	//ID: specifies the mapper ID.
	Spec string `json:"spec"`
}

type DeviceStatus struct {
	DeviceID  string `json:"device_id"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

type DevicesStatus struct {
	DeviceStatus []*DeviceStatus `json:"devices_status"`
}
