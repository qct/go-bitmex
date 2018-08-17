// Code generated by go-swagger; DO NOT EDIT.

package api_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new api key API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api key API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
APIKeyDisable disables an API key
*/
func (a *Client) APIKeyDisable(params *APIKeyDisableParams, authInfo runtime.ClientAuthInfoWriter) (*APIKeyDisableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyDisableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "APIKey.disable",
		Method:             "POST",
		PathPattern:        "/apiKey/disable",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyDisableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIKeyDisableOK), nil

}

/*
APIKeyEnable enables an API key
*/
func (a *Client) APIKeyEnable(params *APIKeyEnableParams, authInfo runtime.ClientAuthInfoWriter) (*APIKeyEnableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyEnableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "APIKey.enable",
		Method:             "POST",
		PathPattern:        "/apiKey/enable",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyEnableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIKeyEnableOK), nil

}

/*
APIKeyGet gets your API keys
*/
func (a *Client) APIKeyGet(params *APIKeyGetParams, authInfo runtime.ClientAuthInfoWriter) (*APIKeyGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "APIKey.get",
		Method:             "GET",
		PathPattern:        "/apiKey",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIKeyGetOK), nil

}

/*
APIKeyNew creates a new API key

API Keys can only be created via the frontend.
*/
func (a *Client) APIKeyNew(params *APIKeyNewParams, authInfo runtime.ClientAuthInfoWriter) (*APIKeyNewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyNewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "APIKey.new",
		Method:             "POST",
		PathPattern:        "/apiKey",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyNewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIKeyNewOK), nil

}

/*
APIKeyRemove removes an API key
*/
func (a *Client) APIKeyRemove(params *APIKeyRemoveParams, authInfo runtime.ClientAuthInfoWriter) (*APIKeyRemoveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIKeyRemoveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "APIKey.remove",
		Method:             "DELETE",
		PathPattern:        "/apiKey",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &APIKeyRemoveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIKeyRemoveOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
