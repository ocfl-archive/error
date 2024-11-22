package error

// Type describes a constant for a known type of error.
type Type string

const (
	// TypeUnknownError is a constant type string for unidentified
	// errors, e.g. the error might be new and yet to be determined
	// within the caller's context.
	TypeUnknownError Type = "unknown"
)
