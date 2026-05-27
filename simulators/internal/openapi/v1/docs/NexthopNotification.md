# NexthopNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]NexthopNotificationItem**](NexthopNotificationItem.md) |  | [optional] 

## Methods

### NewNexthopNotification

`func NewNexthopNotification() *NexthopNotification`

NewNexthopNotification instantiates a new NexthopNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopNotificationWithDefaults

`func NewNexthopNotificationWithDefaults() *NexthopNotification`

NewNexthopNotificationWithDefaults instantiates a new NexthopNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NexthopNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NexthopNotification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NexthopNotification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NexthopNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItems

`func (o *NexthopNotification) GetItems() []NexthopNotificationItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NexthopNotification) GetItemsOk() (*[]NexthopNotificationItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NexthopNotification) SetItems(v []NexthopNotificationItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *NexthopNotification) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


