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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewVirtualizationClusterTypesUpdateParams creates a new VirtualizationClusterTypesUpdateParams object
// with the default values initialized.
func NewVirtualizationClusterTypesUpdateParams() *VirtualizationClusterTypesUpdateParams {
	var ()
	return &VirtualizationClusterTypesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterTypesUpdateParamsWithTimeout creates a new VirtualizationClusterTypesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClusterTypesUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationClusterTypesUpdateParams {
	var ()
	return &VirtualizationClusterTypesUpdateParams{

		timeout: timeout,
	}
}

// NewVirtualizationClusterTypesUpdateParamsWithContext creates a new VirtualizationClusterTypesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClusterTypesUpdateParamsWithContext(ctx context.Context) *VirtualizationClusterTypesUpdateParams {
	var ()
	return &VirtualizationClusterTypesUpdateParams{

		Context: ctx,
	}
}

// NewVirtualizationClusterTypesUpdateParamsWithHTTPClient creates a new VirtualizationClusterTypesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClusterTypesUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationClusterTypesUpdateParams {
	var ()
	return &VirtualizationClusterTypesUpdateParams{
		HTTPClient: client,
	}
}

/*VirtualizationClusterTypesUpdateParams contains all the parameters to send to the API endpoint
for the virtualization cluster types update operation typically these are written to a http.Request
*/
type VirtualizationClusterTypesUpdateParams struct {

	/*Data*/
	Data *models.ClusterType
	/*ID
	  A unique integer value identifying this cluster type.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization cluster types update params
func (o *VirtualizationClusterTypesUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationClusterTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster types update params
func (o *VirtualizationClusterTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster types update params
func (o *VirtualizationClusterTypesUpdateParams) WithContext(ctx context.Context) *VirtualizationClusterTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster types update params
func (o *VirtualizationClusterTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster types update params
func (o *VirtualizationClusterTypesUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationClusterTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster types update params
func (o *VirtualizationClusterTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization cluster types update params
func (o *VirtualizationClusterTypesUpdateParams) WithData(data *models.ClusterType) *VirtualizationClusterTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization cluster types update params
func (o *VirtualizationClusterTypesUpdateParams) SetData(data *models.ClusterType) {
	o.Data = data
}

// WithID adds the id to the virtualization cluster types update params
func (o *VirtualizationClusterTypesUpdateParams) WithID(id int64) *VirtualizationClusterTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization cluster types update params
func (o *VirtualizationClusterTypesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
