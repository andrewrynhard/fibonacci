// Code generated by go-swagger; DO NOT EDIT.

package sequence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetSequenceHandlerFunc turns a function with the right signature into a get sequence handler
type GetSequenceHandlerFunc func(GetSequenceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSequenceHandlerFunc) Handle(params GetSequenceParams) middleware.Responder {
	return fn(params)
}

// GetSequenceHandler interface for that can handle valid get sequence params
type GetSequenceHandler interface {
	Handle(GetSequenceParams) middleware.Responder
}

// NewGetSequence creates a new http.Handler for the get sequence operation
func NewGetSequence(ctx *middleware.Context, handler GetSequenceHandler) *GetSequence {
	return &GetSequence{Context: ctx, Handler: handler}
}

/*GetSequence swagger:route GET /sequence/{n} sequence getSequence

Returns the first n Fibonacci numbers

*/
type GetSequence struct {
	Context *middleware.Context
	Handler GetSequenceHandler
}

func (o *GetSequence) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSequenceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
