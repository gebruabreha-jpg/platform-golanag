# ProducerNotificationMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]QueryPrefixNotificationItem**](QueryPrefixNotificationItem.md) |  | [optional] 
**NetworkInstance** | Pointer to **string** |  | [optional] 

## Methods

### NewProducerNotificationMsg

`func NewProducerNotificationMsg() *ProducerNotificationMsg`

NewProducerNotificationMsg instantiates a new ProducerNotificationMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProducerNotificationMsgWithDefaults

`func NewProducerNotificationMsgWithDefaults() *ProducerNotificationMsg`

NewProducerNotificationMsgWithDefaults instantiates a new ProducerNotificationMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProducerNotificationMsg) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProducerNotificationMsg) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProducerNotificationMsg) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProducerNotificationMsg) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItems

`func (o *ProducerNotificationMsg) GetItems() []QueryPrefixNotificationItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProducerNotificationMsg) GetItemsOk() (*[]QueryPrefixNotificationItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProducerNotificationMsg) SetItems(v []QueryPrefixNotificationItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *ProducerNotificationMsg) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *ProducerNotificationMsg) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *ProducerNotificationMsg) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *ProducerNotificationMsg) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *ProducerNotificationMsg) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


