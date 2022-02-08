"""A simple bazel rule to build a single go_library from a set of zserio schema files.
"""

load("@io_bazel_rules_go//go:def.bzl", "go_library")

def go_zserio_library(name, srcs, rootpackage, pkg, **kwargs):
    native.genrule(
        name = name + "_gen",
        # TODO @aignas 2022-02-08: use something more similar to rules_proto,
        # so that we can encode dependencies between different zs files.
        srcs = srcs,
        outs = [name + "_gen.zs.go"],
        cmd = "{bin} generate --out - --rootpackage {rootpackage} --only {pkg} {srcs} > $@ 2>/dev/null".format(
            bin = "$(execpath //cmd/zserio)",
            rootpackage = rootpackage,
            pkg = pkg,
            srcs = "$(SRCS)",
        ),
        tools = [
            "//cmd/zserio",
        ],
    )

    go_library(
        name = name,
        srcs = [name + "_gen"],
        importpath = "{}/{}".format(rootpackage, pkg.replace(".", "/")),
        deps = [
            "//interface",
            "//ztype",
            "@com_github_icza_bitio//:bitio",
        ] + kwargs.pop("deps", []),
        **kwargs
    )
