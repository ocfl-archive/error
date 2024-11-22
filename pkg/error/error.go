// Package error provides an implementation of an errors factory for
// consistent structuring of errors across multiple libraries that
// invoke logging using this package.

package error

const DefaultWeight = 100
const DefaultMessage = "An unexpected error occurred."

var Errors = map[ID]*Error{
	IDUnknownError: NewErrorStruct(IDUnknownError, TypeUnknownError, 100, "", "An unexpected error occurred."),
}

// NewError returns a baseline error to the caller that can be
// used to initialize or populate the error factory.
func NewError(id ID, additional string, err error) *Error {
	archiveErr, ok := Errors[id]
	if !ok {
		archiveErr = Errors[IDUnknownError]
		additional = string(id) + ": " + additional
	}
	return archiveErr.WithAdditional(additional, runtimeSkipModule, err)
}
