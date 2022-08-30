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

// ListMerchantUsersResponse struct for ListMerchantUsersResponse
type ListMerchantUsersResponse struct {
	Links *PaginationLinks `json:"_links,omitempty"`
	// The list of users.
	Data []User `json:"data,omitempty"`
	// Total number of items.
	ItemsTotal int32 `json:"itemsTotal"`
	// Total number of pages.
	PagesTotal int32 `json:"pagesTotal"`
}

// NewListMerchantUsersResponse instantiates a new ListMerchantUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMerchantUsersResponse(itemsTotal int32, pagesTotal int32) *ListMerchantUsersResponse {
	this := ListMerchantUsersResponse{}
	this.ItemsTotal = itemsTotal
	this.PagesTotal = pagesTotal
	return &this
}

// NewListMerchantUsersResponseWithDefaults instantiates a new ListMerchantUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMerchantUsersResponseWithDefaults() *ListMerchantUsersResponse {
	this := ListMerchantUsersResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListMerchantUsersResponse) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMerchantUsersResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListMerchantUsersResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *ListMerchantUsersResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListMerchantUsersResponse) GetData() []User {
	if o == nil || o.Data == nil {
		var ret []User
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMerchantUsersResponse) GetDataOk() ([]User, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListMerchantUsersResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []User and assigns it to the Data field.
func (o *ListMerchantUsersResponse) SetData(v []User) {
	o.Data = v
}

// GetItemsTotal returns the ItemsTotal field value
func (o *ListMerchantUsersResponse) GetItemsTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsTotal
}

// GetItemsTotalOk returns a tuple with the ItemsTotal field value
// and a boolean to check if the value has been set.
func (o *ListMerchantUsersResponse) GetItemsTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsTotal, true
}

// SetItemsTotal sets field value
func (o *ListMerchantUsersResponse) SetItemsTotal(v int32) {
	o.ItemsTotal = v
}

// GetPagesTotal returns the PagesTotal field value
func (o *ListMerchantUsersResponse) GetPagesTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PagesTotal
}

// GetPagesTotalOk returns a tuple with the PagesTotal field value
// and a boolean to check if the value has been set.
func (o *ListMerchantUsersResponse) GetPagesTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PagesTotal, true
}

// SetPagesTotal sets field value
func (o *ListMerchantUsersResponse) SetPagesTotal(v int32) {
	o.PagesTotal = v
}

func (o ListMerchantUsersResponse) MarshalJSON() ([]byte, error) {
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

type NullableListMerchantUsersResponse struct {
	value *ListMerchantUsersResponse
	isSet bool
}

func (v NullableListMerchantUsersResponse) Get() *ListMerchantUsersResponse {
	return v.value
}

func (v *NullableListMerchantUsersResponse) Set(val *ListMerchantUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMerchantUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMerchantUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMerchantUsersResponse(val *ListMerchantUsersResponse) *NullableListMerchantUsersResponse {
	return &NullableListMerchantUsersResponse{value: val, isSet: true}
}

func (v NullableListMerchantUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMerchantUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


