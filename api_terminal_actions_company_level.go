/*
Management API

Configure and manage your Adyen company and merchant accounts, stores, and payment terminals. ## Authentication Each request to the Management API must be signed with an API key. [Generate your API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key) in the Customer Area and then set this key to the `X-API-Key` header value.  To access the live endpoints, you need to generate a new API key in your live Customer Area. ## Versioning  Management API handles versioning as part of the endpoint URL. For example, to send a request to version 1 of the `/companies/{companyId}/webhooks` endpoint, use:  ```text https://management-test.adyen.com/v1/companies/{companyId}/webhooks ```

API version: 1
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// TerminalActionsCompanyLevelApiService TerminalActionsCompanyLevelApi service
type TerminalActionsCompanyLevelApiService service

type ApiGetCompaniesCompanyIdAndroidAppsRequest struct {
	ctx context.Context
	ApiService *TerminalActionsCompanyLevelApiService
	companyId string
	pageNumber *int32
	pageSize *int32
}

// The number of the page to fetch.
func (r ApiGetCompaniesCompanyIdAndroidAppsRequest) PageNumber(pageNumber int32) ApiGetCompaniesCompanyIdAndroidAppsRequest {
	r.pageNumber = &pageNumber
	return r
}
// The number of items to have on a page, maximum 100. The default is 20 items on a page.
func (r ApiGetCompaniesCompanyIdAndroidAppsRequest) PageSize(pageSize int32) ApiGetCompaniesCompanyIdAndroidAppsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetCompaniesCompanyIdAndroidAppsRequest) Execute() (*AndroidAppsResponse, *http.Response, error) {
	return r.ApiService.GetCompaniesCompanyIdAndroidAppsExecute(r)
}

/*
GetCompaniesCompanyIdAndroidApps Get a list of Android apps

Returns a list of the Android apps that are available for the company identified in the path. 
These apps have been uploaded to Adyen and can be installed or uninstalled on Android payment terminals through [terminal actions](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api).

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal actions read
* Management API—Terminal actions read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return ApiGetCompaniesCompanyIdAndroidAppsRequest
*/
func (a *TerminalActionsCompanyLevelApiService) GetCompaniesCompanyIdAndroidApps(ctx context.Context, companyId string) ApiGetCompaniesCompanyIdAndroidAppsRequest {
	return ApiGetCompaniesCompanyIdAndroidAppsRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return AndroidAppsResponse
func (a *TerminalActionsCompanyLevelApiService) GetCompaniesCompanyIdAndroidAppsExecute(r ApiGetCompaniesCompanyIdAndroidAppsRequest) (*AndroidAppsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AndroidAppsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TerminalActionsCompanyLevelApiService.GetCompaniesCompanyIdAndroidApps")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{companyId}/androidApps"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageNumber != nil {
		localVarQueryParams.Add("pageNumber", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCompaniesCompanyIdAndroidCertificatesRequest struct {
	ctx context.Context
	ApiService *TerminalActionsCompanyLevelApiService
	companyId string
	pageNumber *int32
	pageSize *int32
}

// The number of the page to fetch.
func (r ApiGetCompaniesCompanyIdAndroidCertificatesRequest) PageNumber(pageNumber int32) ApiGetCompaniesCompanyIdAndroidCertificatesRequest {
	r.pageNumber = &pageNumber
	return r
}
// The number of items to have on a page, maximum 100. The default is 20 items on a page.
func (r ApiGetCompaniesCompanyIdAndroidCertificatesRequest) PageSize(pageSize int32) ApiGetCompaniesCompanyIdAndroidCertificatesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetCompaniesCompanyIdAndroidCertificatesRequest) Execute() (*AndroidCertificatesResponse, *http.Response, error) {
	return r.ApiService.GetCompaniesCompanyIdAndroidCertificatesExecute(r)
}

/*
GetCompaniesCompanyIdAndroidCertificates Get a list of Android certificates

Returns a list of the Android certificates that are available for the company identified in the path.
Typically, these certificates enable running apps on Android payment terminals. The certifcates in the list have been uploaded to Adyen and can be installed or uninstalled on Android terminals through [terminal actions](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api).

To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal actions read
* Management API—Terminal actions read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return ApiGetCompaniesCompanyIdAndroidCertificatesRequest
*/
func (a *TerminalActionsCompanyLevelApiService) GetCompaniesCompanyIdAndroidCertificates(ctx context.Context, companyId string) ApiGetCompaniesCompanyIdAndroidCertificatesRequest {
	return ApiGetCompaniesCompanyIdAndroidCertificatesRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return AndroidCertificatesResponse
func (a *TerminalActionsCompanyLevelApiService) GetCompaniesCompanyIdAndroidCertificatesExecute(r ApiGetCompaniesCompanyIdAndroidCertificatesRequest) (*AndroidCertificatesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AndroidCertificatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TerminalActionsCompanyLevelApiService.GetCompaniesCompanyIdAndroidCertificates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{companyId}/androidCertificates"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageNumber != nil {
		localVarQueryParams.Add("pageNumber", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCompaniesCompanyIdTerminalActionsRequest struct {
	ctx context.Context
	ApiService *TerminalActionsCompanyLevelApiService
	companyId string
	pageNumber *int32
	pageSize *int32
	status *string
	type_ *string
}

// The number of the page to fetch.
func (r ApiGetCompaniesCompanyIdTerminalActionsRequest) PageNumber(pageNumber int32) ApiGetCompaniesCompanyIdTerminalActionsRequest {
	r.pageNumber = &pageNumber
	return r
}
// The number of items to have on a page, maximum 100. The default is 20 items on a page.
func (r ApiGetCompaniesCompanyIdTerminalActionsRequest) PageSize(pageSize int32) ApiGetCompaniesCompanyIdTerminalActionsRequest {
	r.pageSize = &pageSize
	return r
}
// Returns terminal actions with the specified status.  Allowed values: **pending**, **successful**, **failed**, **cancelled**, **tryLater**.
func (r ApiGetCompaniesCompanyIdTerminalActionsRequest) Status(status string) ApiGetCompaniesCompanyIdTerminalActionsRequest {
	r.status = &status
	return r
}
// Returns terminal actions of the specified type.  Allowed values: **InstallAndroidApp**, **UninstallAndroidApp**, **InstallAndroidCertificate**, **UninstallAndroidCertificate**.
func (r ApiGetCompaniesCompanyIdTerminalActionsRequest) Type_(type_ string) ApiGetCompaniesCompanyIdTerminalActionsRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetCompaniesCompanyIdTerminalActionsRequest) Execute() (*ListExternalTerminalActionsResponse, *http.Response, error) {
	return r.ApiService.GetCompaniesCompanyIdTerminalActionsExecute(r)
}

/*
GetCompaniesCompanyIdTerminalActions Get a list of terminal actions

Returns the [terminal actions](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api) that have been scheduled for the company identified in the path.The response doesn't include actions that are scheduled by Adyen.
To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal actions read
* Management API—Terminal actions read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @return ApiGetCompaniesCompanyIdTerminalActionsRequest
*/
func (a *TerminalActionsCompanyLevelApiService) GetCompaniesCompanyIdTerminalActions(ctx context.Context, companyId string) ApiGetCompaniesCompanyIdTerminalActionsRequest {
	return ApiGetCompaniesCompanyIdTerminalActionsRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
	}
}

// Execute executes the request
//  @return ListExternalTerminalActionsResponse
func (a *TerminalActionsCompanyLevelApiService) GetCompaniesCompanyIdTerminalActionsExecute(r ApiGetCompaniesCompanyIdTerminalActionsRequest) (*ListExternalTerminalActionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListExternalTerminalActionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TerminalActionsCompanyLevelApiService.GetCompaniesCompanyIdTerminalActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{companyId}/terminalActions"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageNumber != nil {
		localVarQueryParams.Add("pageNumber", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCompaniesCompanyIdTerminalActionsActionIdRequest struct {
	ctx context.Context
	ApiService *TerminalActionsCompanyLevelApiService
	companyId string
	actionId string
}


func (r ApiGetCompaniesCompanyIdTerminalActionsActionIdRequest) Execute() (*ExternalTerminalAction, *http.Response, error) {
	return r.ApiService.GetCompaniesCompanyIdTerminalActionsActionIdExecute(r)
}

/*
GetCompaniesCompanyIdTerminalActionsActionId Get terminal action

Returns the details of the [terminal action](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api) identified in the path.
To make this request, your API credential must have one of the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal actions read
* Management API—Terminal actions read and write

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param companyId The unique identifier of the company account.
 @param actionId The unique identifier of the terminal action.
 @return ApiGetCompaniesCompanyIdTerminalActionsActionIdRequest
*/
func (a *TerminalActionsCompanyLevelApiService) GetCompaniesCompanyIdTerminalActionsActionId(ctx context.Context, companyId string, actionId string) ApiGetCompaniesCompanyIdTerminalActionsActionIdRequest {
	return ApiGetCompaniesCompanyIdTerminalActionsActionIdRequest{
		ApiService: a,
		ctx: ctx,
		companyId: companyId,
		actionId: actionId,
	}
}

// Execute executes the request
//  @return ExternalTerminalAction
func (a *TerminalActionsCompanyLevelApiService) GetCompaniesCompanyIdTerminalActionsActionIdExecute(r ApiGetCompaniesCompanyIdTerminalActionsActionIdRequest) (*ExternalTerminalAction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExternalTerminalAction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TerminalActionsCompanyLevelApiService.GetCompaniesCompanyIdTerminalActionsActionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/companies/{companyId}/terminalActions/{actionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"companyId"+"}", url.PathEscape(parameterToString(r.companyId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"actionId"+"}", url.PathEscape(parameterToString(r.actionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v RestServiceError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
