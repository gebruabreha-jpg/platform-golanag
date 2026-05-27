# ConsumerUpdateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteFilters** | Pointer to [**[]RouteType**](RouteType.md) | for consumer client, used to filter routes as per route type when RIB publish routes to consumer. It shall include all route types that consumer need. Only routes/nexthops that meet both route_filters and network_instance_filters conditions will be downloaded to consumer. | [optional] 
**NetworkInstanceFilters** | Pointer to **[]string** | for consumer client, used to filter routes as per network-instance when RIB publish routes to consumer. It shall include all MWs that consumer need. Only routes/nexthops that meet both route_filters and network_instance_filters conditions will be downloaded to consumer. | [optional] 

## Methods

### NewConsumerUpdateInfo

`func NewConsumerUpdateInfo() *ConsumerUpdateInfo`

NewConsumerUpdateInfo instantiates a new ConsumerUpdateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerUpdateInfoWithDefaults

`func NewConsumerUpdateInfoWithDefaults() *ConsumerUpdateInfo`

NewConsumerUpdateInfoWithDefaults instantiates a new ConsumerUpdateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteFilters

`func (o *ConsumerUpdateInfo) GetRouteFilters() []RouteType`

GetRouteFilters returns the RouteFilters field if non-nil, zero value otherwise.

### GetRouteFiltersOk

`func (o *ConsumerUpdateInfo) GetRouteFiltersOk() (*[]RouteType, bool)`

GetRouteFiltersOk returns a tuple with the RouteFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilters

`func (o *ConsumerUpdateInfo) SetRouteFilters(v []RouteType)`

SetRouteFilters sets RouteFilters field to given value.

### HasRouteFilters

`func (o *ConsumerUpdateInfo) HasRouteFilters() bool`

HasRouteFilters returns a boolean if a field has been set.

### GetNetworkInstanceFilters

`func (o *ConsumerUpdateInfo) GetNetworkInstanceFilters() []string`

GetNetworkInstanceFilters returns the NetworkInstanceFilters field if non-nil, zero value otherwise.

### GetNetworkInstanceFiltersOk

`func (o *ConsumerUpdateInfo) GetNetworkInstanceFiltersOk() (*[]string, bool)`

GetNetworkInstanceFiltersOk returns a tuple with the NetworkInstanceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstanceFilters

`func (o *ConsumerUpdateInfo) SetNetworkInstanceFilters(v []string)`

SetNetworkInstanceFilters sets NetworkInstanceFilters field to given value.

### HasNetworkInstanceFilters

`func (o *ConsumerUpdateInfo) HasNetworkInstanceFilters() bool`

HasNetworkInstanceFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


