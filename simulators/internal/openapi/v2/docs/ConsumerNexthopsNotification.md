# ConsumerNexthopsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Nexthops** | Pointer to [**[]ConsumerNexthopsNotificationItem**](ConsumerNexthopsNotificationItem.md) |  | [optional] 

## Methods

### NewConsumerNexthopsNotification

`func NewConsumerNexthopsNotification() *ConsumerNexthopsNotification`

NewConsumerNexthopsNotification instantiates a new ConsumerNexthopsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerNexthopsNotificationWithDefaults

`func NewConsumerNexthopsNotificationWithDefaults() *ConsumerNexthopsNotification`

NewConsumerNexthopsNotificationWithDefaults instantiates a new ConsumerNexthopsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConsumerNexthopsNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsumerNexthopsNotification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsumerNexthopsNotification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConsumerNexthopsNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNexthops

`func (o *ConsumerNexthopsNotification) GetNexthops() []ConsumerNexthopsNotificationItem`

GetNexthops returns the Nexthops field if non-nil, zero value otherwise.

### GetNexthopsOk

`func (o *ConsumerNexthopsNotification) GetNexthopsOk() (*[]ConsumerNexthopsNotificationItem, bool)`

GetNexthopsOk returns a tuple with the Nexthops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthops

`func (o *ConsumerNexthopsNotification) SetNexthops(v []ConsumerNexthopsNotificationItem)`

SetNexthops sets Nexthops field to given value.

### HasNexthops

`func (o *ConsumerNexthopsNotification) HasNexthops() bool`

HasNexthops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


