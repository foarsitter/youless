/*
Youless energy monitor GO client

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package youless

import (
	"encoding/json"
)

// UploadedValues struct for UploadedValues
type UploadedValues struct {
	// unix-time-format (1489333828 => Sun, 12 Mar 2017 15:50:28 GMT) 
	Tm *int64 `json:"tm,omitempty"`
	// Netto counter, as displayed in the web-interface of the LS-120. It seems equal to: p1 + p2 - n1 - n2 Perhaps also includes some user set offset. 
	Net *float32 `json:"net,omitempty"`
	// Actual power use in Watt (can be negative)
	Pwr *int32 `json:"pwr,omitempty"`
	// P1 consumption counter (low tariff)
	P1 *float32 `json:"p1,omitempty"`
	// P2 consumption counter (high tariff)
	P2 *float32 `json:"p2,omitempty"`
	// N1 production counter (low tariff)
	N1 *float32 `json:"n1,omitempty"`
	// N2 production counter (high tariff)
	N2 *float32 `json:"n2,omitempty"`
	// counter gas-meter (in m^3)
	Gas *float32 `json:"gas,omitempty"`
	// kWh counter of S0 input 
	Cs0 *float32 `json:"cs0,omitempty"`
	// Computed power 
	Ps0 *int32 `json:"ps0,omitempty"`
	// Unix timestamp of the last S0 measurement. 
	Ts0 *int64 `json:"ts0,omitempty"`
	// Last timestamp created by the 'smart meter'. \"1711032100\" = 2017/11/03 21:00 (yyMMddhhmm) Can be used to see if P1 communication fails. 
	Gts *int64 `json:"gts,omitempty"`
}

// NewUploadedValues instantiates a new UploadedValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadedValues() *UploadedValues {
	this := UploadedValues{}
	return &this
}

// NewUploadedValuesWithDefaults instantiates a new UploadedValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadedValuesWithDefaults() *UploadedValues {
	this := UploadedValues{}
	return &this
}

// GetTm returns the Tm field value if set, zero value otherwise.
func (o *UploadedValues) GetTm() int64 {
	if o == nil || o.Tm == nil {
		var ret int64
		return ret
	}
	return *o.Tm
}

// GetTmOk returns a tuple with the Tm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetTmOk() (*int64, bool) {
	if o == nil || o.Tm == nil {
		return nil, false
	}
	return o.Tm, true
}

// HasTm returns a boolean if a field has been set.
func (o *UploadedValues) HasTm() bool {
	if o != nil && o.Tm != nil {
		return true
	}

	return false
}

// SetTm gets a reference to the given int64 and assigns it to the Tm field.
func (o *UploadedValues) SetTm(v int64) {
	o.Tm = &v
}

// GetNet returns the Net field value if set, zero value otherwise.
func (o *UploadedValues) GetNet() float32 {
	if o == nil || o.Net == nil {
		var ret float32
		return ret
	}
	return *o.Net
}

// GetNetOk returns a tuple with the Net field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetNetOk() (*float32, bool) {
	if o == nil || o.Net == nil {
		return nil, false
	}
	return o.Net, true
}

// HasNet returns a boolean if a field has been set.
func (o *UploadedValues) HasNet() bool {
	if o != nil && o.Net != nil {
		return true
	}

	return false
}

// SetNet gets a reference to the given float32 and assigns it to the Net field.
func (o *UploadedValues) SetNet(v float32) {
	o.Net = &v
}

// GetPwr returns the Pwr field value if set, zero value otherwise.
func (o *UploadedValues) GetPwr() int32 {
	if o == nil || o.Pwr == nil {
		var ret int32
		return ret
	}
	return *o.Pwr
}

// GetPwrOk returns a tuple with the Pwr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetPwrOk() (*int32, bool) {
	if o == nil || o.Pwr == nil {
		return nil, false
	}
	return o.Pwr, true
}

// HasPwr returns a boolean if a field has been set.
func (o *UploadedValues) HasPwr() bool {
	if o != nil && o.Pwr != nil {
		return true
	}

	return false
}

// SetPwr gets a reference to the given int32 and assigns it to the Pwr field.
func (o *UploadedValues) SetPwr(v int32) {
	o.Pwr = &v
}

// GetP1 returns the P1 field value if set, zero value otherwise.
func (o *UploadedValues) GetP1() float32 {
	if o == nil || o.P1 == nil {
		var ret float32
		return ret
	}
	return *o.P1
}

// GetP1Ok returns a tuple with the P1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetP1Ok() (*float32, bool) {
	if o == nil || o.P1 == nil {
		return nil, false
	}
	return o.P1, true
}

// HasP1 returns a boolean if a field has been set.
func (o *UploadedValues) HasP1() bool {
	if o != nil && o.P1 != nil {
		return true
	}

	return false
}

// SetP1 gets a reference to the given float32 and assigns it to the P1 field.
func (o *UploadedValues) SetP1(v float32) {
	o.P1 = &v
}

// GetP2 returns the P2 field value if set, zero value otherwise.
func (o *UploadedValues) GetP2() float32 {
	if o == nil || o.P2 == nil {
		var ret float32
		return ret
	}
	return *o.P2
}

// GetP2Ok returns a tuple with the P2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetP2Ok() (*float32, bool) {
	if o == nil || o.P2 == nil {
		return nil, false
	}
	return o.P2, true
}

// HasP2 returns a boolean if a field has been set.
func (o *UploadedValues) HasP2() bool {
	if o != nil && o.P2 != nil {
		return true
	}

	return false
}

// SetP2 gets a reference to the given float32 and assigns it to the P2 field.
func (o *UploadedValues) SetP2(v float32) {
	o.P2 = &v
}

// GetN1 returns the N1 field value if set, zero value otherwise.
func (o *UploadedValues) GetN1() float32 {
	if o == nil || o.N1 == nil {
		var ret float32
		return ret
	}
	return *o.N1
}

// GetN1Ok returns a tuple with the N1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetN1Ok() (*float32, bool) {
	if o == nil || o.N1 == nil {
		return nil, false
	}
	return o.N1, true
}

// HasN1 returns a boolean if a field has been set.
func (o *UploadedValues) HasN1() bool {
	if o != nil && o.N1 != nil {
		return true
	}

	return false
}

// SetN1 gets a reference to the given float32 and assigns it to the N1 field.
func (o *UploadedValues) SetN1(v float32) {
	o.N1 = &v
}

// GetN2 returns the N2 field value if set, zero value otherwise.
func (o *UploadedValues) GetN2() float32 {
	if o == nil || o.N2 == nil {
		var ret float32
		return ret
	}
	return *o.N2
}

// GetN2Ok returns a tuple with the N2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetN2Ok() (*float32, bool) {
	if o == nil || o.N2 == nil {
		return nil, false
	}
	return o.N2, true
}

// HasN2 returns a boolean if a field has been set.
func (o *UploadedValues) HasN2() bool {
	if o != nil && o.N2 != nil {
		return true
	}

	return false
}

// SetN2 gets a reference to the given float32 and assigns it to the N2 field.
func (o *UploadedValues) SetN2(v float32) {
	o.N2 = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *UploadedValues) GetGas() float32 {
	if o == nil || o.Gas == nil {
		var ret float32
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetGasOk() (*float32, bool) {
	if o == nil || o.Gas == nil {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *UploadedValues) HasGas() bool {
	if o != nil && o.Gas != nil {
		return true
	}

	return false
}

// SetGas gets a reference to the given float32 and assigns it to the Gas field.
func (o *UploadedValues) SetGas(v float32) {
	o.Gas = &v
}

// GetCs0 returns the Cs0 field value if set, zero value otherwise.
func (o *UploadedValues) GetCs0() float32 {
	if o == nil || o.Cs0 == nil {
		var ret float32
		return ret
	}
	return *o.Cs0
}

// GetCs0Ok returns a tuple with the Cs0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetCs0Ok() (*float32, bool) {
	if o == nil || o.Cs0 == nil {
		return nil, false
	}
	return o.Cs0, true
}

// HasCs0 returns a boolean if a field has been set.
func (o *UploadedValues) HasCs0() bool {
	if o != nil && o.Cs0 != nil {
		return true
	}

	return false
}

// SetCs0 gets a reference to the given float32 and assigns it to the Cs0 field.
func (o *UploadedValues) SetCs0(v float32) {
	o.Cs0 = &v
}

// GetPs0 returns the Ps0 field value if set, zero value otherwise.
func (o *UploadedValues) GetPs0() int32 {
	if o == nil || o.Ps0 == nil {
		var ret int32
		return ret
	}
	return *o.Ps0
}

// GetPs0Ok returns a tuple with the Ps0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetPs0Ok() (*int32, bool) {
	if o == nil || o.Ps0 == nil {
		return nil, false
	}
	return o.Ps0, true
}

// HasPs0 returns a boolean if a field has been set.
func (o *UploadedValues) HasPs0() bool {
	if o != nil && o.Ps0 != nil {
		return true
	}

	return false
}

// SetPs0 gets a reference to the given int32 and assigns it to the Ps0 field.
func (o *UploadedValues) SetPs0(v int32) {
	o.Ps0 = &v
}

// GetTs0 returns the Ts0 field value if set, zero value otherwise.
func (o *UploadedValues) GetTs0() int64 {
	if o == nil || o.Ts0 == nil {
		var ret int64
		return ret
	}
	return *o.Ts0
}

// GetTs0Ok returns a tuple with the Ts0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetTs0Ok() (*int64, bool) {
	if o == nil || o.Ts0 == nil {
		return nil, false
	}
	return o.Ts0, true
}

// HasTs0 returns a boolean if a field has been set.
func (o *UploadedValues) HasTs0() bool {
	if o != nil && o.Ts0 != nil {
		return true
	}

	return false
}

// SetTs0 gets a reference to the given int64 and assigns it to the Ts0 field.
func (o *UploadedValues) SetTs0(v int64) {
	o.Ts0 = &v
}

// GetGts returns the Gts field value if set, zero value otherwise.
func (o *UploadedValues) GetGts() int64 {
	if o == nil || o.Gts == nil {
		var ret int64
		return ret
	}
	return *o.Gts
}

// GetGtsOk returns a tuple with the Gts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedValues) GetGtsOk() (*int64, bool) {
	if o == nil || o.Gts == nil {
		return nil, false
	}
	return o.Gts, true
}

// HasGts returns a boolean if a field has been set.
func (o *UploadedValues) HasGts() bool {
	if o != nil && o.Gts != nil {
		return true
	}

	return false
}

// SetGts gets a reference to the given int64 and assigns it to the Gts field.
func (o *UploadedValues) SetGts(v int64) {
	o.Gts = &v
}

func (o UploadedValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tm != nil {
		toSerialize["tm"] = o.Tm
	}
	if o.Net != nil {
		toSerialize["net"] = o.Net
	}
	if o.Pwr != nil {
		toSerialize["pwr"] = o.Pwr
	}
	if o.P1 != nil {
		toSerialize["p1"] = o.P1
	}
	if o.P2 != nil {
		toSerialize["p2"] = o.P2
	}
	if o.N1 != nil {
		toSerialize["n1"] = o.N1
	}
	if o.N2 != nil {
		toSerialize["n2"] = o.N2
	}
	if o.Gas != nil {
		toSerialize["gas"] = o.Gas
	}
	if o.Cs0 != nil {
		toSerialize["cs0"] = o.Cs0
	}
	if o.Ps0 != nil {
		toSerialize["ps0"] = o.Ps0
	}
	if o.Ts0 != nil {
		toSerialize["ts0"] = o.Ts0
	}
	if o.Gts != nil {
		toSerialize["gts"] = o.Gts
	}
	return json.Marshal(toSerialize)
}

type NullableUploadedValues struct {
	value *UploadedValues
	isSet bool
}

func (v NullableUploadedValues) Get() *UploadedValues {
	return v.value
}

func (v *NullableUploadedValues) Set(val *UploadedValues) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadedValues) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadedValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadedValues(val *UploadedValues) *NullableUploadedValues {
	return &NullableUploadedValues{value: val, isSet: true}
}

func (v NullableUploadedValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadedValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


