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

// DcimPowerPortTemplatesCreateReader is a Reader for the DcimPowerPortTemplatesCreate structure.
type DcimPowerPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortTemplatesCreateCreated creates a DcimPowerPortTemplatesCreateCreated with default headers values
func NewDcimPowerPortTemplatesCreateCreated() *DcimPowerPortTemplatesCreateCreated {
	return &DcimPowerPortTemplatesCreateCreated{}
}

/*DcimPowerPortTemplatesCreateCreated handles this case with default header values.

DcimPowerPortTemplatesCreateCreated dcim power port templates create created
*/
type DcimPowerPortTemplatesCreateCreated struct {
	Payload *models.PowerPortTemplate
}

func (o *DcimPowerPortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-port-templates/][%d] dcimPowerPortTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerPortTemplatesCreateCreated) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortTemplatesCreateDefault creates a DcimPowerPortTemplatesCreateDefault with default headers values
func NewDcimPowerPortTemplatesCreateDefault(code int) *DcimPowerPortTemplatesCreateDefault {
	return &DcimPowerPortTemplatesCreateDefault{
		_statusCode: code,
	}
}

/*DcimPowerPortTemplatesCreateDefault handles this case with default header values.

DcimPowerPortTemplatesCreateDefault dcim power port templates create default
*/
type DcimPowerPortTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power port templates create default response
func (o *DcimPowerPortTemplatesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortTemplatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/power-port-templates/][%d] dcim_power-port-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
