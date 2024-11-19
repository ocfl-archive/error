package error

var Errors = map[ID]*Error{
	IDUnknownError: NewErrorStruct(IDUnknownError, UnknownErrorType, 0, "", "An unexpected error occurred."),
}

func NewError(id ID, additional string) *Error {
	err, ok := Errors[id]
	if !ok {
		err = Errors[IDUnknownError]
	}
	return err.WithAdditional(additional)
}
