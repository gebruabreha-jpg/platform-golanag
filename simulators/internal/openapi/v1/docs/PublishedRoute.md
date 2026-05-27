# PublishedRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Operation action. | [optional] 
**NetworkInstance** | Pointer to **string** | Network instance. | [optional] 
**Afi** | Pointer to **string** | Address family. | [optional] 
**Prefix** | Pointer to [**IpPrefix**](IpPrefix.md) |  | [optional] 
**Type** | Pointer to [**RouteType**](RouteType.md) |  | [optional] 
**Nh** | Pointer to [**PublishedRouteNh**](PublishedRouteNh.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPublishedRoute

`func NewPublishedRoute() *PublishedRoute`

NewPublishedRoute instantiates a new PublishedRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedRouteWithDefaults

`func NewPublishedRouteWithDefaults() *PublishedRoute`

NewPublishedRouteWithDefaults instantiates a new PublishedRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PublishedRoute) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PublishedRoute) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PublishedRoute) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PublishedRoute) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *PublishedRoute) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *PublishedRoute) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *PublishedRoute) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *PublishedRoute) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetAfi

`func (o *PublishedRoute) GetAfi() string`

GetAfi returns the Afi field if non-nil, zero value otherwise.

### GetAfiOk

`func (o *PublishedRoute) GetAfiOk() (*string, bool)`

GetAfiOk returns a tuple with the Afi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfi

`func (o *PublishedRoute) SetAfi(v string)`

SetAfi sets Afi field to given value.

### HasAfi

`func (o *PublishedRoute) HasAfi() bool`

HasAfi returns a boolean if a field has been set.

### GetPrefix

`func (o *PublishedRoute) GetPrefix() IpPrefix`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PublishedRoute) GetPrefixOk() (*IpPrefix, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PublishedRoute) SetPrefix(v IpPrefix)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PublishedRoute) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetType

`func (o *PublishedRoute) GetType() RouteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublishedRoute) GetTypeOk() (*RouteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublishedRoute) SetType(v RouteType)`

SetType sets Type field to given value.

### HasType

`func (o *PublishedRoute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNh

`func (o *PublishedRoute) GetNh() PublishedRouteNh`

GetNh returns the Nh field if non-nil, zero value otherwise.

### GetNhOk

`func (o *PublishedRoute) GetNhOk() (*PublishedRouteNh, bool)`

GetNhOk returns a tuple with the Nh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNh

`func (o *PublishedRoute) SetNh(v PublishedRouteNh)`

SetNh sets Nh field to given value.

### HasNh

`func (o *PublishedRoute) HasNh() bool`

HasNh returns a boolean if a field has been set.

### GetTimestamp

`func (o *PublishedRoute) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublishedRoute) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublishedRoute) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PublishedRoute) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


