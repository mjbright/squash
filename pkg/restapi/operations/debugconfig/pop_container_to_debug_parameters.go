package debugconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPopContainerToDebugParams creates a new PopContainerToDebugParams object
// with the default values initialized.
func NewPopContainerToDebugParams() PopContainerToDebugParams {
	var ()
	return PopContainerToDebugParams{}
}

// PopContainerToDebugParams contains all the bound params for the pop container to debug operation
// typically these are obtained from a http.Request
//
// swagger:parameters popContainerToDebug
type PopContainerToDebugParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*ID the node we are watching
	  Required: true
	  In: path
	*/
	Node string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PopContainerToDebugParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rNode, rhkNode, _ := route.Params.GetOK("node")
	if err := o.bindNode(rNode, rhkNode, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PopContainerToDebugParams) bindNode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Node = raw

	return nil
}
