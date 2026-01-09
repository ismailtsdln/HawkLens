# HawkLens Development Setup

## Pre-commit Hooks

We recommend using `staticcheck` and `gofmt` for code quality.

```yaml
# .pre-commit-config.yaml
repos:
-   repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
    -   id: trailing-whitespace
    -   id: end-of-file-fixer
-   repo: https://github.com/golangci/golangci-lint
    rev: v1.52.2
    hooks:
    -   id: golangci-lint
```

## CI/CD Pipeline

The project includes a GitHub Actions workflow for automated testing.

```yaml
# .github/workflows/go.yml
name: Go CI

on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    - name: Build
      run: go build -v ./...
    - name: Test
      run: go test -v ./...
```
