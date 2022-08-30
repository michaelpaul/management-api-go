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

// ListMerchantApiCredentialsResponse struct for ListMerchantApiCredentialsResponse
type ListMerchantApiCredentialsResponse struct {
	Links *PaginationLinks `json:"_links,omitempty"`
	// The list of API credentials.
	Data []ApiCredential `json:"data,omitempty"`
	// Total number of items.
	ItemsTotal int32 `json:"itemsTotal"`
	// Total number of pages.
	PagesTotal int32 `json:"pagesTotal"`
}

// NewListMerchantApiCredentialsResponse instantiates a new ListMerchantApiCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMerchantApiCredentialsResponse(itemsTotal int32, pagesTotal int32) *ListMerchantApiCredentialsResponse {
	this := ListMerchantApiCredentialsResponse{}
	this.ItemsTotal = itemsTotal
	this.PagesTotal = pagesTotal
	return &this
}

// NewListMerchantApiCredentialsResponseWithDefaults instantiates a new ListMerchantApiCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMerchantApiCredentialsResponseWithDefaults() *ListMerchantApiCredentialsResponse {
	this := ListMerchantApiCredentialsResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListMerchantApiCredentialsResponse) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMerchantApiCredentialsResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListMerchantApiCredentialsResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *ListMerchantApiCredentialsResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListMerchantApiCredentialsResponse) GetData() []ApiCredential {
	if o == nil || o.Data == nil {
		var ret []ApiCredential
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMerchantApiCredentialsResponse) GetDataOk() ([]ApiCredential, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListMerchantApiCredentialsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ApiCredential and assigns it to the Data field.
func (o *ListMerchantApiCredentialsResponse) SetData(v []ApiCredential) {
	o.Data = v
}

// GetItemsTotal returns the ItemsTotal field value
func (o *ListMerchantApiCredentialsResponse) GetItemsTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsTotal
}

// GetItemsTotalOk returns a tuple with the ItemsTotal field value
// and a boolean to check if the value has been set.
func (o *ListMerchantApiCredentialsResponse) GetItemsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsTotal, true
}

// SetItemsTotal sets field value
func (o *ListMerchantApiCredentialsResponse) SetItemsTotal(v int32) {
	o.ItemsTotal = v
}

// GetPagesTotal returns the PagesTotal field value
func (o *ListMerchantApiCredentialsResponse) GetPagesTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PagesTotal
}

// GetPagesTotalOk returns a tuple with the PagesTotal field value
// and a boolean to check if the value has been set.
func (o *ListMerchantApiCredentialsResponse) GetPagesTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PagesTotal, true
}

// SetPagesTotal sets field value
func (o *ListMerchantApiCredentialsResponse) SetPagesTotal(v int32) {
	o.PagesTotal = v
}

func (o ListMerchantApiCredentialsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["itemsTotal"] = o.ItemsTotal
	}
	if true {
		toSerialize["pagesTotal"] = o.PagesTotal
	}
	return json.Marshal(toSerialize)
}

type NullableListMerchantApiCredentialsResponse struct {
	value *ListMerchantApiCredentialsResponse
	isSet bool
}

func (v NullableListMerchantApiCredentialsResponse) Get() *ListMerchantApiCredentialsResponse {
	return v.value
}

func (v *NullableListMerchantApiCredentialsResponse) Set(val *ListMerchantApiCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMerchantApiCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMerchantApiCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMerchantApiCredentialsResponse(val *ListMerchantApiCredentialsResponse) *NullableListMerchantApiCredentialsResponse {
	return &NullableListMerchantApiCredentialsResponse{value: val, isSet: true}
}

func (v NullableListMerchantApiCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMerchantApiCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


