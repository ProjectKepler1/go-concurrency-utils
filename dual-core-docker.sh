#!/usr/bin/env bash

# spin up a Go development container with:
# 2 CPUs (cpu 0 and cpu 1)
# 1024 file descriptors
# 4 Gb RAM (4096 Mb)

# map the current directory to /work in the container;
# also map the host port 41053 to the container port 41053 for go-trace,
# so that we can see the goroutine traces

docker run --rm -ti \
    --ulimit nofile=1024:1024 \
    --memory="4096m" \
    --cpuset-cpus=0-1 \
    --workdir=/work \
    -v "$PWD":/work \
    -p 41053:41053 \
    golang:1.18.0-bullseye \
    /bin/bash
