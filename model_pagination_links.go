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

// PaginationLinks struct for PaginationLinks
type PaginationLinks struct {
	First LinksElement `json:"first"`
	Last LinksElement `json:"last"`
	Next *LinksElement `json:"next,omitempty"`
	Prev *LinksElement `json:"prev,omitempty"`
	Self LinksElement `json:"self"`
}

// NewPaginationLinks instantiates a new PaginationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationLinks(first LinksElement, last LinksElement, self LinksElement) *PaginationLinks {
	this := PaginationLinks{}
	this.First = first
	this.Last = last
	this.Self = self
	return &this
}

// NewPaginationLinksWithDefaults instantiates a new PaginationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationLinksWithDefaults() *PaginationLinks {
	this := PaginationLinks{}
	return &this
}

// GetFirst returns the First field value
func (o *PaginationLinks) GetFirst() LinksElement {
	if o == nil {
		var ret LinksElement
		return ret
	}

	return o.First
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
func (o *PaginationLinks) GetFirstOk() (*LinksElement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.First, true
}

// SetFirst sets field value
func (o *PaginationLinks) SetFirst(v LinksElement) {
	o.First = v
}

// GetLast returns the Last field value
func (o *PaginationLinks) GetLast() LinksElement {
	if o == nil {
		var ret LinksElement
		return ret
	}

	return o.Last
}

// GetLastOk returns a tuple with the Last field value
// and a boolean to check if the value has been set.
func (o *PaginationLinks) GetLastOk() (*LinksElement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last, true
}

// SetLast sets field value
func (o *PaginationLinks) SetLast(v LinksElement) {
	o.Last = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginationLinks) GetNext() LinksElement {
	if o == nil || o.Next == nil {
		var ret LinksElement
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationLinks) GetNextOk() (*LinksElement, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginationLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given LinksElement and assigns it to the Next field.
func (o *PaginationLinks) SetNext(v LinksElement) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *PaginationLinks) GetPrev() LinksElement {
	if o == nil || o.Prev == nil {
		var ret LinksElement
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationLinks) GetPrevOk() (*LinksElement, bool) {
	if o == nil || o.Prev == nil {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *PaginationLinks) HasPrev() bool {
	if o != nil && o.Prev != nil {
		return true
	}

	return false
}

// SetPrev gets a reference to the given LinksElement and assigns it to the Prev field.
func (o *PaginationLinks) SetPrev(v LinksElement) {
	o.Prev = &v
}

// GetSelf returns the Self field value
func (o *PaginationLinks) GetSelf() LinksElement {
	if o == nil {
		var ret LinksElement
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *PaginationLinks) GetSelfOk() (*LinksElement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *PaginationLinks) SetSelf(v LinksElement) {
	o.Self = v
}

func (o PaginationLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["first"] = o.First
	}
	if true {
		toSerialize["last"] = o.Last
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Prev != nil {
		toSerialize["prev"] = o.Prev
	}
	if true {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullablePaginationLinks struct {
	value *PaginationLinks
	isSet bool
}

func (v NullablePaginationLinks) Get() *PaginationLinks {
	return v.value
}

func (v *NullablePaginationLinks) Set(val *PaginationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationLinks(val *PaginationLinks) *NullablePaginationLinks {
	return &NullablePaginationLinks{value: val, isSet: true}
}

func (v NullablePaginationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


