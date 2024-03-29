[![Go Reference](https://pkg.go.dev/badge/github.com/woven-planet/go-zserio.svg)](https://pkg.go.dev/github.com/woven-planet/go-zserio)
[![Build Status](https://github.com/woven-planet/go-zserio/actions/workflows/test.yaml/badge.svg)](https://github.com/woven-planet/go-zserio/actions/workflows/test.yaml)

# go-zserio

[zserio](https://github.com/ndsev/zserio) code generation in Go.

## Versioning

go-zserio is still in active development, and not completely stable yet. Until
the first stable release all versions will have a major version number of `0`.
The minor version number will be updated for (API) breaking changes.

## Usage

First you need to install the `go-zserio` command:

```shell
go install github.com/woven-planet/go-zserio/cmd/go-zserio@latest
```

With this command you can generate the Go interface files to read/write zserio:

```shell
go-zserio generate <zserio_directory> --rootpackage <root_package> --out <output_directory>
```

This command compiles the zserio files, and generates Go files to read and write
zserio encoded data. `zserio_directory` is the directory where the zserio
definitions are stored. Supported file extensions are `.zs` and `.zserio`.
`rootpackage` specifies the root package name of the generated Go files.
`output_directory` specifies the path where the Go files will be generated.

## Limitations

- Implicit length arrays are deprecated, and are not supported.
- The ternary operator only works when used in functions.
- Cyclic imports, while not recommended, are supported by zserio, but not supported by go-zserio.
- Offsets and indexed offsets are not implemented yet.
- Constraints are not implemented yet.
- The `bazel run //internal/parser:update` command to rerun the `antlr` parser does not work on Windows. See https://github.com/woven-planet/go-zserio/issues/128 for more details.

## Links

- [zserio](https://github.com/ndsev/zserio)
- [Woven by Toyota](https://woven.toyota/en/)
