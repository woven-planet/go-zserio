load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_antlr",
    patch_args = [
        "-p1",
    ],
    patches = [
        "//third_party/patches:rules_antlr_no_gofmt.patch",
    ],
    sha256 = "489225cfbaf4e4c6a2ae3b41af95cf6c3fa746148db17e4057e645652e231edc",
    strip_prefix = "rules_antlr-0.5.8",
    urls = ["https://github.com/GorNishanov/rules_antlr/archive/0.5.8.tar.gz"],
)

http_archive(
    name = "zserio",
    sha256 = "e1b5824836405635fc5a22639394458949d358e6450c4d874f45a56ed5186f3c",
    url = "https://github.com/ndsev/zserio/releases/download/v2.11.0/zserio-2.11.0-bin.zip",
)

load("@rules_antlr//antlr:repositories.bzl", "rules_antlr_dependencies")

rules_antlr_dependencies("4.11.1")
