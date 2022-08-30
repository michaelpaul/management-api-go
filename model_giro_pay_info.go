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

// GiroPayInfo struct for GiroPayInfo
type GiroPayInfo struct {
	// The email address of merchant support.
	SupportEmail string `json:"supportEmail"`
}

// NewGiroPayInfo instantiates a new GiroPayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiroPayInfo(supportEmail string) *GiroPayInfo {
	this := GiroPayInfo{}
	this.SupportEmail = supportEmail
	return &this
}

// NewGiroPayInfoWithDefaults instantiates a new GiroPayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiroPayInfoWithDefaults() *GiroPayInfo {
	this := GiroPayInfo{}
	return &this
}

// GetSupportEmail returns the SupportEmail field value
func (o *GiroPayInfo) GetSupportEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportEmail
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value
// and a boolean to check if the value has been set.
func (o *GiroPayInfo) GetSupportEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupportEmail, true
}

// SetSupportEmail sets field value
func (o *GiroPayInfo) SetSupportEmail(v string) {
	o.SupportEmail = v
}

func (o GiroPayInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supportEmail"] = o.SupportEmail
	}
	return json.Marshal(toSerialize)
}

type NullableGiroPayInfo struct {
	value *GiroPayInfo
	isSet bool
}

func (v NullableGiroPayInfo) Get() *GiroPayInfo {
	return v.value
}

func (v *NullableGiroPayInfo) Set(val *GiroPayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGiroPayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGiroPayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiroPayInfo(val *GiroPayInfo) *NullableGiroPayInfo {
	return &NullableGiroPayInfo{value: val, isSet: true}
}

func (v NullableGiroPayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiroPayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

