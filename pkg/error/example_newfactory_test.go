package error_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ocfl-archive/error/pkg/error"
)

func ExampleNewFactory_toml() {

	// Example read from TOML register and output to stderr log.

	seedErrorsToml := []byte(`
[[errors]]
id = "GOCFL::1"
type = "FileSystem"
default_weight = 20
message = "cannot access filesystem"

[[errors]]
id = "INDEXER::1"
type = "Preservation"
default_weight = 100
message = "cannot determine content type for archival object"
`)

	errs, err := error.LoadTOMLData(seedErrorsToml)
	if err != nil {
		panic(err)
	}
	errorFactory := error.NewFactory("OCFLError")
	if err := errorFactory.RegisterErrors(errs); err != nil {
		panic(err)
	}

	factoryErr := errorFactory.NewError("xGOCFL::1", "(add)ing to GOCFL archive", fmt.Errorf("disk is readonly"))
	if factoryErr.ID == "IDUnknownError" {
		// error could not be retrieved from the factory.
		// we can handle it differently here, or fall through
		// to log it anyway.
	}

	log.Println(factoryErr)
}

func ExampleNewFactory_yaml() {

	// Example read from YAML register and output to stdout as JSON.

	seedErrorsYaml := []byte(`
- id: "GOCFL::1"
  type: "FileSystem"
  weight: 20
  message: "cannot access filesystem"
- id: "INDEXER::1"
  type: "Preservation"
  default_weight: 100
  message: "cannot determine content type for archival object"
`)

	errs, err := error.LoadYAMLData(seedErrorsYaml)
	if err != nil {
		panic(err)
	}

	errorFactory := error.NewFactory("OCFLError")
	if err := errorFactory.RegisterErrors(errs); err != nil {
		panic(err)
	}

	factoryErr := errorFactory.NewError("GOCFL::1", "(add)ing to GOCFL archive", fmt.Errorf("disk is readonly"))
	if factoryErr.ID == "IDUnknownError" {
		// error could not be retrieved from the factory.
		// we can handle it differently here, or fall through
		// to log it anyway.
	}

	// Output JSON to stdout.
	jsonBytes, err := json.MarshalIndent(factoryErr, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}
