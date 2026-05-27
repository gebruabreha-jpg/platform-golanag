# InterfaceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IfName** | Pointer to **string** | For subnet route, needs to carry interface name for display. for other routes, interface name is not needed. | [optional] 
**IfType** | Pointer to [**IfType**](IfType.md) |  | [optional] 

## Methods

### NewInterfaceInfo

`func NewInterfaceInfo() *InterfaceInfo`

NewInterfaceInfo instantiates a new InterfaceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceInfoWithDefaults

`func NewInterfaceInfoWithDefaults() *InterfaceInfo`

NewInterfaceInfoWithDefaults instantiates a new InterfaceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIfName

`func (o *InterfaceInfo) GetIfName() string`

GetIfName returns the IfName field if non-nil, zero value otherwise.

### GetIfNameOk

`func (o *InterfaceInfo) GetIfNameOk() (*string, bool)`

GetIfNameOk returns a tuple with the IfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfName

`func (o *InterfaceInfo) SetIfName(v string)`

SetIfName sets IfName field to given value.

### HasIfName

`func (o *InterfaceInfo) HasIfName() bool`

HasIfName returns a boolean if a field has been set.

### GetIfType

`func (o *InterfaceInfo) GetIfType() IfType`

GetIfType returns the IfType field if non-nil, zero value otherwise.

### GetIfTypeOk

`func (o *InterfaceInfo) GetIfTypeOk() (*IfType, bool)`

GetIfTypeOk returns a tuple with the IfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfType

`func (o *InterfaceInfo) SetIfType(v IfType)`

SetIfType sets IfType field to given value.

### HasIfType

`func (o *InterfaceInfo) HasIfType() bool`

HasIfType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


