package metadata

import (
	"time"
)

// DeviceTwin provides a logical representation of A series of control properties on a device(writable
// properties in the device model). These properties can have a Desired state and a Reported state.
// The cloud configures the `Desired`state of a device's properties and this configuration update is pushed
// to the edge node. The mapper sends a command to the device to change these properties value as per
// the desired state . It receives the `Reported` state of the property once the previous operation is
// complete and sends the reported state to the cloud.
// Offline device interaction in the edge is possible via twin properties for control/command operations.
type DeviceTwin struct {
	// Required: the desired property list.
	Desired []*TwinProperty `json:"desired,omitempty"`
	// Required: the reported property list.
	Reported []*TwinProperty `json:"reported,omitempty"`
}

func NewDeviceTwin(desired, reported []*TwinProperty) *DeviceTwin {
	return &DeviceTwin{
		Desired:  desired,
		Reported: reported,
	}
}

// TwinProperty represents the device property for which an Expected/Actual state can be defined.
type TwinProperty struct {
	Service string `json:"service"`
	// Required: The property name for which the desired/reported values are specified.
	// This property should be present in the device model.
	PropertyName string `json:"property_name"`
	// Required: The value for this property.
	Value interface{} `json:"value,omitempty"`
	//the timestamp of the property collecting value
	// +optional
	Timestamp int64 `json:"timestamp,omitempty"`
	// some error message info when collecting this property.
	// +optional, just for reported.
	ErrorMessage string `json:"error_message,omitempty"`
}

//New device twin property.
func NewTwinProperty(svc, propName, errMsg string, val interface{}) *TwinProperty {
	now := time.Now().Unix()

	return &TwinProperty{
		Service:      svc,
		PropertyName: propName,
		Value:        val,
		Timestamp:    now,
		ErrorMessage: errMsg,
	}
}

// DeviceSpec represents a single device instance. It is an instantation of a device model.
type DeviceSpec struct {
	//Required: device name, changed by user
	Name string `json:"Name"`

	//Required: the device belong to which edge.
	EdgeID                   string `form:"edgeID" json:"edgeID"`
	DeviceOS                 string `form:"deviceOS" json:"deviceOS,omitempty"`
	DeviceCatagory           string `form:"deviceCatagory" json:"deviceCatagory,omitempty"`
	DeviceVersion            int    `form:"deviceVersion" json:"deviceVersion,omitempty"`
	DeviceIdentificationCode string `form:"deviceIdentificationCode" json:"deviceIdentificationCode,omitempty"`
	Description              string `form:"description" json:"description"`

	//group.
	//TODO: reserved in future
	GroupName string `form:"groupName" json:"groupName,omitempty"`
	//who create the device by ID.
	Creator string `form:"creator" json:"creator,omitempty"`

	// +optional
	//TODO: reserved in future
	DeviceAuthType string `form:"deviceAuthType" json:"deviceAuthType,omitempty"`
	Secret         string `form:"secret" json:"secret,omitempty"`

	//Device Type
	//Kind: Direct, GateWay, SubDevcie.
	DeviceType  string `form:"deviceType" json:"deviceType,omitempty"`
	GatewayID   string `form:"gatewayId" json:"gatewayId,omitempty"`
	GatewayName string `form:"gatewayName" json:"gatewayName,omitempty"`
	// Additional metadata like tags.
	// +optional
	Tags map[string]string `form:"tags" json:"tags,omitempty"`

	// Required: DeviceModelRef is reference to the device model used as a template
	// to create the device instance.
	//this should be the Name of device model.
	DeviceModelRef string `form:"deviceModelRef" json:"deviceModelRef,omitempty"`

	ProtocolType string `form:"protocolType" json:"protocolType,omitempty"`
	// Required: The protocol configuration used to connect to the device.
	//this should be a json string
	Protocol string `form:"protocol" json:"protocol,omitempty"`

	// ExtensionConfig which describe how to access the device properties,command, and events.
	// +optional
	ExtensionConfig *ExtensionConfig `form:"extensionConfig" json:"extensionConfig,omitempty"`
}

type ExtensionConfig struct {
	// Required: List of device services.
	Services []*DeviceServiceSpec `json:"services,omitempty"`
}

func (ec *ExtensionConfig) FindDeviceServiceSpec(name string) *DeviceServiceSpec {
	for _, dss := range ec.Services {
		if dss == nil {
			continue
		}

		if dss.Name == name {
			return dss
		}
	}

	return nil
}

// DeviceServiceSpec is the  an instantation of a DeviceServiceModel.
type DeviceServiceSpec struct {
	Name       string                `json:"name"`
	Properties []*DevicePropertySpec `json:"properties,omitempty"`
	Events     []*DeviceEventSpec    `json:"events,omitempty"`
	Commands   []*DeviceCommandSpec  `json:"commands,omitempty"`
}

func (dss *DeviceServiceSpec) FindDevicePropertySpec(name string) *DevicePropertySpec {
	for _, p := range dss.Properties {
		if p.Name == name {
			return p
		}
	}

	return nil
}

func (dss *DeviceServiceSpec) FindDeviceEventSpec(name string) *DeviceEventSpec {
	for _, e := range dss.Events {
		if e.Name == name {
			return e
		}
	}

	return nil
}

func (dss *DeviceServiceSpec) FindDeviceCommandSpec(name string) *DeviceCommandSpec {
	for _, c := range dss.Commands {
		if c.Name == name {
			return c
		}
	}

	return nil
}

// DevicePropertySpec is an instantation of a DevicePropertyModel.
type DevicePropertySpec struct {
	*DevicePropertyModel `json:",inline"`
	// List of AccessConfig which describe how to access the device properties,command, and events.
	// AccessConfig must unique by AccessConfig.propertyName.
	// +optional
	//this should be a json string
	// AccessConfig must unique by AccessConfig.propertyName.
	AccessConfig string `json:"accessConfig,omitempty"`
}

// DeviceEventSpec is an instantation of a DeviceEventModel.
type DeviceEventSpec struct {
	*DeviceEventModel `json:",inline"`
	// List of AccessConfig which describe how to access the device properties,command, and events.
	// AccessConfig must unique by AccessConfig.propertyName.
	// +optional
	//this should be a json string
	// AccessConfig must unique by AccessConfig.propertyName.
	AccessConfig string `json:"accessConfig,omitempty"`
}

// DeviceCommandSpec is an instantation of a DeviceCommandModel.
type DeviceCommandSpec struct {
	*DeviceCommandModel `json:",inline"`
	// List of AccessConfig which describe how to access the device properties,command, and events.
	// AccessConfig must unique by AccessConfig.propertyName.
	// +optional
	//this should be a json string
	// AccessConfig must unique by AccessConfig.propertyName.
	AccessConfig string `json:"accessConfig,omitempty"`
}

// DeviceStatus reports the device state and the desired/reported values of twin attributes.
type DeviceStatus struct {
	//device status
	// inactive, active, online, offline, error etc.
	DeviceStatus string `json:"deviceStatus"`
	//start/stop collecting flag.
	Collecting bool `json:"collecting"`
	// DeviceTwin provides a logical representation of A series of control properties on a device
	// Required:
	Twins *DeviceTwin `json:"twins"`
}

func NewDeviceStatus() *DeviceStatus {
	return &DeviceStatus{
		DeviceStatus: "inactive",
		Collecting:   false,
	}
}

func (ds *DeviceStatus) StartCollect() error {
	ds.Collecting = true

	return nil
}

// Device is the Schema for the devices API
type Device struct {
	DeviceID        string `json:"DeviceID"`
	GroupID         string `json:"groupId,omitempty"`
	CreateTimeStamp int64  `json:"createTimeStamp,omitempty"`
	UpdateTimeStamp int64  `json:"updateTimeStamp,omitempty"`

	Spec   *DeviceSpec   `json:"spec,omitempty"`
	Status *DeviceStatus `json:"status,omitempty"`
}

// DeviceList contains a list of Device
type DeviceList struct {
	Items []*Device `json:"items"`
}
