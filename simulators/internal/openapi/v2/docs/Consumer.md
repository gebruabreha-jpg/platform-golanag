# Consumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Indicates client name, for either Producer or Consumer. For route-engine, client name is like {cre-pod-name} + {static/bgp/ifm}, for data-plane, client name is like {dp-pod-name}, for client that is service, client name is like {service-name}. | [optional] 
**KeepaliveTimeout** | Pointer to **int32** | Unit is second. Indicates how long RIB has not received the client\&quot;s heartbeat before it considers the client to be dead. If not present, RIB will give a default value. | [optional] 
**RouteFilters** | [**[]RouteType**](RouteType.md) | Currently not implemented. Cor consumer client, used to filter routes as per route type when RIB publish routes to consumer. Only routes/nexthops that meet both routeFilters and networkInstanceFilters conditions will be downloaded to consumer. | 
**NetworkInstanceFilters** | Pointer to **[]string** | for consumer client, used to filter routes as per networkInstance when RIB publish routes to consumer. Only routes/nexthops that meet both routeFilters and networkInstanceFilters conditions will be downloaded to consumer. | [optional] 
**UrlCallback** | **string** | The location will be called to notify EOF information to consumer client. The location will be called to publish routes to consumer client. | 

## Methods

### NewConsumer

`func NewConsumer(routeFilters []RouteType, urlCallback string, ) *Consumer`

NewConsumer instantiates a new Consumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerWithDefaults

`func NewConsumerWithDefaults() *Consumer`

NewConsumerWithDefaults instantiates a new Consumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Consumer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Consumer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Consumer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Consumer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeepaliveTimeout

`func (o *Consumer) GetKeepaliveTimeout() int32`

GetKeepaliveTimeout returns the KeepaliveTimeout field if non-nil, zero value otherwise.

### GetKeepaliveTimeoutOk

`func (o *Consumer) GetKeepaliveTimeoutOk() (*int32, bool)`

GetKeepaliveTimeoutOk returns a tuple with the KeepaliveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveTimeout

`func (o *Consumer) SetKeepaliveTimeout(v int32)`

SetKeepaliveTimeout sets KeepaliveTimeout field to given value.

### HasKeepaliveTimeout

`func (o *Consumer) HasKeepaliveTimeout() bool`

HasKeepaliveTimeout returns a boolean if a field has been set.

### GetRouteFilters

`func (o *Consumer) GetRouteFilters() []RouteType`

GetRouteFilters returns the RouteFilters field if non-nil, zero value otherwise.

### GetRouteFiltersOk

`func (o *Consumer) GetRouteFiltersOk() (*[]RouteType, bool)`

GetRouteFiltersOk returns a tuple with the RouteFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilters

`func (o *Consumer) SetRouteFilters(v []RouteType)`

SetRouteFilters sets RouteFilters field to given value.


### GetNetworkInstanceFilters

`func (o *Consumer) GetNetworkInstanceFilters() []string`

GetNetworkInstanceFilters returns the NetworkInstanceFilters field if non-nil, zero value otherwise.

### GetNetworkInstanceFiltersOk

`func (o *Consumer) GetNetworkInstanceFiltersOk() (*[]string, bool)`

GetNetworkInstanceFiltersOk returns a tuple with the NetworkInstanceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstanceFilters

`func (o *Consumer) SetNetworkInstanceFilters(v []string)`

SetNetworkInstanceFilters sets NetworkInstanceFilters field to given value.

### HasNetworkInstanceFilters

`func (o *Consumer) HasNetworkInstanceFilters() bool`

HasNetworkInstanceFilters returns a boolean if a field has been set.

### GetUrlCallback

`func (o *Consumer) GetUrlCallback() string`

GetUrlCallback returns the UrlCallback field if non-nil, zero value otherwise.

### GetUrlCallbackOk

`func (o *Consumer) GetUrlCallbackOk() (*string, bool)`

GetUrlCallbackOk returns a tuple with the UrlCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCallback

`func (o *Consumer) SetUrlCallback(v string)`

SetUrlCallback sets UrlCallback field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


