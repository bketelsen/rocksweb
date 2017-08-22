package models

import (
	"encoding/json"

	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
)

// String is not required by pop and may be deleted
func (v Version) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Versions is not required by pop and may be deleted
type Versions []Version

// String is not required by pop and may be deleted
func (v Versions) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Validate gets run every time you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (v *Version) Validate() (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: v.Tag, Name: "Tag"},
		&validators.TimeIsPresent{Field: v.ReleasedOn, Name: "ReleasedOn"},
	), nil
}

// ValidateSave gets run every time you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (v *Version) ValidateSave() (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (v *Version) ValidateUpdate() (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
