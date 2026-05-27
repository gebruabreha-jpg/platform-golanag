# PublishedNexthopEntryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViaNh** | Pointer to [**PublishedNexthopEntryInnerViaNh**](PublishedNexthopEntryInnerViaNh.md) |  | [optional] 
**ViaNexthopId** | Pointer to [**PublishedNexthopEntryInnerViaNexthopId**](PublishedNexthopEntryInnerViaNexthopId.md) |  | [optional] 
**NhAddr** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**Label** | Pointer to **int32** | For vpn and mpls label forwarding. | [optional] 
**NetworkInstance** | Pointer to **string** | The networkInstance which nh_addr belongs to. | [optional] 
**IfType** | Pointer to [**IfType**](IfType.md) |  | [optional] 

## Methods

### NewPublishedNexthopEntryInner

`func NewPublishedNexthopEntryInner() *PublishedNexthopEntryInner`

NewPublishedNexthopEntryInner instantiates a new PublishedNexthopEntryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedNexthopEntryInnerWithDefaults

`func NewPublishedNexthopEntryInnerWithDefaults() *PublishedNexthopEntryInner`

NewPublishedNexthopEntryInnerWithDefaults instantiates a new PublishedNexthopEntryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViaNh

`func (o *PublishedNexthopEntryInner) GetViaNh() PublishedNexthopEntryInnerViaNh`

GetViaNh returns the ViaNh field if non-nil, zero value otherwise.

### GetViaNhOk

`func (o *PublishedNexthopEntryInner) GetViaNhOk() (*PublishedNexthopEntryInnerViaNh, bool)`

GetViaNhOk returns a tuple with the ViaNh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaNh

`func (o *PublishedNexthopEntryInner) SetViaNh(v PublishedNexthopEntryInnerViaNh)`

SetViaNh sets ViaNh field to given value.

### HasViaNh

`func (o *PublishedNexthopEntryInner) HasViaNh() bool`

HasViaNh returns a boolean if a field has been set.

### GetViaNexthopId

`func (o *PublishedNexthopEntryInner) GetViaNexthopId() PublishedNexthopEntryInnerViaNexthopId`

GetViaNexthopId returns the ViaNexthopId field if non-nil, zero value otherwise.

### GetViaNexthopIdOk

`func (o *PublishedNexthopEntryInner) GetViaNexthopIdOk() (*PublishedNexthopEntryInnerViaNexthopId, bool)`

GetViaNexthopIdOk returns a tuple with the ViaNexthopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaNexthopId

`func (o *PublishedNexthopEntryInner) SetViaNexthopId(v PublishedNexthopEntryInnerViaNexthopId)`

SetViaNexthopId sets ViaNexthopId field to given value.

### HasViaNexthopId

`func (o *PublishedNexthopEntryInner) HasViaNexthopId() bool`

HasViaNexthopId returns a boolean if a field has been set.

### GetNhAddr

`func (o *PublishedNexthopEntryInner) GetNhAddr() IpAddress`

GetNhAddr returns the NhAddr field if non-nil, zero value otherwise.

### GetNhAddrOk

`func (o *PublishedNexthopEntryInner) GetNhAddrOk() (*IpAddress, bool)`

GetNhAddrOk returns a tuple with the NhAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNhAddr

`func (o *PublishedNexthopEntryInner) SetNhAddr(v IpAddress)`

SetNhAddr sets NhAddr field to given value.

### HasNhAddr

`func (o *PublishedNexthopEntryInner) HasNhAddr() bool`

HasNhAddr returns a boolean if a field has been set.

### GetLabel

`func (o *PublishedNexthopEntryInner) GetLabel() int32`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PublishedNexthopEntryInner) GetLabelOk() (*int32, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PublishedNexthopEntryInner) SetLabel(v int32)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PublishedNexthopEntryInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *PublishedNexthopEntryInner) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *PublishedNexthopEntryInner) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *PublishedNexthopEntryInner) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *PublishedNexthopEntryInner) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetIfType

`func (o *PublishedNexthopEntryInner) GetIfType() IfType`

GetIfType returns the IfType field if non-nil, zero value otherwise.

### GetIfTypeOk

`func (o *PublishedNexthopEntryInner) GetIfTypeOk() (*IfType, bool)`

GetIfTypeOk returns a tuple with the IfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfType

`func (o *PublishedNexthopEntryInner) SetIfType(v IfType)`

SetIfType sets IfType field to given value.

### HasIfType

`func (o *PublishedNexthopEntryInner) HasIfType() bool`

HasIfType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


