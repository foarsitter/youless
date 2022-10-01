# UploadedValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tm** | Pointer to **int32** | unix-time-format (1489333828 &#x3D;&gt; Sun, 12 Mar 2017 15:50:28 GMT)  | [optional] 
**Net** | Pointer to **float32** | Netto counter, as displayed in the web-interface of the LS-120. It seems equal to: p1 + p2 - n1 - n2 Perhaps also includes some user set offset.  | [optional] 
**Pwr** | Pointer to **int32** | Actual power use in Watt (can be negative) | [optional] 
**P1** | Pointer to **float32** | P1 consumption counter (low tariff) | [optional] 
**P2** | Pointer to **float32** | P2 consumption counter (high tariff) | [optional] 
**N1** | Pointer to **float32** | N1 production counter (low tariff) | [optional] 
**N2** | Pointer to **float32** | N2 production counter (high tariff) | [optional] 
**Gas** | Pointer to **float32** | counter gas-meter (in m^3) | [optional] 
**Cs0** | Pointer to **float32** | kWh counter of S0 input  | [optional] 
**Ps0** | Pointer to **int32** | Computed power  | [optional] 
**Ts0** | Pointer to **int32** | Unix timestamp of the last S0 measurement.  | [optional] 
**Gts** | Pointer to **int32** | Last timestamp created by the &#39;smart meter&#39;. \&quot;1711032100\&quot; &#x3D; 2017/11/03 21:00 (yyMMddhhmm) Can be used to see if P1 communication fails.  | [optional] 

## Methods

### NewUploadedValues

`func NewUploadedValues() *UploadedValues`

NewUploadedValues instantiates a new UploadedValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadedValuesWithDefaults

`func NewUploadedValuesWithDefaults() *UploadedValues`

NewUploadedValuesWithDefaults instantiates a new UploadedValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTm

`func (o *UploadedValues) GetTm() int32`

GetTm returns the Tm field if non-nil, zero value otherwise.

### GetTmOk

`func (o *UploadedValues) GetTmOk() (*int32, bool)`

GetTmOk returns a tuple with the Tm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTm

`func (o *UploadedValues) SetTm(v int32)`

SetTm sets Tm field to given value.

### HasTm

`func (o *UploadedValues) HasTm() bool`

HasTm returns a boolean if a field has been set.

### GetNet

`func (o *UploadedValues) GetNet() float32`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *UploadedValues) GetNetOk() (*float32, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *UploadedValues) SetNet(v float32)`

SetNet sets Net field to given value.

### HasNet

`func (o *UploadedValues) HasNet() bool`

HasNet returns a boolean if a field has been set.

### GetPwr

`func (o *UploadedValues) GetPwr() int32`

GetPwr returns the Pwr field if non-nil, zero value otherwise.

### GetPwrOk

`func (o *UploadedValues) GetPwrOk() (*int32, bool)`

GetPwrOk returns a tuple with the Pwr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwr

`func (o *UploadedValues) SetPwr(v int32)`

SetPwr sets Pwr field to given value.

### HasPwr

`func (o *UploadedValues) HasPwr() bool`

HasPwr returns a boolean if a field has been set.

### GetP1

`func (o *UploadedValues) GetP1() float32`

GetP1 returns the P1 field if non-nil, zero value otherwise.

### GetP1Ok

`func (o *UploadedValues) GetP1Ok() (*float32, bool)`

GetP1Ok returns a tuple with the P1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP1

`func (o *UploadedValues) SetP1(v float32)`

SetP1 sets P1 field to given value.

### HasP1

`func (o *UploadedValues) HasP1() bool`

HasP1 returns a boolean if a field has been set.

### GetP2

`func (o *UploadedValues) GetP2() float32`

GetP2 returns the P2 field if non-nil, zero value otherwise.

### GetP2Ok

`func (o *UploadedValues) GetP2Ok() (*float32, bool)`

GetP2Ok returns a tuple with the P2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2

`func (o *UploadedValues) SetP2(v float32)`

SetP2 sets P2 field to given value.

### HasP2

`func (o *UploadedValues) HasP2() bool`

HasP2 returns a boolean if a field has been set.

### GetN1

`func (o *UploadedValues) GetN1() float32`

GetN1 returns the N1 field if non-nil, zero value otherwise.

### GetN1Ok

`func (o *UploadedValues) GetN1Ok() (*float32, bool)`

GetN1Ok returns a tuple with the N1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1

`func (o *UploadedValues) SetN1(v float32)`

SetN1 sets N1 field to given value.

### HasN1

`func (o *UploadedValues) HasN1() bool`

HasN1 returns a boolean if a field has been set.

### GetN2

`func (o *UploadedValues) GetN2() float32`

GetN2 returns the N2 field if non-nil, zero value otherwise.

### GetN2Ok

`func (o *UploadedValues) GetN2Ok() (*float32, bool)`

GetN2Ok returns a tuple with the N2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2

`func (o *UploadedValues) SetN2(v float32)`

SetN2 sets N2 field to given value.

### HasN2

`func (o *UploadedValues) HasN2() bool`

HasN2 returns a boolean if a field has been set.

### GetGas

`func (o *UploadedValues) GetGas() float32`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *UploadedValues) GetGasOk() (*float32, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *UploadedValues) SetGas(v float32)`

SetGas sets Gas field to given value.

### HasGas

`func (o *UploadedValues) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetCs0

`func (o *UploadedValues) GetCs0() float32`

GetCs0 returns the Cs0 field if non-nil, zero value otherwise.

### GetCs0Ok

`func (o *UploadedValues) GetCs0Ok() (*float32, bool)`

GetCs0Ok returns a tuple with the Cs0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCs0

`func (o *UploadedValues) SetCs0(v float32)`

SetCs0 sets Cs0 field to given value.

### HasCs0

`func (o *UploadedValues) HasCs0() bool`

HasCs0 returns a boolean if a field has been set.

### GetPs0

`func (o *UploadedValues) GetPs0() int32`

GetPs0 returns the Ps0 field if non-nil, zero value otherwise.

### GetPs0Ok

`func (o *UploadedValues) GetPs0Ok() (*int32, bool)`

GetPs0Ok returns a tuple with the Ps0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPs0

`func (o *UploadedValues) SetPs0(v int32)`

SetPs0 sets Ps0 field to given value.

### HasPs0

`func (o *UploadedValues) HasPs0() bool`

HasPs0 returns a boolean if a field has been set.

### GetTs0

`func (o *UploadedValues) GetTs0() int32`

GetTs0 returns the Ts0 field if non-nil, zero value otherwise.

### GetTs0Ok

`func (o *UploadedValues) GetTs0Ok() (*int32, bool)`

GetTs0Ok returns a tuple with the Ts0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs0

`func (o *UploadedValues) SetTs0(v int32)`

SetTs0 sets Ts0 field to given value.

### HasTs0

`func (o *UploadedValues) HasTs0() bool`

HasTs0 returns a boolean if a field has been set.

### GetGts

`func (o *UploadedValues) GetGts() int32`

GetGts returns the Gts field if non-nil, zero value otherwise.

### GetGtsOk

`func (o *UploadedValues) GetGtsOk() (*int32, bool)`

GetGtsOk returns a tuple with the Gts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGts

`func (o *UploadedValues) SetGts(v int32)`

SetGts sets Gts field to given value.

### HasGts

`func (o *UploadedValues) HasGts() bool`

HasGts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


