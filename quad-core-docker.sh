#!/usr/bin/env bash

# spin up a Go development container with:
# 4 CPUs (cpu 0, 1, 2, 3)
# 4000 file descriptors
# 8 Gb RAM (8192 Mb)

# map the current directory to /work in the container;
# also map the host port 41053 to the container port 41053 for go-trace,
# so that we can see the goroutine traces

docker run --rm -ti \
    --ulimit nofile=4000:4000 \
    --memory="8192m" \
    --cpuset-cpus=0-3 \
    --workdir=/work \
    -v "$PWD":/work \
    -p 41053:41053 \
    golang:1.18.0-bullseye \
    /bin/bash
