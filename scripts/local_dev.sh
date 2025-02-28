#!/bin/bash

set -e

cd fe && npm run build && cd ..
docker compose -f docker/docker-compose.yml down
docker compose -f docker/docker-compose.yml up --build

