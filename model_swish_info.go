/*
Management API

Configure and manage your Adyen company and merchant accounts, stores, and payment terminals. ## Authentication Each request to the Management API must be signed with an API key. [Generate your API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key) in the Customer Area and then set this key to the `X-API-Key` header value.  To access the live endpoints, you need to generate a new API key in your live Customer Area. ## Versioning  Management API handles versioning as part of the endpoint URL. For example, to send a request to version 1 of the `/companies/{companyId}/webhooks` endpoint, use:  ```text https://management-test.adyen.com/v1/companies/{companyId}/webhooks ```

API version: 1
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SwishInfo struct for SwishInfo
type SwishInfo struct {
	// Swish number. Format: 10 digits without spaces. For example, **1231111111**.
	SwishNumber string `json:"swishNumber"`
}

// NewSwishInfo instantiates a new SwishInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwishInfo(swishNumber string) *SwishInfo {
	this := SwishInfo{}
	this.SwishNumber = swishNumber
	return &this
}

// NewSwishInfoWithDefaults instantiates a new SwishInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwishInfoWithDefaults() *SwishInfo {
	this := SwishInfo{}
	return &this
}

// GetSwishNumber returns the SwishNumber field value
func (o *SwishInfo) GetSwishNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SwishNumber
}

// GetSwishNumberOk returns a tuple with the SwishNumber field value
// and a boolean to check if the value has been set.
func (o *SwishInfo) GetSwishNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SwishNumber, true
}

// SetSwishNumber sets field value
func (o *SwishInfo) SetSwishNumber(v string) {
	o.SwishNumber = v
}

func (o SwishInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["swishNumber"] = o.SwishNumber
	}
	return json.Marshal(toSerialize)
}

type NullableSwishInfo struct {
	value *SwishInfo
	isSet bool
}

func (v NullableSwishInfo) Get() *SwishInfo {
	return v.value
}

func (v *NullableSwishInfo) Set(val *SwishInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSwishInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSwishInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwishInfo(val *SwishInfo) *NullableSwishInfo {
	return &NullableSwishInfo{value: val, isSet: true}
}

func (v NullableSwishInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwishInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


