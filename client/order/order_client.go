// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new order API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for order API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
OrderAmend amends the quantity or price of an open order

Send an `orderID` or `origClOrdID` to identify the order you wish to amend.

Both order quantity and price can be amended. Only one `qty` field can be used to amend.

Use the `leavesQty` field to specify how much of the order you wish to remain open. This can be useful
if you want to adjust your position's delta by a certain amount, regardless of how much of the order has
already filled.

> A `leavesQty` can be used to make a "Filled" order live again, if it is received within 60 seconds of the fill.

Use the `simpleOrderQty` and `simpleLeavesQty` fields to specify order size in Bitcoin, rather than contracts.
These fields will round up to the nearest contract.

Like order placement, amending can be done in bulk. Simply send a request to `PUT /api/v1/order/bulk` with
a JSON body of the shape: `{"orders": [{...}, {...}]}`, each object containing the fields used in this endpoint.

*/
func (a *Client) OrderAmend(params *OrderAmendParams, authInfo runtime.ClientAuthInfoWriter) (*OrderAmendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderAmendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Order.amend",
		Method:             "PUT",
		PathPattern:        "/order",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrderAmendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrderAmendOK), nil

}

/*
OrderAmendBulk amends multiple orders for the same symbol

Similar to POST /amend, but with multiple orders. `application/json` only. Ratelimited at 10%.
*/
func (a *Client) OrderAmendBulk(params *OrderAmendBulkParams, authInfo runtime.ClientAuthInfoWriter) (*OrderAmendBulkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderAmendBulkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Order.amendBulk",
		Method:             "PUT",
		PathPattern:        "/order/bulk",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrderAmendBulkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrderAmendBulkOK), nil

}

/*
OrderCancel cancels order s send multiple order ids to cancel in bulk

Either an orderID or a clOrdID must be provided.
*/
func (a *Client) OrderCancel(params *OrderCancelParams, authInfo runtime.ClientAuthInfoWriter) (*OrderCancelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderCancelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Order.cancel",
		Method:             "DELETE",
		PathPattern:        "/order",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrderCancelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrderCancelOK), nil

}

/*
OrderCancelAll cancels all of your orders
*/
func (a *Client) OrderCancelAll(params *OrderCancelAllParams, authInfo runtime.ClientAuthInfoWriter) (*OrderCancelAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderCancelAllParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Order.cancelAll",
		Method:             "DELETE",
		PathPattern:        "/order/all",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrderCancelAllReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrderCancelAllOK), nil

}

/*
OrderCancelAllAfter automaticallies cancel all your orders after a specified timeout

Useful as a dead-man's switch to ensure your orders are canceled in case of an outage.
If called repeatedly, the existing offset will be canceled and a new one will be inserted in its place.

Example usage: call this route at 15s intervals with an offset of 60000 (60s).
If this route is not called within 60 seconds, all your orders will be automatically canceled.

This is also available via [WebSocket](https://www.bitmex.com/app/wsAPI#Dead-Mans-Switch-Auto-Cancel).

*/
func (a *Client) OrderCancelAllAfter(params *OrderCancelAllAfterParams, authInfo runtime.ClientAuthInfoWriter) (*OrderCancelAllAfterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderCancelAllAfterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Order.cancelAllAfter",
		Method:             "POST",
		PathPattern:        "/order/cancelAllAfter",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrderCancelAllAfterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrderCancelAllAfterOK), nil

}

/*
OrderClosePosition closes a position deprecated use p o s t order with exec inst close

If no `price` is specified, a market order will be submitted to close the whole of your position. This will also close all other open orders in this symbol.
*/
func (a *Client) OrderClosePosition(params *OrderClosePositionParams, authInfo runtime.ClientAuthInfoWriter) (*OrderClosePositionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderClosePositionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Order.closePosition",
		Method:             "POST",
		PathPattern:        "/order/closePosition",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrderClosePositionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrderClosePositionOK), nil

}

/*
OrderGetOrders gets your orders

To get open orders only, send {"open": true} in the filter param.

See <a href="http://www.onixs.biz/fix-dictionary/5.0.SP2/msgType_D_68.html">the FIX Spec</a> for explanations of these fields.
*/
func (a *Client) OrderGetOrders(params *OrderGetOrdersParams, authInfo runtime.ClientAuthInfoWriter) (*OrderGetOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderGetOrdersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Order.getOrders",
		Method:             "GET",
		PathPattern:        "/order",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrderGetOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrderGetOrdersOK), nil

}

/*
OrderNew creates a new order

## Placing Orders

This endpoint is used for placing orders. See individual fields below for more details on their use.

#### Order Types

All orders require a `symbol`. All other fields are optional except when otherwise specified.

These are the valid `ordType`s:

* **Limit**: The default order type. Specify an `orderQty` and `price`.
* **Market**: A traditional Market order. A Market order will execute until filled or your bankruptcy price is reached, at
  which point it will cancel.
* **MarketWithLeftOverAsLimit**: A market order that, after eating through the order book as far as
  permitted by available margin, will become a limit order. The difference between this type and `Market` only
  affects the behavior in thin books. Upon reaching the deepest possible price, if there is quantity left over,
  a `Market` order will cancel the remaining quantity. `MarketWithLeftOverAsLimit` will keep the remaining
  quantity in the books as a `Limit`.
* **Stop**: A Stop Market order. Specify an `orderQty` and `stopPx`. When the `stopPx` is reached, the order will be entered
  into the book.
  * On sell orders, the order will trigger if the triggering price is lower than the `stopPx`. On buys, higher.
  * Note: Stop orders do not consume margin until triggered. Be sure that the required margin is available in your
    account so that it may trigger fully.
  * `Close` Stops don't require an `orderQty`. See Execution Instructions below.
* **StopLimit**: Like a Stop Market, but enters a Limit order instead of a Market order. Specify an `orderQty`, `stopPx`,
  and `price`.
* **MarketIfTouched**: Similar to a Stop, but triggers are done in the opposite direction. Useful for Take Profit orders.
* **LimitIfTouched**: As above; use for Take Profit Limit orders.

#### Execution Instructions

The following `execInst`s are supported. If using multiple, separate with a comma (e.g. `LastPrice,Close`).

* **ParticipateDoNotInitiate**: Also known as a Post-Only order. If this order would have executed on placement,
  it will cancel instead.
* **MarkPrice, LastPrice, IndexPrice**: Used by stop and if-touched orders to determine the triggering price.
  Use only one. By default, `'MarkPrice'` is used. Also used for Pegged orders to define the value of `'LastPeg'`.
* **ReduceOnly**: A `'ReduceOnly'` order can only reduce your position, not increase it. If you have a `'ReduceOnly'`
  limit order that rests in the order book while the position is reduced by other orders, then its order quantity will
  be amended down or canceled. If there are multiple `'ReduceOnly'` orders the least aggressive will be amended first.
* **Close**: `'Close'` implies `'ReduceOnly'`. A `'Close'` order will cancel other active limit orders with the same side
  and symbol if the open quantity exceeds the current position. This is useful for stops: by canceling these orders, a
  `'Close'` Stop is ensured to have the margin required to execute, and can only execute up to the full size of your
  position. If `orderQty` is not specified, a `'Close'` order has an `orderQty` equal to your current position's size.
  * Note that a `Close` order without an `orderQty` requires a `side`, so that BitMEX knows if it should trigger
  above or below the `stopPx`.

#### Linked Orders

Linked Orders are an advanced capability. It is very powerful, but its use requires careful coding and testing.
Please follow this document carefully and use the [Testnet Exchange](https://testnet.bitmex.com) while developing.

BitMEX offers four advanced Linked Order types:

* **OCO**: *One Cancels the Other*. A very flexible version of the standard Stop / Take Profit technique.
  Multiple orders may be linked together using a single `clOrdLinkID`. Send a `contingencyType` of
  `OneCancelsTheOther` on the orders. The first order that fully or partially executes (or activates
  for `Stop` orders) will cancel all other orders with the same `clOrdLinkID`.
* **OTO**: *One Triggers the Other*. Send a `contingencyType` of `'OneTriggersTheOther'` on the primary order and
  then subsequent orders with the same `clOrdLinkID` will be not be triggered until the primary order fully executes.
* **OUOA**: *One Updates the Other Absolute*. Send a `contingencyType` of `'OneUpdatesTheOtherAbsolute'` on the orders. Then
  as one order has a execution, other orders with the same `clOrdLinkID` will have their order quantity amended
  down by the execution quantity.
* **OUOP**: *One Updates the Other Proportional*. Send a `contingencyType` of `'OneUpdatesTheOtherProportional'` on the orders. Then
  as one order has a execution, other orders with the same `clOrdLinkID` will have their order quantity reduced proportionally
  by the fill percentage.

#### Trailing Stops

You may use `pegPriceType` of `'TrailingStopPeg'` to create Trailing Stops. The pegged `stopPx` will move as the market
moves away from the peg, and freeze as the market moves toward it.

To use, combine with `pegOffsetValue` to set the `stopPx` of your order. The peg is set to the triggering price
specified in the `execInst` (default `'MarkPrice'`). Use a negative offset for stop-sell and buy-if-touched orders.

Requires `ordType`: `'Stop', 'StopLimit', 'MarketIfTouched', 'LimitIfTouched'`.

#### Simple Quantities

Send a `simpleOrderQty` instead of an `orderQty` to create an order denominated in the underlying currency.
This is useful for opening up a position with 1 XBT of exposure without having to calculate how many contracts it is.

#### Rate Limits

See the [Bulk Order Documentation](#!/Order/Order_newBulk) if you need to place multiple orders at the same time.
Bulk orders require fewer risk checks in the trading engine and thus are ratelimited at **1/10** the normal rate.

You can also improve your reactivity to market movements while staying under your ratelimit by using the
[Amend](#!/Order/Order_amend) and [Amend Bulk](#!/Order/Order_amendBulk) endpoints. This allows you to stay
in the market and avoids the cancel/replace cycle.

#### Tracking Your Orders

If you want to keep track of order IDs yourself, set a unique `clOrdID` per order.
This `clOrdID` will come back as a property on the order and any related executions (including on the WebSocket),
and can be used to get or cancel the order. Max length is 36 characters.

You can also change the `clOrdID` by amending an order, supplying an `origClOrdID`, and your desired new
ID as the `clOrdID` param, like so:

```
# Amends an order's leavesQty, and updates its clOrdID to "def-456"
PUT /api/v1/order {"origClOrdID": "abc-123", "clOrdID": "def-456", "leavesQty": 1000}
```

*/
func (a *Client) OrderNew(params *OrderNewParams, authInfo runtime.ClientAuthInfoWriter) (*OrderNewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderNewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Order.new",
		Method:             "POST",
		PathPattern:        "/order",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrderNewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrderNewOK), nil

}

/*
OrderNewBulk creates multiple new orders for the same symbol

This endpoint is used for placing bulk orders. Valid order types are Market, Limit, Stop, StopLimit, MarketIfTouched, LimitIfTouched, MarketWithLeftOverAsLimit, and Pegged.

Each individual order object in the array should have the same properties as an individual POST /order call.

This endpoint is much faster for getting many orders into the book at once. Because it reduces load on BitMEX
systems, this endpoint is ratelimited at `ceil(0.1 * orders)`. Submitting 10 orders via a bulk order call
will only count as 1 request, 15 as 2, 32 as 4, and so on.

For now, only `application/json` is supported on this endpoint.

*/
func (a *Client) OrderNewBulk(params *OrderNewBulkParams, authInfo runtime.ClientAuthInfoWriter) (*OrderNewBulkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrderNewBulkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Order.newBulk",
		Method:             "POST",
		PathPattern:        "/order/bulk",
		ProducesMediaTypes: []string{"application/javascript", "application/json", "application/xml", "text/javascript", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrderNewBulkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrderNewBulkOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
