/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// DeleteNonEmployeeRecordsInBulkRequest struct for DeleteNonEmployeeRecordsInBulkRequest
type DeleteNonEmployeeRecordsInBulkRequest struct {
	// List of non-employee ids.
	Ids []string `json:"ids"`
	AdditionalProperties map[string]interface{}
}

type _DeleteNonEmployeeRecordsInBulkRequest DeleteNonEmployeeRecordsInBulkRequest

// NewDeleteNonEmployeeRecordsInBulkRequest instantiates a new DeleteNonEmployeeRecordsInBulkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteNonEmployeeRecordsInBulkRequest(ids []string) *DeleteNonEmployeeRecordsInBulkRequest {
	this := DeleteNonEmployeeRecordsInBulkRequest{}
	this.Ids = ids
	return &this
}

// NewDeleteNonEmployeeRecordsInBulkRequestWithDefaults instantiates a new DeleteNonEmployeeRecordsInBulkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteNonEmployeeRecordsInBulkRequestWithDefaults() *DeleteNonEmployeeRecordsInBulkRequest {
	this := DeleteNonEmployeeRecordsInBulkRequest{}
	return &this
}

// GetIds returns the Ids field value
func (o *DeleteNonEmployeeRecordsInBulkRequest) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *DeleteNonEmployeeRecordsInBulkRequest) GetIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *DeleteNonEmployeeRecordsInBulkRequest) SetIds(v []string) {
	o.Ids = v
}

func (o DeleteNonEmployeeRecordsInBulkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeleteNonEmployeeRecordsInBulkRequest) UnmarshalJSON(bytes []byte) (err error) {
	varDeleteNonEmployeeRecordsInBulkRequest := _DeleteNonEmployeeRecordsInBulkRequest{}

	if err = json.Unmarshal(bytes, &varDeleteNonEmployeeRecordsInBulkRequest); err == nil {
		*o = DeleteNonEmployeeRecordsInBulkRequest(varDeleteNonEmployeeRecordsInBulkRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteNonEmployeeRecordsInBulkRequest struct {
	value *DeleteNonEmployeeRecordsInBulkRequest
	isSet bool
}

func (v NullableDeleteNonEmployeeRecordsInBulkRequest) Get() *DeleteNonEmployeeRecordsInBulkRequest {
	return v.value
}

func (v *NullableDeleteNonEmployeeRecordsInBulkRequest) Set(val *DeleteNonEmployeeRecordsInBulkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteNonEmployeeRecordsInBulkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteNonEmployeeRecordsInBulkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteNonEmployeeRecordsInBulkRequest(val *DeleteNonEmployeeRecordsInBulkRequest) *NullableDeleteNonEmployeeRecordsInBulkRequest {
	return &NullableDeleteNonEmployeeRecordsInBulkRequest{value: val, isSet: true}
}

func (v NullableDeleteNonEmployeeRecordsInBulkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteNonEmployeeRecordsInBulkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

