# gourmet

Write build scripts for your Go projects in Go

## Installation

```shell
brew install @tsukinoko-kun/tap/gourmet
```

```shell
go install github.com/tsukinoko-kun/gourmet@latest
```

## Usage

`gourmet` is a wrapper around the `go` command. You can use the same parameters as you would with the `go` command.

### Build

`gourmet` is equivalent to `go build .` in the root of your project.  
`gourmet ./cmd/myapp/main.go` is equivalent to `go build ./cmd/myapp/main.go`.

Order of execution:
1. `cmd/prebuild/`
2. `cmd/build/` or `go build`
3. `cmd/postbuild/`

### Run

`gourmet run` is equivalent to `go run .` in the root of your project.  
`gourmet run ./cmd/myapp/main.go` is equivalent to `go run ./cmd/myapp/main.go`.

Order of execution:
1. `cmd/prebuild/`
2. `cmd/prerun/`
3. `cmd/run/` or `go run`
4. `cmd/postrun/`

### Test

`gourmet test` is equivalent to `go test ./...` in the root of your project.  
`gourmet test ./pkg/mypkg` is equivalent to `go test ./pkg/mypkg`.

Order of execution:
1. `cmd/prebuild/`
2. `cmd/pretest/`
3. `go test`
4. `cmd/posttest/`
