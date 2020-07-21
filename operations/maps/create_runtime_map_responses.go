// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models/v2"
)

// CreateRuntimeMapCreatedCode is the HTTP code returned for type CreateRuntimeMapCreated
const CreateRuntimeMapCreatedCode int = 201

/*CreateRuntimeMapCreated Map file created with its entries

swagger:response createRuntimeMapCreated
*/
type CreateRuntimeMapCreated struct {

	/*
	  In: Body
	*/
	Payload models.MapEntries `json:"body,omitempty"`
}

// NewCreateRuntimeMapCreated creates CreateRuntimeMapCreated with default headers values
func NewCreateRuntimeMapCreated() *CreateRuntimeMapCreated {

	return &CreateRuntimeMapCreated{}
}

// WithPayload adds the payload to the create runtime map created response
func (o *CreateRuntimeMapCreated) WithPayload(payload models.MapEntries) *CreateRuntimeMapCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create runtime map created response
func (o *CreateRuntimeMapCreated) SetPayload(payload models.MapEntries) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRuntimeMapCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.MapEntries{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateRuntimeMapAcceptedCode is the HTTP code returned for type CreateRuntimeMapAccepted
const CreateRuntimeMapAcceptedCode int = 202

/*CreateRuntimeMapAccepted Configuration change accepted and reload requested

swagger:response createRuntimeMapAccepted
*/
type CreateRuntimeMapAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload models.MapEntries `json:"body,omitempty"`
}

// NewCreateRuntimeMapAccepted creates CreateRuntimeMapAccepted with default headers values
func NewCreateRuntimeMapAccepted() *CreateRuntimeMapAccepted {

	return &CreateRuntimeMapAccepted{}
}

// WithReloadID adds the reloadId to the create runtime map accepted response
func (o *CreateRuntimeMapAccepted) WithReloadID(reloadID string) *CreateRuntimeMapAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create runtime map accepted response
func (o *CreateRuntimeMapAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create runtime map accepted response
func (o *CreateRuntimeMapAccepted) WithPayload(payload models.MapEntries) *CreateRuntimeMapAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create runtime map accepted response
func (o *CreateRuntimeMapAccepted) SetPayload(payload models.MapEntries) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRuntimeMapAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.MapEntries{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateRuntimeMapBadRequestCode is the HTTP code returned for type CreateRuntimeMapBadRequest
const CreateRuntimeMapBadRequestCode int = 400

/*CreateRuntimeMapBadRequest Bad request

swagger:response createRuntimeMapBadRequest
*/
type CreateRuntimeMapBadRequest struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateRuntimeMapBadRequest creates CreateRuntimeMapBadRequest with default headers values
func NewCreateRuntimeMapBadRequest() *CreateRuntimeMapBadRequest {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateRuntimeMapBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create runtime map bad request response
func (o *CreateRuntimeMapBadRequest) WithConfigurationVersion(configurationVersion int64) *CreateRuntimeMapBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create runtime map bad request response
func (o *CreateRuntimeMapBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create runtime map bad request response
func (o *CreateRuntimeMapBadRequest) WithPayload(payload *models.Error) *CreateRuntimeMapBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create runtime map bad request response
func (o *CreateRuntimeMapBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRuntimeMapBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRuntimeMapConflictCode is the HTTP code returned for type CreateRuntimeMapConflict
const CreateRuntimeMapConflictCode int = 409

/*CreateRuntimeMapConflict The specified resource already exists

swagger:response createRuntimeMapConflict
*/
type CreateRuntimeMapConflict struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateRuntimeMapConflict creates CreateRuntimeMapConflict with default headers values
func NewCreateRuntimeMapConflict() *CreateRuntimeMapConflict {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateRuntimeMapConflict{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create runtime map conflict response
func (o *CreateRuntimeMapConflict) WithConfigurationVersion(configurationVersion int64) *CreateRuntimeMapConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create runtime map conflict response
func (o *CreateRuntimeMapConflict) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create runtime map conflict response
func (o *CreateRuntimeMapConflict) WithPayload(payload *models.Error) *CreateRuntimeMapConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create runtime map conflict response
func (o *CreateRuntimeMapConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRuntimeMapConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateRuntimeMapDefault General Error

swagger:response createRuntimeMapDefault
*/
type CreateRuntimeMapDefault struct {
	_statusCode int
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateRuntimeMapDefault creates CreateRuntimeMapDefault with default headers values
func NewCreateRuntimeMapDefault(code int) *CreateRuntimeMapDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateRuntimeMapDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the create runtime map default response
func (o *CreateRuntimeMapDefault) WithStatusCode(code int) *CreateRuntimeMapDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create runtime map default response
func (o *CreateRuntimeMapDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create runtime map default response
func (o *CreateRuntimeMapDefault) WithConfigurationVersion(configurationVersion int64) *CreateRuntimeMapDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create runtime map default response
func (o *CreateRuntimeMapDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create runtime map default response
func (o *CreateRuntimeMapDefault) WithPayload(payload *models.Error) *CreateRuntimeMapDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create runtime map default response
func (o *CreateRuntimeMapDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRuntimeMapDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
