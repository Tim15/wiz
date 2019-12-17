// Code generated by go-swagger; DO NOT EDIT.

package processors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// AddDataHandlerFunc turns a function with the right signature into a add data handler
type AddDataHandlerFunc func(AddDataParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddDataHandlerFunc) Handle(params AddDataParams) middleware.Responder {
	return fn(params)
}

// AddDataHandler interface for that can handle valid add data params
type AddDataHandler interface {
	Handle(AddDataParams) middleware.Responder
}

// NewAddData creates a new http.Handler for the add data operation
func NewAddData(ctx *middleware.Context, handler AddDataHandler) *AddData {
	return &AddData{Context: ctx, Handler: handler}
}

/*AddData swagger:route POST /processor/{id}/run/{runID}/data/input/{chunkID} Processors addData

AddData

TODO: add example data to values for this request, also:
figure out if we need to specifically distinguish between raw and file data at API level: yes

*/
type AddData struct {
	Context *middleware.Context
	Handler AddDataHandler
}

func (o *AddData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddDataParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddDataBody add data body
// swagger:model AddDataBody
type AddDataBody struct {

	// the ID of the chunk
	// Required: true
	ChunkID *string `json:"chunk_id"`

	// data
	Data interface{} `json:"data,omitempty"`

	// error
	Error *AddDataParamsBodyError `json:"error,omitempty"`

	// the output chunk ID association
	OutputChunkID string `json:"output_chunk_id,omitempty"`

	// the state of the data chunk.
	// Required: true
	// Enum: [Syncing Determining Validating Running Failed Succeeded]
	State *string `json:"state"`
}

// Validate validates this add data body
func (o *AddDataBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChunkID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddDataBody) validateChunkID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"chunk_id", "body", o.ChunkID); err != nil {
		return err
	}

	return nil
}

func (o *AddDataBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "error")
			}
			return err
		}
	}

	return nil
}

var addDataBodyTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Syncing","Determining","Validating","Running","Failed","Succeeded"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addDataBodyTypeStatePropEnum = append(addDataBodyTypeStatePropEnum, v)
	}
}

const (

	// AddDataBodyStateSyncing captures enum value "Syncing"
	AddDataBodyStateSyncing string = "Syncing"

	// AddDataBodyStateDetermining captures enum value "Determining"
	AddDataBodyStateDetermining string = "Determining"

	// AddDataBodyStateValidating captures enum value "Validating"
	AddDataBodyStateValidating string = "Validating"

	// AddDataBodyStateRunning captures enum value "Running"
	AddDataBodyStateRunning string = "Running"

	// AddDataBodyStateFailed captures enum value "Failed"
	AddDataBodyStateFailed string = "Failed"

	// AddDataBodyStateSucceeded captures enum value "Succeeded"
	AddDataBodyStateSucceeded string = "Succeeded"
)

// prop value enum
func (o *AddDataBody) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addDataBodyTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddDataBody) validateState(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"state", "body", o.State); err != nil {
		return err
	}

	// value enum
	if err := o.validateStateEnum("body"+"."+"state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddDataBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddDataBody) UnmarshalBinary(b []byte) error {
	var res AddDataBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddDataOKBody add data o k body
// swagger:model AddDataOKBody
type AddDataOKBody struct {

	// the ID of the chunk
	// Required: true
	ChunkID *string `json:"chunk_id"`

	// error
	Error *AddDataOKBodyError `json:"error,omitempty"`

	// the state of the data chunk.
	// Required: true
	// Enum: [Syncing Determining Validating Running Failed Succeeded]
	State *string `json:"state"`
}

// Validate validates this add data o k body
func (o *AddDataOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChunkID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddDataOKBody) validateChunkID(formats strfmt.Registry) error {

	if err := validate.Required("addDataOK"+"."+"chunk_id", "body", o.ChunkID); err != nil {
		return err
	}

	return nil
}

func (o *AddDataOKBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addDataOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

var addDataOKBodyTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Syncing","Determining","Validating","Running","Failed","Succeeded"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addDataOKBodyTypeStatePropEnum = append(addDataOKBodyTypeStatePropEnum, v)
	}
}

const (

	// AddDataOKBodyStateSyncing captures enum value "Syncing"
	AddDataOKBodyStateSyncing string = "Syncing"

	// AddDataOKBodyStateDetermining captures enum value "Determining"
	AddDataOKBodyStateDetermining string = "Determining"

	// AddDataOKBodyStateValidating captures enum value "Validating"
	AddDataOKBodyStateValidating string = "Validating"

	// AddDataOKBodyStateRunning captures enum value "Running"
	AddDataOKBodyStateRunning string = "Running"

	// AddDataOKBodyStateFailed captures enum value "Failed"
	AddDataOKBodyStateFailed string = "Failed"

	// AddDataOKBodyStateSucceeded captures enum value "Succeeded"
	AddDataOKBodyStateSucceeded string = "Succeeded"
)

// prop value enum
func (o *AddDataOKBody) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addDataOKBodyTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddDataOKBody) validateState(formats strfmt.Registry) error {

	if err := validate.Required("addDataOK"+"."+"state", "body", o.State); err != nil {
		return err
	}

	// value enum
	if err := o.validateStateEnum("addDataOK"+"."+"state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddDataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddDataOKBody) UnmarshalBinary(b []byte) error {
	var res AddDataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddDataOKBodyError the error which caused the chunk to fail if it is in the Failed state
// swagger:model AddDataOKBodyError
type AddDataOKBodyError struct {

	// The full trace of the error. This may only be available in a debug mode
	Trace string `json:"trace,omitempty"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this add data o k body error
func (o *AddDataOKBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddDataOKBodyError) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("addDataOK"+"."+"error"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddDataOKBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddDataOKBodyError) UnmarshalBinary(b []byte) error {
	var res AddDataOKBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddDataParamsBodyError the error which caused the chunk to fail if it is in the Failed state
// swagger:model AddDataParamsBodyError
type AddDataParamsBodyError struct {

	// The full trace of the error. This may only be available in a debug mode
	Trace string `json:"trace,omitempty"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this add data params body error
func (o *AddDataParamsBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddDataParamsBodyError) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"error"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddDataParamsBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddDataParamsBodyError) UnmarshalBinary(b []byte) error {
	var res AddDataParamsBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}