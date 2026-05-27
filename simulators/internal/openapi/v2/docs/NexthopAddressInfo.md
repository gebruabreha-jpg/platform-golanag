# NexthopAddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Address** | **string** | A valid IP address. | 
**AddressFamily** | [**AddressFamily**](AddressFamily.md) |  | 
**ServiceName** | **string** | Name of the service. | 
**InstanceId** | Pointer to **string** | Pod unique ID. | [optional] 

## Methods

### NewNexthopAddressInfo

`func NewNexthopAddressInfo(type_ string, address string, addressFamily AddressFamily, serviceName string, ) *NexthopAddressInfo`

NewNexthopAddressInfo instantiates a new NexthopAddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopAddressInfoWithDefaults

`func NewNexthopAddressInfoWithDefaults() *NexthopAddressInfo`

NewNexthopAddressInfoWithDefaults instantiates a new NexthopAddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NexthopAddressInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NexthopAddressInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NexthopAddressInfo) SetType(v string)`

SetType sets Type field to given value.


### GetAddress

`func (o *NexthopAddressInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NexthopAddressInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NexthopAddressInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressFamily

`func (o *NexthopAddressInfo) GetAddressFamily() AddressFamily`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *NexthopAddressInfo) GetAddressFamilyOk() (*AddressFamily, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *NexthopAddressInfo) SetAddressFamily(v AddressFamily)`

SetAddressFamily sets AddressFamily field to given value.


### GetServiceName

`func (o *NexthopAddressInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *NexthopAddressInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *NexthopAddressInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetInstanceId

`func (o *NexthopAddressInfo) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *NexthopAddressInfo) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *NexthopAddressInfo) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *NexthopAddressInfo) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


