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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ObjectChangeAction Action
//
// swagger:model objectChangeAction
type ObjectChangeAction struct {

	// label
	// Required: true
	// Enum: [Created Updated Deleted]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [create update delete]
	Value *string `json:"value"`
}

// Validate validates this object change action
func (m *ObjectChangeAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var objectChangeActionTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Created","Updated","Deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectChangeActionTypeLabelPropEnum = append(objectChangeActionTypeLabelPropEnum, v)
	}
}

const (

	// ObjectChangeActionLabelCreated captures enum value "Created"
	ObjectChangeActionLabelCreated string = "Created"

	// ObjectChangeActionLabelUpdated captures enum value "Updated"
	ObjectChangeActionLabelUpdated string = "Updated"

	// ObjectChangeActionLabelDeleted captures enum value "Deleted"
	ObjectChangeActionLabelDeleted string = "Deleted"
)

// prop value enum
func (m *ObjectChangeAction) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, objectChangeActionTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ObjectChangeAction) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var objectChangeActionTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["create","update","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectChangeActionTypeValuePropEnum = append(objectChangeActionTypeValuePropEnum, v)
	}
}

const (

	// ObjectChangeActionValueCreate captures enum value "create"
	ObjectChangeActionValueCreate string = "create"

	// ObjectChangeActionValueUpdate captures enum value "update"
	ObjectChangeActionValueUpdate string = "update"

	// ObjectChangeActionValueDelete captures enum value "delete"
	ObjectChangeActionValueDelete string = "delete"
)

// prop value enum
func (m *ObjectChangeAction) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, objectChangeActionTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ObjectChangeAction) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectChangeAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectChangeAction) UnmarshalBinary(b []byte) error {
	var res ObjectChangeAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}