package error

import (
	"fmt"
	"runtime"
)

func NewErrorStruct(id ID, t Type, defaultWeight int64, source, message string) *Error {
	return &Error{
		ID:            id,
		Type:          t,
		DefaultWeight: defaultWeight,
		SourceFile:    source,
		Message:       message,
	}
}

type Error struct {
	ID            ID     `json:"id"`
	Type          Type   `json:"type"`
	DefaultWeight int64  `json:"default_weight"`
	SourceFile    string `json:"source_file"`
	SourceFunc    string `json:"source_func"`
	Message       string `json:"message"`
	Additional    string `json:"additional"`
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	return fmt.Sprintf("[%s] #%d (%s): %s - %s", e.Type, e.ID, e.SourceFile, e.Message, e.Additional)
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
		ID:            e.ID,
		Type:          e.Type,
		DefaultWeight: e.DefaultWeight,
		SourceFile:    source,
		SourceFunc:    funcName,
		Message:       e.Message,
		Additional:    additional,
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
