# RouteArrayItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | **string** | A valid IP prefix. | 
**AddressFamily** | [**AddressFamily**](AddressFamily.md) |  | 
**Action** | **string** | Operation action. | 
**Lsp** | [**RouteLspType**](RouteLspType.md) | Route lsp flag. | 
**Dist** | Pointer to **int32** | Distance. If not present, RIB will grant the deault distance of the protocol. | [optional] 
**Metric** | Pointer to **int32** | if not present, RIB will grant zero to the metric. | [optional] 
**RouteType** | [**RouteType**](RouteType.md) | Route type If the route entry carries routeType, it will be classified by this field. | 
**Protocol** | Pointer to [**[]AddRouteItemProtocolInner**](AddRouteItemProtocolInner.md) | sub protocol type of the route. | [optional] 
**Argument** | Pointer to **string** | Optional private argument, used to carry additional routing attributes. This private argument will not be involved in route calculation. | [optional] 
**NexthopNumber** | **int32** | Nexthop number. | 
**Nexthops** | [**[]RouteNexthopInfo**](RouteNexthopInfo.md) | Nexthop information. | 

## Methods

### NewRouteArrayItem

`func NewRouteArrayItem(prefix string, addressFamily AddressFamily, action string, lsp RouteLspType, routeType RouteType, nexthopNumber int32, nexthops []RouteNexthopInfo, ) *RouteArrayItem`

NewRouteArrayItem instantiates a new RouteArrayItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteArrayItemWithDefaults

`func NewRouteArrayItemWithDefaults() *RouteArrayItem`

NewRouteArrayItemWithDefaults instantiates a new RouteArrayItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *RouteArrayItem) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteArrayItem) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteArrayItem) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetAddressFamily

`func (o *RouteArrayItem) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *RouteArrayItem) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *RouteArrayItem) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.


### GetAction

`func (o *RouteArrayItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RouteArrayItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RouteArrayItem) SetAction(v string)`

SetAction sets Action field to given value.


### GetLsp

`func (o *RouteArrayItem) GetLsp() RouteLspType`

GetLsp returns the Lsp field if non-nil, zero value otherwise.

### GetLspOk

`func (o *RouteArrayItem) GetLspOk() (*RouteLspType, bool)`

GetLspOk returns a tuple with the Lsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLsp

`func (o *RouteArrayItem) SetLsp(v RouteLspType)`

SetLsp sets Lsp field to given value.


### GetDist

`func (o *RouteArrayItem) GetDist() int32`

GetDist returns the Dist field if non-nil, zero value otherwise.

### GetDistOk

`func (o *RouteArrayItem) GetDistOk() (*int32, bool)`

GetDistOk returns a tuple with the Dist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDist

`func (o *RouteArrayItem) SetDist(v int32)`

SetDist sets Dist field to given value.

### HasDist

`func (o *RouteArrayItem) HasDist() bool`

HasDist returns a boolean if a field has been set.

### GetMetric

`func (o *RouteArrayItem) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *RouteArrayItem) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *RouteArrayItem) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *RouteArrayItem) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetRouteType

`func (o *RouteArrayItem) GetRouteType() RouteType`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *RouteArrayItem) GetRouteTypeOk() (*RouteType, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *RouteArrayItem) SetRouteType(v RouteType)`

SetRouteType sets RouteType field to given value.


### GetProtocol

`func (o *RouteArrayItem) GetProtocol() []AddRouteItemProtocolInner`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RouteArrayItem) GetProtocolOk() (*[]AddRouteItemProtocolInner, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RouteArrayItem) SetProtocol(v []AddRouteItemProtocolInner)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RouteArrayItem) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetArgument

`func (o *RouteArrayItem) GetArgument() string`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *RouteArrayItem) GetArgumentOk() (*string, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *RouteArrayItem) SetArgument(v string)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *RouteArrayItem) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### GetNexthopNumber

`func (o *RouteArrayItem) GetNexthopNumber() int32`

GetNexthopNumber returns the NexthopNumber field if non-nil, zero value otherwise.

### GetNexthopNumberOk

`func (o *RouteArrayItem) GetNexthopNumberOk() (*int32, bool)`

GetNexthopNumberOk returns a tuple with the NexthopNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthopNumber

`func (o *RouteArrayItem) SetNexthopNumber(v int32)`

SetNexthopNumber sets NexthopNumber field to given value.


### GetNexthops

`func (o *RouteArrayItem) GetNexthops() []RouteNexthopInfo`

GetNexthops returns the Nexthops field if non-nil, zero value otherwise.

### GetNexthopsOk

`func (o *RouteArrayItem) GetNexthopsOk() (*[]RouteNexthopInfo, bool)`

GetNexthopsOk returns a tuple with the Nexthops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthops

`func (o *RouteArrayItem) SetNexthops(v []RouteNexthopInfo)`

SetNexthops sets Nexthops field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


