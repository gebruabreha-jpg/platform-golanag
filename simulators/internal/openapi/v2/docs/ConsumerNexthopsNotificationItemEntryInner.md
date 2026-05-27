# ConsumerNexthopsNotificationItemEntryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViaNexthop** | Pointer to [**ConsumerNexthopsNotificationItemEntryInnerViaNexthop**](ConsumerNexthopsNotificationItemEntryInnerViaNexthop.md) |  | [optional] 
**AddressInfo** | Pointer to [**NexthopAddressInfo**](NexthopAddressInfo.md) | Nexthop address for consumer. If above \&quot;viaNexthop\&quot; id is present, nexthopAddress can be ignored during forwarding. Only meaningful when adding operations. | [optional] 
**Label** | Pointer to **int32** | For vpn and mpls label forwarding. | [optional] 
**InterfaceType** | Pointer to [**InterfaceType**](InterfaceType.md) |  | [optional] 
**Priority** | Pointer to **int32** | Nexthop priority. The higher the value, the lower the priority | [optional] 

## Methods

### NewConsumerNexthopsNotificationItemEntryInner

`func NewConsumerNexthopsNotificationItemEntryInner() *ConsumerNexthopsNotificationItemEntryInner`

NewConsumerNexthopsNotificationItemEntryInner instantiates a new ConsumerNexthopsNotificationItemEntryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerNexthopsNotificationItemEntryInnerWithDefaults

`func NewConsumerNexthopsNotificationItemEntryInnerWithDefaults() *ConsumerNexthopsNotificationItemEntryInner`

NewConsumerNexthopsNotificationItemEntryInnerWithDefaults instantiates a new ConsumerNexthopsNotificationItemEntryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViaNexthop

`func (o *ConsumerNexthopsNotificationItemEntryInner) GetViaNexthop() ConsumerNexthopsNotificationItemEntryInnerViaNexthop`

GetViaNexthop returns the ViaNexthop field if non-nil, zero value otherwise.

### GetViaNexthopOk

`func (o *ConsumerNexthopsNotificationItemEntryInner) GetViaNexthopOk() (*ConsumerNexthopsNotificationItemEntryInnerViaNexthop, bool)`

GetViaNexthopOk returns a tuple with the ViaNexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaNexthop

`func (o *ConsumerNexthopsNotificationItemEntryInner) SetViaNexthop(v ConsumerNexthopsNotificationItemEntryInnerViaNexthop)`

SetViaNexthop sets ViaNexthop field to given value.

### HasViaNexthop

`func (o *ConsumerNexthopsNotificationItemEntryInner) HasViaNexthop() bool`

HasViaNexthop returns a boolean if a field has been set.

### GetAddressInfo

`func (o *ConsumerNexthopsNotificationItemEntryInner) GetAddressInfo() NexthopAddressInfo`

GetAddressInfo returns the AddressInfo field if non-nil, zero value otherwise.

### GetAddressInfoOk

`func (o *ConsumerNexthopsNotificationItemEntryInner) GetAddressInfoOk() (*NexthopAddressInfo, bool)`

GetAddressInfoOk returns a tuple with the AddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInfo

`func (o *ConsumerNexthopsNotificationItemEntryInner) SetAddressInfo(v NexthopAddressInfo)`

SetAddressInfo sets AddressInfo field to given value.

### HasAddressInfo

`func (o *ConsumerNexthopsNotificationItemEntryInner) HasAddressInfo() bool`

HasAddressInfo returns a boolean if a field has been set.

### GetLabel

`func (o *ConsumerNexthopsNotificationItemEntryInner) GetLabel() int32`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConsumerNexthopsNotificationItemEntryInner) GetLabelOk() (*int32, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConsumerNexthopsNotificationItemEntryInner) SetLabel(v int32)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConsumerNexthopsNotificationItemEntryInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetInterfaceType

`func (o *ConsumerNexthopsNotificationItemEntryInner) GetInterfaceType() InterfaceType`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *ConsumerNexthopsNotificationItemEntryInner) GetInterfaceTypeOk() (*InterfaceType, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *ConsumerNexthopsNotificationItemEntryInner) SetInterfaceType(v InterfaceType)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *ConsumerNexthopsNotificationItemEntryInner) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetPriority

`func (o *ConsumerNexthopsNotificationItemEntryInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ConsumerNexthopsNotificationItemEntryInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ConsumerNexthopsNotificationItemEntryInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ConsumerNexthopsNotificationItemEntryInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


