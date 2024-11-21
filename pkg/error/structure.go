package error

import (
	"fmt"
	"runtime"

	"emperror.dev/emperror"
	"emperror.dev/errors"
)

func getErrorStacktrace(err error) errors.StackTrace {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	var stack errors.StackTrace

	errors.UnwrapEach(err, func(err error) bool {
		e := emperror.ExposeStackTrace(err)
		st, ok := e.(stackTracer)
		if !ok {
			return true
		}

		stack = st.StackTrace()
		return true
	})

	if len(stack) > 2 {
		stack = stack[:len(stack)-2]
	}
	return stack
	// fmt.Printf("%+v", st[0:2]) // top two frames
}

func NewErrorStruct(id ID, t Type, weight int64, source, message string) *Error {
	return &Error{
		ID:         id,
		Type:       t,
		Weight:     weight,
		SourceFile: source,
		Message:    message,
		ErrorData:  &ErrorData{},
	}
}

type ErrorData struct {
	Message string `json:"message" toml:"message" yaml:"message"`
	Stack   string `json:"stack" toml:"stack" yaml:"stack"`
}

type Error struct {
	ID         ID         `json:"id" toml:"id" yaml:"id"`
	Type       Type       `json:"type" toml:"type" yaml:"type"`
	Weight     int64      `json:"weight" toml:"weight" yaml:"weight"`
	SourceFile string     `json:"source_file" toml:"-" yaml:"-"`
	SourceFunc string     `json:"source_func" toml:"-" yaml:"-"`
	Message    string     `json:"message" toml:"message" yaml:"message"`
	Additional string     `json:"additional,omitempty" toml:"-" yaml:"-"`
	ErrorData  *ErrorData `json:"error_data,omitempty" toml:"-" yaml:"-"`
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	return fmt.Sprintf("[%s] #%s (%s): %s - %s", e.Type, e.ID, e.SourceFile, e.Message, e.Additional)
}

func (e *Error) WithAdditional(additional string, skip int, err error) *Error {
	if skip <= 0 {
		skip = 1
	}
	var funcName string
	var errorData *ErrorData
	pc, file, line, ok := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)
	if !ok {
		file = "???"
		line = 0
		funcName = "???"
	} else {
		if details != nil {
			funcName = details.Name()
		}
	}
	if err != nil {
		stack := getErrorStacktrace(err)
		errorData = &ErrorData{
			Message: err.Error(),
			Stack:   fmt.Sprintf("%+v", stack),
		}
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
		ErrorData:  errorData,
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
