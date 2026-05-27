# AddRouteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | **string** | A valid IP prefix. | 
**AddressFamily** | [**AddressFamily**](AddressFamily.md) |  | 
**Action** | **string** | Operation action. | 
**Lsp** | [**RouteLspType**](RouteLspType.md) | Route lsp type. | 
**Dist** | Pointer to **int32** | Distance. If not present, RIB will grant the deault distance of the protocol. | [optional] 
**Metric** | Pointer to **int32** | if not present, RIB will grant zero to the metric. | [optional] 
**RouteType** | [**RouteType**](RouteType.md) | Route type. the route entry must carry routeType, which will be classified by this field. | 
**Protocol** | Pointer to [**[]AddRouteItemProtocolInner**](AddRouteItemProtocolInner.md) | sub protocol type of the route. | [optional] 
**Argument** | Pointer to **string** | Optional private argument, used to carry additional routing attributes. This private argument will not be involved in route calculation. | [optional] 
**NexthopNumber** | **int32** | Nexthop number. | 
**Nexthops** | [**[]RouteNexthopInfo**](RouteNexthopInfo.md) | Nexthop information. | 

## Methods

### NewAddRouteItem

`func NewAddRouteItem(prefix string, addressFamily AddressFamily, action string, lsp RouteLspType, routeType RouteType, nexthopNumber int32, nexthops []RouteNexthopInfo, ) *AddRouteItem`

NewAddRouteItem instantiates a new AddRouteItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRouteItemWithDefaults

`func NewAddRouteItemWithDefaults() *AddRouteItem`

NewAddRouteItemWithDefaults instantiates a new AddRouteItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *AddRouteItem) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AddRouteItem) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AddRouteItem) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetAddressFamily

`func (o *AddRouteItem) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *AddRouteItem) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *AddRouteItem) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.


### GetAction

`func (o *AddRouteItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AddRouteItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AddRouteItem) SetAction(v string)`

SetAction sets Action field to given value.


### GetLsp

`func (o *AddRouteItem) GetLsp() RouteLspType`

GetLsp returns the Lsp field if non-nil, zero value otherwise.

### GetLspOk

`func (o *AddRouteItem) GetLspOk() (*RouteLspType, bool)`

GetLspOk returns a tuple with the Lsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLsp

`func (o *AddRouteItem) SetLsp(v RouteLspType)`

SetLsp sets Lsp field to given value.


### GetDist

`func (o *AddRouteItem) GetDist() int32`

GetDist returns the Dist field if non-nil, zero value otherwise.

### GetDistOk

`func (o *AddRouteItem) GetDistOk() (*int32, bool)`

GetDistOk returns a tuple with the Dist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDist

`func (o *AddRouteItem) SetDist(v int32)`

SetDist sets Dist field to given value.

### HasDist

`func (o *AddRouteItem) HasDist() bool`

HasDist returns a boolean if a field has been set.

### GetMetric

`func (o *AddRouteItem) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AddRouteItem) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AddRouteItem) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *AddRouteItem) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetRouteType

`func (o *AddRouteItem) GetRouteType() RouteType`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *AddRouteItem) GetRouteTypeOk() (*RouteType, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *AddRouteItem) SetRouteType(v RouteType)`

SetRouteType sets RouteType field to given value.


### GetProtocol

`func (o *AddRouteItem) GetProtocol() []AddRouteItemProtocolInner`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AddRouteItem) GetProtocolOk() (*[]AddRouteItemProtocolInner, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AddRouteItem) SetProtocol(v []AddRouteItemProtocolInner)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AddRouteItem) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetArgument

`func (o *AddRouteItem) GetArgument() string`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *AddRouteItem) GetArgumentOk() (*string, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *AddRouteItem) SetArgument(v string)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *AddRouteItem) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### GetNexthopNumber

`func (o *AddRouteItem) GetNexthopNumber() int32`

GetNexthopNumber returns the NexthopNumber field if non-nil, zero value otherwise.

### GetNexthopNumberOk

`func (o *AddRouteItem) GetNexthopNumberOk() (*int32, bool)`

GetNexthopNumberOk returns a tuple with the NexthopNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthopNumber

`func (o *AddRouteItem) SetNexthopNumber(v int32)`

SetNexthopNumber sets NexthopNumber field to given value.


### GetNexthops

`func (o *AddRouteItem) GetNexthops() []RouteNexthopInfo`

GetNexthops returns the Nexthops field if non-nil, zero value otherwise.

### GetNexthopsOk

`func (o *AddRouteItem) GetNexthopsOk() (*[]RouteNexthopInfo, bool)`

GetNexthopsOk returns a tuple with the Nexthops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthops

`func (o *AddRouteItem) SetNexthops(v []RouteNexthopInfo)`

SetNexthops sets Nexthops field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


