package error

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"emperror.dev/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Compare slices where the slice count is two.
func compareSlices(slice1 []*Error, slice2 []*Error) bool {
	for _, slice1value := range slice1 {
		if !reflect.DeepEqual(slice1value, slice2[0]) &&
			!reflect.DeepEqual(slice1value, slice2[1]) {
			fmt.Println("argh!")
			return false
		}
	}
	return true
}

func TestLogging(t *testing.T) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	var buf bytes.Buffer
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:     &buf,
		NoColor: true,
	})
	factory := NewFactory("OCFLError")
	if err := factory.RegisterError("Test", TypeUnknownError, 50, "Testing for error"); err != nil {
		t.Errorf("factory.RegisterError() failed: %v", err)
	}

	// Output an example log line.
	log.Info().Any(factory.LogError("Test", "additional", errors.New("Testing 123"))).Msg("hello world")

	// Rudimentary constants from log line. We could try and be more dynamic here
	// but it would require mocking time, and understanding different path behavior.
	const testStr1 = "INF hello world OCFLError={\"additional\":\"additional\",\"error_data\":{\"message\":\"Testing 123\",\"stack\""
	const testStr2 = "\"id\":\"Test\",\"message\":\"Testing for error\",\"source_file\""
	const testStr3 = "\"source_func\":\"github.com/ocfl-archive/error/pkg/error.TestLogging\",\"type\":\"unknown\",\"weight\":50}"

	if !strings.Contains(buf.String(), testStr1) &&
		!strings.Contains(buf.String(), testStr1) &&
		!strings.Contains(buf.String(), testStr1) {
		t.Errorf("log not output as expected: '%s'", buf.String())
	}
}

// TestFactoryInitAndRoundTrip ensures that data consistency is protected
// through different cycles of initializing the factory and roundtrip
// via export.
func TestFactoryInitAndRoundTrip(t *testing.T) {

	factory := NewFactory("OCFLError")
	if factory == nil {
		t.Errorf("factory is nil")
	}
	if factory.errors == nil {
		t.Errorf("factory.errors is nil")
	}
	if len(factory.errors) != 1 {
		t.Errorf("len(factory.errors) = %d, want 1", len(factory.errors))
	}
	if err := factory.RegisterError("Test", TypeUnknownError, 50, "Testing for error"); err != nil {
		t.Errorf("factory.RegisterError() failed: %v", err)
	}
	if len(factory.errors) != 2 {
		t.Errorf("len(factory.errors) = %d, want 2", len(factory.errors))
	}
	if err := factory.RegisterError(IDUnknownError, TypeUnknownError, 50, "Testing for error"); err == nil {
		t.Errorf("factory.RegisterError() should have failed")
	}
	err := errors.New("Test")
	err = errors.Wrap(err, "additional")

	testErr := factory.NewError("Test", "additional", err)
	pc, file, line, ok := runtime.Caller(0)
	details := runtime.FuncForPC(pc)
	if !ok {
		t.Errorf("runtime.Caller(0) failed")
	}

	sourceFile := fmt.Sprintf("%s:%d", file, line-1)
	if testErr.ID != "Test" {
		t.Errorf("testErr.ID = %s, want Test", testErr.ID)
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

	if err := factory.RegisterError("Test2", TypeUnknownError, 50, "Testing two for error"); err != nil {
		t.Errorf("factory.RegisterError() failed: %v", err)
	}

	// Roundtrip Factory data to TOML and back.

	// Export to TOML from the existing factory.
	toml1, err := factory.TOML()
	if err != nil {
		t.Errorf("factory.TOML() failed: %v", err)
	}

	// Create a new factory and create errors slice.
	factory2 := NewFactory("OCFLError")
	errors2, err := LoadTOMLData(toml1)
	if err != nil {
		t.Errorf("LoadTOMLData() failed: %v", err)
	}
	// Register errors with the factory.
	if err := factory2.RegisterErrors(errors2); err != nil {
		t.Errorf("factory2.RegisterErrors() failed: %v", err)
	}

	// Export to toml from our second factory.
	toml2, err := factory2.TOML()
	if err != nil {
		t.Errorf("factory2.TOML() failed: %v", err)
	}

	// Load the exported strings into new TOML structures to test for
	// equivalence. TOML doesn't have guraanteed order and so needs to
	// be compared structurally.
	toml1Compare, _ := LoadTOMLData(toml1)
	toml2Compare, _ := LoadTOMLData(toml2)

	if !compareSlices(toml1Compare, toml2Compare) {
		t.Errorf("toml1 != toml2:\ntoml1:\n%s,\ntoml2:\n%s)", toml1, toml2)
	}

	// Roundtrip Factory data to TOML and back.

	// Export to TOML from the existing factory.
	yaml1, err := factory.YAML()
	if err != nil {
		t.Errorf("factory.YAML() failed: %v", err)
	}

	// Create a new factory and create errors slice.
	factory3 := NewFactory("OCFLError")
	errors3, err := LoadYAMLData(yaml1)
	if err != nil {
		t.Errorf("LoadYAMLData() failed: %v", err)
	}

	// Register errors with the factory.
	if err := factory3.RegisterErrors(errors3); err != nil {
		t.Errorf("factory3.RegisterErrors() failed: %v", err)
	}

	// Export to toml from our second factory.
	yaml2, err := factory3.YAML()
	if err != nil {
		t.Errorf("factory3.YAML() failed: %v", err)
	}

	// Load the exported strings into new TOML structures to test for
	// equivalence. TOML doesn't have guraanteed order and so needs to
	// be compared structurally.
	yaml1Compare, _ := LoadYAMLData(yaml1)
	yaml2Compare, _ := LoadYAMLData(yaml2)

	if !compareSlices(yaml1Compare, yaml2Compare) {
		t.Errorf("yaml != yaml3:\nyaml1:\n%s,\nyaml2:\n%s)", yaml1, yaml2)
	}
}
