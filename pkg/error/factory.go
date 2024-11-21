// Package error provides an implementation of an errors factory for
// consistent structuring of errors across multiple libraries that
// invoke logging using this package.

package error

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

// NewFactory Initializes a new error factory with a single
// IDUnknownError.
func NewFactory() *Factory {
	return &Factory{
		errors: map[ID]*Error{IDUnknownError: &Error{
			ID:      IDUnknownError,
			Type:    TypeUnknownError,
			Weight:  DefaultWeight,
			Message: DefaultMessage,
		},
		},
	}
}

// Factory object for recording all structured error objects for an
// application using this module.
type Factory struct {
	errors map[ID]*Error
}

// ExportGOConstants will export constants registered with this module.
// If no constants have been registered then a single IDUnknown error
// will be returned in the output requested by the caller.
func (f *Factory) ExportGOConstants() string {
	var result = "const (\n"
	for _, err := range f.errors {
		result += fmt.Sprintf(`ERROR%s = "%s"`+"\n", err.ID, err.ID)
	}
	return result + ")\n"
}

// RegisterError provides a way of registering a single correctly
// formatted errory with the factory object.
func (f *Factory) RegisterError(id ID, t Type, defaultWeight int64, message string) error {
	if _, ok := f.errors[id]; ok {
		return fmt.Errorf("error with id %s already exists", id)
	}
	f.errors[id] = &Error{
		ID:      id,
		Type:    t,
		Weight:  defaultWeight,
		Message: message,
	}
	return nil
}

// RegisterErrors records a slice of properly structured errors with
// the factory object returning nil on success and error on failure.
func (f *Factory) RegisterErrors(errors []*Error) error {
	for _, newErr := range errors {
		if err := f.RegisterError(newErr.ID, newErr.Type, newErr.Weight, newErr.Message); err != nil {
			return err
		}
	}
	return nil
}

// NewError retrieves an error from the factory. If the error is known
// its additional context appended and returned to the caller. If the
// erorr is unknown it is returned with additional context appended and
// ID unknown.
func (f *Factory) NewError(id ID, additional string, err error) *Error {
	archiveErr, ok := f.errors[id]
	if !ok {
		archiveErr = f.errors[IDUnknownError]
		additional = string(id) + ": " + additional
	}
	return archiveErr.WithAdditional(additional, 2, err)
}

// TOML will return a byte array to the caller containing all
// registered errors with the exception of IDUnknownError.
func (f *Factory) TOML() ([]byte, error) {
	var errs []*Error
	for _, err := range f.errors {
		if err.ID == IDUnknownError {
			continue
		}
		errs = append(errs, err)
	}
	return toml.Marshal(_tomlErrors{Errors: errs})
}

// YAML will return a byte array to the caller containing all
// registered errors with the exception of IDUnknownError.
func (f *Factory) YAML() ([]byte, error) {
	var errs []*Error
	for _, err := range f.errors {
		if err.ID == IDUnknownError {
			continue
		}
		errs = append(errs, err)
	}
	return yaml.Marshal(errs)
}
