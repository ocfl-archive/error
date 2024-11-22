// Package error provides an implementation of an errors factory for
// consistent structuring of errors across multiple libraries that
// invoke logging using this package.

package error

// DefaultWeight provides a default for weighting errors.
const DefaultWeight = 100

// DefaultMessage is provided for errors unknown to the factory.
const DefaultMessage = "An unexpected error occurred."

// Errors map persists the errors in the factory.
var Errors = map[ID]*Error{
	IDUnknownError: NewErrorStruct(
		IDUnknownError,
		TypeUnknownError,
		DefaultWeight,
		"",
		"An unexpected error occurred.",
	),
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
