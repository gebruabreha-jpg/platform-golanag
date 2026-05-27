# PublishedNexthop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Operation action. | [optional] 
**Id** | Pointer to **int32** | Nexthop id, identify a nexthop. | [optional] 
**Type** | Pointer to **string** | Nexthop type. | [optional] 
**Argument** | Pointer to **string** | Optional private argument, used to carry additional routing attributes to consumer. This private argument will not be involved in route calculation. | [optional] 
**NetworkInstance** | Pointer to **string** | The network instance which nexthop belongs to. | [optional] 
**Afi** | Pointer to **string** | Address family. | [optional] 
**Entry** | Pointer to [**[]PublishedNexthopEntryInner**](PublishedNexthopEntryInner.md) | Nexthop entry, used to indicate the address information of the nexthop or whether there is a via nexthop. If above &#39;action&#39; is \&quot;add\&quot;, the item is required. | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPublishedNexthop

`func NewPublishedNexthop() *PublishedNexthop`

NewPublishedNexthop instantiates a new PublishedNexthop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedNexthopWithDefaults

`func NewPublishedNexthopWithDefaults() *PublishedNexthop`

NewPublishedNexthopWithDefaults instantiates a new PublishedNexthop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PublishedNexthop) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PublishedNexthop) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PublishedNexthop) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PublishedNexthop) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetId

`func (o *PublishedNexthop) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublishedNexthop) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublishedNexthop) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PublishedNexthop) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PublishedNexthop) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublishedNexthop) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublishedNexthop) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PublishedNexthop) HasType() bool`

HasType returns a boolean if a field has been set.

### GetArgument

`func (o *PublishedNexthop) GetArgument() string`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *PublishedNexthop) GetArgumentOk() (*string, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *PublishedNexthop) SetArgument(v string)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *PublishedNexthop) HasArgument() bool`

HasArgument returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *PublishedNexthop) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *PublishedNexthop) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *PublishedNexthop) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *PublishedNexthop) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetAfi

`func (o *PublishedNexthop) GetAfi() string`

GetAfi returns the Afi field if non-nil, zero value otherwise.

### GetAfiOk

`func (o *PublishedNexthop) GetAfiOk() (*string, bool)`

GetAfiOk returns a tuple with the Afi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfi

`func (o *PublishedNexthop) SetAfi(v string)`

SetAfi sets Afi field to given value.

### HasAfi

`func (o *PublishedNexthop) HasAfi() bool`

HasAfi returns a boolean if a field has been set.

### GetEntry

`func (o *PublishedNexthop) GetEntry() []PublishedNexthopEntryInner`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *PublishedNexthop) GetEntryOk() (*[]PublishedNexthopEntryInner, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *PublishedNexthop) SetEntry(v []PublishedNexthopEntryInner)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *PublishedNexthop) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetTimestamp

`func (o *PublishedNexthop) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublishedNexthop) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublishedNexthop) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PublishedNexthop) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


