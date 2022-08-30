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

// TerminalOrder struct for TerminalOrder
type TerminalOrder struct {
	BillingEntity *BillingEntity `json:"billingEntity,omitempty"`
	// The merchant-defined purchase order number. This will be printed on the packing list.
	CustomerOrderReference *string `json:"customerOrderReference,omitempty"`
	// The unique identifier of the order.
	Id *string `json:"id,omitempty"`
	// The products included in the order.
	Items []OrderItem `json:"items,omitempty"`
	// The date and time that the order was placed, in UTC ISO 8601 format. For example, \"2011-12-03T10:15:30Z\".
	OrderDate *string `json:"orderDate,omitempty"`
	ShippingLocation *ShippingLocation `json:"shippingLocation,omitempty"`
	// The processing status of the order.
	Status *string `json:"status,omitempty"`
	// The URL, provided by the carrier company, where the shipment can be tracked.
	TrackingUrl *string `json:"trackingUrl,omitempty"`
}

// NewTerminalOrder instantiates a new TerminalOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalOrder() *TerminalOrder {
	this := TerminalOrder{}
	return &this
}

// NewTerminalOrderWithDefaults instantiates a new TerminalOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalOrderWithDefaults() *TerminalOrder {
	this := TerminalOrder{}
	return &this
}

// GetBillingEntity returns the BillingEntity field value if set, zero value otherwise.
func (o *TerminalOrder) GetBillingEntity() BillingEntity {
	if o == nil || o.BillingEntity == nil {
		var ret BillingEntity
		return ret
	}
	return *o.BillingEntity
}

// GetBillingEntityOk returns a tuple with the BillingEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrder) GetBillingEntityOk() (*BillingEntity, bool) {
	if o == nil || o.BillingEntity == nil {
		return nil, false
	}
	return o.BillingEntity, true
}

// HasBillingEntity returns a boolean if a field has been set.
func (o *TerminalOrder) HasBillingEntity() bool {
	if o != nil && o.BillingEntity != nil {
		return true
	}

	return false
}

// SetBillingEntity gets a reference to the given BillingEntity and assigns it to the BillingEntity field.
func (o *TerminalOrder) SetBillingEntity(v BillingEntity) {
	o.BillingEntity = &v
}

// GetCustomerOrderReference returns the CustomerOrderReference field value if set, zero value otherwise.
func (o *TerminalOrder) GetCustomerOrderReference() string {
	if o == nil || o.CustomerOrderReference == nil {
		var ret string
		return ret
	}
	return *o.CustomerOrderReference
}

// GetCustomerOrderReferenceOk returns a tuple with the CustomerOrderReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrder) GetCustomerOrderReferenceOk() (*string, bool) {
	if o == nil || o.CustomerOrderReference == nil {
		return nil, false
	}
	return o.CustomerOrderReference, true
}

// HasCustomerOrderReference returns a boolean if a field has been set.
func (o *TerminalOrder) HasCustomerOrderReference() bool {
	if o != nil && o.CustomerOrderReference != nil {
		return true
	}

	return false
}

// SetCustomerOrderReference gets a reference to the given string and assigns it to the CustomerOrderReference field.
func (o *TerminalOrder) SetCustomerOrderReference(v string) {
	o.CustomerOrderReference = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TerminalOrder) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrder) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TerminalOrder) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TerminalOrder) SetId(v string) {
	o.Id = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *TerminalOrder) GetItems() []OrderItem {
	if o == nil || o.Items == nil {
		var ret []OrderItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrder) GetItemsOk() ([]OrderItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TerminalOrder) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrderItem and assigns it to the Items field.
func (o *TerminalOrder) SetItems(v []OrderItem) {
	o.Items = v
}

// GetOrderDate returns the OrderDate field value if set, zero value otherwise.
func (o *TerminalOrder) GetOrderDate() string {
	if o == nil || o.OrderDate == nil {
		var ret string
		return ret
	}
	return *o.OrderDate
}

// GetOrderDateOk returns a tuple with the OrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrder) GetOrderDateOk() (*string, bool) {
	if o == nil || o.OrderDate == nil {
		return nil, false
	}
	return o.OrderDate, true
}

// HasOrderDate returns a boolean if a field has been set.
func (o *TerminalOrder) HasOrderDate() bool {
	if o != nil && o.OrderDate != nil {
		return true
	}

	return false
}

// SetOrderDate gets a reference to the given string and assigns it to the OrderDate field.
func (o *TerminalOrder) SetOrderDate(v string) {
	o.OrderDate = &v
}

// GetShippingLocation returns the ShippingLocation field value if set, zero value otherwise.
func (o *TerminalOrder) GetShippingLocation() ShippingLocation {
	if o == nil || o.ShippingLocation == nil {
		var ret ShippingLocation
		return ret
	}
	return *o.ShippingLocation
}

// GetShippingLocationOk returns a tuple with the ShippingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrder) GetShippingLocationOk() (*ShippingLocation, bool) {
	if o == nil || o.ShippingLocation == nil {
		return nil, false
	}
	return o.ShippingLocation, true
}

// HasShippingLocation returns a boolean if a field has been set.
func (o *TerminalOrder) HasShippingLocation() bool {
	if o != nil && o.ShippingLocation != nil {
		return true
	}

	return false
}

// SetShippingLocation gets a reference to the given ShippingLocation and assigns it to the ShippingLocation field.
func (o *TerminalOrder) SetShippingLocation(v ShippingLocation) {
	o.ShippingLocation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TerminalOrder) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrder) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TerminalOrder) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TerminalOrder) SetStatus(v string) {
	o.Status = &v
}

// GetTrackingUrl returns the TrackingUrl field value if set, zero value otherwise.
func (o *TerminalOrder) GetTrackingUrl() string {
	if o == nil || o.TrackingUrl == nil {
		var ret string
		return ret
	}
	return *o.TrackingUrl
}

// GetTrackingUrlOk returns a tuple with the TrackingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalOrder) GetTrackingUrlOk() (*string, bool) {
	if o == nil || o.TrackingUrl == nil {
		return nil, false
	}
	return o.TrackingUrl, true
}

// HasTrackingUrl returns a boolean if a field has been set.
func (o *TerminalOrder) HasTrackingUrl() bool {
	if o != nil && o.TrackingUrl != nil {
		return true
	}

	return false
}

// SetTrackingUrl gets a reference to the given string and assigns it to the TrackingUrl field.
func (o *TerminalOrder) SetTrackingUrl(v string) {
	o.TrackingUrl = &v
}

func (o TerminalOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingEntity != nil {
		toSerialize["billingEntity"] = o.BillingEntity
	}
	if o.CustomerOrderReference != nil {
		toSerialize["customerOrderReference"] = o.CustomerOrderReference
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.OrderDate != nil {
		toSerialize["orderDate"] = o.OrderDate
	}
	if o.ShippingLocation != nil {
		toSerialize["shippingLocation"] = o.ShippingLocation
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TrackingUrl != nil {
		toSerialize["trackingUrl"] = o.TrackingUrl
	}
	return json.Marshal(toSerialize)
}

type NullableTerminalOrder struct {
	value *TerminalOrder
	isSet bool
}

func (v NullableTerminalOrder) Get() *TerminalOrder {
	return v.value
}

func (v *NullableTerminalOrder) Set(val *TerminalOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalOrder(val *TerminalOrder) *NullableTerminalOrder {
	return &NullableTerminalOrder{value: val, isSet: true}
}

func (v NullableTerminalOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

