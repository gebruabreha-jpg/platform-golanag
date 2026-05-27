# NexthopIPInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | A valid IP address. | 
**AddressFamily** | [**AddressFamily**](AddressFamily.md) |  | 
**Type** | **string** |  | 

## Methods

### NewNexthopIPInfo

`func NewNexthopIPInfo(address string, addressFamily AddressFamily, type_ string, ) *NexthopIPInfo`

NewNexthopIPInfo instantiates a new NexthopIPInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopIPInfoWithDefaults

`func NewNexthopIPInfoWithDefaults() *NexthopIPInfo`

NewNexthopIPInfoWithDefaults instantiates a new NexthopIPInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NexthopIPInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NexthopIPInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NexthopIPInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressFamily

`func (o *NexthopIPInfo) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *NexthopIPInfo) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *NexthopIPInfo) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.


### GetType

`func (o *NexthopIPInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NexthopIPInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NexthopIPInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


