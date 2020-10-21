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

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// SecretsSecretRolesUpdateReader is a Reader for the SecretsSecretRolesUpdate structure.
type SecretsSecretRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsSecretRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecretsSecretRolesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecretsSecretRolesUpdateOK creates a SecretsSecretRolesUpdateOK with default headers values
func NewSecretsSecretRolesUpdateOK() *SecretsSecretRolesUpdateOK {
	return &SecretsSecretRolesUpdateOK{}
}

/*SecretsSecretRolesUpdateOK handles this case with default header values.

SecretsSecretRolesUpdateOK secrets secret roles update o k
*/
type SecretsSecretRolesUpdateOK struct {
	Payload *models.SecretRole
}

func (o *SecretsSecretRolesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /secrets/secret-roles/{id}/][%d] secretsSecretRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *SecretsSecretRolesUpdateOK) GetPayload() *models.SecretRole {
	return o.Payload
}

func (o *SecretsSecretRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsSecretRolesUpdateDefault creates a SecretsSecretRolesUpdateDefault with default headers values
func NewSecretsSecretRolesUpdateDefault(code int) *SecretsSecretRolesUpdateDefault {
	return &SecretsSecretRolesUpdateDefault{
		_statusCode: code,
	}
}

/*SecretsSecretRolesUpdateDefault handles this case with default header values.

SecretsSecretRolesUpdateDefault secrets secret roles update default
*/
type SecretsSecretRolesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the secrets secret roles update default response
func (o *SecretsSecretRolesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *SecretsSecretRolesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /secrets/secret-roles/{id}/][%d] secrets_secret-roles_update default  %+v", o._statusCode, o.Payload)
}

func (o *SecretsSecretRolesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *SecretsSecretRolesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
