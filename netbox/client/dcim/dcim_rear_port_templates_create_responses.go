// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimRearPortTemplatesCreateReader is a Reader for the DcimRearPortTemplatesCreate structure.
type DcimRearPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRearPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortTemplatesCreateCreated creates a DcimRearPortTemplatesCreateCreated with default headers values
func NewDcimRearPortTemplatesCreateCreated() *DcimRearPortTemplatesCreateCreated {
	return &DcimRearPortTemplatesCreateCreated{}
}

/*DcimRearPortTemplatesCreateCreated handles this case with default header values.

DcimRearPortTemplatesCreateCreated dcim rear port templates create created
*/
type DcimRearPortTemplatesCreateCreated struct {
	Payload *models.RearPortTemplate
}

func (o *DcimRearPortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rear-port-templates/][%d] dcimRearPortTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRearPortTemplatesCreateCreated) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortTemplatesCreateDefault creates a DcimRearPortTemplatesCreateDefault with default headers values
func NewDcimRearPortTemplatesCreateDefault(code int) *DcimRearPortTemplatesCreateDefault {
	return &DcimRearPortTemplatesCreateDefault{
		_statusCode: code,
	}
}

/*DcimRearPortTemplatesCreateDefault handles this case with default header values.

DcimRearPortTemplatesCreateDefault dcim rear port templates create default
*/
type DcimRearPortTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rear port templates create default response
func (o *DcimRearPortTemplatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortTemplatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/rear-port-templates/][%d] dcim_rear-port-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
