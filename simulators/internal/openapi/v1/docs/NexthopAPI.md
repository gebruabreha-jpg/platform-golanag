# \NexthopAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeregisterNexthop**](NexthopAPI.md#DeregisterNexthop) | **Delete** /v1/network-instances/{network-instance-name}/nexthops | 
[**RegisterNexthop**](NexthopAPI.md#RegisterNexthop) | **Post** /v1/network-instances/{network-instance-name}/nexthops | 



## DeregisterNexthop

> DeregisterNexthop(ctx, networkInstanceName).NexthopRegister(nexthopRegister).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | network-instance to which the nexthop belongs to.
	nexthopRegister := *openapiclient.NewNexthopRegister() // NexthopRegister | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NexthopAPI.DeregisterNexthop(context.Background(), networkInstanceName).NexthopRegister(nexthopRegister).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexthopAPI.DeregisterNexthop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | network-instance to which the nexthop belongs to. | 

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


## RegisterNexthop

> RegisterNexthop(ctx, networkInstanceName).NexthopRegister(nexthopRegister).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | network-instance to which the nexthop belongs to.
	nexthopRegister := *openapiclient.NewNexthopRegister() // NexthopRegister | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NexthopAPI.RegisterNexthop(context.Background(), networkInstanceName).NexthopRegister(nexthopRegister).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexthopAPI.RegisterNexthop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | network-instance to which the nexthop belongs to. | 

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

