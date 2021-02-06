// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// RunMetricFormat  - UNSPECIFIED: Default value if not present.
//  - RAW: Display value as its raw format.
//  - PERCENTAGE: Display value in percentage format.
// swagger:model RunMetricFormat
type RunMetricFormat string

const (

	// RunMetricFormatUNSPECIFIED captures enum value "UNSPECIFIED"
	RunMetricFormatUNSPECIFIED RunMetricFormat = "UNSPECIFIED"

	// RunMetricFormatRAW captures enum value "RAW"
	RunMetricFormatRAW RunMetricFormat = "RAW"

	// RunMetricFormatPERCENTAGE captures enum value "PERCENTAGE"
	RunMetricFormatPERCENTAGE RunMetricFormat = "PERCENTAGE"
)

// for schema
var runMetricFormatEnum []interface{}

func init() {
	var res []RunMetricFormat
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","RAW","PERCENTAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		runMetricFormatEnum = append(runMetricFormatEnum, v)
	}
}

func (m RunMetricFormat) validateRunMetricFormatEnum(path, location string, value RunMetricFormat) error {
	if err := validate.Enum(path, location, value, runMetricFormatEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this run metric format
func (m RunMetricFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRunMetricFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
