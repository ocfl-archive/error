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
		Source:        source,
		Message:       message,
	}
}

type Error struct {
	ID            ID     `json:"id"`
	Type          Type   `json:"type"`
	DefaultWeight int64  `json:"default_weight"`
	Source        string `json:"source"`
	Message       string `json:"message"`
	Additional    string `json:"additional"`
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	return fmt.Sprintf("[%s] #%d (%s): %s - %s", e.Type, e.ID, e.Source, e.Message, e.Additional)
}

func (e *Error) WithAdditional(additional string) *Error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	source := fmt.Sprintf("%s:%d", file, line)
	return &Error{
		ID:            e.ID,
		Type:          e.Type,
		DefaultWeight: e.DefaultWeight,
		Source:        source,
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
