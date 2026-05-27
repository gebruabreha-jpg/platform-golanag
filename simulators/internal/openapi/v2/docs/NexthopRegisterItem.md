# NexthopRegisterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | A valid IP address. | 
**AddressFamily** | [**AddressFamily**](AddressFamily.md) |  | 

## Methods

### NewNexthopRegisterItem

`func NewNexthopRegisterItem(address string, addressFamily AddressFamily, ) *NexthopRegisterItem`

NewNexthopRegisterItem instantiates a new NexthopRegisterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopRegisterItemWithDefaults

`func NewNexthopRegisterItemWithDefaults() *NexthopRegisterItem`

NewNexthopRegisterItemWithDefaults instantiates a new NexthopRegisterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NexthopRegisterItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NexthopRegisterItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NexthopRegisterItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressFamily

`func (o *NexthopRegisterItem) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *NexthopRegisterItem) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *NexthopRegisterItem) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


