package error

const DefaultWeight = 100
const DefaultMessage = "An unexpected error occurred."

var Errors = map[ID]*Error{
	IDUnknownError: NewErrorStruct(IDUnknownError, TypeUnknownError, 100, "", "An unexpected error occurred."),
}

func NewError(id ID, additional string, err error) *Error {
	archiveErr, ok := Errors[id]
	if !ok {
		archiveErr = Errors[IDUnknownError]
		additional = string(id) + ": " + additional
	}
	return archiveErr.WithAdditional(additional, 2, err)
}
