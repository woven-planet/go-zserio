# syntax=docker/dockerfile:1.2

# Setup initial apt-get infrastructure and CA certificates which are going to
# be needed in later layers.
FROM ubuntu:focal AS base

RUN apt-get update && \
    env DEBIAN_FRONTEND=noninteractive apt-get upgrade -y \
    && \
    env DEBIAN_FRONTEND=noninteractive apt-get install -y \
        apt-transport-https \
        ca-certificates \
    && \
    update-ca-certificates \
    && \
    rm -rf /var/lib/apt/lists/* && \
    rm -f /var/log/dpkg.log &&\
    rm -f /var/log/alternatives.log &&\
    rm -f /var/cache/ldconfig/aux-cache &&\
    rm -f /var/cache/apt/pkgcache.bin

# Create a final image used in our build process.
FROM base AS builder

ARG bazelisk_version=1.11.0

# Extra dependencies:
# Pacakges required by go-zerio:
# * default-jdk is required for ANTLR
#
# Packages required by bazel, see (https://docs.bazel.build/versions/main/install-ubuntu.html#step-1-install-required-packages)
# * g++
# * zip/unzip are required by bazel when archiving test results
#
# Packages required by our setup:
# * git is assumed to be present on all development machines
# * direnv is used to maintain the development environment
#
# NOTE: Bazel is not installed via apt. Instead we use bazelisk.
RUN apt-get update && \
    env DEBIAN_FRONTEND=noninteractive apt-get install -y \
        default-jdk \
        direnv \
        git \
        g++ \
        unzip \
        zip \
    && \
    rm -rf /var/lib/apt/lists/* && \
    rm -f /var/log/dpkg.log &&\
    rm -f /var/log/alternatives.log &&\
    rm -f /var/cache/ldconfig/aux-cache &&\
    rm -f /var/cache/apt/pkgcache.bin

# Downloads bazelisk for development.
ADD https://github.com/bazelbuild/bazelisk/releases/download/v${bazelisk_version}/bazelisk-linux-amd64 /usr/bin/bazel
RUN chmod 755 /usr/bin/bazel

# Set up a non-root user for bazel build.
RUN useradd -u 1001 -ms /bin/bash builder \
    && \
    groupadd -g 998 robots \
    && \
    mkdir -p /home/builder/.cache /home/builder/.local \
    && \
    chown -R builder:robots /home/builder/.cache \
    && \
    chown -R builder:robots /home/builder/.local

USER builder:robots
