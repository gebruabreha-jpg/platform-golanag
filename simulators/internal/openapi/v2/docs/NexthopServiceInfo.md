# NexthopServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** | Name of the service. | 
**InstanceId** | Pointer to **string** | Pod unique ID. | [optional] 
**Type** | **string** |  | 

## Methods

### NewNexthopServiceInfo

`func NewNexthopServiceInfo(serviceName string, type_ string, ) *NexthopServiceInfo`

NewNexthopServiceInfo instantiates a new NexthopServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexthopServiceInfoWithDefaults

`func NewNexthopServiceInfoWithDefaults() *NexthopServiceInfo`

NewNexthopServiceInfoWithDefaults instantiates a new NexthopServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *NexthopServiceInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *NexthopServiceInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *NexthopServiceInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetInstanceId

`func (o *NexthopServiceInfo) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *NexthopServiceInfo) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *NexthopServiceInfo) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *NexthopServiceInfo) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetType

`func (o *NexthopServiceInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NexthopServiceInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NexthopServiceInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


