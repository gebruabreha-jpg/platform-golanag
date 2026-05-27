# ConsumerRoutesNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ConsumerRoutesNotificationItemAction**](ConsumerRoutesNotificationItemAction.md) |  | [optional] 
**NetworkInstance** | Pointer to **string** | Network instance. | [optional] 
**AddressFamily** | Pointer to [**AddressFamily**](AddressFamily.md) |  | [optional] 
**Prefix** | Pointer to **string** | A valid IP prefix. | [optional] 
**Type** | Pointer to [**RouteType**](RouteType.md) | Route type. | [optional] 
**Nexthop** | Pointer to [**ConsumerRoutesNotificationItemNexthop**](ConsumerRoutesNotificationItemNexthop.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Argument** | Pointer to **string** | Optional private argument, used to carry additional routing attributes to consumer. This private argument will not be involved in route calculation. | [optional] 

## Methods

### NewConsumerRoutesNotificationItem

`func NewConsumerRoutesNotificationItem() *ConsumerRoutesNotificationItem`

NewConsumerRoutesNotificationItem instantiates a new ConsumerRoutesNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerRoutesNotificationItemWithDefaults

`func NewConsumerRoutesNotificationItemWithDefaults() *ConsumerRoutesNotificationItem`

NewConsumerRoutesNotificationItemWithDefaults instantiates a new ConsumerRoutesNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ConsumerRoutesNotificationItem) GetAction() ConsumerRoutesNotificationItemAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ConsumerRoutesNotificationItem) GetActionOk() (*ConsumerRoutesNotificationItemAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ConsumerRoutesNotificationItem) SetAction(v ConsumerRoutesNotificationItemAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ConsumerRoutesNotificationItem) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *ConsumerRoutesNotificationItem) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *ConsumerRoutesNotificationItem) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *ConsumerRoutesNotificationItem) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *ConsumerRoutesNotificationItem) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetAddressFamily

`func (o *ConsumerRoutesNotificationItem) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *ConsumerRoutesNotificationItem) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *ConsumerRoutesNotificationItem) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *ConsumerRoutesNotificationItem) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetPrefix

`func (o *ConsumerRoutesNotificationItem) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ConsumerRoutesNotificationItem) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ConsumerRoutesNotificationItem) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ConsumerRoutesNotificationItem) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetType

`func (o *ConsumerRoutesNotificationItem) GetType() RouteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsumerRoutesNotificationItem) GetTypeOk() (*RouteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsumerRoutesNotificationItem) SetType(v RouteType)`

SetType sets Type field to given value.

### HasType

`func (o *ConsumerRoutesNotificationItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNexthop

`func (o *ConsumerRoutesNotificationItem) GetNexthop() ConsumerRoutesNotificationItemNexthop`

GetNexthop returns the Nexthop field if non-nil, zero value otherwise.

### GetNexthopOk

`func (o *ConsumerRoutesNotificationItem) GetNexthopOk() (*ConsumerRoutesNotificationItemNexthop, bool)`

GetNexthopOk returns a tuple with the Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthop

`func (o *ConsumerRoutesNotificationItem) SetNexthop(v ConsumerRoutesNotificationItemNexthop)`

SetNexthop sets Nexthop field to given value.

### HasNexthop

`func (o *ConsumerRoutesNotificationItem) HasNexthop() bool`

HasNexthop returns a boolean if a field has been set.

### GetTimestamp

`func (o *ConsumerRoutesNotificationItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ConsumerRoutesNotificationItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ConsumerRoutesNotificationItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ConsumerRoutesNotificationItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetArgument

`func (o *ConsumerRoutesNotificationItem) GetArgument() string`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *ConsumerRoutesNotificationItem) GetArgumentOk() (*string, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *ConsumerRoutesNotificationItem) SetArgument(v string)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *ConsumerRoutesNotificationItem) HasArgument() bool`

HasArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


