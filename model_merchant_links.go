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

// MerchantLinks struct for MerchantLinks
type MerchantLinks struct {
	ApiCredentials *LinksElement `json:"apiCredentials,omitempty"`
	Self LinksElement `json:"self"`
	Users *LinksElement `json:"users,omitempty"`
	Webhooks *LinksElement `json:"webhooks,omitempty"`
}

// NewMerchantLinks instantiates a new MerchantLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantLinks(self LinksElement) *MerchantLinks {
	this := MerchantLinks{}
	this.Self = self
	return &this
}

// NewMerchantLinksWithDefaults instantiates a new MerchantLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantLinksWithDefaults() *MerchantLinks {
	this := MerchantLinks{}
	return &this
}

// GetApiCredentials returns the ApiCredentials field value if set, zero value otherwise.
func (o *MerchantLinks) GetApiCredentials() LinksElement {
	if o == nil || o.ApiCredentials == nil {
		var ret LinksElement
		return ret
	}
	return *o.ApiCredentials
}

// GetApiCredentialsOk returns a tuple with the ApiCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantLinks) GetApiCredentialsOk() (*LinksElement, bool) {
	if o == nil || o.ApiCredentials == nil {
		return nil, false
	}
	return o.ApiCredentials, true
}

// HasApiCredentials returns a boolean if a field has been set.
func (o *MerchantLinks) HasApiCredentials() bool {
	if o != nil && o.ApiCredentials != nil {
		return true
	}

	return false
}

// SetApiCredentials gets a reference to the given LinksElement and assigns it to the ApiCredentials field.
func (o *MerchantLinks) SetApiCredentials(v LinksElement) {
	o.ApiCredentials = &v
}

// GetSelf returns the Self field value
func (o *MerchantLinks) GetSelf() LinksElement {
	if o == nil {
		var ret LinksElement
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *MerchantLinks) GetSelfOk() (*LinksElement, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *MerchantLinks) SetSelf(v LinksElement) {
	o.Self = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *MerchantLinks) GetUsers() LinksElement {
	if o == nil || o.Users == nil {
		var ret LinksElement
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantLinks) GetUsersOk() (*LinksElement, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *MerchantLinks) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given LinksElement and assigns it to the Users field.
func (o *MerchantLinks) SetUsers(v LinksElement) {
	o.Users = &v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *MerchantLinks) GetWebhooks() LinksElement {
	if o == nil || o.Webhooks == nil {
		var ret LinksElement
		return ret
	}
	return *o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantLinks) GetWebhooksOk() (*LinksElement, bool) {
	if o == nil || o.Webhooks == nil {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *MerchantLinks) HasWebhooks() bool {
	if o != nil && o.Webhooks != nil {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given LinksElement and assigns it to the Webhooks field.
func (o *MerchantLinks) SetWebhooks(v LinksElement) {
	o.Webhooks = &v
}

func (o MerchantLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiCredentials != nil {
		toSerialize["apiCredentials"] = o.ApiCredentials
	}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Webhooks != nil {
		toSerialize["webhooks"] = o.Webhooks
	}
	return json.Marshal(toSerialize)
}

type NullableMerchantLinks struct {
	value *MerchantLinks
	isSet bool
}

func (v NullableMerchantLinks) Get() *MerchantLinks {
	return v.value
}

func (v *NullableMerchantLinks) Set(val *MerchantLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantLinks(val *MerchantLinks) *NullableMerchantLinks {
	return &NullableMerchantLinks{value: val, isSet: true}
}

func (v NullableMerchantLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


