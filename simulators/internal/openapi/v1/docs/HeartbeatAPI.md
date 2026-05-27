# \HeartbeatAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Heartbeat**](HeartbeatAPI.md#Heartbeat) | **Post** /v1/heartbeats | 



## Heartbeat

> Heartbeat(ctx).Heartbeat(heartbeat).Execute()





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
	heartbeat := *openapiclient.NewHeartbeat() // Heartbeat | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HeartbeatAPI.Heartbeat(context.Background()).Heartbeat(heartbeat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeartbeatAPI.Heartbeat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHeartbeatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **heartbeat** | [**Heartbeat**](Heartbeat.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

