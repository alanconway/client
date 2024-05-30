// Copyright: This file is part of korrel8r, released under https://github.com/korrel8r/korrel8r/blob/main/LICENSE

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/korrel8r/client/pkg/swagger/models"
)

// GetDomainsDomainClassesReader is a Reader for the GetDomainsDomainClasses structure.
type GetDomainsDomainClassesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainsDomainClassesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomainsDomainClassesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDomainsDomainClassesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDomainsDomainClassesOK creates a GetDomainsDomainClassesOK with default headers values
func NewGetDomainsDomainClassesOK() *GetDomainsDomainClassesOK {
	return &GetDomainsDomainClassesOK{}
}

/*
GetDomainsDomainClassesOK describes a response with status code 200, with default header values.

OK
*/
type GetDomainsDomainClassesOK struct {
	Payload models.Classes
}

// IsSuccess returns true when this get domains domain classes o k response has a 2xx status code
func (o *GetDomainsDomainClassesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get domains domain classes o k response has a 3xx status code
func (o *GetDomainsDomainClassesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domains domain classes o k response has a 4xx status code
func (o *GetDomainsDomainClassesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domains domain classes o k response has a 5xx status code
func (o *GetDomainsDomainClassesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get domains domain classes o k response a status code equal to that given
func (o *GetDomainsDomainClassesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get domains domain classes o k response
func (o *GetDomainsDomainClassesOK) Code() int {
	return 200
}

func (o *GetDomainsDomainClassesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains/{domain}/classes][%d] getDomainsDomainClassesOK %s", 200, payload)
}

func (o *GetDomainsDomainClassesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains/{domain}/classes][%d] getDomainsDomainClassesOK %s", 200, payload)
}

func (o *GetDomainsDomainClassesOK) GetPayload() models.Classes {
	return o.Payload
}

func (o *GetDomainsDomainClassesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainsDomainClassesDefault creates a GetDomainsDomainClassesDefault with default headers values
func NewGetDomainsDomainClassesDefault(code int) *GetDomainsDomainClassesDefault {
	return &GetDomainsDomainClassesDefault{
		_statusCode: code,
	}
}

/*
GetDomainsDomainClassesDefault describes a response with status code -1, with default header values.

GetDomainsDomainClassesDefault get domains domain classes default
*/
type GetDomainsDomainClassesDefault struct {
	_statusCode int

	Payload string
}

// IsSuccess returns true when this get domains domain classes default response has a 2xx status code
func (o *GetDomainsDomainClassesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get domains domain classes default response has a 3xx status code
func (o *GetDomainsDomainClassesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get domains domain classes default response has a 4xx status code
func (o *GetDomainsDomainClassesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get domains domain classes default response has a 5xx status code
func (o *GetDomainsDomainClassesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get domains domain classes default response a status code equal to that given
func (o *GetDomainsDomainClassesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get domains domain classes default response
func (o *GetDomainsDomainClassesDefault) Code() int {
	return o._statusCode
}

func (o *GetDomainsDomainClassesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains/{domain}/classes][%d] GetDomainsDomainClasses default %s", o._statusCode, payload)
}

func (o *GetDomainsDomainClassesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /domains/{domain}/classes][%d] GetDomainsDomainClasses default %s", o._statusCode, payload)
}

func (o *GetDomainsDomainClassesDefault) GetPayload() string {
	return o.Payload
}

func (o *GetDomainsDomainClassesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
