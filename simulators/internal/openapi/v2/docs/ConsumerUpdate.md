# ConsumerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteFilters** | [**[]RouteType**](RouteType.md) | Currently not implemented. For consumer client, used to filter routes as per route type when RIB publish routes to consumer. It shall include all route types that consumer need. Only routes/nexthops that meet both routeFilters and networkInstanceFilters conditions will be downloaded to consumer. | 
**NetworkInstanceFilters** | **[]string** | for consumer client, used to filter routes as per networkInstance when RIB publish routes to consumer. It shall include all MWs that consumer need. Only routes/nexthops that meet both routeFilters and networkInstanceFilters conditions will be downloaded to consumer. | 

## Methods

### NewConsumerUpdate

`func NewConsumerUpdate(routeFilters []RouteType, networkInstanceFilters []string, ) *ConsumerUpdate`

NewConsumerUpdate instantiates a new ConsumerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerUpdateWithDefaults

`func NewConsumerUpdateWithDefaults() *ConsumerUpdate`

NewConsumerUpdateWithDefaults instantiates a new ConsumerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteFilters

`func (o *ConsumerUpdate) GetRouteFilters() []RouteType`

GetRouteFilters returns the RouteFilters field if non-nil, zero value otherwise.

### GetRouteFiltersOk

`func (o *ConsumerUpdate) GetRouteFiltersOk() (*[]RouteType, bool)`

GetRouteFiltersOk returns a tuple with the RouteFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilters

`func (o *ConsumerUpdate) SetRouteFilters(v []RouteType)`

SetRouteFilters sets RouteFilters field to given value.


### GetNetworkInstanceFilters

`func (o *ConsumerUpdate) GetNetworkInstanceFilters() []string`

GetNetworkInstanceFilters returns the NetworkInstanceFilters field if non-nil, zero value otherwise.

### GetNetworkInstanceFiltersOk

`func (o *ConsumerUpdate) GetNetworkInstanceFiltersOk() (*[]string, bool)`

GetNetworkInstanceFiltersOk returns a tuple with the NetworkInstanceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstanceFilters

`func (o *ConsumerUpdate) SetNetworkInstanceFilters(v []string)`

SetNetworkInstanceFilters sets NetworkInstanceFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


