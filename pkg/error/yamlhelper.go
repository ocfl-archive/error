package error

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"emperror.dev/errors"
	"gopkg.in/yaml.v3"
)

func LoadYAMLData(data []byte) ([]*Error, error) {
	var errs = []*Error{}
	if err := yaml.Unmarshal(data, &errs); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal yaml data")
	}
	return errs, nil
}

func LoadYAMLReader(r io.Reader) ([]*Error, error) {
	dec := yaml.NewDecoder(r)
	var errs = []*Error{}
	if err := dec.Decode(&errs); err != nil {
		return nil, errors.Wrapf(err, "failed to decode yaml data")
	}
	return errs, nil
}

func LoadYAMLFile(fp string) ([]*Error, error) {
	fp = filepath.Clean(fp)
	f, err := os.Open(fp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", fp)
	}
	defer f.Close()
	return LoadYAMLReader(f)
}

func LoadYAMLFileFS(fSys fs.FS, fp string) ([]*Error, error) {
	fp = filepath.Clean(fp)
	f, err := fSys.Open(fp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file %s", fp)
	}
	defer f.Close()
	return LoadYAMLReader(f)
}
