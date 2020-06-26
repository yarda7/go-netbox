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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualizationVirtualMachinesUpdateReader is a Reader for the VirtualizationVirtualMachinesUpdate structure.
type VirtualizationVirtualMachinesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualizationVirtualMachinesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesUpdateOK creates a VirtualizationVirtualMachinesUpdateOK with default headers values
func NewVirtualizationVirtualMachinesUpdateOK() *VirtualizationVirtualMachinesUpdateOK {
	return &VirtualizationVirtualMachinesUpdateOK{}
}

/*VirtualizationVirtualMachinesUpdateOK handles this case with default header values.

VirtualizationVirtualMachinesUpdateOK virtualization virtual machines update o k
*/
type VirtualizationVirtualMachinesUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

func (o *VirtualizationVirtualMachinesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationVirtualMachinesUpdateBadRequest creates a VirtualizationVirtualMachinesUpdateBadRequest with default headers values
func NewVirtualizationVirtualMachinesUpdateBadRequest() *VirtualizationVirtualMachinesUpdateBadRequest {
	return &VirtualizationVirtualMachinesUpdateBadRequest{}
}

/*VirtualizationVirtualMachinesUpdateBadRequest handles this case with default header values.

test
*/
type VirtualizationVirtualMachinesUpdateBadRequest struct {
	Payload *VirtualizationVirtualMachinesUpdateBadRequestBody
}

func (o *VirtualizationVirtualMachinesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateBadRequest) GetPayload() *VirtualizationVirtualMachinesUpdateBadRequestBody {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VirtualizationVirtualMachinesUpdateBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*VirtualizationVirtualMachinesUpdateBadRequestBody virtualization virtual machines update bad request body
swagger:model VirtualizationVirtualMachinesUpdateBadRequestBody
*/
type VirtualizationVirtualMachinesUpdateBadRequestBody struct {

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// virtualization virtual machines update bad request body additional properties
	VirtualizationVirtualMachinesUpdateBadRequestBodyAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (o *VirtualizationVirtualMachinesUpdateBadRequestBody) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// error message
		ErrorMessage string `json:"errorMessage,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv VirtualizationVirtualMachinesUpdateBadRequestBody

	rcv.ErrorMessage = stage1.ErrorMessage
	*o = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "errorMessage")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		o.VirtualizationVirtualMachinesUpdateBadRequestBodyAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (o VirtualizationVirtualMachinesUpdateBadRequestBody) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// error message
		ErrorMessage string `json:"errorMessage,omitempty"`
	}

	stage1.ErrorMessage = o.ErrorMessage

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(o.VirtualizationVirtualMachinesUpdateBadRequestBodyAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(o.VirtualizationVirtualMachinesUpdateBadRequestBodyAdditionalProperties)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this virtualization virtual machines update bad request body
func (o *VirtualizationVirtualMachinesUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VirtualizationVirtualMachinesUpdateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VirtualizationVirtualMachinesUpdateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res VirtualizationVirtualMachinesUpdateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
