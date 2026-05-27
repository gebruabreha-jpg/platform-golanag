# NexthopRegisterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 

## Methods

### NewNexthopRegisterItem

`func NewNexthopRegisterItem() *NexthopRegisterItem`

NewNexthopRegisterItem instantiates a new NexthopRegisterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopRegisterItemWithDefaults

`func NewNexthopRegisterItemWithDefaults() *NexthopRegisterItem`

NewNexthopRegisterItemWithDefaults instantiates a new NexthopRegisterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *NexthopRegisterItem) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *NexthopRegisterItem) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *NexthopRegisterItem) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *NexthopRegisterItem) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAddress

`func (o *NexthopRegisterItem) GetAddress() IpAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NexthopRegisterItem) GetAddressOk() (*IpAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NexthopRegisterItem) SetAddress(v IpAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NexthopRegisterItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


