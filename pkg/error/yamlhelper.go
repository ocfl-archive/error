package error

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"emperror.dev/errors"
	"gopkg.in/yaml.v3"
)

// LoadYAMLData returns a slice of errors from an input reade
// that can be used to to initialize an error factory.
func LoadYAMLData(data []byte) ([]*Error, error) {
	var errs = []*Error{}
	if err := yaml.Unmarshal(data, &errs); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal yaml data")
	}
	return errs, nil
}

// LoadYAMLReader returns a slice of errors from a filepath pointing at
// a valid errors toml file. The slice can be used to initialize an
// error factory.
func LoadYAMLReader(r io.Reader) ([]*Error, error) {
	dec := yaml.NewDecoder(r)
	var errs = []*Error{}
	if err := dec.Decode(&errs); err != nil {
		return nil, errors.Wrapf(err, "failed to decode yaml data")
	}
	return errs, nil
}

// LoadYAMLFile returns a slice of errors from a filepath pointing at
// a valid errors toml file. The slice can be used to initialize an
// error factory.
func LoadYAMLFile(fp string) ([]*Error, error) {
	fp = filepath.Clean(fp)
	f, err := os.Open(fp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", fp)
	}
	defer f.Close()
	return LoadYAMLReader(f)
}

// LoadYAMLFileFS returns a slice of errors from a default filesystem.
// The slice can be used to initialize an error factory.
func LoadYAMLFileFS(fSys fs.FS, fp string) ([]*Error, error) {
	fp = filepath.Clean(fp)
	f, err := fSys.Open(fp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", fp)
	}
	defer f.Close()
	return LoadYAMLReader(f)
}
