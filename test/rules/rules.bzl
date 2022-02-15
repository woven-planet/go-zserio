"""Generate Zserio Python library and facilitate generation of test data.
"""

load("@pip//:requirements.bzl", "requirement")
load("@rules_python//python:defs.bzl", "py_binary", "py_library")

def _dirname(f):
    return "/".join(f.dirname.split("/")[:-1])

def _impl(ctx):
    root = _dirname(ctx.files.proto[0])
    if root == None:
        fail("TODO")

    for s in ctx.files.deps:
        candidate = _dirname(s)
        if len(root) > len(candidate):
            root = candidate

    args = ctx.actions.args()
    dir_name = ctx.label.name
    prefix = ctx.attr.prefix

    out = ctx.actions.declare_directory(dir_name)

    if prefix:
        dir_name = "{}/{}".format(dir_name, prefix)
        args.add("-setTopLevelPackage", prefix.replace("/", "."))

    args.add("-src", root)
    args.add("-python", out.path)
    args.add(ctx.files.proto[0].short_path.replace(root + "/", ""))

    outs = [out] + [ctx.actions.declare_file(dir_name + "/" + o) for o in ctx.attr.outs]

    ctx.actions.run(
        executable = ctx.executable._tool,
        mnemonic = "PyZserio",
        outputs = outs,
        inputs = ctx.files.deps + ctx.files.proto,
        arguments = [args],
        progress_message = "Generating Python bindings for zserio schema",
    )

    return [
        DefaultInfo(files = depset(outs)),
    ]

py_zserio_compile = rule(
    implementation = _impl,
    attrs = {
        "deps": attr.label_list(
            allow_empty = True,
            allow_files = [".zs"],
            doc = "Any dependencies for the Zserio file",
        ),
        "outs": attr.string_list(
            mandatory = True,
            doc = "The list of files that are output by the Zserio compiler",
        ),
        "prefix": attr.string(
            doc = "The import prefix for the generated code",
        ),
        "proto": attr.label(
            allow_single_file = [".zs"],
            mandatory = True,
            doc = "The Zserio file to pass to the compiler",
        ),
        "_tool": attr.label(
            cfg = "host",
            default = Label("//test/rules:zserio_compiler"),
            doc = "The zserio compiler tool",
            executable = True,
        ),
    },
)

def py_zserio_library(name, proto, outs, proto_deps = [], prefix = None, **kwargs):
    """py_zserio_library generates Python code from Zserio files and provides a python library.

    Args:
        name: The name of the resultant zserio Python library.
        proto: The file to use to output the zserio files.
        outs: The python files that are created during the compilation.
        proto_deps: The list of Zserio files that are dependencies of the 'proto' file.
        prefix: The prefix for the output package name.
        **kwargs: Any extra arguments to be passed to the underlying py_library.
    """
    gen = name + "_zs"
    py_zserio_compile(
        name = gen,
        proto = proto,
        deps = proto_deps,
        outs = outs,
        prefix = prefix,
    )

    py_library(
        name = name,
        srcs = [gen] + kwargs.pop("srcs", []),
        imports = [gen] + kwargs.pop("imports", []),
        deps = [requirement("zserio")] + kwargs.pop("deps", []),
        **kwargs
    )

def zs_payload(name, srcs, out, **kwargs):
    """gen_data makes it easier to generate test data binary for running tests.

    It assums that there will be a file named 'data.py' and it will have a
    'new' constructor for creating the data to be serialized.

    Args:
        name: The name the genrule target, that is doing the generation.
        srcs: The source files used to create an object to be serialized.
        out: The output filename.
        **kwargs: Any extra arguments to be passed to the py_binary rule.
    """
    gen = name + "_gen"
    py_binary(
        name = gen,
        srcs = ["//test/rules:write_data.py"] + srcs,
        main = "//test/rules:write_data.py",
        imports = ["."],
        **kwargs
    )

    native.genrule(
        name = name,
        outs = [out],
        cmd = "$(execpath {}) >$@".format(gen),
        message = "Generating zserio payload binary using official Python bindings",
        tools = [gen],
    )
