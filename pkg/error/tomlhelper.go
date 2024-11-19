package error

import (
	"emperror.dev/errors"
	"github.com/BurntSushi/toml"
	"io"
	"io/fs"
	"os"
)

type _tomlErrors struct {
	Errors []*Error `toml:"errors"`
}

func LoadTOMLData(data []byte) ([]*Error, error) {
	var errs = _tomlErrors{Errors: []*Error{}}
	if _, err := toml.Decode(string(data), &errs); err != nil {
		return nil, errors.Wrapf(err, "failed to decode toml data")
	}
	return errs.Errors, nil
}

func LoadTOMLReader(r io.Reader) ([]*Error, error) {
	var errs = _tomlErrors{Errors: []*Error{}}
	if _, err := toml.NewDecoder(r).Decode(&errs); err != nil {
		return nil, errors.Wrapf(err, "failed to decode toml data")
	}
	return errs.Errors, nil
}

func LoadTOMLFile(fp string) ([]*Error, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", fp)
	}
	defer f.Close()
	return LoadTOMLReader(f)
}

// LoadTOMLFileFS loads a toml file from a fs.FS
func LoadTOMLFileFS(fSys fs.FS, fp string) ([]*Error, error) {
	f, err := fSys.Open(fp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", fp)
	}
	defer f.Close()
	return LoadTOMLReader(f)
}
