# docker images

## go-zserio:dev

This docker image is intended to be used as a development environment of `go-zserio`.
The image has minimum dependencies to build `go-zserio` code such as `bazel` and `java`.

To use it you can use `bazel-dev` script in `bin` directory.
The script behaves differently as follows:

```bash
# Build go-zserio:dev image.
$ make build

# If the 1st argument is one of bazel commands, it directly executes the bazel command inside a container.
$ bazel-dev build //...

# Without arguments, the command let you enter a container.
$ bazel-dev
```

