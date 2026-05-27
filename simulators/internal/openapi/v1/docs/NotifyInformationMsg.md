# NotifyInformationMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]PrefixNotificationItem**](PrefixNotificationItem.md) |  | [optional] 
**Header** | Pointer to [**RedistributionNotificationHeader**](RedistributionNotificationHeader.md) |  | [optional] 
**Routes** | Pointer to [**[]PublishedRoute**](PublishedRoute.md) |  | [optional] 
**NetworkInstance** | Pointer to **string** |  | [optional] 
**Nexthops** | Pointer to [**[]PublishedNexthop**](PublishedNexthop.md) |  | [optional] 

## Methods

### NewNotifyInformationMsg

`func NewNotifyInformationMsg() *NotifyInformationMsg`

NewNotifyInformationMsg instantiates a new NotifyInformationMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyInformationMsgWithDefaults

`func NewNotifyInformationMsgWithDefaults() *NotifyInformationMsg`

NewNotifyInformationMsgWithDefaults instantiates a new NotifyInformationMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotifyInformationMsg) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotifyInformationMsg) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotifyInformationMsg) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotifyInformationMsg) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItems

`func (o *NotifyInformationMsg) GetItems() []PrefixNotificationItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NotifyInformationMsg) GetItemsOk() (*[]PrefixNotificationItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NotifyInformationMsg) SetItems(v []PrefixNotificationItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *NotifyInformationMsg) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetHeader

`func (o *NotifyInformationMsg) GetHeader() RedistributionNotificationHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *NotifyInformationMsg) GetHeaderOk() (*RedistributionNotificationHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *NotifyInformationMsg) SetHeader(v RedistributionNotificationHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *NotifyInformationMsg) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetRoutes

`func (o *NotifyInformationMsg) GetRoutes() []PublishedRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NotifyInformationMsg) GetRoutesOk() (*[]PublishedRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NotifyInformationMsg) SetRoutes(v []PublishedRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *NotifyInformationMsg) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetNetworkInstance

`func (o *NotifyInformationMsg) GetNetworkInstance() string`

GetNetworkInstance returns the NetworkInstance field if non-nil, zero value otherwise.

### GetNetworkInstanceOk

`func (o *NotifyInformationMsg) GetNetworkInstanceOk() (*string, bool)`

GetNetworkInstanceOk returns a tuple with the NetworkInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstance

`func (o *NotifyInformationMsg) SetNetworkInstance(v string)`

SetNetworkInstance sets NetworkInstance field to given value.

### HasNetworkInstance

`func (o *NotifyInformationMsg) HasNetworkInstance() bool`

HasNetworkInstance returns a boolean if a field has been set.

### GetNexthops

`func (o *NotifyInformationMsg) GetNexthops() []PublishedNexthop`

GetNexthops returns the Nexthops field if non-nil, zero value otherwise.

### GetNexthopsOk

`func (o *NotifyInformationMsg) GetNexthopsOk() (*[]PublishedNexthop, bool)`

GetNexthopsOk returns a tuple with the Nexthops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthops

`func (o *NotifyInformationMsg) SetNexthops(v []PublishedNexthop)`

SetNexthops sets Nexthops field to given value.

### HasNexthops

`func (o *NotifyInformationMsg) HasNexthops() bool`

HasNexthops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


