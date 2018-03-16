// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ResizeOptions options of resizing container tty size
// swagger:model ResizeOptions

type ResizeOptions struct {

	// height
	Height int64 `json:"Height,omitempty"`

	// width
	Width int64 `json:"Width,omitempty"`
}

/* polymorph ResizeOptions Height false */

/* polymorph ResizeOptions Width false */

// Validate validates this resize options
func (m *ResizeOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ResizeOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResizeOptions) UnmarshalBinary(b []byte) error {
	var res ResizeOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}