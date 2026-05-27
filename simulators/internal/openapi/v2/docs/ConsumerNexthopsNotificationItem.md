# ConsumerNexthopsNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ConsumerRoutesNotificationItemAction**](ConsumerRoutesNotificationItemAction.md) |  | [optional] 
**Id** | Pointer to **int32** | Nexthop id, identify a nexthop. | [optional] 
**Type** | Pointer to [**ConsumerNexthopsNotificationItemType**](ConsumerNexthopsNotificationItemType.md) |  | [optional] 
**NetworkInstance** | Pointer to **string** | The network instance which nexthop belongs to. | [optional] 
**AddressFamily** | Pointer to [**AddressFamily**](AddressFamily.md) |  | [optional] 
**Entry** | Pointer to [**[]ConsumerNexthopsNotificationItemEntryInner**](ConsumerNexthopsNotificationItemEntryInner.md) | Nexthop entry, used to indicate the address information of the nexthop or whether there is a via nexthop. If above \&quot;action\&quot; is \&quot;add\&quot;, the item is required. | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewConsumerNexthopsNotificationItem

`func NewConsumerNexthopsNotificationItem() *ConsumerNexthopsNotificationItem`

NewConsumerNexthopsNotificationItem instantiates a new ConsumerNexthopsNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerNexthopsNotificationItemWithDefaults

`func NewConsumerNexthopsNotificationItemWithDefaults() *ConsumerNexthopsNotificationItem`

NewConsumerNexthopsNotificationItemWithDefaults instantiates a new ConsumerNexthopsNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ConsumerNexthopsNotificationItem) GetAction() ConsumerRoutesNotificationItemAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ConsumerNexthopsNotificationItem) GetActionOk() (*ConsumerRoutesNotificationItemAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ConsumerNexthopsNotificationItem) SetAction(v ConsumerRoutesNotificationItemAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ConsumerNexthopsNotificationItem) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetId

`func (o *ConsumerNexthopsNotificationItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumerNexthopsNotificationItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumerNexthopsNotificationItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConsumerNexthopsNotificationItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ConsumerNexthopsNotificationItem) GetType() ConsumerNexthopsNotificationItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsumerNexthopsNotificationItem) GetTypeOk() (*ConsumerNexthopsNotificationItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsumerNexthopsNotificationItem) SetType(v ConsumerNexthopsNotificationItemType)`

SetType sets Type field to given value.

### HasType

`func (o *ConsumerNexthopsNotificationItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *ConsumerNexthopsNotificationItem) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *ConsumerNexthopsNotificationItem) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *ConsumerNexthopsNotificationItem) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *ConsumerNexthopsNotificationItem) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetAddressFamily

`func (o *ConsumerNexthopsNotificationItem) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *ConsumerNexthopsNotificationItem) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *ConsumerNexthopsNotificationItem) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *ConsumerNexthopsNotificationItem) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetEntry

`func (o *ConsumerNexthopsNotificationItem) GetEntry() []ConsumerNexthopsNotificationItemEntryInner`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *ConsumerNexthopsNotificationItem) GetEntryOk() (*[]ConsumerNexthopsNotificationItemEntryInner, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *ConsumerNexthopsNotificationItem) SetEntry(v []ConsumerNexthopsNotificationItemEntryInner)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *ConsumerNexthopsNotificationItem) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetTimestamp

`func (o *ConsumerNexthopsNotificationItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ConsumerNexthopsNotificationItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ConsumerNexthopsNotificationItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ConsumerNexthopsNotificationItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


