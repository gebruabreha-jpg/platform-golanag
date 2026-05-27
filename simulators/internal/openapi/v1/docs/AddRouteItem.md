# AddRouteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | [**IpPrefix**](IpPrefix.md) |  | 
**Afi** | **string** | Address family. | 
**Action** | **string** | Operation action. | 
**Lsp** | [**RouteLspType**](RouteLspType.md) |  | 
**Dist** | Pointer to **int32** | Distance; if not present, RIB will grant the deault distance of the protocol. | [optional] 
**Metric** | Pointer to **int32** | if not present, RIB will grant zero to the metric. | [optional] 
**RouteType** | [**RouteType**](RouteType.md) |  | 
**Protocol** | Pointer to **[]string** | sub protocol type of the route. | [optional] 
**Argument** | Pointer to **string** | Optional private argument, used to carry additional routing attributes. This private argument will not be involved in route calculation. | [optional] 
**NhNum** | **int32** | Nexthop number. | 
**Nh** | [**[]NexthopInformation**](NexthopInformation.md) | Nexthop information. | 

## Methods

### NewAddRouteItem

`func NewAddRouteItem(prefix IpPrefix, afi string, action string, lsp RouteLspType, routeType RouteType, nhNum int32, nh []NexthopInformation, ) *AddRouteItem`

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

`func (o *AddRouteItem) GetPrefix() IpPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AddRouteItem) GetPrefixOk() (*IpPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AddRouteItem) SetPrefix(v IpPrefix)`

SetPrefix sets Prefix field to given value.


### GetAfi

`func (o *AddRouteItem) GetAfi() string`

GetAfi returns the Afi field if non-nil, zero value otherwise.

### GetAfiOk

`func (o *AddRouteItem) GetAfiOk() (*string, bool)`

GetAfiOk returns a tuple with the Afi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfi

`func (o *AddRouteItem) SetAfi(v string)`

SetAfi sets Afi field to given value.


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

`func (o *AddRouteItem) GetProtocol() []string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AddRouteItem) GetProtocolOk() (*[]string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AddRouteItem) SetProtocol(v []string)`

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

### GetNhNum

`func (o *AddRouteItem) GetNhNum() int32`

GetNhNum returns the NhNum field if non-nil, zero value otherwise.

### GetNhNumOk

`func (o *AddRouteItem) GetNhNumOk() (*int32, bool)`

GetNhNumOk returns a tuple with the NhNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNhNum

`func (o *AddRouteItem) SetNhNum(v int32)`

SetNhNum sets NhNum field to given value.


### GetNh

`func (o *AddRouteItem) GetNh() []NexthopInformation`

GetNh returns the Nh field if non-nil, zero value otherwise.

### GetNhOk

`func (o *AddRouteItem) GetNhOk() (*[]NexthopInformation, bool)`

GetNhOk returns a tuple with the Nh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNh

`func (o *AddRouteItem) SetNh(v []NexthopInformation)`

SetNh sets Nh field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


