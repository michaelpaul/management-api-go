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

// Signature struct for Signature
type Signature struct {
	// If `skipSignature` is false, indicates whether the shopper should provide a signature on the display (**true**) or on the merchant receipt (**false**).
	AskSignatureOnScreen *bool `json:"askSignatureOnScreen,omitempty"`
	// Name that identifies the terminal.
	DeviceName *string `json:"deviceName,omitempty"`
	// Skip asking for a signature. This is possible because all global card schemes (American Express, Diners, Discover, JCB, MasterCard, VISA, and UnionPay) regard a signature as optional.
	SkipSignature *bool `json:"skipSignature,omitempty"`
}

// NewSignature instantiates a new Signature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignature() *Signature {
	this := Signature{}
	return &this
}

// NewSignatureWithDefaults instantiates a new Signature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureWithDefaults() *Signature {
	this := Signature{}
	return &this
}

// GetAskSignatureOnScreen returns the AskSignatureOnScreen field value if set, zero value otherwise.
func (o *Signature) GetAskSignatureOnScreen() bool {
	if o == nil || o.AskSignatureOnScreen == nil {
		var ret bool
		return ret
	}
	return *o.AskSignatureOnScreen
}

// GetAskSignatureOnScreenOk returns a tuple with the AskSignatureOnScreen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Signature) GetAskSignatureOnScreenOk() (*bool, bool) {
	if o == nil || o.AskSignatureOnScreen == nil {
		return nil, false
	}
	return o.AskSignatureOnScreen, true
}

// HasAskSignatureOnScreen returns a boolean if a field has been set.
func (o *Signature) HasAskSignatureOnScreen() bool {
	if o != nil && o.AskSignatureOnScreen != nil {
		return true
	}

	return false
}

// SetAskSignatureOnScreen gets a reference to the given bool and assigns it to the AskSignatureOnScreen field.
func (o *Signature) SetAskSignatureOnScreen(v bool) {
	o.AskSignatureOnScreen = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *Signature) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Signature) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *Signature) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *Signature) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetSkipSignature returns the SkipSignature field value if set, zero value otherwise.
func (o *Signature) GetSkipSignature() bool {
	if o == nil || o.SkipSignature == nil {
		var ret bool
		return ret
	}
	return *o.SkipSignature
}

// GetSkipSignatureOk returns a tuple with the SkipSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Signature) GetSkipSignatureOk() (*bool, bool) {
	if o == nil || o.SkipSignature == nil {
		return nil, false
	}
	return o.SkipSignature, true
}

// HasSkipSignature returns a boolean if a field has been set.
func (o *Signature) HasSkipSignature() bool {
	if o != nil && o.SkipSignature != nil {
		return true
	}

	return false
}

// SetSkipSignature gets a reference to the given bool and assigns it to the SkipSignature field.
func (o *Signature) SetSkipSignature(v bool) {
	o.SkipSignature = &v
}

func (o Signature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AskSignatureOnScreen != nil {
		toSerialize["askSignatureOnScreen"] = o.AskSignatureOnScreen
	}
	if o.DeviceName != nil {
		toSerialize["deviceName"] = o.DeviceName
	}
	if o.SkipSignature != nil {
		toSerialize["skipSignature"] = o.SkipSignature
	}
	return json.Marshal(toSerialize)
}

type NullableSignature struct {
	value *Signature
	isSet bool
}

func (v NullableSignature) Get() *Signature {
	return v.value
}

func (v *NullableSignature) Set(val *Signature) {
	v.value = val
	v.isSet = true
}

func (v NullableSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignature(val *Signature) *NullableSignature {
	return &NullableSignature{value: val, isSet: true}
}

func (v NullableSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


