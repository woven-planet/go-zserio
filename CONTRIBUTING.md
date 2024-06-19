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

You will need to use [bazel](http://bazel.build) if you want to build go-zserio
or run its tests. We use bazel to install required build tools, generate test data,
compile go-zserio itself and run the tests.

To run all tests:

```shell
bazel test --test_output=all //...
```

To run the go-zserio command:

```shell
bazel run //cmd/go-zserio
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

In order for [VS Code](https://code.visualstudio.com) you will need to configure
it to use a bazel-aware _gopackagesdriver_ with the gopls language server. To do
that use the following for `.vscode/settings.json`, replacing
`/Users/wichert/Code/go-zserio` with the correct location on your machine.

```json
{
  "go.goroot": "/Users/wichert/Code/go-zserio/bazel-go-zserio/external/go_sdk",
  "go.toolsEnvVars": {
    "GOPACKAGESDRIVER": "/Users/wichert/Code/go-zserio/bazel-out/bazel_env-opt/bin/bazel_env/bin/gopackagesdriver.sh"
  },
  "go.enableCodeLens": {
    "references": false,
    "runtest": false
  },
  "gopls": {
    "build.directoryFilters": [
      "-bazel-bin",
      "-bazel-out",
      "-bazel-testlogs",
      "-bazel-go-zserio",
    ],
    "formatting.gofumpt": true,
    "formatting.local": "github.com/woven-planet/go-zserio",
    "ui.completion.usePlaceholders": true,
    "ui.semanticTokens": true,
    "ui.codelenses": {
      "gc_details": false,
      "regenerate_cgo": false,
      "generate": false,
      "test": false,
      "tidy": false,
      "upgrade_dependency": false,
      "vendor": false
    },
  },
  "go.useLanguageServer": true,
  "go.buildOnSave": "off",
  "go.lintOnSave": "off",
  "go.vetOnSave": "off",
}
```
