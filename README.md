# go-zserio

[zserio] code generation in Go.

## Usage

To generate the Go interface files to read/write zserio, a zserio scheme can be compiled by running:

```shell
zserio generate <zserio_directory> --rootpackage <root_package> --out <output_directory>
```

This command compiles the zserio files, and generates Go files to read and write zserio encoded data.
`zserio_directory` is the directory where the zserio definitions are stored. Supported file extensions are `*.zs` and `*.zserio`.
`rootpackage` specifies the root package name of the generated Go files.
`output_directory` specifies the path where the Go files will be generated.

## Limitations

- Implicit length arrays are deprecated, and are not supported.
- The ternary operator only works when used in functions.
- Offsets and indexed offsets are not implemented yet.
- Constraints are not implemented yet.

## Developement

### Direnv

Install `direnv` tool to setup your PATH automatically to include the `bin`
directory to be able to use the extra helper scripts for managing go
dependencies, otherwise, add it manually to your shell.

### Bazel

Update `BUILD.bazel` files:

```shell
gazelle
```

Run all tests:

```shell
bazel test --test_output=all //...
```

Run command:

```shell
bazel run //cmd:hello
```

Update modules:

```shell
gazelle-update-repos
```

[zserio]: https://github.com/ndsev/zserio
