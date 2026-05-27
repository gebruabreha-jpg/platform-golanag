# ProducerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** |  | 
**RtHoldTime** | **int32** | for producer client, unit is second. RIB will hold the routes for the specified time after the producer client get dead. If not present, RIB will give a default value. | 
**KeepaliveTimeout** | **int32** | Unit is second. Indicates how long RIB has not received the client&#39;s heartbeat before it considers the client to be dead. If not present, RIB will give a default value. If value &#x3D; 0, indicates that the RIB will never delete the producer based on heartbeat(For service producer scenario). | 
**StaleRoute** | Pointer to **string** | Optional field. This field is used to tell RIB when the producer registers whether to trigger the route stale operation for this registration. If value &#x3D; true, it means that the route aging operation will be performed on the producer registered this time; if value &#x3D; false, it means that the route aging operation will not be performed on the producer registered this time. Without this field, the default is stale_route &#x3D; true, that is, the normal route aging process will be performed on the producer when re-registering. | [optional] 
**UrlCallback** | Pointer to **string** | The location will be called to notify nexthop change to producer client The location will be called to notify prefix change to producer client. The location will be called to notify routes that need to be redistributed to producer client(routing protocol, for example, BGP). The location will be called to notify EOF information to producer client. | [optional] 

## Methods

### NewProducerInfo

`func NewProducerInfo(serviceName string, rtHoldTime int32, keepaliveTimeout int32, ) *ProducerInfo`

NewProducerInfo instantiates a new ProducerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProducerInfoWithDefaults

`func NewProducerInfoWithDefaults() *ProducerInfo`

NewProducerInfoWithDefaults instantiates a new ProducerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *ProducerInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ProducerInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ProducerInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetRtHoldTime

`func (o *ProducerInfo) GetRtHoldTime() int32`

GetRtHoldTime returns the RtHoldTime field if non-nil, zero value otherwise.

### GetRtHoldTimeOk

`func (o *ProducerInfo) GetRtHoldTimeOk() (*int32, bool)`

GetRtHoldTimeOk returns a tuple with the RtHoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtHoldTime

`func (o *ProducerInfo) SetRtHoldTime(v int32)`

SetRtHoldTime sets RtHoldTime field to given value.


### GetKeepaliveTimeout

`func (o *ProducerInfo) GetKeepaliveTimeout() int32`

GetKeepaliveTimeout returns the KeepaliveTimeout field if non-nil, zero value otherwise.

### GetKeepaliveTimeoutOk

`func (o *ProducerInfo) GetKeepaliveTimeoutOk() (*int32, bool)`

GetKeepaliveTimeoutOk returns a tuple with the KeepaliveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveTimeout

`func (o *ProducerInfo) SetKeepaliveTimeout(v int32)`

SetKeepaliveTimeout sets KeepaliveTimeout field to given value.


### GetStaleRoute

`func (o *ProducerInfo) GetStaleRoute() string`

GetStaleRoute returns the StaleRoute field if non-nil, zero value otherwise.

### GetStaleRouteOk

`func (o *ProducerInfo) GetStaleRouteOk() (*string, bool)`

GetStaleRouteOk returns a tuple with the StaleRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleRoute

`func (o *ProducerInfo) SetStaleRoute(v string)`

SetStaleRoute sets StaleRoute field to given value.

### HasStaleRoute

`func (o *ProducerInfo) HasStaleRoute() bool`

HasStaleRoute returns a boolean if a field has been set.

### GetUrlCallback

`func (o *ProducerInfo) GetUrlCallback() string`

GetUrlCallback returns the UrlCallback field if non-nil, zero value otherwise.

### GetUrlCallbackOk

`func (o *ProducerInfo) GetUrlCallbackOk() (*string, bool)`

GetUrlCallbackOk returns a tuple with the UrlCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCallback

`func (o *ProducerInfo) SetUrlCallback(v string)`

SetUrlCallback sets UrlCallback field to given value.

### HasUrlCallback

`func (o *ProducerInfo) HasUrlCallback() bool`

HasUrlCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


