package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"
)

/*DeleteEventByIDNoContent Successful response

swagger:response deleteEventByIdNoContent
*/
type DeleteEventByIDNoContent struct {
}

// WriteResponse to the client
func (o *DeleteEventByIDNoContent) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(204)
}

/*DeleteEventByIDDefault Generic Error

swagger:response deleteEventByIdDefault
*/
type DeleteEventByIDDefault struct {
}

// WriteResponse to the client
func (o *DeleteEventByIDDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(500)
}
