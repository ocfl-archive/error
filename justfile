# CLI helpers.

# help
help:
 @echo "Command line helpers for this project.\n"
 @just -l

# Run go linting
linting:
 - goimports -w .
 - go fmt ./...
 - go vet ./...
 - staticcheck ./...

# Run pre-commit
all-checks:
  pre-commit run --all-files

# Setup linting
setup:
  go install golang.org/x/tools/cmd/godoc@latest
  go install golang.org/x/tools/cmd/goimports@latest
  go install honnef.co/go/tools/cmd/staticcheck@latest

# Fix imports
fix-imports:
  goimports -w .

# Run tests
test:
  go test ./...

# Docs
docs:
  godoc -http 0.0.0.0:8000
