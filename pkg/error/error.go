package error

const DefaultWeight = 100
const DefaultMessage = "An unexpected error occurred."

var Errors = map[ID]*Error{
	IDUnknownError: NewErrorStruct(IDUnknownError, TypeUnknownError, 100, "", "An unexpected error occurred."),
}

func NewError(id ID, additional string) *Error {
	err, ok := Errors[id]
	if !ok {
		err = Errors[IDUnknownError]
	}
	return err.WithAdditional(additional, 2)
}
