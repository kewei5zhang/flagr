// Code generated by go-swagger; DO NOT EDIT.

package variant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/checkr/flagr/swagger_gen/models"
)

// CreateVariantOKCode is the HTTP code returned for type CreateVariantOK
const CreateVariantOKCode int = 200

/*CreateVariantOK variant just created

swagger:response createVariantOK
*/
type CreateVariantOK struct {

	/*
	  In: Body
	*/
	Payload *models.Variant `json:"body,omitempty"`
}

// NewCreateVariantOK creates CreateVariantOK with default headers values
func NewCreateVariantOK() *CreateVariantOK {
	return &CreateVariantOK{}
}

// WithPayload adds the payload to the create variant o k response
func (o *CreateVariantOK) WithPayload(payload *models.Variant) *CreateVariantOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create variant o k response
func (o *CreateVariantOK) SetPayload(payload *models.Variant) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVariantOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateVariantDefault generic error response

swagger:response createVariantDefault
*/
type CreateVariantDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateVariantDefault creates CreateVariantDefault with default headers values
func NewCreateVariantDefault(code int) *CreateVariantDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateVariantDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create variant default response
func (o *CreateVariantDefault) WithStatusCode(code int) *CreateVariantDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create variant default response
func (o *CreateVariantDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create variant default response
func (o *CreateVariantDefault) WithPayload(payload *models.Error) *CreateVariantDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create variant default response
func (o *CreateVariantDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVariantDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}