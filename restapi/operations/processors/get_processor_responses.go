// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/alexkreidler/wiz/models"
)

// GetProcessorOKCode is the HTTP code returned for type GetProcessorOK
const GetProcessorOKCode int = 200

/*GetProcessorOK OK

swagger:response getProcessorOK
*/
type GetProcessorOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProcessorObject `json:"body,omitempty"`
}

// NewGetProcessorOK creates GetProcessorOK with default headers values
func NewGetProcessorOK() *GetProcessorOK {

	return &GetProcessorOK{}
}

// WithPayload adds the payload to the get processor o k response
func (o *GetProcessorOK) WithPayload(payload *models.ProcessorObject) *GetProcessorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get processor o k response
func (o *GetProcessorOK) SetPayload(payload *models.ProcessorObject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProcessorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProcessorNotFoundCode is the HTTP code returned for type GetProcessorNotFound
const GetProcessorNotFoundCode int = 404

/*GetProcessorNotFound Not Found

swagger:response getProcessorNotFound
*/
type GetProcessorNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetProcessorNotFound creates GetProcessorNotFound with default headers values
func NewGetProcessorNotFound() *GetProcessorNotFound {

	return &GetProcessorNotFound{}
}

// WithPayload adds the payload to the get processor not found response
func (o *GetProcessorNotFound) WithPayload(payload *models.Error) *GetProcessorNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get processor not found response
func (o *GetProcessorNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProcessorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
