# \AppsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppById**](AppsApi.md#GetAppById) | **Get** /apps/{appId} | Get app by appId
[**GetApps**](AppsApi.md#GetApps) | **Get** /apps | Retrieve list of apps
[**InitAppFromDevice**](AppsApi.md#InitAppFromDevice) | **Post** /apps/{appId}/init | Init call from sdk
[**UpdateApp**](AppsApi.md#UpdateApp) | **Put** /apps/{appId} | Update app by id


# **GetAppById**
> App GetAppById(ctx, appId)
Get app by appId

Retrieve all information for a single app including all child information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**| The id for the app that needs to be fetched. | 

### Return type

[**App**](App.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApps**
> Apps GetApps(ctx, )
Retrieve list of apps

Returns root level information for all apps 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Apps**](Apps.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitAppFromDevice**
> AppInitResponse InitAppFromDevice(ctx, appId, body)
Init call from sdk

Capture metrics from device and return if the app version they are using is disabled and has a set disabled message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**| The id for the app. | 
  **body** | [**AppInit**](AppInit.md)| Updated app object | 

### Return type

[**AppInitResponse**](AppInitResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApp**
> UpdateApp(ctx, appId, body)
Update app by id

Update a single app using the app id, including updating version information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**| The id for the app that needs to be updated. | 
  **body** | [**App**](App.md)| Updated app object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

