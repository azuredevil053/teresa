package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*AutoScale horizontal auto scale params

swagger:model AutoScale
*/
type AutoScale struct {

	/* target average CPU utilization (represented as a percentage of requested CPU) over all the pods

	Maximum: 100
	Minimum: 0
	*/
	CPUTargetUtilization *int64 `json:"cpuTargetUtilization,omitempty"`

	/* maximum number of PODs running the app

	Minimum: 1
	*/
	Max int64 `json:"max,omitempty"`

	/* minimum number of PODs running the app

	Minimum: 1
	*/
	Min int64 `json:"min,omitempty"`
}

// Validate validates this auto scale
func (m *AutoScale) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUTargetUtilization(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMax(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMin(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoScale) validateCPUTargetUtilization(formats strfmt.Registry) error {

	if swag.IsZero(m.CPUTargetUtilization) { // not required
		return nil
	}

	if err := validate.MinimumInt("cpuTargetUtilization", "body", int64(*m.CPUTargetUtilization), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("cpuTargetUtilization", "body", int64(*m.CPUTargetUtilization), 100, false); err != nil {
		return err
	}

	return nil
}

func (m *AutoScale) validateMax(formats strfmt.Registry) error {

	if swag.IsZero(m.Max) { // not required
		return nil
	}

	if err := validate.MinimumInt("max", "body", int64(m.Max), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *AutoScale) validateMin(formats strfmt.Registry) error {

	if swag.IsZero(m.Min) { // not required
		return nil
	}

	if err := validate.MinimumInt("min", "body", int64(m.Min), 1, false); err != nil {
		return err
	}

	return nil
}