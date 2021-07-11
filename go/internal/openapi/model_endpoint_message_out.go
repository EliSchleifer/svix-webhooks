/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the Svix API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each of your users. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// EndpointMessageOut struct for EndpointMessageOut
type EndpointMessageOut struct {
	EventId *string `json:"eventId,omitempty"`
	EventType string `json:"eventType"`
	Id string `json:"id"`
	Payload map[string]interface{} `json:"payload"`
	Status MessageStatus `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// NewEndpointMessageOut instantiates a new EndpointMessageOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointMessageOut(eventType string, id string, payload map[string]interface{}, status MessageStatus, timestamp time.Time, ) *EndpointMessageOut {
	this := EndpointMessageOut{}
	this.EventType = eventType
	this.Id = id
	this.Payload = payload
	this.Status = status
	this.Timestamp = timestamp
	return &this
}

// NewEndpointMessageOutWithDefaults instantiates a new EndpointMessageOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointMessageOutWithDefaults() *EndpointMessageOut {
	this := EndpointMessageOut{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *EndpointMessageOut) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *EndpointMessageOut) HasEventId() bool {
	if o != nil && o.EventId != nil {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *EndpointMessageOut) SetEventId(v string) {
	o.EventId = &v
}

// GetEventType returns the EventType field value
func (o *EndpointMessageOut) GetEventType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *EndpointMessageOut) SetEventType(v string) {
	o.EventType = v
}

// GetId returns the Id field value
func (o *EndpointMessageOut) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EndpointMessageOut) SetId(v string) {
	o.Id = v
}

// GetPayload returns the Payload field value
func (o *EndpointMessageOut) GetPayload() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *EndpointMessageOut) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetStatus returns the Status field value
func (o *EndpointMessageOut) GetStatus() MessageStatus {
	if o == nil  {
		var ret MessageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetStatusOk() (*MessageStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EndpointMessageOut) SetStatus(v MessageStatus) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value
func (o *EndpointMessageOut) GetTimestamp() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *EndpointMessageOut) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *EndpointMessageOut) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o EndpointMessageOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventId != nil {
		toSerialize["eventId"] = o.EventId
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointMessageOut struct {
	value *EndpointMessageOut
	isSet bool
}

func (v NullableEndpointMessageOut) Get() *EndpointMessageOut {
	return v.value
}

func (v *NullableEndpointMessageOut) Set(val *EndpointMessageOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointMessageOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointMessageOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointMessageOut(val *EndpointMessageOut) *NullableEndpointMessageOut {
	return &NullableEndpointMessageOut{value: val, isSet: true}
}

func (v NullableEndpointMessageOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointMessageOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


