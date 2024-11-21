package error

import (
	"io"
	"io/fs"
	"os"

	"emperror.dev/errors"
	"github.com/BurntSushi/toml"
)

type _tomlErrors struct {
	Errors []*Error `toml:"errors"`
}

// LoadTOMLData returns a slice of errors from an input byte array
// that can be used to to initialize an error factory.
func LoadTOMLData(data []byte) ([]*Error, error) {
	var errs = _tomlErrors{Errors: []*Error{}}
	if _, err := toml.Decode(string(data), &errs); err != nil {
		return nil, errors.Wrapf(err, "failed to decode toml data")
	}
	return errs.Errors, nil
}

// LoadTOMLReader returns a slice of errors from an input reade
// that can be used to to initialize an error factory.
func LoadTOMLReader(r io.Reader) ([]*Error, error) {
	var errs = _tomlErrors{Errors: []*Error{}}
	if _, err := toml.NewDecoder(r).Decode(&errs); err != nil {
		return nil, errors.Wrapf(err, "failed to decode toml data")
	}
	return errs.Errors, nil
}

// LoadTOMLFile returns a slice of errors from a filepath pointing at
// a valid errors toml file. The slice can be used to initialize an
// error factory.
func LoadTOMLFile(fp string) ([]*Error, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", fp)
	}
	defer f.Close()
	return LoadTOMLReader(f)
}

// LoadTOMLFile returns a slice of errors from a default filesystem.
// The slice can be used to initialize an error factory.
func LoadTOMLFileFS(fSys fs.FS, fp string) ([]*Error, error) {
	f, err := fSys.Open(fp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", fp)
	}
	defer f.Close()
	return LoadTOMLReader(f)
}
