package metadata

/*
* DeviceSpecMeta
* this is for edge device or edge gateway
 */
type DeviceSpecMeta struct {
	DeviceID                 string `json:"DeviceID"`
	DeviceOS                 string `json:"deviceOS,omitempty"`
	DeviceCatagory           string `json:"deviceCatagory,omitempty"`
	DeviceIdentificationCode string `json:"deviceIdentificationCode,omitempty"`

	State string `json:state,omitempty`
	// Additional metadata like tags.
	// +optional
	Tags map[string]string `json:"tags,omitempty"`

	// Required: The protocol configuration used to connect to the device.
	//this should be a json string
	Protocol string `json:"protocol,omitempty"`

	// Required: List of device services.
	Services []*DeviceServiceSpec `json:"Services,omitempty"`
}

/*
* Single device report message.
 */
type ReportDeviceMessage struct {
	DeviceID string `json:"device_id"`
	// Required: List of device services.
	Services []*TwinProperty `json:"services"`
}

/*
* Devices report message.
 */
type ReportDevicesMessage struct {
	Devices []*ReportDeviceMessage `json:"devices,omitempty"`
}

const (
	DEVICE_STATUS_ONLINE  = "online"
	DEVICE_STATUS_OFFLINE = "offline"
)

// single device status message
type DeviceStatusMessage struct {
	DeviceID string `json:"device_id"`
	//Required:
	//device status
	// offline: device offline online: device online
	Status string `json:"status"`
	// some error message info of this device.
	// +optional
	ErrorMessage string `json:"error_message,omitempty"`
}

// devices status message
type DevicesStatusMessage struct {
	DevicesStatus []*DeviceStatusMessage `json:"devices_status"`
}
