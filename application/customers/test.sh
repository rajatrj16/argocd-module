#!/bin/bash
set -e
minikube ssh "docker run --rm --entrypoint /usr/local/go/bin/go rajatrj16/backend-image:2023 test ./..."