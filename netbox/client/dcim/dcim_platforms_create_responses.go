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

// DcimPlatformsCreateReader is a Reader for the DcimPlatformsCreate structure.
type DcimPlatformsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPlatformsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPlatformsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPlatformsCreateCreated creates a DcimPlatformsCreateCreated with default headers values
func NewDcimPlatformsCreateCreated() *DcimPlatformsCreateCreated {
	return &DcimPlatformsCreateCreated{}
}

/*DcimPlatformsCreateCreated handles this case with default header values.

DcimPlatformsCreateCreated dcim platforms create created
*/
type DcimPlatformsCreateCreated struct {
	Payload *models.Platform
}

func (o *DcimPlatformsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/platforms/][%d] dcimPlatformsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPlatformsCreateCreated) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPlatformsCreateDefault creates a DcimPlatformsCreateDefault with default headers values
func NewDcimPlatformsCreateDefault(code int) *DcimPlatformsCreateDefault {
	return &DcimPlatformsCreateDefault{
		_statusCode: code,
	}
}

/*DcimPlatformsCreateDefault handles this case with default header values.

DcimPlatformsCreateDefault dcim platforms create default
*/
type DcimPlatformsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim platforms create default response
func (o *DcimPlatformsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPlatformsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/platforms/][%d] dcim_platforms_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
