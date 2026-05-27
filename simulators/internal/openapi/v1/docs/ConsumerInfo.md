# ConsumerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepaliveTimeout** | **int32** | Unit is second. Indicates how long RIB has not received the client&#39;s heartbeat before it considers the client to be dead. If not present, RIB will give a default value. | 
**RouteFilters** | Pointer to **[]string** | for consumer client, used to filter routes as per route type when RIB publish routes to consumer. Only routes/nexthops that meet both route_filters and network_instance_filters conditions will be downloaded to consumer. | [optional] 
**NetworkInstanceFilters** | Pointer to **[]string** | for consumer client, used to filter routes as per network-instance when RIB publish routes to consumer. Only routes/nexthops that meet both route_filters and network_instance_filters conditions will be downloaded to consumer. | [optional] 
**UrlCallback** | **string** | The location will be called to notify EOF information to consumer client. The location will be called to publish routes to consumer client. | 

## Methods

### NewConsumerInfo

`func NewConsumerInfo(keepaliveTimeout int32, urlCallback string, ) *ConsumerInfo`

NewConsumerInfo instantiates a new ConsumerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerInfoWithDefaults

`func NewConsumerInfoWithDefaults() *ConsumerInfo`

NewConsumerInfoWithDefaults instantiates a new ConsumerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeepaliveTimeout

`func (o *ConsumerInfo) GetKeepaliveTimeout() int32`

GetKeepaliveTimeout returns the KeepaliveTimeout field if non-nil, zero value otherwise.

### GetKeepaliveTimeoutOk

`func (o *ConsumerInfo) GetKeepaliveTimeoutOk() (*int32, bool)`

GetKeepaliveTimeoutOk returns a tuple with the KeepaliveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveTimeout

`func (o *ConsumerInfo) SetKeepaliveTimeout(v int32)`

SetKeepaliveTimeout sets KeepaliveTimeout field to given value.


### GetRouteFilters

`func (o *ConsumerInfo) GetRouteFilters() []string`

GetRouteFilters returns the RouteFilters field if non-nil, zero value otherwise.

### GetRouteFiltersOk

`func (o *ConsumerInfo) GetRouteFiltersOk() (*[]string, bool)`

GetRouteFiltersOk returns a tuple with the RouteFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilters

`func (o *ConsumerInfo) SetRouteFilters(v []string)`

SetRouteFilters sets RouteFilters field to given value.

### HasRouteFilters

`func (o *ConsumerInfo) HasRouteFilters() bool`

HasRouteFilters returns a boolean if a field has been set.

### GetNetworkInstanceFilters

`func (o *ConsumerInfo) GetNetworkInstanceFilters() []string`

GetNetworkInstanceFilters returns the NetworkInstanceFilters field if non-nil, zero value otherwise.

### GetNetworkInstanceFiltersOk

`func (o *ConsumerInfo) GetNetworkInstanceFiltersOk() (*[]string, bool)`

GetNetworkInstanceFiltersOk returns a tuple with the NetworkInstanceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstanceFilters

`func (o *ConsumerInfo) SetNetworkInstanceFilters(v []string)`

SetNetworkInstanceFilters sets NetworkInstanceFilters field to given value.

### HasNetworkInstanceFilters

`func (o *ConsumerInfo) HasNetworkInstanceFilters() bool`

HasNetworkInstanceFilters returns a boolean if a field has been set.

### GetUrlCallback

`func (o *ConsumerInfo) GetUrlCallback() string`

GetUrlCallback returns the UrlCallback field if non-nil, zero value otherwise.

### GetUrlCallbackOk

`func (o *ConsumerInfo) GetUrlCallbackOk() (*string, bool)`

GetUrlCallbackOk returns a tuple with the UrlCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCallback

`func (o *ConsumerInfo) SetUrlCallback(v string)`

SetUrlCallback sets UrlCallback field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


