# error

Error management for archival workflows

## Examples

* Godoc: [NewFactory initialization and retrieval][docs-1].

[docs-1]: https://pkg.go.dev/github.com/ocfl-archive/error/pkg/error#NewFactory

### Example output

Given an example error described in YAML and registered with the error factory
on startup:

```yaml
- id: "GOCFL::1"
  type: "FileSystem"
  weight: 20
  message: "cannot access filesystem"
```

Output is structured as shown below.

#### Via the default logger

<!--markdownlint-disable-->

```log
2024/11/22 13:07:46 [FileSystem] #GOCFL::1 (/tmp/logtest/logtest.go:41): cannot access filesystem - (add)ing to GOCFL archive
```

<!--markdownlint-enable-->

#### Via zerolog

```json
{
  "level": "error",
  "archive": {
    "id": "GOCFL::1",
    "type": "FileSystem",
    "weight": 20,
    "source_file": "/tmp/logtest/logtest.go:39",
    "source_func": "main.main",
    "message": "cannot access filesystem",
    "additional": "(add)ing to GOCFL archive",
    "error_data": {
      "message": "disk is readonly",
      "stack": ""
    }
  },
  "time": 1732277218,
  "message": "An error occurred"
}
```

#### As JSON

```json
{
  "id": "GOCFL::1",
  "type": "FileSystem",
  "weight": 20,
  "source_file": "/tmp/logtest/logtest.go:36",
  "source_func": "main.main",
  "message": "cannot access filesystem",
  "additional": "(add)ing to GOCFL archive",
  "error_data": {
    "message": "disk is readonly",
    "stack": ""
  }
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
