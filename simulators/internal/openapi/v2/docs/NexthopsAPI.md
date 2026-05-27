# \NexthopsAPI

All URIs are relative to *https://eric-pc-routing-information-base:443/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeregisterNexthop**](NexthopsAPI.md#DeregisterNexthop) | **Delete** /network-instances/{networkInstanceName}/producers/{producerId}/nexthops | 
[**ReceiveNexthopRegistrationEOF**](NexthopsAPI.md#ReceiveNexthopRegistrationEOF) | **Post** /network-instances/{networkInstanceName}/producers/{producerId}/nexthop-reg-eof | 
[**RegisterNexthop**](NexthopsAPI.md#RegisterNexthop) | **Post** /network-instances/{networkInstanceName}/producers/{producerId}/nexthops | 



## DeregisterNexthop

> DeregisterNexthop(ctx, networkInstanceName, producerId).NexthopRegister(nexthopRegister).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkInstanceName := "networkInstanceName_example" // string | networkInstance to which the nexthop belongs to.
	producerId := "producerId_example" // string | producer name, please refer to the clientIdentifier provided by producer when it register to RIB.
	nexthopRegister := *openapiclient.NewNexthopRegister() // NexthopRegister | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NexthopsAPI.DeregisterNexthop(context.Background(), networkInstanceName, producerId).NexthopRegister(nexthopRegister).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexthopsAPI.DeregisterNexthop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | networkInstance to which the nexthop belongs to. | 
**producerId** | **string** | producer name, please refer to the clientIdentifier provided by producer when it register to RIB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterNexthopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nexthopRegister** | [**NexthopRegister**](NexthopRegister.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReceiveNexthopRegistrationEOF

> ReceiveNexthopRegistrationEOF(ctx, networkInstanceName, producerId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkInstanceName := "networkInstanceName_example" // string | networkInstance to which the EOF belongs to.
	producerId := "producerId_example" // string | producer name, please refer to the clientIdentifier provided by producer when it register to RIB.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NexthopsAPI.ReceiveNexthopRegistrationEOF(context.Background(), networkInstanceName, producerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexthopsAPI.ReceiveNexthopRegistrationEOF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | networkInstance to which the EOF belongs to. | 
**producerId** | **string** | producer name, please refer to the clientIdentifier provided by producer when it register to RIB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReceiveNexthopRegistrationEOFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterNexthop

> RegisterNexthop(ctx, networkInstanceName, producerId).NexthopRegister(nexthopRegister).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkInstanceName := "networkInstanceName_example" // string | networkInstance to which the nexthop belongs to.
	producerId := "producerId_example" // string | producer name, please refer to the clientIdentifier provided by producer when it register to RIB.
	nexthopRegister := *openapiclient.NewNexthopRegister() // NexthopRegister | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NexthopsAPI.RegisterNexthop(context.Background(), networkInstanceName, producerId).NexthopRegister(nexthopRegister).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexthopsAPI.RegisterNexthop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | networkInstance to which the nexthop belongs to. | 
**producerId** | **string** | producer name, please refer to the clientIdentifier provided by producer when it register to RIB. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterNexthopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nexthopRegister** | [**NexthopRegister**](NexthopRegister.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

