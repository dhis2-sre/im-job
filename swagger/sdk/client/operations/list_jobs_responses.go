// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dhis2-sre/im-job/swagger/sdk/models"
)

// ListJobsReader is a Reader for the ListJobs structure.
type ListJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewListJobsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListJobsOK creates a ListJobsOK with default headers values
func NewListJobsOK() *ListJobsOK {
	return &ListJobsOK{}
}

/* ListJobsOK describes a response with status code 200, with default header values.

ListJobsOK list jobs o k
*/
type ListJobsOK struct {
	Payload *models.Job
}

func (o *ListJobsOK) Error() string {
	return fmt.Sprintf("[GET /jobs][%d] listJobsOK  %+v", 200, o.Payload)
}
func (o *ListJobsOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *ListJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobsUnauthorized creates a ListJobsUnauthorized with default headers values
func NewListJobsUnauthorized() *ListJobsUnauthorized {
	return &ListJobsUnauthorized{}
}

/* ListJobsUnauthorized describes a response with status code 401, with default header values.

ListJobsUnauthorized list jobs unauthorized
*/
type ListJobsUnauthorized struct {
}

func (o *ListJobsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /jobs][%d] listJobsUnauthorized ", 401)
}

func (o *ListJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListJobsForbidden creates a ListJobsForbidden with default headers values
func NewListJobsForbidden() *ListJobsForbidden {
	return &ListJobsForbidden{}
}

/* ListJobsForbidden describes a response with status code 403, with default header values.

ListJobsForbidden list jobs forbidden
*/
type ListJobsForbidden struct {
}

func (o *ListJobsForbidden) Error() string {
	return fmt.Sprintf("[GET /jobs][%d] listJobsForbidden ", 403)
}

func (o *ListJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListJobsUnsupportedMediaType creates a ListJobsUnsupportedMediaType with default headers values
func NewListJobsUnsupportedMediaType() *ListJobsUnsupportedMediaType {
	return &ListJobsUnsupportedMediaType{}
}

/* ListJobsUnsupportedMediaType describes a response with status code 415, with default header values.

ListJobsUnsupportedMediaType list jobs unsupported media type
*/
type ListJobsUnsupportedMediaType struct {
}

func (o *ListJobsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /jobs][%d] listJobsUnsupportedMediaType ", 415)
}

func (o *ListJobsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
