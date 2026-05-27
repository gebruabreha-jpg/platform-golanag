# Producer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Indicates client name, for either Producer or Consumer. For route-engine, client name is like {cre-pod-name} + {static/bgp/ifm}, for data-plane, client name is like {dp-pod-name}, for client that is service, client name is like {service-name}. | [optional] 
**RouteHoldTime** | Pointer to **int32** | for producer client, unit is second. RIB will hold the routes for the specified time after the producer client get dead. If not present, RIB will give a default value. | [optional] 
**KeepaliveTimeout** | Pointer to **int32** | Unit is second. Indicates how long RIB has not received the client&#39;s heartbeat before it considers the client to be dead. If not present, RIB will give a default value. If value &#x3D; 0, indicates that the RIB will never delete the producer based on heartbeat(For service producer scenario). | [optional] 
**StaleRoute** | Pointer to **bool** | Optional field. This field is used to tell RIB when the producer registers whether to trigger the route stale operation for this registration. If value &#x3D; true, it means that the route aging operation will be performed on the producer registered this time. If value &#x3D; false, it means that the route aging operation will not be performed on the producer registered this time. Without this field, the default is staleRoute &#x3D; true, that is, the normal route aging process will be performed on the producer when re-registering. | [optional] 
**UrlCallback** | Pointer to **string** | The location will be called to notify nexthop change to producer client The location will be called to notify prefix change to producer client. The location will be called to notify routes that need to be redistributed to producer client(routing protocol, for example, BGP). The location will be called to notify EOF information to producer client. | [optional] 

## Methods

### NewProducer

`func NewProducer() *Producer`

NewProducer instantiates a new Producer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProducerWithDefaults

`func NewProducerWithDefaults() *Producer`

NewProducerWithDefaults instantiates a new Producer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Producer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Producer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Producer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Producer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRouteHoldTime

`func (o *Producer) GetRouteHoldTime() int32`

GetRouteHoldTime returns the RouteHoldTime field if non-nil, zero value otherwise.

### GetRouteHoldTimeOk

`func (o *Producer) GetRouteHoldTimeOk() (*int32, bool)`

GetRouteHoldTimeOk returns a tuple with the RouteHoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteHoldTime

`func (o *Producer) SetRouteHoldTime(v int32)`

SetRouteHoldTime sets RouteHoldTime field to given value.

### HasRouteHoldTime

`func (o *Producer) HasRouteHoldTime() bool`

HasRouteHoldTime returns a boolean if a field has been set.

### GetKeepaliveTimeout

`func (o *Producer) GetKeepaliveTimeout() int32`

GetKeepaliveTimeout returns the KeepaliveTimeout field if non-nil, zero value otherwise.

### GetKeepaliveTimeoutOk

`func (o *Producer) GetKeepaliveTimeoutOk() (*int32, bool)`

GetKeepaliveTimeoutOk returns a tuple with the KeepaliveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveTimeout

`func (o *Producer) SetKeepaliveTimeout(v int32)`

SetKeepaliveTimeout sets KeepaliveTimeout field to given value.

### HasKeepaliveTimeout

`func (o *Producer) HasKeepaliveTimeout() bool`

HasKeepaliveTimeout returns a boolean if a field has been set.

### GetStaleRoute

`func (o *Producer) GetStaleRoute() bool`

GetStaleRoute returns the StaleRoute field if non-nil, zero value otherwise.

### GetStaleRouteOk

`func (o *Producer) GetStaleRouteOk() (*bool, bool)`

GetStaleRouteOk returns a tuple with the StaleRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleRoute

`func (o *Producer) SetStaleRoute(v bool)`

SetStaleRoute sets StaleRoute field to given value.

### HasStaleRoute

`func (o *Producer) HasStaleRoute() bool`

HasStaleRoute returns a boolean if a field has been set.

### GetUrlCallback

`func (o *Producer) GetUrlCallback() string`

GetUrlCallback returns the UrlCallback field if non-nil, zero value otherwise.

### GetUrlCallbackOk

`func (o *Producer) GetUrlCallbackOk() (*string, bool)`

GetUrlCallbackOk returns a tuple with the UrlCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCallback

`func (o *Producer) SetUrlCallback(v string)`

SetUrlCallback sets UrlCallback field to given value.

### HasUrlCallback

`func (o *Producer) HasUrlCallback() bool`

HasUrlCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


