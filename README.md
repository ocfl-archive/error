# error
Error management for archival workflows

### YAML

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

### TOML

```go
package main

import (
	"encoding/json"
	archiveerror "github.com/ocfl-archive/error/pkg/error"
	"os"
)

func main() {
	errs, err := archiveerror.LoadTOMLFileFS(os.DirFS("data"), "errors.toml")
	if err != nil {
		panic(err)
	}
	errorFactory2 := archiveerror.NewFactory()
	if err := errorFactory2.RegisterErrors(errs); err != nil {
		panic(err)
	}

	archiveError := errorFactory2.NewError("TestError2", "additional data")
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

