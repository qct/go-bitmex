// Code generated by go-swagger; DO NOT EDIT.

package chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewChatGetConnectedParams creates a new ChatGetConnectedParams object
// with the default values initialized.
func NewChatGetConnectedParams() *ChatGetConnectedParams {

	return &ChatGetConnectedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChatGetConnectedParamsWithTimeout creates a new ChatGetConnectedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChatGetConnectedParamsWithTimeout(timeout time.Duration) *ChatGetConnectedParams {

	return &ChatGetConnectedParams{

		timeout: timeout,
	}
}

// NewChatGetConnectedParamsWithContext creates a new ChatGetConnectedParams object
// with the default values initialized, and the ability to set a context for a request
func NewChatGetConnectedParamsWithContext(ctx context.Context) *ChatGetConnectedParams {

	return &ChatGetConnectedParams{

		Context: ctx,
	}
}

// NewChatGetConnectedParamsWithHTTPClient creates a new ChatGetConnectedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChatGetConnectedParamsWithHTTPClient(client *http.Client) *ChatGetConnectedParams {

	return &ChatGetConnectedParams{
		HTTPClient: client,
	}
}

/*ChatGetConnectedParams contains all the parameters to send to the API endpoint
for the chat get connected operation typically these are written to a http.Request
*/
type ChatGetConnectedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the chat get connected params
func (o *ChatGetConnectedParams) WithTimeout(timeout time.Duration) *ChatGetConnectedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chat get connected params
func (o *ChatGetConnectedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chat get connected params
func (o *ChatGetConnectedParams) WithContext(ctx context.Context) *ChatGetConnectedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chat get connected params
func (o *ChatGetConnectedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chat get connected params
func (o *ChatGetConnectedParams) WithHTTPClient(client *http.Client) *ChatGetConnectedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chat get connected params
func (o *ChatGetConnectedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ChatGetConnectedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
