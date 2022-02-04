# docker images

## go-zserio:dev

This docker image is intended to be used as a development environment of `go-zserio`.
The image has minimum dependencies to build `go-zserio` code such as `bazel` and `java`.

The image can be built by `make`.

To use it you can:
\`\`\`bash
$ make build
$ bazel-dev build //...
\`\`\`
