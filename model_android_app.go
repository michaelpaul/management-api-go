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

// AndroidApp struct for AndroidApp
type AndroidApp struct {
	// The description that was provided when uploading the app. The description is not shown on the terminal.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the app.
	Id string `json:"id"`
	// The app name that is shown on the terminal.
	Label *string `json:"label,omitempty"`
	// The package name of the app.
	PackageName *string `json:"packageName,omitempty"`
	// The status of the app. Possible values:  * `processing`: The app is being signed and converted to a format that the terminal can handle. * `error`: Something went wrong. Check that the app matches the [requirements](https://docs.adyen.com/point-of-sale/android-terminals/app-requirements). * `invalid`: There is something wrong with the APK file of the app. * `ready`: The app has been signed and converted. * `archived`: The app is no longer available.
	Status string `json:"status"`
	// The internal version number of the app.
	VersionCode *int32 `json:"versionCode,omitempty"`
	// The app version number that is shown on the terminal.
	VersionName *string `json:"versionName,omitempty"`
}

// NewAndroidApp instantiates a new AndroidApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAndroidApp(id string, status string) *AndroidApp {
	this := AndroidApp{}
	this.Id = id
	this.Status = status
	return &this
}

// NewAndroidAppWithDefaults instantiates a new AndroidApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAndroidAppWithDefaults() *AndroidApp {
	this := AndroidApp{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AndroidApp) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AndroidApp) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AndroidApp) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *AndroidApp) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AndroidApp) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AndroidApp) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AndroidApp) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AndroidApp) SetLabel(v string) {
	o.Label = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *AndroidApp) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *AndroidApp) HasPackageName() bool {
	if o != nil && o.PackageName != nil {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *AndroidApp) SetPackageName(v string) {
	o.PackageName = &v
}

// GetStatus returns the Status field value
func (o *AndroidApp) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AndroidApp) SetStatus(v string) {
	o.Status = v
}

// GetVersionCode returns the VersionCode field value if set, zero value otherwise.
func (o *AndroidApp) GetVersionCode() int32 {
	if o == nil || o.VersionCode == nil {
		var ret int32
		return ret
	}
	return *o.VersionCode
}

// GetVersionCodeOk returns a tuple with the VersionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetVersionCodeOk() (*int32, bool) {
	if o == nil || o.VersionCode == nil {
		return nil, false
	}
	return o.VersionCode, true
}

// HasVersionCode returns a boolean if a field has been set.
func (o *AndroidApp) HasVersionCode() bool {
	if o != nil && o.VersionCode != nil {
		return true
	}

	return false
}

// SetVersionCode gets a reference to the given int32 and assigns it to the VersionCode field.
func (o *AndroidApp) SetVersionCode(v int32) {
	o.VersionCode = &v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *AndroidApp) GetVersionName() string {
	if o == nil || o.VersionName == nil {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AndroidApp) GetVersionNameOk() (*string, bool) {
	if o == nil || o.VersionName == nil {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *AndroidApp) HasVersionName() bool {
	if o != nil && o.VersionName != nil {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *AndroidApp) SetVersionName(v string) {
	o.VersionName = &v
}

func (o AndroidApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.PackageName != nil {
		toSerialize["packageName"] = o.PackageName
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.VersionCode != nil {
		toSerialize["versionCode"] = o.VersionCode
	}
	if o.VersionName != nil {
		toSerialize["versionName"] = o.VersionName
	}
	return json.Marshal(toSerialize)
}

type NullableAndroidApp struct {
	value *AndroidApp
	isSet bool
}

func (v NullableAndroidApp) Get() *AndroidApp {
	return v.value
}

func (v *NullableAndroidApp) Set(val *AndroidApp) {
	v.value = val
	v.isSet = true
}

func (v NullableAndroidApp) IsSet() bool {
	return v.isSet
}

func (v *NullableAndroidApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAndroidApp(val *AndroidApp) *NullableAndroidApp {
	return &NullableAndroidApp{value: val, isSet: true}
}

func (v NullableAndroidApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAndroidApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


