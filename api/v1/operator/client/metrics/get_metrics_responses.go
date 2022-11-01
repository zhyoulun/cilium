// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/operator/models"
)

// GetMetricsReader is a Reader for the GetMetrics structure.
type GetMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetMetricsFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMetricsOK creates a GetMetricsOK with default headers values
func NewGetMetricsOK() *GetMetricsOK {
	return &GetMetricsOK{}
}

/*
GetMetricsOK describes a response with status code 200, with default header values.

Success
*/
type GetMetricsOK struct {
	Payload []*models.Metric
}

// IsSuccess returns true when this get metrics o k response has a 2xx status code
func (o *GetMetricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get metrics o k response has a 3xx status code
func (o *GetMetricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metrics o k response has a 4xx status code
func (o *GetMetricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get metrics o k response has a 5xx status code
func (o *GetMetricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get metrics o k response a status code equal to that given
func (o *GetMetricsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMetricsOK) Error() string {
	return fmt.Sprintf("[GET /metrics/][%d] getMetricsOK  %+v", 200, o.Payload)
}

func (o *GetMetricsOK) String() string {
	return fmt.Sprintf("[GET /metrics/][%d] getMetricsOK  %+v", 200, o.Payload)
}

func (o *GetMetricsOK) GetPayload() []*models.Metric {
	return o.Payload
}

func (o *GetMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsFailed creates a GetMetricsFailed with default headers values
func NewGetMetricsFailed() *GetMetricsFailed {
	return &GetMetricsFailed{}
}

/*
GetMetricsFailed describes a response with status code 500, with default header values.

Metrics cannot be retrieved
*/
type GetMetricsFailed struct {
}

// IsSuccess returns true when this get metrics failed response has a 2xx status code
func (o *GetMetricsFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get metrics failed response has a 3xx status code
func (o *GetMetricsFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metrics failed response has a 4xx status code
func (o *GetMetricsFailed) IsClientError() bool {
	return false
}

// IsServerError returns true when this get metrics failed response has a 5xx status code
func (o *GetMetricsFailed) IsServerError() bool {
	return true
}

// IsCode returns true when this get metrics failed response a status code equal to that given
func (o *GetMetricsFailed) IsCode(code int) bool {
	return code == 500
}

func (o *GetMetricsFailed) Error() string {
	return fmt.Sprintf("[GET /metrics/][%d] getMetricsFailed ", 500)
}

func (o *GetMetricsFailed) String() string {
	return fmt.Sprintf("[GET /metrics/][%d] getMetricsFailed ", 500)
}

func (o *GetMetricsFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
