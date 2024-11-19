package error

import (
	"fmt"
	"runtime"
	"testing"
)

func TestError(t *testing.T) {
	testError := NewError(IDUnknownError, "additional")
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
		t.Errorf("error.ID = %d, want %d", testError.ID, IDUnknownError)
	}
	if testError.Type != UnknownErrorType {
		t.Errorf("error.Type = %s, want %s", testError.Type, UnknownErrorType)
	}
	if testError.DefaultWeight != 100 {
		t.Errorf("error.DefaultWeight = %d, want 100", testError.DefaultWeight)
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
