# RouteArrayItem

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

### NewRouteArrayItem

`func NewRouteArrayItem(prefix IpPrefix, afi string, action string, lsp RouteLspType, routeType RouteType, nhNum int32, nh []NexthopInformation, ) *RouteArrayItem`

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

`func (o *RouteArrayItem) GetPrefix() IpPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteArrayItem) GetPrefixOk() (*IpPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteArrayItem) SetPrefix(v IpPrefix)`

SetPrefix sets Prefix field to given value.


### GetAfi

`func (o *RouteArrayItem) GetAfi() string`

GetAfi returns the Afi field if non-nil, zero value otherwise.

### GetAfiOk

`func (o *RouteArrayItem) GetAfiOk() (*string, bool)`

GetAfiOk returns a tuple with the Afi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfi

`func (o *RouteArrayItem) SetAfi(v string)`

SetAfi sets Afi field to given value.


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

`func (o *RouteArrayItem) GetProtocol() []string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RouteArrayItem) GetProtocolOk() (*[]string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RouteArrayItem) SetProtocol(v []string)`

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

### GetNhNum

`func (o *RouteArrayItem) GetNhNum() int32`

GetNhNum returns the NhNum field if non-nil, zero value otherwise.

### GetNhNumOk

`func (o *RouteArrayItem) GetNhNumOk() (*int32, bool)`

GetNhNumOk returns a tuple with the NhNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNhNum

`func (o *RouteArrayItem) SetNhNum(v int32)`

SetNhNum sets NhNum field to given value.


### GetNh

`func (o *RouteArrayItem) GetNh() []NexthopInformation`

GetNh returns the Nh field if non-nil, zero value otherwise.

### GetNhOk

`func (o *RouteArrayItem) GetNhOk() (*[]NexthopInformation, bool)`

GetNhOk returns a tuple with the Nh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNh

`func (o *RouteArrayItem) SetNh(v []NexthopInformation)`

SetNh sets Nh field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


