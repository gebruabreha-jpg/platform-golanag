# ConsumerNotificationMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Routes** | Pointer to [**[]ConsumerRoutesNotificationItem**](ConsumerRoutesNotificationItem.md) |  | [optional] 
**Nexthops** | Pointer to [**[]ConsumerNexthopsNotificationItem**](ConsumerNexthopsNotificationItem.md) |  | [optional] 

## Methods

### NewConsumerNotificationMsg

`func NewConsumerNotificationMsg() *ConsumerNotificationMsg`

NewConsumerNotificationMsg instantiates a new ConsumerNotificationMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerNotificationMsgWithDefaults

`func NewConsumerNotificationMsgWithDefaults() *ConsumerNotificationMsg`

NewConsumerNotificationMsgWithDefaults instantiates a new ConsumerNotificationMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConsumerNotificationMsg) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsumerNotificationMsg) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsumerNotificationMsg) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConsumerNotificationMsg) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoutes

`func (o *ConsumerNotificationMsg) GetRoutes() []ConsumerRoutesNotificationItem`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ConsumerNotificationMsg) GetRoutesOk() (*[]ConsumerRoutesNotificationItem, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ConsumerNotificationMsg) SetRoutes(v []ConsumerRoutesNotificationItem)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *ConsumerNotificationMsg) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetNexthops

`func (o *ConsumerNotificationMsg) GetNexthops() []ConsumerNexthopsNotificationItem`

GetNexthops returns the Nexthops field if non-nil, zero value otherwise.

### GetNexthopsOk

`func (o *ConsumerNotificationMsg) GetNexthopsOk() (*[]ConsumerNexthopsNotificationItem, bool)`

GetNexthopsOk returns a tuple with the Nexthops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthops

`func (o *ConsumerNotificationMsg) SetNexthops(v []ConsumerNexthopsNotificationItem)`

SetNexthops sets Nexthops field to given value.

### HasNexthops

`func (o *ConsumerNotificationMsg) HasNexthops() bool`

HasNexthops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


