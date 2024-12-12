package error

import (
	"fmt"
	"runtime"

	"emperror.dev/emperror"
	"emperror.dev/errors"
)

// Stack frame constants determine how many stack frames to skip when
// accessing information about the current runtime.
const runtimeSkipInvalid = 0
const runtimeSkipDefault = 1
const runtimeSkipModule = 2
const runtimeSkipExternalCall = 3

// getErrorStacktrace is an internal function used to return the
// stack trace to the caller.
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

// NewErrorStruct returns an initialized factory error object to the
// caller.
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

// Error data provides a structured way to wrap additional error data.
type ErrorData struct {
	Message string `json:"message,omitempty" toml:"message,omitempty" yaml:"message,omitempty"`
	Stack   string `json:"stack,omitempty" toml:"stack,omitempty" yaml:"stack,omitempty"`
}

// Error is the underlying type to be initialized in the error factory
// with fields available for additional caller context.
type Error struct {
	// ID is an identifier unique to the caller's context. It is
	// initialized in the factory.
	ID ID `json:"id" toml:"id" yaml:"id"`
	// Type provides additional taxonomic information about an error
	// within the caller's context. It is initialized in the factory.
	Type Type `json:"type" toml:"type" yaml:"type"`
	// Weight provides the caller a way to determine the severity
	// of a weight in the error factory compared to the factory's
	// other errors. It is initialized in the factory.
	Weight int64 `json:"weight" toml:"weight" yaml:"weight"`
	// SourceFile describes the file from which the error was raised.
	SourceFile string `json:"source_file" toml:"-" yaml:"-"`
	// SourceFunc describes the function and line from which the error
	// was raised. It is initialized in the factory.
	SourceFunc string `json:"source_func" toml:"-" yaml:"-"`
	// Message describes the error in more detail and is initialized in
	// the factory.
	Message string `json:"message" toml:"message" yaml:"message"`
	// Additional provides caller specific additional information
	// about the error when it is retrieved and reported on.
	Additional string `json:"additional,omitempty" toml:"-" yaml:"-"`
	// ErrorData provides a mechanism to access wrapped error data.
	ErrorData *ErrorData `json:"error_data,omitempty" toml:"-" yaml:"-"`
}

// Error provides a formatted string representation of the error
// structure.
func (e *Error) Error() string {
	return e.String()
}

// String provides a formatted string representation of the erroor
// structure.
func (e *Error) String() string {
	return fmt.Sprintf("[%s] #%s (%s): %s - %s", e.Type, e.ID, e.SourceFile, e.Message, e.Additional)
}

// WithAdditional builds an error structure with additional context
// provided by the caller.
func (e *Error) WithAdditional(additional string, skip int, err error) *Error {
	if skip <= runtimeSkipInvalid {
		skip = runtimeSkipModule
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

// Unwrap returns the underlying error from the factory if it exists.
func (e *Error) Unwrap() error {
	err, ok := Errors[e.ID]
	if !ok {
		return nil
	}
	return err
}

var _ error = &Error{}
var _ fmt.Stringer = &Error{}
