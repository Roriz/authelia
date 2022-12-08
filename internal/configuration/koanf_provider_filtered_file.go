package configuration

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// FilteredFile implements a koanf.Provider.
type FilteredFile struct {
	path    string
	filters []FileFilterFunc
}

// FilteredFileProvider returns a koanf.Provider which provides filtered file output.
func FilteredFileProvider(path string, filters ...FileFilterFunc) *FilteredFile {
	return &FilteredFile{
		path:    filepath.Clean(path),
		filters: filters,
	}
}

// ReadBytes reads the contents of a file on disk, passes it through any configured filters, and returns the bytes.
func (f *FilteredFile) ReadBytes() (data []byte, err error) {
	if data, err = os.ReadFile(f.path); err != nil {
		return nil, err
	}

	if len(data) == 0 || len(f.filters) == 0 {
		return data, nil
	}

	for _, filter := range f.filters {
		if data, err = filter(data); err != nil {
			return nil, err
		}
	}

	return data, nil
}

// Read is not supported by the filtered file koanf.Provider.
func (f *FilteredFile) Read() (map[string]interface{}, error) {
	return nil, errors.New("filtered file provider does not support this method")
}

// FileFilterFunc describes a func used to filter files.
type FileFilterFunc func(in []byte) (out []byte, err error)

// NewExpandEnvFileFilter is a FileFilterFunc which passes the bytes through os.ExpandEnv.
func NewExpandEnvFileFilter() FileFilterFunc {
	return func(in []byte) (out []byte, err error) {
		return []byte(os.ExpandEnv(string(in))), nil
	}
}

// NewTemplateFileFilter is a FileFilterFunc which passes the bytes through text/template.
func NewTemplateFileFilter() FileFilterFunc {
	data := &templated{
		Env: map[string]string{},
	}

	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)

		if len(kv) != 2 {
			continue
		}

		data.Env[kv[0]] = kv[1]
	}

	var t *template.Template

	t = template.New("config.template").Funcs(map[string]any{
		"env": newTemplatedEnvFunc(data),
	})

	return func(in []byte) (out []byte, err error) {
		if t, err = t.Parse(string(in)); err != nil {
			return nil, err
		}

		buf := &bytes.Buffer{}

		if err = t.Execute(buf, data); err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}
}

type templated struct {
	Env map[string]string
}

func newTemplatedEnvFunc(data *templated) func(key string) (value string, err error) {
	return func(key string) (value string, err error) {
		var ok bool

		if value, ok = data.Env[key]; !ok {
			return "", fmt.Errorf("environment variable %s does not exist", key)
		}

		return value, nil
	}
}
