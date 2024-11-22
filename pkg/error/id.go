package error

type ID string

const (
	// IDUnknownError is a constant identifier for unidentified errors,
	// e.g. the error might be new and yet to be determined within the
	// caller's context.
	IDUnknownError = "IDUnknownError"
)
