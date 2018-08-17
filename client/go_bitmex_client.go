// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/qct/go-bitmex/client/announcement"
	"github.com/qct/go-bitmex/client/api_key"
	"github.com/qct/go-bitmex/client/chat"
	"github.com/qct/go-bitmex/client/execution"
	"github.com/qct/go-bitmex/client/funding"
	"github.com/qct/go-bitmex/client/instrument"
	"github.com/qct/go-bitmex/client/insurance"
	"github.com/qct/go-bitmex/client/leaderboard"
	"github.com/qct/go-bitmex/client/liquidation"
	"github.com/qct/go-bitmex/client/notification"
	"github.com/qct/go-bitmex/client/order"
	"github.com/qct/go-bitmex/client/order_book"
	"github.com/qct/go-bitmex/client/position"
	"github.com/qct/go-bitmex/client/quote"
	"github.com/qct/go-bitmex/client/schema"
	"github.com/qct/go-bitmex/client/settlement"
	"github.com/qct/go-bitmex/client/stats"
	"github.com/qct/go-bitmex/client/trade"
	"github.com/qct/go-bitmex/client/user"
)

// Default go bitmex HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new go bitmex HTTP client.
func NewHTTPClient(formats strfmt.Registry) *GoBitmex {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new go bitmex HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *GoBitmex {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new go bitmex client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *GoBitmex {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(GoBitmex)
	cli.Transport = transport

	cli.Announcement = announcement.New(transport, formats)

	cli.APIKey = api_key.New(transport, formats)

	cli.Chat = chat.New(transport, formats)

	cli.Execution = execution.New(transport, formats)

	cli.Funding = funding.New(transport, formats)

	cli.Instrument = instrument.New(transport, formats)

	cli.Insurance = insurance.New(transport, formats)

	cli.Leaderboard = leaderboard.New(transport, formats)

	cli.Liquidation = liquidation.New(transport, formats)

	cli.Notification = notification.New(transport, formats)

	cli.Order = order.New(transport, formats)

	cli.OrderBook = order_book.New(transport, formats)

	cli.Position = position.New(transport, formats)

	cli.Quote = quote.New(transport, formats)

	cli.Schema = schema.New(transport, formats)

	cli.Settlement = settlement.New(transport, formats)

	cli.Stats = stats.New(transport, formats)

	cli.Trade = trade.New(transport, formats)

	cli.User = user.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// GoBitmex is a client for go bitmex
type GoBitmex struct {
	Announcement *announcement.Client

	APIKey *api_key.Client

	Chat *chat.Client

	Execution *execution.Client

	Funding *funding.Client

	Instrument *instrument.Client

	Insurance *insurance.Client

	Leaderboard *leaderboard.Client

	Liquidation *liquidation.Client

	Notification *notification.Client

	Order *order.Client

	OrderBook *order_book.Client

	Position *position.Client

	Quote *quote.Client

	Schema *schema.Client

	Settlement *settlement.Client

	Stats *stats.Client

	Trade *trade.Client

	User *user.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *GoBitmex) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Announcement.SetTransport(transport)

	c.APIKey.SetTransport(transport)

	c.Chat.SetTransport(transport)

	c.Execution.SetTransport(transport)

	c.Funding.SetTransport(transport)

	c.Instrument.SetTransport(transport)

	c.Insurance.SetTransport(transport)

	c.Leaderboard.SetTransport(transport)

	c.Liquidation.SetTransport(transport)

	c.Notification.SetTransport(transport)

	c.Order.SetTransport(transport)

	c.OrderBook.SetTransport(transport)

	c.Position.SetTransport(transport)

	c.Quote.SetTransport(transport)

	c.Schema.SetTransport(transport)

	c.Settlement.SetTransport(transport)

	c.Stats.SetTransport(transport)

	c.Trade.SetTransport(transport)

	c.User.SetTransport(transport)

}