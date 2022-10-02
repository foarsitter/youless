/*
Youless energy monitor GO client

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package youless

import (
	"encoding/json"
	"fmt"
)

// OutputFormat the model 'OutputFormat'
type OutputFormat string

// List of OutputFormat
const (
	J OutputFormat = "j"
)

// All allowed values of OutputFormat enum
var AllowedOutputFormatEnumValues = []OutputFormat{
	"j",
}

func (v *OutputFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutputFormat(value)
	for _, existing := range AllowedOutputFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutputFormat", value)
}

// NewOutputFormatFromValue returns a pointer to a valid OutputFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOutputFormatFromValue(v string) (*OutputFormat, error) {
	ev := OutputFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OutputFormat: valid values are %v", v, AllowedOutputFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OutputFormat) IsValid() bool {
	for _, existing := range AllowedOutputFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutputFormat value
func (v OutputFormat) Ptr() *OutputFormat {
	return &v
}

type NullableOutputFormat struct {
	value *OutputFormat
	isSet bool
}

func (v NullableOutputFormat) Get() *OutputFormat {
	return v.value
}

func (v *NullableOutputFormat) Set(val *OutputFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputFormat(val *OutputFormat) *NullableOutputFormat {
	return &NullableOutputFormat{value: val, isSet: true}
}

func (v NullableOutputFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
