# DeleteRouteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | [**IpPrefix**](IpPrefix.md) |  | 
**Afi** | **string** | Address family. | 
**Action** | **string** | Operation action. | 
**RouteType** | [**RouteType**](RouteType.md) |  | 
**Lsp** | [**RouteLspType**](RouteLspType.md) |  | 

## Methods

### NewDeleteRouteItem

`func NewDeleteRouteItem(prefix IpPrefix, afi string, action string, routeType RouteType, lsp RouteLspType, ) *DeleteRouteItem`

NewDeleteRouteItem instantiates a new DeleteRouteItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRouteItemWithDefaults

`func NewDeleteRouteItemWithDefaults() *DeleteRouteItem`

NewDeleteRouteItemWithDefaults instantiates a new DeleteRouteItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *DeleteRouteItem) GetPrefix() IpPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DeleteRouteItem) GetPrefixOk() (*IpPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DeleteRouteItem) SetPrefix(v IpPrefix)`

SetPrefix sets Prefix field to given value.


### GetAfi

`func (o *DeleteRouteItem) GetAfi() string`

GetAfi returns the Afi field if non-nil, zero value otherwise.

### GetAfiOk

`func (o *DeleteRouteItem) GetAfiOk() (*string, bool)`

GetAfiOk returns a tuple with the Afi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfi

`func (o *DeleteRouteItem) SetAfi(v string)`

SetAfi sets Afi field to given value.


### GetAction

`func (o *DeleteRouteItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeleteRouteItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeleteRouteItem) SetAction(v string)`

SetAction sets Action field to given value.


### GetRouteType

`func (o *DeleteRouteItem) GetRouteType() RouteType`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *DeleteRouteItem) GetRouteTypeOk() (*RouteType, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *DeleteRouteItem) SetRouteType(v RouteType)`

SetRouteType sets RouteType field to given value.


### GetLsp

`func (o *DeleteRouteItem) GetLsp() RouteLspType`

GetLsp returns the Lsp field if non-nil, zero value otherwise.

### GetLspOk

`func (o *DeleteRouteItem) GetLspOk() (*RouteLspType, bool)`

GetLspOk returns a tuple with the Lsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLsp

`func (o *DeleteRouteItem) SetLsp(v RouteLspType)`

SetLsp sets Lsp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


