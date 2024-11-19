# error
Error management for archival workflows

### TOML & Zerolog

Playground: https://go.dev/play/p/5AZUWzxEgVk

```go
package main

import (
	archiveerror "github.com/ocfl-archive/error/pkg/error"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	tomlData := []byte(`
[[errors]]
id = "TestError2"
type = "unknown"
default_weight = 50
message = "Testing two for error"

[[errors]]
id = "TestError"
type = "unknown"
default_weight = 50
message = "Testing for error"
`)

	errs, err := archiveerror.LoadTOMLData(tomlData)
	if err != nil {
		panic(err)
	}
	errorFactory := archiveerror.NewFactory()
	if err := errorFactory.RegisterErrors(errs); err != nil {
		panic(err)
	}

	archiveError := errorFactory.NewError("TestError2", "additional data")
	if archiveError == nil {
		panic("error is nil")
	}

	log.Error().Any("archive", archiveError).Msg("An error occurred")
}
```

### YAML

Playground: https://go.dev/play/p/heFWPrPpYgv

`data/errors.yaml`
```yaml
- id: TestError
  type: unknown
  default_weight: 50
  message: Testing for error
- id: TestError2
  type: unknown
  default_weight: 50
  message: Testing two for error
``` 

```go
package main

import (
	"encoding/json"
	archiveerror "github.com/ocfl-archive/error/pkg/error"
	"os"
)

func main() {
	errs, err := archiveerror.LoadYAMLFile("data/errors.yaml")
	if err != nil {
		panic(err)
	}
	errorFactory := archiveerror.NewFactory()
	if err := errorFactory.RegisterErrors(errs); err != nil {
		panic(err)
	}

	archiveError := errorFactory.NewError("TestError2", "additional data")
	if archiveError == nil {
		panic("error is nil")
	}
	jsonBytes, err := json.MarshalIndent(archiveError, "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(jsonBytes))
}
```

