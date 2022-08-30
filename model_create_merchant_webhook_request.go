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

// CreateMerchantWebhookRequest struct for CreateMerchantWebhookRequest
type CreateMerchantWebhookRequest struct {
	// Indicates if expired SSL certificates are accepted. Default value: **false**.
	AcceptsExpiredCertificate *bool `json:"acceptsExpiredCertificate,omitempty"`
	// Indicates if self-signed SSL certificates are accepted. Default value: **false**.
	AcceptsSelfSignedCertificate *bool `json:"acceptsSelfSignedCertificate,omitempty"`
	// Indicates if untrusted SSL certificates are accepted. Default value: **false**.
	AcceptsUntrustedRootCertificate *bool `json:"acceptsUntrustedRootCertificate,omitempty"`
	// Indicates if the webhook configuration is active. The field must be **true** for us to send webhooks about events related an account.
	Active bool `json:"active"`
	AdditionalSettings *AdditionalSettings `json:"additionalSettings,omitempty"`
	// Format or protocol for receiving webhooks. Possible values: * **soap** * **http** * **json** 
	CommunicationFormat string `json:"communicationFormat"`
	// Your description for this webhook configuration.
	Description *string `json:"description,omitempty"`
	// Network type for Terminal API notification webhooks. Possible values: * **public** * **local**  Default Value: **public**.
	NetworkType *string `json:"networkType,omitempty"`
	// Password to access the webhook URL.
	Password *string `json:"password,omitempty"`
	// Indicates if the SOAP action header needs to be populated. Default value: **false**.  Only applies if `communicationFormat`: **soap**.
	PopulateSoapActionHeader *bool `json:"populateSoapActionHeader,omitempty"`
	// SSL version to access the public webhook URL specified in the `url` field. Possible values: * **TLSv1.2** * **HTTP** - Only allowed on Test environment.  If not specified, the webhook will use `sslVersion`: **TLSv1.2**.
	SslVersion *string `json:"sslVersion,omitempty"`
	// The type of webhook that is being created. Possible values are:  - **standard** - **account-settings-notification** - **banktransfer-notification** - **boletobancario-notification** - **directdebit-notification** - **pending-notification** - **ideal-notification** - **ideal-pending-notification** - **report-notification** - **rreq-notification**  Find out more about [standard notification webhooks](https://docs.adyen.com/development-resources/webhooks/understand-notifications#event-codes) and [other types of notifications](https://docs.adyen.com/development-resources/webhooks/understand-notifications#other-notifications).
	Type string `json:"type"`
	// Public URL where webhooks will be sent, for example **https://www.domain.com/webhook-endpoint**.
	Url string `json:"url"`
	// Username to access the webhook URL.
	Username *string `json:"username,omitempty"`
}

// NewCreateMerchantWebhookRequest instantiates a new CreateMerchantWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMerchantWebhookRequest(active bool, communicationFormat string, type_ string, url string) *CreateMerchantWebhookRequest {
	this := CreateMerchantWebhookRequest{}
	this.Active = active
	this.CommunicationFormat = communicationFormat
	this.Type = type_
	this.Url = url
	return &this
}

// NewCreateMerchantWebhookRequestWithDefaults instantiates a new CreateMerchantWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMerchantWebhookRequestWithDefaults() *CreateMerchantWebhookRequest {
	this := CreateMerchantWebhookRequest{}
	return &this
}

// GetAcceptsExpiredCertificate returns the AcceptsExpiredCertificate field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetAcceptsExpiredCertificate() bool {
	if o == nil || o.AcceptsExpiredCertificate == nil {
		var ret bool
		return ret
	}
	return *o.AcceptsExpiredCertificate
}

// GetAcceptsExpiredCertificateOk returns a tuple with the AcceptsExpiredCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetAcceptsExpiredCertificateOk() (*bool, bool) {
	if o == nil || o.AcceptsExpiredCertificate == nil {
		return nil, false
	}
	return o.AcceptsExpiredCertificate, true
}

// HasAcceptsExpiredCertificate returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasAcceptsExpiredCertificate() bool {
	if o != nil && o.AcceptsExpiredCertificate != nil {
		return true
	}

	return false
}

// SetAcceptsExpiredCertificate gets a reference to the given bool and assigns it to the AcceptsExpiredCertificate field.
func (o *CreateMerchantWebhookRequest) SetAcceptsExpiredCertificate(v bool) {
	o.AcceptsExpiredCertificate = &v
}

// GetAcceptsSelfSignedCertificate returns the AcceptsSelfSignedCertificate field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetAcceptsSelfSignedCertificate() bool {
	if o == nil || o.AcceptsSelfSignedCertificate == nil {
		var ret bool
		return ret
	}
	return *o.AcceptsSelfSignedCertificate
}

// GetAcceptsSelfSignedCertificateOk returns a tuple with the AcceptsSelfSignedCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetAcceptsSelfSignedCertificateOk() (*bool, bool) {
	if o == nil || o.AcceptsSelfSignedCertificate == nil {
		return nil, false
	}
	return o.AcceptsSelfSignedCertificate, true
}

// HasAcceptsSelfSignedCertificate returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasAcceptsSelfSignedCertificate() bool {
	if o != nil && o.AcceptsSelfSignedCertificate != nil {
		return true
	}

	return false
}

// SetAcceptsSelfSignedCertificate gets a reference to the given bool and assigns it to the AcceptsSelfSignedCertificate field.
func (o *CreateMerchantWebhookRequest) SetAcceptsSelfSignedCertificate(v bool) {
	o.AcceptsSelfSignedCertificate = &v
}

// GetAcceptsUntrustedRootCertificate returns the AcceptsUntrustedRootCertificate field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetAcceptsUntrustedRootCertificate() bool {
	if o == nil || o.AcceptsUntrustedRootCertificate == nil {
		var ret bool
		return ret
	}
	return *o.AcceptsUntrustedRootCertificate
}

// GetAcceptsUntrustedRootCertificateOk returns a tuple with the AcceptsUntrustedRootCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetAcceptsUntrustedRootCertificateOk() (*bool, bool) {
	if o == nil || o.AcceptsUntrustedRootCertificate == nil {
		return nil, false
	}
	return o.AcceptsUntrustedRootCertificate, true
}

// HasAcceptsUntrustedRootCertificate returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasAcceptsUntrustedRootCertificate() bool {
	if o != nil && o.AcceptsUntrustedRootCertificate != nil {
		return true
	}

	return false
}

// SetAcceptsUntrustedRootCertificate gets a reference to the given bool and assigns it to the AcceptsUntrustedRootCertificate field.
func (o *CreateMerchantWebhookRequest) SetAcceptsUntrustedRootCertificate(v bool) {
	o.AcceptsUntrustedRootCertificate = &v
}

// GetActive returns the Active field value
func (o *CreateMerchantWebhookRequest) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *CreateMerchantWebhookRequest) SetActive(v bool) {
	o.Active = v
}

// GetAdditionalSettings returns the AdditionalSettings field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetAdditionalSettings() AdditionalSettings {
	if o == nil || o.AdditionalSettings == nil {
		var ret AdditionalSettings
		return ret
	}
	return *o.AdditionalSettings
}

// GetAdditionalSettingsOk returns a tuple with the AdditionalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetAdditionalSettingsOk() (*AdditionalSettings, bool) {
	if o == nil || o.AdditionalSettings == nil {
		return nil, false
	}
	return o.AdditionalSettings, true
}

// HasAdditionalSettings returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasAdditionalSettings() bool {
	if o != nil && o.AdditionalSettings != nil {
		return true
	}

	return false
}

// SetAdditionalSettings gets a reference to the given AdditionalSettings and assigns it to the AdditionalSettings field.
func (o *CreateMerchantWebhookRequest) SetAdditionalSettings(v AdditionalSettings) {
	o.AdditionalSettings = &v
}

// GetCommunicationFormat returns the CommunicationFormat field value
func (o *CreateMerchantWebhookRequest) GetCommunicationFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommunicationFormat
}

// GetCommunicationFormatOk returns a tuple with the CommunicationFormat field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetCommunicationFormatOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommunicationFormat, true
}

// SetCommunicationFormat sets field value
func (o *CreateMerchantWebhookRequest) SetCommunicationFormat(v string) {
	o.CommunicationFormat = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateMerchantWebhookRequest) SetDescription(v string) {
	o.Description = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetNetworkType() string {
	if o == nil || o.NetworkType == nil {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetNetworkTypeOk() (*string, bool) {
	if o == nil || o.NetworkType == nil {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasNetworkType() bool {
	if o != nil && o.NetworkType != nil {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *CreateMerchantWebhookRequest) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CreateMerchantWebhookRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPopulateSoapActionHeader returns the PopulateSoapActionHeader field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetPopulateSoapActionHeader() bool {
	if o == nil || o.PopulateSoapActionHeader == nil {
		var ret bool
		return ret
	}
	return *o.PopulateSoapActionHeader
}

// GetPopulateSoapActionHeaderOk returns a tuple with the PopulateSoapActionHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetPopulateSoapActionHeaderOk() (*bool, bool) {
	if o == nil || o.PopulateSoapActionHeader == nil {
		return nil, false
	}
	return o.PopulateSoapActionHeader, true
}

// HasPopulateSoapActionHeader returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasPopulateSoapActionHeader() bool {
	if o != nil && o.PopulateSoapActionHeader != nil {
		return true
	}

	return false
}

// SetPopulateSoapActionHeader gets a reference to the given bool and assigns it to the PopulateSoapActionHeader field.
func (o *CreateMerchantWebhookRequest) SetPopulateSoapActionHeader(v bool) {
	o.PopulateSoapActionHeader = &v
}

// GetSslVersion returns the SslVersion field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetSslVersion() string {
	if o == nil || o.SslVersion == nil {
		var ret string
		return ret
	}
	return *o.SslVersion
}

// GetSslVersionOk returns a tuple with the SslVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetSslVersionOk() (*string, bool) {
	if o == nil || o.SslVersion == nil {
		return nil, false
	}
	return o.SslVersion, true
}

// HasSslVersion returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasSslVersion() bool {
	if o != nil && o.SslVersion != nil {
		return true
	}

	return false
}

// SetSslVersion gets a reference to the given string and assigns it to the SslVersion field.
func (o *CreateMerchantWebhookRequest) SetSslVersion(v string) {
	o.SslVersion = &v
}

// GetType returns the Type field value
func (o *CreateMerchantWebhookRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateMerchantWebhookRequest) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *CreateMerchantWebhookRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreateMerchantWebhookRequest) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateMerchantWebhookRequest) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantWebhookRequest) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateMerchantWebhookRequest) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateMerchantWebhookRequest) SetUsername(v string) {
	o.Username = &v
}

func (o CreateMerchantWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptsExpiredCertificate != nil {
		toSerialize["acceptsExpiredCertificate"] = o.AcceptsExpiredCertificate
	}
	if o.AcceptsSelfSignedCertificate != nil {
		toSerialize["acceptsSelfSignedCertificate"] = o.AcceptsSelfSignedCertificate
	}
	if o.AcceptsUntrustedRootCertificate != nil {
		toSerialize["acceptsUntrustedRootCertificate"] = o.AcceptsUntrustedRootCertificate
	}
	if true {
		toSerialize["active"] = o.Active
	}
	if o.AdditionalSettings != nil {
		toSerialize["additionalSettings"] = o.AdditionalSettings
	}
	if true {
		toSerialize["communicationFormat"] = o.CommunicationFormat
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.NetworkType != nil {
		toSerialize["networkType"] = o.NetworkType
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PopulateSoapActionHeader != nil {
		toSerialize["populateSoapActionHeader"] = o.PopulateSoapActionHeader
	}
	if o.SslVersion != nil {
		toSerialize["sslVersion"] = o.SslVersion
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMerchantWebhookRequest struct {
	value *CreateMerchantWebhookRequest
	isSet bool
}

func (v NullableCreateMerchantWebhookRequest) Get() *CreateMerchantWebhookRequest {
	return v.value
}

func (v *NullableCreateMerchantWebhookRequest) Set(val *CreateMerchantWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMerchantWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMerchantWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMerchantWebhookRequest(val *CreateMerchantWebhookRequest) *NullableCreateMerchantWebhookRequest {
	return &NullableCreateMerchantWebhookRequest{value: val, isSet: true}
}

func (v NullableCreateMerchantWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMerchantWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


