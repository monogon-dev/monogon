#!/usr/bin/env bash
set -euo pipefail

IMAGE=gcr.io/monogon-infra/monogon-builder:$(date +%s)

docker build -t "$IMAGE" .
gcloud docker --authorize-only
docker push "$IMAGE"
