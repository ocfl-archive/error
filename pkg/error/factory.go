package error

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

func NewFactory() *Factory {
	return &Factory{
		errors: map[ID]*Error{IDUnknownError: &Error{
			ID:            IDUnknownError,
			Type:          TypeUnknownError,
			DefaultWeight: DefaultWeight,
			Message:       DefaultMessage,
		},
		},
	}
}

type Factory struct {
	errors map[ID]*Error
}

func (f *Factory) RegisterError(id ID, t Type, defaultWeight int64, message string) error {
	if _, ok := f.errors[id]; ok {
		return fmt.Errorf("error with id %s already exists", id)
	}
	f.errors[id] = &Error{
		ID:            id,
		Type:          t,
		DefaultWeight: defaultWeight,
		Message:       message,
	}
	return nil
}

func (f *Factory) RegisterErrors(errors []*Error) error {
	for _, newErr := range errors {
		if err := f.RegisterError(newErr.ID, newErr.Type, newErr.DefaultWeight, newErr.Message); err != nil {
			return err
		}
	}
	return nil
}

func (f *Factory) NewError(id ID, additional string) *Error {
	err, ok := f.errors[id]
	if !ok {
		err = f.errors[IDUnknownError]
	}
	return err.WithAdditional(additional, 2)
}

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
