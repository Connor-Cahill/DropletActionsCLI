#!/bin/bash

echo "Starting Setup Script..."
curl -fsSL https://get.docker.com -o get-docker.sh \
        && sh get-docker.sh \
        && apt-get update && apt-get install -y docker-compose
