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
	"github.com/go-openapi/validate"
)

// Container The `container` option describes a minimally namespaced container.
//
// swagger:model container
type Container struct {

	// command
	// Required: true
	Command *string `json:"command"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// logfile
	// Read Only: true
	Logfile string `json:"logfile,omitempty"`

	// mount
	// Required: true
	Mount *Mount `json:"mount"`

	// A list of Linux namespaces to use.
	//
	// Note: This is currently unused.  All containers currently get `mnt` and `pid`.
	//       It's here as a placeholder for future use.
	//
	Namespaces []ContainerNamespace `json:"namespaces"`

	// When read, this contains the current container state. On creation, this requests the initial state (valid options: `created` or `running`). The default is `created`.
	//
	State ContainerState `json:"state,omitempty"`

	// When `systemd` is set to `true`, we will assume that this container will run `systemd`, and perform the necessary magic dance to make systemd run inside of the container. The default is `false`.
	//
	Systemd bool `json:"systemd,omitempty"`
}

// Validate validates this container
func (m *Container) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Container) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	return nil
}

func (m *Container) validateMount(formats strfmt.Registry) error {

	if err := validate.Required("mount", "body", m.Mount); err != nil {
		return err
	}

	if m.Mount != nil {
		if err := m.Mount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mount")
			}
			return err
		}
	}

	return nil
}

func (m *Container) validateNamespaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Namespaces); i++ {

		if err := m.Namespaces[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespaces" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Container) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// ContextValidate validate this container based on the context it is used
func (m *Container) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Container) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Container) contextValidateLogfile(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "logfile", "body", string(m.Logfile)); err != nil {
		return err
	}

	return nil
}

func (m *Container) contextValidateMount(ctx context.Context, formats strfmt.Registry) error {

	if m.Mount != nil {
		if err := m.Mount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mount")
			}
			return err
		}
	}

	return nil
}

func (m *Container) contextValidateNamespaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Namespaces); i++ {

		if err := m.Namespaces[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespaces" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Container) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Container) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Container) UnmarshalBinary(b []byte) error {
	var res Container
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}