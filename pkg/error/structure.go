package error

import (
	"fmt"
	"runtime"
)

func NewErrorStruct(id ID, t Type, weight int64, source, message string) *Error {
	return &Error{
		ID:         id,
		Type:       t,
		Weight:     weight,
		SourceFile: source,
		Message:    message,
	}
}

type Error struct {
	ID         ID     `json:"id" toml:"id" yaml:"id"`
	Type       Type   `json:"type" toml:"type" yaml:"type"`
	Weight     int64  `json:"weight" toml:"weight" yaml:"weight"`
	SourceFile string `json:"source_file" toml:"-" yaml:"-"`
	SourceFunc string `json:"source_func" toml:"-" yaml:"-"`
	Message    string `json:"message" toml:"message" yaml:"message"`
	Additional string `json:"additional" toml:"-" yaml:"-"`
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	return fmt.Sprintf("[%s] #%s (%s): %s - %s", e.Type, e.ID, e.SourceFile, e.Message, e.Additional)
}

func (e *Error) WithAdditional(additional string, skip int) *Error {
	var funcName string
	if skip <= 0 {
		skip = 1
	}
	pc, file, line, ok := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)
	if !ok {
		file = "???"
		line = 0
		funcName = "???"
	} else if details != nil {
		funcName = details.Name()
	}
	source := fmt.Sprintf("%s:%d", file, line)
	return &Error{
		ID:         e.ID,
		Type:       e.Type,
		Weight:     e.Weight,
		SourceFile: source,
		SourceFunc: funcName,
		Message:    e.Message,
		Additional: additional,
	}
}

func (e *Error) Unwrap() error {
	err, ok := Errors[e.ID]
	if !ok {
		return nil
	}
	return err
}

var _ error = &Error{}
var _ fmt.Stringer = &Error{}
