load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "ab21448cef298740765f33a7f5acee0607203e4ea321219f2a4c85a6e0fb0a27",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.32.0/rules_go-v0.32.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.32.0/rules_go-v0.32.0.zip",
    ],
)

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "b8b69615e8d9ade79f3612311b8d0c4dfe01017420c90eed11db15e9e7c9ff3c",
    strip_prefix = "buildtools-4.2.1",
    urls = [
        "https://github.com/bazelbuild/buildtools/archive/refs/tags/4.2.1.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "5982e5463f171da99e3bdaeff8c0f48283a7a5f396ec5282910b9e8a49c0dd7e",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.25.0/bazel-gazelle-v0.25.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.25.0/bazel-gazelle-v0.25.0.tar.gz",
    ],
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "9b4ee22c250fe31b16f1a24d61467e40780a3fbb9b91c3b65be2a376ed913a1a",
    strip_prefix = "protobuf-3.13.0",
    urls = [
        "https://github.com/protocolbuffers/protobuf/archive/v3.13.0.tar.gz",
    ],
)

http_archive(
    name = "rules_antlr",
    patch_args = [
        "-p1",
    ],
    patches = [
        "//third_party/patches:rules_antlr_no_gofmt.patch",
    ],
    sha256 = "26e6a83c665cf6c1093b628b3a749071322f0f70305d12ede30909695ed85591",
    strip_prefix = "rules_antlr-0.5.0",
    urls = ["https://github.com/marcohu/rules_antlr/archive/0.5.0.tar.gz"],
)

http_archive(
    name = "rules_python",
    sha256 = "cdf6b84084aad8f10bf20b46b77cb48d83c319ebe6458a18e9d2cebf57807cdd",
    strip_prefix = "rules_python-0.8.1",
    url = "https://github.com/bazelbuild/rules_python/archive/refs/tags/0.8.1.tar.gz",
)

http_archive(
    name = "zserio",
    sha256 = "dd23369dca96514fcab91ae9cefe15aaf27a54fc9fe8cb613c5d6b1c742d46fc",
    url = "https://github.com/ndsev/zserio/releases/download/v2.6.0/zserio-2.6.0-bin.zip",
)

http_archive(
    name = "bazel_skylib",
    sha256 = "f7be3474d42aae265405a592bb7da8e171919d74c16f082a5457840f06054728",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.2.1/bazel-skylib-1.2.1.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.2.1/bazel-skylib-1.2.1.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@rules_antlr//antlr:repositories.bzl", "rules_antlr_dependencies")
load("//:deps.bzl", "go_dependencies")

bazel_skylib_workspace()

protobuf_deps()

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

rules_antlr_dependencies("4.8")

go_rules_dependencies()

go_register_toolchains(version = "1.18.2")

gazelle_dependencies()

load("@rules_python//python:pip.bzl", "pip_install")

pip_install(
    requirements = "//test/rules:requirements.txt",
)

RULES_JVM_EXTERNAL_TAG = "4.2"

RULES_JVM_EXTERNAL_SHA = "cd1a77b7b02e8e008439ca76fd34f5b07aecb8c752961f9640dea15e9e5ba1ca"

http_archive(
    name = "rules_jvm_external",
    sha256 = RULES_JVM_EXTERNAL_SHA,
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
    url = "https://github.com/bazelbuild/rules_jvm_external/archive/%s.zip" % RULES_JVM_EXTERNAL_TAG,
)

load("@rules_jvm_external//:repositories.bzl", "rules_jvm_external_deps")

rules_jvm_external_deps()

load("@rules_jvm_external//:setup.bzl", "rules_jvm_external_setup")

rules_jvm_external_setup()

load("@rules_jvm_external//:defs.bzl", "maven_install")

maven_install(
    artifacts = [
        "io.github.ndsev:zserio:2.6.0",
    ],
    repositories = [
        "https://repo.maven.apache.org/maven2/",
    ],
)
