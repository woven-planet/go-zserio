load("@rules_python//python:repositories.bzl", "python_register_toolchains")
load("@rules_python//python:pip.bzl", "pip_parse")


def initialize_python_toolchain():
    python_register_toolchains(
        name = "python3_12",
        ignore_root_user_error = True,
        # Available versions are listed in @rules_python//python:versions.bzl.
        python_version = "3.12.1",
    )
