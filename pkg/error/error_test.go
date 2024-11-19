package error

import (
	"testing"
)

func TestError(t *testing.T) {
	testError := NewError(IDUnknownError, "additional")
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
	if testError.SourceFile != "C:/daten/go/dev/error/pkg/error/error_test.go:9" {
		t.Errorf("error.SourceFile = %s, want 'C:/daten/go/dev/error/pkg/error/error_test.go:9'", testError.SourceFile)
	}
	if testError.SourceFunc != "github.com/ocfl-archive/error/pkg/error.TestError" {
		t.Errorf("error.SourceFunc = %s, want 'github.com/ocfl-archive/error/pkg/error.TestError'", testError.SourceFunc)
	}
	if testError.Message != "An unexpected error occurred." {
		t.Errorf("error.Message = %s, want 'An unexpected error occurred.'", testError.Message)
	}
	if testError.Additional != "additional" {
		t.Errorf("error.Additional = %s, want 'additional'", testError.Additional)
	}
}
