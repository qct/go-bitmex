// Code generated by go-swagger; DO NOT EDIT.

package position

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/qct/go-bitmex/models"
)

// PositionGetReader is a Reader for the PositionGet structure.
type PositionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PositionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPositionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPositionGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPositionGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPositionGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPositionGetOK creates a PositionGetOK with default headers values
func NewPositionGetOK() *PositionGetOK {
	return &PositionGetOK{}
}

/*PositionGetOK handles this case with default header values.

Request was successful
*/
type PositionGetOK struct {
	Payload []*models.Position
}

func (o *PositionGetOK) Error() string {
	return fmt.Sprintf("[GET /position][%d] positionGetOK  %+v", 200, o.Payload)
}

func (o *PositionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPositionGetBadRequest creates a PositionGetBadRequest with default headers values
func NewPositionGetBadRequest() *PositionGetBadRequest {
	return &PositionGetBadRequest{}
}

/*PositionGetBadRequest handles this case with default header values.

Parameter Error
*/
type PositionGetBadRequest struct {
	Payload *models.Error
}

func (o *PositionGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /position][%d] positionGetBadRequest  %+v", 400, o.Payload)
}

func (o *PositionGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPositionGetUnauthorized creates a PositionGetUnauthorized with default headers values
func NewPositionGetUnauthorized() *PositionGetUnauthorized {
	return &PositionGetUnauthorized{}
}

/*PositionGetUnauthorized handles this case with default header values.

Unauthorized
*/
type PositionGetUnauthorized struct {
	Payload *models.Error
}

func (o *PositionGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /position][%d] positionGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PositionGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPositionGetNotFound creates a PositionGetNotFound with default headers values
func NewPositionGetNotFound() *PositionGetNotFound {
	return &PositionGetNotFound{}
}

/*PositionGetNotFound handles this case with default header values.

Not Found
*/
type PositionGetNotFound struct {
	Payload *models.Error
}

func (o *PositionGetNotFound) Error() string {
	return fmt.Sprintf("[GET /position][%d] positionGetNotFound  %+v", 404, o.Payload)
}

func (o *PositionGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
