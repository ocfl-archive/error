# error

Error management for archival workflows

## Examples

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
weight = 50
message = "Testing two for error"

[[errors]]
id = "TestError"
type = "unknown"
weight = 50
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
  weight: 50
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

## Developer guide

### justfile

`just` can be installed using cargo.

```sh
curl -sSf https://static.rust-lang.org/rustup.sh | sh
cargo install just
```

> NB. rustup.sh will install rust and cargo at the same time. It is the most
convenient way to keep cargo up to date. `rustup` can be run in future to do
that.

`just` is just like `make` but for modern development environments.

### Linting and formatting

#### GitHub CI

Continuous integration tasks are expected to pass before merging. Take a look
at `.github/workflows/` to see how they mirror the commands below.

#### CLI

Linting can be run on the command line. Run:

```sh
just linting
```

The majority of tooling is installed with the go standard library. The other
tools need to be installed manually as follows.

* goimports

```sh
go install golang.org/x/tools/cmd/goimports@latest
```

* staticcheck

```sh
go install honnef.co/go/tools/cmd/staticcheck@latest
```

You can also run:

```sh
just setup
```
