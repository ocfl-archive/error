package error

import (
	"emperror.dev/errors"
	"fmt"
	"runtime"
	"testing"
)

func TestError(t *testing.T) {
	var err = errors.New("test error")
	err = errors.Wrap(err, "wrap test error")
	testError := NewError(IDUnknownError, "additional", err)
	pc, file, line, ok := runtime.Caller(0)
	details := runtime.FuncForPC(pc)
	if !ok {
		t.Errorf("runtime.Caller(0) failed")
	}
	sourceFile := fmt.Sprintf("%s:%d", file, line-1)

	sourceFunc := details.Name()
	if testError == nil {
		t.Errorf("error is nil")
	}
	if testError.ID != IDUnknownError {
		t.Errorf("error.ID = %s, want %s", testError.ID, IDUnknownError)
	}
	if testError.Type != TypeUnknownError {
		t.Errorf("error.Type = %s, want %s", testError.Type, TypeUnknownError)
	}
	if testError.Weight != 100 {
		t.Errorf("error.Weight = %d, want 100", testError.Weight)
	}
	if testError.SourceFile != sourceFile {
		t.Errorf("error.SourceFile = %s, want %s", testError.SourceFile, sourceFile)
	}
	if testError.SourceFunc != sourceFunc {
		t.Errorf("error.SourceFunc = %s, want %s", testError.SourceFunc, sourceFunc)
	}
	if testError.Message != "An unexpected error occurred." {
		t.Errorf("error.Message = %s, want 'An unexpected error occurred.'", testError.Message)
	}
	if testError.Additional != "additional" {
		t.Errorf("error.Additional = %s, want 'additional'", testError.Additional)
	}

}
