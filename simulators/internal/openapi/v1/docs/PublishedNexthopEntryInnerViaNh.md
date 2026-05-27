# PublishedNexthopEntryInnerViaNh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInstance** | Pointer to **string** | If \&quot;via_nh\&quot; exists, \&quot;network_instance\&quot; is required. The network_instance which via nexthop belongs to. | [optional] 
**Id** | Pointer to **int32** | Nexthop id, identify a nexthop. Optional, if \&quot;id\&quot; is empty(none), it means that the current nexthop is the final nexthop. | [optional] 

## Methods

### NewPublishedNexthopEntryInnerViaNh

`func NewPublishedNexthopEntryInnerViaNh() *PublishedNexthopEntryInnerViaNh`

NewPublishedNexthopEntryInnerViaNh instantiates a new PublishedNexthopEntryInnerViaNh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedNexthopEntryInnerViaNhWithDefaults

`func NewPublishedNexthopEntryInnerViaNhWithDefaults() *PublishedNexthopEntryInnerViaNh`

NewPublishedNexthopEntryInnerViaNhWithDefaults instantiates a new PublishedNexthopEntryInnerViaNh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInstance

`func (o *PublishedNexthopEntryInnerViaNh) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *PublishedNexthopEntryInnerViaNh) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *PublishedNexthopEntryInnerViaNh) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *PublishedNexthopEntryInnerViaNh) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetId

`func (o *PublishedNexthopEntryInnerViaNh) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublishedNexthopEntryInnerViaNh) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublishedNexthopEntryInnerViaNh) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PublishedNexthopEntryInnerViaNh) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


