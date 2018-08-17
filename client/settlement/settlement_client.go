// Code generated by go-swagger; DO NOT EDIT.

package settlement

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new settlement API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for settlement API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SettlementGet gets settlement history
*/
func (a *Client) SettlementGet(params *SettlementGetParams) (*SettlementGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettlementGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Settlement.get",
		Method:             "GET",
		PathPattern:        "/settlement",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SettlementGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SettlementGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
