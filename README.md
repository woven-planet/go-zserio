# go-zerio

[zserio] code generation in Go.

## Setup
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
