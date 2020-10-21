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

// DcimRearPortTemplatesUpdateReader is a Reader for the DcimRearPortTemplatesUpdate structure.
type DcimRearPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortTemplatesUpdateOK creates a DcimRearPortTemplatesUpdateOK with default headers values
func NewDcimRearPortTemplatesUpdateOK() *DcimRearPortTemplatesUpdateOK {
	return &DcimRearPortTemplatesUpdateOK{}
}

/*DcimRearPortTemplatesUpdateOK handles this case with default header values.

DcimRearPortTemplatesUpdateOK dcim rear port templates update o k
*/
type DcimRearPortTemplatesUpdateOK struct {
	Payload *models.RearPortTemplate
}

func (o *DcimRearPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-port-templates/{id}/][%d] dcimRearPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortTemplatesUpdateOK) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortTemplatesUpdateDefault creates a DcimRearPortTemplatesUpdateDefault with default headers values
func NewDcimRearPortTemplatesUpdateDefault(code int) *DcimRearPortTemplatesUpdateDefault {
	return &DcimRearPortTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*DcimRearPortTemplatesUpdateDefault handles this case with default header values.

DcimRearPortTemplatesUpdateDefault dcim rear port templates update default
*/
type DcimRearPortTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rear port templates update default response
func (o *DcimRearPortTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRearPortTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/rear-port-templates/{id}/][%d] dcim_rear-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
