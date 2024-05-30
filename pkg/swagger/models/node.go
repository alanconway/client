// Copyright: This file is part of korrel8r, released under https://github.com/korrel8r/korrel8r/blob/main/LICENSE

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Node node
//
// swagger:model Node
type Node struct {

	// Class is the full class name in "DOMAIN:CLASS" form.
	// Example: domain:class
	Class string `json:"class,omitempty"`

	// Count of results found for this class, after de-duplication.
	Count int64 `json:"count,omitempty"`

	// Queries yielding results for this class.
	Queries []*QueryCount `json:"queries"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateQueries(formats strfmt.Registry) error {
	if swag.IsZero(m.Queries) { // not required
		return nil
	}

	for i := 0; i < len(m.Queries); i++ {
		if swag.IsZero(m.Queries[i]) { // not required
			continue
		}

		if m.Queries[i] != nil {
			if err := m.Queries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("queries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this node based on the context it is used
func (m *Node) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQueries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) contextValidateQueries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Queries); i++ {

		if m.Queries[i] != nil {

			if swag.IsZero(m.Queries[i]) { // not required
				return nil
			}

			if err := m.Queries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("queries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}