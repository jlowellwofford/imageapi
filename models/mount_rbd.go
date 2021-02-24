// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MountRbd mount_rbd describes an RBD mount.  This must have at least and RBD ID associated with it
// (which becomes the mount's ID), and a provided filesystem type.
//
// Either `rbd_id` or `rbd` must be specified.  If both are specified, `rbd` will be ignored.
//
// If `rbd` is specified and `rbd_id` is omitted, the RBD will first be attached, and will be
// detached on deletion.
//
//
// swagger:model mount_rbd
type MountRbd struct {

	// fs type
	// Required: true
	FsType *string `json:"fs_type"`

	// id
	// Read Only: true
	ID ID `json:"id,omitempty"`

	// mount options
	MountOptions []string `json:"mount_options"`

	// mountpoint
	// Read Only: true
	Mountpoint string `json:"mountpoint,omitempty"`

	// rbd
	Rbd *Rbd `json:"rbd,omitempty"`

	// rbd id
	RbdID ID `json:"rbd_id,omitempty"`

	// ref
	// Read Only: true
	Ref int64 `json:"ref,omitempty"`
}

// Validate validates this mount rbd
func (m *MountRbd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRbd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRbdID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountRbd) validateFsType(formats strfmt.Registry) error {

	if err := validate.Required("fs_type", "body", m.FsType); err != nil {
		return err
	}

	return nil
}

func (m *MountRbd) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *MountRbd) validateRbd(formats strfmt.Registry) error {
	if swag.IsZero(m.Rbd) { // not required
		return nil
	}

	if m.Rbd != nil {
		if err := m.Rbd.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rbd")
			}
			return err
		}
	}

	return nil
}

func (m *MountRbd) validateRbdID(formats strfmt.Registry) error {
	if swag.IsZero(m.RbdID) { // not required
		return nil
	}

	if err := m.RbdID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rbd_id")
		}
		return err
	}

	return nil
}

// ContextValidate validate this mount rbd based on the context it is used
func (m *MountRbd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMountpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRbd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRbdID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountRbd) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *MountRbd) contextValidateMountpoint(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mountpoint", "body", string(m.Mountpoint)); err != nil {
		return err
	}

	return nil
}

func (m *MountRbd) contextValidateRbd(ctx context.Context, formats strfmt.Registry) error {

	if m.Rbd != nil {
		if err := m.Rbd.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rbd")
			}
			return err
		}
	}

	return nil
}

func (m *MountRbd) contextValidateRbdID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RbdID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rbd_id")
		}
		return err
	}

	return nil
}

func (m *MountRbd) contextValidateRef(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ref", "body", int64(m.Ref)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MountRbd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MountRbd) UnmarshalBinary(b []byte) error {
	var res MountRbd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
