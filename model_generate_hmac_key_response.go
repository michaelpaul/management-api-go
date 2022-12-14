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

// GenerateHmacKeyResponse struct for GenerateHmacKeyResponse
type GenerateHmacKeyResponse struct {
	// The HMAC key generated for this webhook.
	HmacKey string `json:"hmacKey"`
}

// NewGenerateHmacKeyResponse instantiates a new GenerateHmacKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateHmacKeyResponse(hmacKey string) *GenerateHmacKeyResponse {
	this := GenerateHmacKeyResponse{}
	this.HmacKey = hmacKey
	return &this
}

// NewGenerateHmacKeyResponseWithDefaults instantiates a new GenerateHmacKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateHmacKeyResponseWithDefaults() *GenerateHmacKeyResponse {
	this := GenerateHmacKeyResponse{}
	return &this
}

// GetHmacKey returns the HmacKey field value
func (o *GenerateHmacKeyResponse) GetHmacKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HmacKey
}

// GetHmacKeyOk returns a tuple with the HmacKey field value
// and a boolean to check if the value has been set.
func (o *GenerateHmacKeyResponse) GetHmacKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HmacKey, true
}

// SetHmacKey sets field value
func (o *GenerateHmacKeyResponse) SetHmacKey(v string) {
	o.HmacKey = v
}

func (o GenerateHmacKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hmacKey"] = o.HmacKey
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateHmacKeyResponse struct {
	value *GenerateHmacKeyResponse
	isSet bool
}

func (v NullableGenerateHmacKeyResponse) Get() *GenerateHmacKeyResponse {
	return v.value
}

func (v *NullableGenerateHmacKeyResponse) Set(val *GenerateHmacKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateHmacKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateHmacKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateHmacKeyResponse(val *GenerateHmacKeyResponse) *NullableGenerateHmacKeyResponse {
	return &NullableGenerateHmacKeyResponse{value: val, isSet: true}
}

func (v NullableGenerateHmacKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateHmacKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


