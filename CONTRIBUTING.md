# Constributing

## Direnv

Install [direnv](https://direnv.net) tool to setup your PATH automatically to
include the `bin` directory to be able to use the extra helper scripts for
managing go dependencies, otherwise, add it manually to your shell.

## Bazel

We use [bazel](https://bazel.build) for building and running all tests. The
following bazel commands will be useful for go-zserio:

To run all tests:

```shell
bazel test --test_output=all //...
```

To run the go-zserio command:

```shell
bazel run //cmd/zserio
```

To update the go dependencies for bazel using `go.mod`:

```shell
gazelle-update-repos
```
