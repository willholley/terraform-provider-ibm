// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostInstancesParamsBodyBootVolumeAttachment VolumeAttachmentTemplateinstanceContext
//
// The boot volume interface for the instance
// swagger:model postInstancesParamsBodyBootVolumeAttachment
type PostInstancesParamsBodyBootVolumeAttachment struct {

	// delete volume on instance delete
	DeleteVolumeOnInstanceDelete bool `json:"delete_volume_on_instance_delete,omitempty"`

	// The user-defined name for this volume interface
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// volume
	Volume *PostInstancesParamsBodyBootVolumeAttachmentVolume `json:"volume,omitempty"`
}

// Validate validates this post instances params body boot volume attachment
func (m *PostInstancesParamsBodyBootVolumeAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostInstancesParamsBodyBootVolumeAttachment) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PostInstancesParamsBodyBootVolumeAttachment) validateVolume(formats strfmt.Registry) error {

	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostInstancesParamsBodyBootVolumeAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostInstancesParamsBodyBootVolumeAttachment) UnmarshalBinary(b []byte) error {
	var res PostInstancesParamsBodyBootVolumeAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
