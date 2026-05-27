# QueryPrefixNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]QueryPrefixNotificationItem**](QueryPrefixNotificationItem.md) |  | [optional] 

## Methods

### NewQueryPrefixNotification

`func NewQueryPrefixNotification() *QueryPrefixNotification`

NewQueryPrefixNotification instantiates a new QueryPrefixNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPrefixNotificationWithDefaults

`func NewQueryPrefixNotificationWithDefaults() *QueryPrefixNotification`

NewQueryPrefixNotificationWithDefaults instantiates a new QueryPrefixNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QueryPrefixNotification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryPrefixNotification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryPrefixNotification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryPrefixNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItems

`func (o *QueryPrefixNotification) GetItems() []QueryPrefixNotificationItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *QueryPrefixNotification) GetItemsOk() (*[]QueryPrefixNotificationItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *QueryPrefixNotification) SetItems(v []QueryPrefixNotificationItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *QueryPrefixNotification) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


