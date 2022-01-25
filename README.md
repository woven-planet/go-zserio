# Bazel

Update `BUILD.bazel` files:

```shell
bazel run //:gazelle
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
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=repositories.bzl%go_repositories
```
