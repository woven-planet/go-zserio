[![Go Reference](https://pkg.go.dev/badge/github.com/woven-planet/go-zserio.svg)](https://pkg.go.dev/github.com/woven-planet/go-zserio)
[![Build Status](https://github.com/woven-planet/go-zserio/actions/workflows/test.yaml/badge.svg)](https://github.com/woven-planet/go-zserio/actions/workflows/test.yaml)

# go-zserio

[zserio](https://github.com/ndsev/zserio) code generation in Go.

## Usage

To generate the Go interface files to read/write zserio, a zserio scheme can be compiled by running:

```shell
zserio generate <zserio_directory> --rootpackage <root_package> --out <output_directory>
```

This command compiles the zserio files, and generates Go files to read and write
zserio encoded data. `zserio_directory` is the directory where the zserio
definitions are stored. Supported file extensions are `.zs` and `.zserio`.
`rootpackage` specifies the root package name of the generated Go files.
`output_directory` specifies the path where the Go files will be generated.

## Limitations

- Implicit length arrays are deprecated, and are not supported.
- The ternary operator only works when used in functions.
- Offsets and indexed offsets are not implemented yet.
- Constraints are not implemented yet.

## Links

- [zserio](https://github.com/ndsev/zserio)
- [Woven Planet](https://www.woven-planet.global/en)
