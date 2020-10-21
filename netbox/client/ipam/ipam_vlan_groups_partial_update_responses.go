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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamVlanGroupsPartialUpdateReader is a Reader for the IpamVlanGroupsPartialUpdate structure.
type IpamVlanGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsPartialUpdateOK creates a IpamVlanGroupsPartialUpdateOK with default headers values
func NewIpamVlanGroupsPartialUpdateOK() *IpamVlanGroupsPartialUpdateOK {
	return &IpamVlanGroupsPartialUpdateOK{}
}

/*IpamVlanGroupsPartialUpdateOK handles this case with default header values.

IpamVlanGroupsPartialUpdateOK ipam vlan groups partial update o k
*/
type IpamVlanGroupsPartialUpdateOK struct {
	Payload *models.VLANGroup
}

func (o *IpamVlanGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsPartialUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsPartialUpdateDefault creates a IpamVlanGroupsPartialUpdateDefault with default headers values
func NewIpamVlanGroupsPartialUpdateDefault(code int) *IpamVlanGroupsPartialUpdateDefault {
	return &IpamVlanGroupsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*IpamVlanGroupsPartialUpdateDefault handles this case with default header values.

IpamVlanGroupsPartialUpdateDefault ipam vlan groups partial update default
*/
type IpamVlanGroupsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlan groups partial update default response
func (o *IpamVlanGroupsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlanGroupsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlan-groups/{id}/][%d] ipam_vlan-groups_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
