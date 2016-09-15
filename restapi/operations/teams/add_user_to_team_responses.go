package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/teresa-api/models"
)

/*AddUserToTeamOK Team details

swagger:response addUserToTeamOK
*/
type AddUserToTeamOK struct {

	// In: body
	Payload *models.Team `json:"body,omitempty"`
}

// NewAddUserToTeamOK creates AddUserToTeamOK with default headers values
func NewAddUserToTeamOK() *AddUserToTeamOK {
	return &AddUserToTeamOK{}
}

// WithPayload adds the payload to the add user to team o k response
func (o *AddUserToTeamOK) WithPayload(payload *models.Team) *AddUserToTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to team o k response
func (o *AddUserToTeamOK) SetPayload(payload *models.Team) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddUserToTeamDefault Error

swagger:response addUserToTeamDefault
*/
type AddUserToTeamDefault struct {
	_statusCode int

	// In: body
	Payload *models.GenericError `json:"body,omitempty"`
}

// NewAddUserToTeamDefault creates AddUserToTeamDefault with default headers values
func NewAddUserToTeamDefault(code int) *AddUserToTeamDefault {
	if code <= 0 {
		code = 500
	}

	return &AddUserToTeamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add user to team default response
func (o *AddUserToTeamDefault) WithStatusCode(code int) *AddUserToTeamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add user to team default response
func (o *AddUserToTeamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add user to team default response
func (o *AddUserToTeamDefault) WithPayload(payload *models.GenericError) *AddUserToTeamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add user to team default response
func (o *AddUserToTeamDefault) SetPayload(payload *models.GenericError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddUserToTeamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
