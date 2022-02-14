# Constributing

## Prerequisites

You will need to install a number of tools to develop go-zserio:

- bazel
- a Java runtime

## Direnv

Install [direnv](https://direnv.net) tool to setup your PATH automatically to
include the `bin` directory to be able to use the extra helper scripts for
managing go dependencies, otherwise, add it manually to your shell.

## Bazel

You will need to use  [bazel](http://bazel.build) if you want to build go-zserio
or run its tests. We use bazel to install required build tools, generate test data,
compile go-zserio itself and run the tests.

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

## pre-commit

It is strongly recommended to run tests locally before you commit or a push a
change. The simplest way to do that is to use
[pre-commit](https://pre-commit.com/). This will automatically run all linters
and tests as a git pre-commit hook.

```shell
$ pre-commit install --install-hook
pre-commit installed at .git/hooks/pre-commit
[INFO] Initializing environment for https://github.com/executablebooks/mdformat.
[INFO] Initializing environment for https://github.com/executablebooks/mdformat:mdformat-black,mdformat-gfm.
[INFO] Initializing environment for https://github.com/pre-commit/pre-commit-hooks.
[INFO] Installing environment for https://github.com/executablebooks/mdformat.
[INFO] Once installed this environment will be reused.
[INFO] This may take a few minutes...
[INFO] Installing environment for https://github.com/pre-commit/pre-commit-hooks.
[INFO] Once installed this environment will be reused.
[INFO] This may take a few minutes...
```

## Visual Studio Code

In order for [VS Code](https://code.visualstudio.com) to find generated files
you will need to create a virtual GOPATH and point VS Code to it. First run this
command:

```shell
bazel build :gopath
```

This will create a new gopath folder in `bazel-out/gopath` with symlinks
pointing to your real GOPATH and any generated files. To point VSCode to it open
up `settings.json` for the workspace and set `go.path`.

```json
{
    "go.gopath": "/Users/wichert.akkerman/Code/go-zserio/bazel-out/gopath",
}
```
