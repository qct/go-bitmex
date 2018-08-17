// Code generated by go-swagger; DO NOT EDIT.

package execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/qct/go-bitmex/models"
)

// ExecutionGetTradeHistoryReader is a Reader for the ExecutionGetTradeHistory structure.
type ExecutionGetTradeHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecutionGetTradeHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExecutionGetTradeHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewExecutionGetTradeHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewExecutionGetTradeHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewExecutionGetTradeHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExecutionGetTradeHistoryOK creates a ExecutionGetTradeHistoryOK with default headers values
func NewExecutionGetTradeHistoryOK() *ExecutionGetTradeHistoryOK {
	return &ExecutionGetTradeHistoryOK{}
}

/*ExecutionGetTradeHistoryOK handles this case with default header values.

Request was successful
*/
type ExecutionGetTradeHistoryOK struct {
	Payload []*models.Execution
}

func (o *ExecutionGetTradeHistoryOK) Error() string {
	return fmt.Sprintf("[GET /execution/tradeHistory][%d] executionGetTradeHistoryOK  %+v", 200, o.Payload)
}

func (o *ExecutionGetTradeHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionGetTradeHistoryBadRequest creates a ExecutionGetTradeHistoryBadRequest with default headers values
func NewExecutionGetTradeHistoryBadRequest() *ExecutionGetTradeHistoryBadRequest {
	return &ExecutionGetTradeHistoryBadRequest{}
}

/*ExecutionGetTradeHistoryBadRequest handles this case with default header values.

Parameter Error
*/
type ExecutionGetTradeHistoryBadRequest struct {
	Payload *models.Error
}

func (o *ExecutionGetTradeHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /execution/tradeHistory][%d] executionGetTradeHistoryBadRequest  %+v", 400, o.Payload)
}

func (o *ExecutionGetTradeHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionGetTradeHistoryUnauthorized creates a ExecutionGetTradeHistoryUnauthorized with default headers values
func NewExecutionGetTradeHistoryUnauthorized() *ExecutionGetTradeHistoryUnauthorized {
	return &ExecutionGetTradeHistoryUnauthorized{}
}

/*ExecutionGetTradeHistoryUnauthorized handles this case with default header values.

Unauthorized
*/
type ExecutionGetTradeHistoryUnauthorized struct {
	Payload *models.Error
}

func (o *ExecutionGetTradeHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /execution/tradeHistory][%d] executionGetTradeHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *ExecutionGetTradeHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionGetTradeHistoryNotFound creates a ExecutionGetTradeHistoryNotFound with default headers values
func NewExecutionGetTradeHistoryNotFound() *ExecutionGetTradeHistoryNotFound {
	return &ExecutionGetTradeHistoryNotFound{}
}

/*ExecutionGetTradeHistoryNotFound handles this case with default header values.

Not Found
*/
type ExecutionGetTradeHistoryNotFound struct {
	Payload *models.Error
}

func (o *ExecutionGetTradeHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /execution/tradeHistory][%d] executionGetTradeHistoryNotFound  %+v", 404, o.Payload)
}

func (o *ExecutionGetTradeHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
