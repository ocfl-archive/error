package error

import (
	"fmt"
	"runtime"
	"testing"
)

func TestName(t *testing.T) {

	factory := NewFactory()
	if factory == nil {
		t.Errorf("factory is nil")
	}
	if factory.errors == nil {
		t.Errorf("factory.errors is nil")
	}
	if len(factory.errors) != 1 {
		t.Errorf("len(factory.errors) = %d, want 1", len(factory.errors))
	}
	if err := factory.RegisterError("TestError", TypeUnknownError, 50, "Testing for error"); err != nil {
		t.Errorf("factory.RegisterError() failed: %v", err)
	}
	if len(factory.errors) != 2 {
		t.Errorf("len(factory.errors) = %d, want 2", len(factory.errors))
	}
	if err := factory.RegisterError(IDUnknownError, TypeUnknownError, 50, "Testing for error"); err == nil {
		t.Errorf("factory.RegisterError() should have failed")
	}
	testErr := factory.NewError("TestError", "additional")
	pc, file, line, ok := runtime.Caller(0)
	details := runtime.FuncForPC(pc)
	if !ok {
		t.Errorf("runtime.Caller(0) failed")
	}
	sourceFile := fmt.Sprintf("%s:%d", file, line-1)
	if testErr.ID != "TestError" {
		t.Errorf("testErr.ID = %s, want TestError", testErr.ID)
	}
	if testErr.Type != TypeUnknownError {
		t.Errorf("testErr.Type = %s, want %s", testErr.Type, TypeUnknownError)
	}
	if testErr.Weight != 50 {
		t.Errorf("testErr.Weight = %d, want 50", testErr.Weight)
	}
	if testErr.SourceFile != sourceFile {
		t.Errorf("testErr.SourceFile = %s, want %s", testErr.SourceFile, sourceFile)
	}
	if testErr.SourceFunc != details.Name() {
		t.Errorf("testErr.SourceFunc = %s, want %s", testErr.SourceFunc, details.Name())
	}
	if testErr.Message != "Testing for error" {
		t.Errorf("testErr.Message = %s, want 'Testing for error'", testErr.Message)
	}
	if testErr.Additional != "additional" {
		t.Errorf("testErr.Additional = %s, want 'additional'", testErr.Additional)
	}

	if err := factory.RegisterError("TestError2", TypeUnknownError, 50, "Testing two for error"); err != nil {
		t.Errorf("factory.RegisterError() failed: %v", err)
	}

	toml, err := factory.TOML()
	if err != nil {
		t.Errorf("factory.TOML() failed: %v", err)
	}
	yaml, err := factory.YAML()
	if err != nil {
		t.Errorf("factory.YAML() failed: %v", err)
	}
	fmt.Println(string(toml))
	fmt.Println(string(yaml))

	factory2 := NewFactory()
	errors2, err := LoadTOMLData(toml)
	if err != nil {
		t.Errorf("LoadTOMLData() failed: %v", err)
	}
	if err := factory2.RegisterErrors(errors2); err != nil {
		t.Errorf("factory2.RegisterErrors() failed: %v", err)
	}
	toml2, err := factory2.TOML()
	if err != nil {
		t.Errorf("factory2.TOML() failed: %v", err)
	}
	if string(toml) != string(toml2) {
		t.Errorf("toml != toml2")
	}
	factory3 := NewFactory()
	errors3, err := LoadYAMLData(yaml)
	if err != nil {
		t.Errorf("LoadYAMLData() failed: %v", err)
	}
	if err := factory3.RegisterErrors(errors3); err != nil {
		t.Errorf("factory3.RegisterErrors() failed: %v", err)
	}
	yaml3, err := factory3.YAML()
	if err != nil {
		t.Errorf("factory3.YAML() failed: %v", err)
	}
	if string(yaml) != string(yaml3) {
		t.Errorf("yaml != yaml3")
	}
}
