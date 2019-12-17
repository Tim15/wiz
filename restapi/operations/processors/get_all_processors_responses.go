// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/alexkreidler/wiz/models"
)

// GetAllProcessorsOKCode is the HTTP code returned for type GetAllProcessorsOK
const GetAllProcessorsOKCode int = 200

/*GetAllProcessorsOK OK

swagger:response getAllProcessorsOK
*/
type GetAllProcessorsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ProcessorObject `json:"body,omitempty"`
}

// NewGetAllProcessorsOK creates GetAllProcessorsOK with default headers values
func NewGetAllProcessorsOK() *GetAllProcessorsOK {

	return &GetAllProcessorsOK{}
}

// WithPayload adds the payload to the get all processors o k response
func (o *GetAllProcessorsOK) WithPayload(payload []*models.ProcessorObject) *GetAllProcessorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all processors o k response
func (o *GetAllProcessorsOK) SetPayload(payload []*models.ProcessorObject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllProcessorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ProcessorObject, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
