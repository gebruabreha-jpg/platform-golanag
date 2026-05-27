# \BulkDeleteRoutesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteRoutes**](BulkDeleteRoutesAPI.md#BulkDeleteRoutes) | **Delete** /v1/network-instances/{network-instance-name}/routes | 



## BulkDeleteRoutes

> BulkDeleteRoutes(ctx, networkInstanceName).BulkDeleteRoutes(bulkDeleteRoutes).Execute()





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
	networkInstanceName := "networkInstanceName_example" // string | name of network-instance that the routes belongs to.
	bulkDeleteRoutes := *openapiclient.NewBulkDeleteRoutes() // BulkDeleteRoutes | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BulkDeleteRoutesAPI.BulkDeleteRoutes(context.Background(), networkInstanceName).BulkDeleteRoutes(bulkDeleteRoutes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkDeleteRoutesAPI.BulkDeleteRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInstanceName** | **string** | name of network-instance that the routes belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkDeleteRoutes** | [**BulkDeleteRoutes**](BulkDeleteRoutes.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

