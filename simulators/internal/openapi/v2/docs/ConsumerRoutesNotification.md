# ConsumerRoutesNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Routes** | Pointer to [**[]ConsumerRoutesNotificationItem**](ConsumerRoutesNotificationItem.md) |  | [optional] 

## Methods

### NewConsumerRoutesNotification

`func NewConsumerRoutesNotification() *ConsumerRoutesNotification`

NewConsumerRoutesNotification instantiates a new ConsumerRoutesNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerRoutesNotificationWithDefaults

`func NewConsumerRoutesNotificationWithDefaults() *ConsumerRoutesNotification`

NewConsumerRoutesNotificationWithDefaults instantiates a new ConsumerRoutesNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConsumerRoutesNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsumerRoutesNotification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsumerRoutesNotification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConsumerRoutesNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoutes

`func (o *ConsumerRoutesNotification) GetRoutes() []ConsumerRoutesNotificationItem`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ConsumerRoutesNotification) GetRoutesOk() (*[]ConsumerRoutesNotificationItem, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ConsumerRoutesNotification) SetRoutes(v []ConsumerRoutesNotificationItem)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *ConsumerRoutesNotification) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


