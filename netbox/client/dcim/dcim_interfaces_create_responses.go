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

// DcimInterfacesCreateReader is a Reader for the DcimInterfacesCreate structure.
type DcimInterfacesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimInterfacesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfacesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfacesCreateCreated creates a DcimInterfacesCreateCreated with default headers values
func NewDcimInterfacesCreateCreated() *DcimInterfacesCreateCreated {
	return &DcimInterfacesCreateCreated{}
}

/*DcimInterfacesCreateCreated handles this case with default header values.

DcimInterfacesCreateCreated dcim interfaces create created
*/
type DcimInterfacesCreateCreated struct {
	Payload *models.Interface
}

func (o *DcimInterfacesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/interfaces/][%d] dcimInterfacesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimInterfacesCreateCreated) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfacesCreateDefault creates a DcimInterfacesCreateDefault with default headers values
func NewDcimInterfacesCreateDefault(code int) *DcimInterfacesCreateDefault {
	return &DcimInterfacesCreateDefault{
		_statusCode: code,
	}
}

/*DcimInterfacesCreateDefault handles this case with default header values.

DcimInterfacesCreateDefault dcim interfaces create default
*/
type DcimInterfacesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interfaces create default response
func (o *DcimInterfacesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInterfacesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/interfaces/][%d] dcim_interfaces_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfacesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
