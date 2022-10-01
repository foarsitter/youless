# DeviceSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceSchema

`func NewDeviceSchema() *DeviceSchema`

NewDeviceSchema instantiates a new DeviceSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSchemaWithDefaults

`func NewDeviceSchemaWithDefaults() *DeviceSchema`

NewDeviceSchemaWithDefaults instantiates a new DeviceSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *DeviceSchema) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceSchema) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceSchema) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceSchema) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMac

`func (o *DeviceSchema) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceSchema) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceSchema) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceSchema) HasMac() bool`

HasMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


