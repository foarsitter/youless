# \DefaultApi

All URIs are relative to *http://192.168.178.206*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBasicInfo**](DefaultApi.md#GetBasicInfo) | **Get** /a | 
[**GetDeviceInfo**](DefaultApi.md#GetDeviceInfo) | **Get** /d | 
[**GetUploadedValues**](DefaultApi.md#GetUploadedValues) | **Get** /e | Uploaded Values



## GetBasicInfo

> BasicStatusInfo GetBasicInfo(ctx).F(f).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    f := openapiclient.OutputFormat("j") // OutputFormat |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetBasicInfo(context.Background()).F(f).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBasicInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicInfo`: BasicStatusInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBasicInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **f** | [**OutputFormat**](OutputFormat.md) |  | 

### Return type

[**BasicStatusInfo**](BasicStatusInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: appliation/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceInfo

> DeviceSchema GetDeviceInfo(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetDeviceInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDeviceInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceInfo`: DeviceSchema
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDeviceInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInfoRequest struct via the builder pattern


### Return type

[**DeviceSchema**](DeviceSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: appliation/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUploadedValues

> []UploadedValues GetUploadedValues(ctx).Execute()

Uploaded Values



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetUploadedValues(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUploadedValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUploadedValues`: []UploadedValues
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUploadedValues`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUploadedValuesRequest struct via the builder pattern


### Return type

[**[]UploadedValues**](UploadedValues.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: appliation/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

