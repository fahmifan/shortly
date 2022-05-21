// Code generated by go-swagger; DO NOT EDIT.

package urls

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/fahmifan/shortly/gen/models"
)

// ListURLsOKCode is the HTTP code returned for type ListURLsOK
const ListURLsOKCode int = 200

/*ListURLsOK list shorten urls

swagger:response listURLsOK
*/
type ListURLsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.URL `json:"body,omitempty"`
}

// NewListURLsOK creates ListURLsOK with default headers values
func NewListURLsOK() *ListURLsOK {

	return &ListURLsOK{}
}

// WithPayload adds the payload to the list u r ls o k response
func (o *ListURLsOK) WithPayload(payload []*models.URL) *ListURLsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list u r ls o k response
func (o *ListURLsOK) SetPayload(payload []*models.URL) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListURLsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.URL, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*ListURLsDefault generic error response

swagger:response listURLsDefault
*/
type ListURLsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListURLsDefault creates ListURLsDefault with default headers values
func NewListURLsDefault(code int) *ListURLsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListURLsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list u r ls default response
func (o *ListURLsDefault) WithStatusCode(code int) *ListURLsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list u r ls default response
func (o *ListURLsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list u r ls default response
func (o *ListURLsDefault) WithPayload(payload *models.Error) *ListURLsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list u r ls default response
func (o *ListURLsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListURLsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}