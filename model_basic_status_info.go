/*
Youless energy monitor

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package youless

import (
	"encoding/json"
)

// BasicStatusInfo struct for BasicStatusInfo
type BasicStatusInfo struct {
	// counter in kWh
	Cnt *string `json:"cnt,omitempty"`
	// Power consumption in Watt
	Pwr *int32 `json:"pwr,omitempty"`
	// moving average level (intensity of reflected light on analog meters)
	Lvl *int32 `json:"lvl,omitempty"`
	// deviation of reflection
	Dev *string `json:"dev,omitempty"`
	Det *string `json:"det,omitempty"`
	// connection status
	Con *string `json:"con,omitempty"`
	// Time until next status update with online monitoring
	Sts *string `json:"sts,omitempty"`
	// kWh counter of S0 input
	Cs0 *string `json:"cs0,omitempty"`
	// Computed power
	Ps0 *int32 `json:"ps0,omitempty"`
	// raw 10-bit light reflection level (without averaging)
	Raw *int32 `json:"raw,omitempty"`
}

// NewBasicStatusInfo instantiates a new BasicStatusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicStatusInfo() *BasicStatusInfo {
	this := BasicStatusInfo{}
	return &this
}

// NewBasicStatusInfoWithDefaults instantiates a new BasicStatusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicStatusInfoWithDefaults() *BasicStatusInfo {
	this := BasicStatusInfo{}
	return &this
}

// GetCnt returns the Cnt field value if set, zero value otherwise.
func (o *BasicStatusInfo) GetCnt() string {
	if o == nil || o.Cnt == nil {
		var ret string
		return ret
	}
	return *o.Cnt
}

// GetCntOk returns a tuple with the Cnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicStatusInfo) GetCntOk() (*string, bool) {
	if o == nil || o.Cnt == nil {
		return nil, false
	}
	return o.Cnt, true
}

// HasCnt returns a boolean if a field has been set.
func (o *BasicStatusInfo) HasCnt() bool {
	if o != nil && o.Cnt != nil {
		return true
	}

	return false
}

// SetCnt gets a reference to the given string and assigns it to the Cnt field.
func (o *BasicStatusInfo) SetCnt(v string) {
	o.Cnt = &v
}

// GetPwr returns the Pwr field value if set, zero value otherwise.
func (o *BasicStatusInfo) GetPwr() int32 {
	if o == nil || o.Pwr == nil {
		var ret int32
		return ret
	}
	return *o.Pwr
}

// GetPwrOk returns a tuple with the Pwr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicStatusInfo) GetPwrOk() (*int32, bool) {
	if o == nil || o.Pwr == nil {
		return nil, false
	}
	return o.Pwr, true
}

// HasPwr returns a boolean if a field has been set.
func (o *BasicStatusInfo) HasPwr() bool {
	if o != nil && o.Pwr != nil {
		return true
	}

	return false
}

// SetPwr gets a reference to the given int32 and assigns it to the Pwr field.
func (o *BasicStatusInfo) SetPwr(v int32) {
	o.Pwr = &v
}

// GetLvl returns the Lvl field value if set, zero value otherwise.
func (o *BasicStatusInfo) GetLvl() int32 {
	if o == nil || o.Lvl == nil {
		var ret int32
		return ret
	}
	return *o.Lvl
}

// GetLvlOk returns a tuple with the Lvl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicStatusInfo) GetLvlOk() (*int32, bool) {
	if o == nil || o.Lvl == nil {
		return nil, false
	}
	return o.Lvl, true
}

// HasLvl returns a boolean if a field has been set.
func (o *BasicStatusInfo) HasLvl() bool {
	if o != nil && o.Lvl != nil {
		return true
	}

	return false
}

// SetLvl gets a reference to the given int32 and assigns it to the Lvl field.
func (o *BasicStatusInfo) SetLvl(v int32) {
	o.Lvl = &v
}

// GetDev returns the Dev field value if set, zero value otherwise.
func (o *BasicStatusInfo) GetDev() string {
	if o == nil || o.Dev == nil {
		var ret string
		return ret
	}
	return *o.Dev
}

// GetDevOk returns a tuple with the Dev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicStatusInfo) GetDevOk() (*string, bool) {
	if o == nil || o.Dev == nil {
		return nil, false
	}
	return o.Dev, true
}

// HasDev returns a boolean if a field has been set.
func (o *BasicStatusInfo) HasDev() bool {
	if o != nil && o.Dev != nil {
		return true
	}

	return false
}

// SetDev gets a reference to the given string and assigns it to the Dev field.
func (o *BasicStatusInfo) SetDev(v string) {
	o.Dev = &v
}

// GetDet returns the Det field value if set, zero value otherwise.
func (o *BasicStatusInfo) GetDet() string {
	if o == nil || o.Det == nil {
		var ret string
		return ret
	}
	return *o.Det
}

// GetDetOk returns a tuple with the Det field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicStatusInfo) GetDetOk() (*string, bool) {
	if o == nil || o.Det == nil {
		return nil, false
	}
	return o.Det, true
}

// HasDet returns a boolean if a field has been set.
func (o *BasicStatusInfo) HasDet() bool {
	if o != nil && o.Det != nil {
		return true
	}

	return false
}

// SetDet gets a reference to the given string and assigns it to the Det field.
func (o *BasicStatusInfo) SetDet(v string) {
	o.Det = &v
}

// GetCon returns the Con field value if set, zero value otherwise.
func (o *BasicStatusInfo) GetCon() string {
	if o == nil || o.Con == nil {
		var ret string
		return ret
	}
	return *o.Con
}

// GetConOk returns a tuple with the Con field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicStatusInfo) GetConOk() (*string, bool) {
	if o == nil || o.Con == nil {
		return nil, false
	}
	return o.Con, true
}

// HasCon returns a boolean if a field has been set.
func (o *BasicStatusInfo) HasCon() bool {
	if o != nil && o.Con != nil {
		return true
	}

	return false
}

// SetCon gets a reference to the given string and assigns it to the Con field.
func (o *BasicStatusInfo) SetCon(v string) {
	o.Con = &v
}

// GetSts returns the Sts field value if set, zero value otherwise.
func (o *BasicStatusInfo) GetSts() string {
	if o == nil || o.Sts == nil {
		var ret string
		return ret
	}
	return *o.Sts
}

// GetStsOk returns a tuple with the Sts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicStatusInfo) GetStsOk() (*string, bool) {
	if o == nil || o.Sts == nil {
		return nil, false
	}
	return o.Sts, true
}

// HasSts returns a boolean if a field has been set.
func (o *BasicStatusInfo) HasSts() bool {
	if o != nil && o.Sts != nil {
		return true
	}

	return false
}

// SetSts gets a reference to the given string and assigns it to the Sts field.
func (o *BasicStatusInfo) SetSts(v string) {
	o.Sts = &v
}

// GetCs0 returns the Cs0 field value if set, zero value otherwise.
func (o *BasicStatusInfo) GetCs0() string {
	if o == nil || o.Cs0 == nil {
		var ret string
		return ret
	}
	return *o.Cs0
}

// GetCs0Ok returns a tuple with the Cs0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicStatusInfo) GetCs0Ok() (*string, bool) {
	if o == nil || o.Cs0 == nil {
		return nil, false
	}
	return o.Cs0, true
}

// HasCs0 returns a boolean if a field has been set.
func (o *BasicStatusInfo) HasCs0() bool {
	if o != nil && o.Cs0 != nil {
		return true
	}

	return false
}

// SetCs0 gets a reference to the given string and assigns it to the Cs0 field.
func (o *BasicStatusInfo) SetCs0(v string) {
	o.Cs0 = &v
}

// GetPs0 returns the Ps0 field value if set, zero value otherwise.
func (o *BasicStatusInfo) GetPs0() int32 {
	if o == nil || o.Ps0 == nil {
		var ret int32
		return ret
	}
	return *o.Ps0
}

// GetPs0Ok returns a tuple with the Ps0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicStatusInfo) GetPs0Ok() (*int32, bool) {
	if o == nil || o.Ps0 == nil {
		return nil, false
	}
	return o.Ps0, true
}

// HasPs0 returns a boolean if a field has been set.
func (o *BasicStatusInfo) HasPs0() bool {
	if o != nil && o.Ps0 != nil {
		return true
	}

	return false
}

// SetPs0 gets a reference to the given int32 and assigns it to the Ps0 field.
func (o *BasicStatusInfo) SetPs0(v int32) {
	o.Ps0 = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *BasicStatusInfo) GetRaw() int32 {
	if o == nil || o.Raw == nil {
		var ret int32
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicStatusInfo) GetRawOk() (*int32, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *BasicStatusInfo) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given int32 and assigns it to the Raw field.
func (o *BasicStatusInfo) SetRaw(v int32) {
	o.Raw = &v
}

func (o BasicStatusInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cnt != nil {
		toSerialize["cnt"] = o.Cnt
	}
	if o.Pwr != nil {
		toSerialize["pwr"] = o.Pwr
	}
	if o.Lvl != nil {
		toSerialize["lvl"] = o.Lvl
	}
	if o.Dev != nil {
		toSerialize["dev"] = o.Dev
	}
	if o.Det != nil {
		toSerialize["det"] = o.Det
	}
	if o.Con != nil {
		toSerialize["con"] = o.Con
	}
	if o.Sts != nil {
		toSerialize["sts"] = o.Sts
	}
	if o.Cs0 != nil {
		toSerialize["cs0"] = o.Cs0
	}
	if o.Ps0 != nil {
		toSerialize["ps0"] = o.Ps0
	}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	return json.Marshal(toSerialize)
}

type NullableBasicStatusInfo struct {
	value *BasicStatusInfo
	isSet bool
}

func (v NullableBasicStatusInfo) Get() *BasicStatusInfo {
	return v.value
}

func (v *NullableBasicStatusInfo) Set(val *BasicStatusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicStatusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicStatusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicStatusInfo(val *BasicStatusInfo) *NullableBasicStatusInfo {
	return &NullableBasicStatusInfo{value: val, isSet: true}
}

func (v NullableBasicStatusInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicStatusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

