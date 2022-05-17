package metadata

type EventDefType int32

const (
	EventDef_InfoEvent    EventDefType = 0
	EventDef_WarningEvent EventDefType = 1
	EventDef_AlertEvent   EventDefType = 2
)

// Represents the type and data validation of a property.
// Only one of its members may be specified.
type PropertyType struct {
	// +optional
	Int *PropertyTypeInt64 `json:"int,omitempty"`
	// +optional
	String *PropertyTypeString `json:"string,omitempty"`
	// +optional
	Double *PropertyTypeDouble `json:"double,omitempty"`
	// +optional
	Float *PropertyTypeFloat `json:"float,omitempty"`
	// +optional
	Boolean *PropertyTypeBoolean `json:"boolean,omitempty"`
	// +optional
	Bytes *PropertyTypeBytes `json:"bytes,omitempty"`
}

type PropertyTypeInt64 struct {
	// +optional
	DefaultValue int64 `json:"defaultValue,omitempty"`
	// +optional
	Minimum int64 `json:"minimum,omitempty"`
	// +optional
	Maximum int64 `json:"maximum,omitempty"`
	// The unit of the property
	// +optional
	Unit string `json:"unit,omitempty"`
}

type PropertyTypeString struct {
	// +optional
	DefaultValue string `json:"defaultValue,omitempty"`
}

type PropertyTypeDouble struct {
	// +optional
	DefaultValue float64 `json:"defaultValue,omitempty"`
	// +optional
	Minimum float64 `json:"minimum,omitempty"`
	// +optional
	Maximum float64 `json:"maximum,omitempty"`
	// The unit of the property
	// +optional
	Unit string `json:"unit,omitempty"`
}

type PropertyTypeFloat struct {
	// +optional
	DefaultValue float32 `json:"defaultValue,omitempty"`
	// +optional
	Minimum float32 `json:"minimum,omitempty"`
	// +optional
	Maximum float32 `json:"maximum,omitempty"`
	// The unit of the property
	// +optional
	Unit string `json:"unit,omitempty"`
}

type PropertyTypeBoolean struct {
	// +optional
	DefaultValue bool `json:"defaultValue,omitempty"`
}

type PropertyTypeBytes struct {
}

// device command param defination
type RequestParamDef struct {
	DataType PropertyType `json:"DataType,omitempty"`
}

// The access mode for  a device property.
// +kubebuilder:validation:Enum=ReadWrite;ReadOnly
type PropertyAccessMode string

// Access mode constants for a device property.
const (
	ReadWrite PropertyAccessMode = "ReadWrite"
	ReadOnly  PropertyAccessMode = "ReadOnly"
)

// device property defination
type DevicePropertyModel struct {
	// Required: The device property name.
	Name        string  `form:"name" json:"name" binding:"required"`
	WriteAble   bool    `form:"writeAble" json:"writeAble"`
	MaxValue    float64 `form:"maxValue" json:"maxValue"`
	MinValue    float64 `form:"minValue" json:"minValue"`
	Unit        string  `form:"unit" json:"unit"`
	DataType    string  `form:"dataType" json:"dataType"`
	Description string  `form:"description" json:"description"`
}

type DeviceCommandModel struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	// params
	RequestParam map[string]string `json:"RequestParam,omitempty"`
}

// device event defination
type DeviceEventModel struct {
	Name        string  `form:"name" json:"name"  binding:"required"`
	EventType   string  `form:"eventType" json:"eventType"  binding:"required"`
	MaxValue    float64 `form:"maxValue" json:"maxValue,omitempty"`
	MinValue    float64 `form:"minValue" json:"minValue,omitempty"`
	Unit        string  `form:"unit" json:"unit,omitempty"`
	DataType    string  `form:"dataType" json:"dataType,omitempty"`
	Description string  `form:"description" json:"description,omitempty"`
}

//service describe a based module contains some properties, events
// and commands.
type DeviceServiceModel struct {
	Name        string                 `form:"name" json:"name" binding:"required"`
	Description string                 `form:"description" json:"description"`
	Properties  []*DevicePropertyModel `json:"properties,omitempty"`
	Events      []*DeviceEventModel    `json:"events,omitempty"`
	Commands    []*DeviceCommandModel  `json:"commands,omitempty"`
}

// DeviceModelSpec defines the model / template for a device.It is a blueprint which describes the device
// capabilities.
type DeviceModelSpec struct {
	// Required: List of device services.
	Services []*DeviceServiceModel `json:"Services,omitempty"`
}

// DeviceModel is the Schema for the device model API
type DeviceModel struct {
	Name         string `form:"name" json:"name" binding:"required"`
	Description  string `form:"description" json:"description,omitempty"`
	Manufacturer string `form:"manufacturer" json:"manufacturer,omitempty"`
	Industry     string `form:"industry" json:"industry,omitempty"`
	DataFormat   string `form:"dataFormat" json:"dataFormat,omitempty"`
	DeviceNumber int64  `json:"deviceNumber,omitempty"`
	TagNumber    int64  `json:"tagNumber,omitempty"`
	GroupID      string `json:"groupId,omitempty"`
	//who create the device by ID.
	Creator         string `form:"creator" json:"creator"`
	CreateTimeStamp int64  `json:"createTimeStamp,omitempty"`
	UpdateTimeStamp int64  `json:"updateTimeStamp,omitempty"`
	//Spec.
	Spec DeviceModelSpec `json:"spec,omitempty"`
}

// DeviceModelList contains a list of DeviceModel
type DeviceModelList struct {
	Items []DeviceModel `json:"items"`
}

// Device instance access config
type InstanceConfig struct {
	// Required: The device property name.
	//device Instance property instance event instance command instance
	Name         string `form:"name" json:"name" binding:"required"`
	DeviceName   string `form:"deviceName" json:"deviceName"`
	ServiceName  string `form:"serviceName" json:"serviceName`
	AccessConfig string `form:"accessConfig" json:"accessConfig"`
	InstanceType string `form:"type" json:"type"`
}
