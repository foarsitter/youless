# BasicStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cnt** | Pointer to **string** | counter in kWh | [optional] 
**Pwr** | Pointer to **int32** | Power consumption in Watt | [optional] 
**Lvl** | Pointer to **int32** | moving average level (intensity of reflected light on analog meters) | [optional] 
**Dev** | Pointer to **string** | deviation of reflection | [optional] 
**Det** | Pointer to **string** |  | [optional] 
**Con** | Pointer to **string** | connection status | [optional] 
**Sts** | Pointer to **string** | Time until next status update with online monitoring | [optional] 
**Cs0** | Pointer to **string** | kWh counter of S0 input | [optional] 
**Ps0** | Pointer to **int32** | Computed power | [optional] 
**Raw** | Pointer to **int32** | raw 10-bit light reflection level (without averaging) | [optional] 

## Methods

### NewBasicStatusInfo

`func NewBasicStatusInfo() *BasicStatusInfo`

NewBasicStatusInfo instantiates a new BasicStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicStatusInfoWithDefaults

`func NewBasicStatusInfoWithDefaults() *BasicStatusInfo`

NewBasicStatusInfoWithDefaults instantiates a new BasicStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnt

`func (o *BasicStatusInfo) GetCnt() string`

GetCnt returns the Cnt field if non-nil, zero value otherwise.

### GetCntOk

`func (o *BasicStatusInfo) GetCntOk() (*string, bool)`

GetCntOk returns a tuple with the Cnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnt

`func (o *BasicStatusInfo) SetCnt(v string)`

SetCnt sets Cnt field to given value.

### HasCnt

`func (o *BasicStatusInfo) HasCnt() bool`

HasCnt returns a boolean if a field has been set.

### GetPwr

`func (o *BasicStatusInfo) GetPwr() int32`

GetPwr returns the Pwr field if non-nil, zero value otherwise.

### GetPwrOk

`func (o *BasicStatusInfo) GetPwrOk() (*int32, bool)`

GetPwrOk returns a tuple with the Pwr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwr

`func (o *BasicStatusInfo) SetPwr(v int32)`

SetPwr sets Pwr field to given value.

### HasPwr

`func (o *BasicStatusInfo) HasPwr() bool`

HasPwr returns a boolean if a field has been set.

### GetLvl

`func (o *BasicStatusInfo) GetLvl() int32`

GetLvl returns the Lvl field if non-nil, zero value otherwise.

### GetLvlOk

`func (o *BasicStatusInfo) GetLvlOk() (*int32, bool)`

GetLvlOk returns a tuple with the Lvl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvl

`func (o *BasicStatusInfo) SetLvl(v int32)`

SetLvl sets Lvl field to given value.

### HasLvl

`func (o *BasicStatusInfo) HasLvl() bool`

HasLvl returns a boolean if a field has been set.

### GetDev

`func (o *BasicStatusInfo) GetDev() string`

GetDev returns the Dev field if non-nil, zero value otherwise.

### GetDevOk

`func (o *BasicStatusInfo) GetDevOk() (*string, bool)`

GetDevOk returns a tuple with the Dev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev

`func (o *BasicStatusInfo) SetDev(v string)`

SetDev sets Dev field to given value.

### HasDev

`func (o *BasicStatusInfo) HasDev() bool`

HasDev returns a boolean if a field has been set.

### GetDet

`func (o *BasicStatusInfo) GetDet() string`

GetDet returns the Det field if non-nil, zero value otherwise.

### GetDetOk

`func (o *BasicStatusInfo) GetDetOk() (*string, bool)`

GetDetOk returns a tuple with the Det field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDet

`func (o *BasicStatusInfo) SetDet(v string)`

SetDet sets Det field to given value.

### HasDet

`func (o *BasicStatusInfo) HasDet() bool`

HasDet returns a boolean if a field has been set.

### GetCon

`func (o *BasicStatusInfo) GetCon() string`

GetCon returns the Con field if non-nil, zero value otherwise.

### GetConOk

`func (o *BasicStatusInfo) GetConOk() (*string, bool)`

GetConOk returns a tuple with the Con field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCon

`func (o *BasicStatusInfo) SetCon(v string)`

SetCon sets Con field to given value.

### HasCon

`func (o *BasicStatusInfo) HasCon() bool`

HasCon returns a boolean if a field has been set.

### GetSts

`func (o *BasicStatusInfo) GetSts() string`

GetSts returns the Sts field if non-nil, zero value otherwise.

### GetStsOk

`func (o *BasicStatusInfo) GetStsOk() (*string, bool)`

GetStsOk returns a tuple with the Sts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSts

`func (o *BasicStatusInfo) SetSts(v string)`

SetSts sets Sts field to given value.

### HasSts

`func (o *BasicStatusInfo) HasSts() bool`

HasSts returns a boolean if a field has been set.

### GetCs0

`func (o *BasicStatusInfo) GetCs0() string`

GetCs0 returns the Cs0 field if non-nil, zero value otherwise.

### GetCs0Ok

`func (o *BasicStatusInfo) GetCs0Ok() (*string, bool)`

GetCs0Ok returns a tuple with the Cs0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCs0

`func (o *BasicStatusInfo) SetCs0(v string)`

SetCs0 sets Cs0 field to given value.

### HasCs0

`func (o *BasicStatusInfo) HasCs0() bool`

HasCs0 returns a boolean if a field has been set.

### GetPs0

`func (o *BasicStatusInfo) GetPs0() int32`

GetPs0 returns the Ps0 field if non-nil, zero value otherwise.

### GetPs0Ok

`func (o *BasicStatusInfo) GetPs0Ok() (*int32, bool)`

GetPs0Ok returns a tuple with the Ps0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPs0

`func (o *BasicStatusInfo) SetPs0(v int32)`

SetPs0 sets Ps0 field to given value.

### HasPs0

`func (o *BasicStatusInfo) HasPs0() bool`

HasPs0 returns a boolean if a field has been set.

### GetRaw

`func (o *BasicStatusInfo) GetRaw() int32`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *BasicStatusInfo) GetRawOk() (*int32, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *BasicStatusInfo) SetRaw(v int32)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *BasicStatusInfo) HasRaw() bool`

HasRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


