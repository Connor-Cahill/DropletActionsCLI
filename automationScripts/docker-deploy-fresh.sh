#!/bin/bash

VPS_IP=$1
PROJECT=$2
REPO_NAME=$3
ENV_FILE=$4

# sets up Droplet with docker and docker-compose CLI
setupDroplet() {
    curl -fsSL https://get.docker.com -o get-docker.sh \
        && sh get-docker.sh \
        && apt-get update && apt-get install -y docker-compose < "/dev/null" \
        git clone $PROJECT $REPO_NAME
}

ssh root@VPS_IP 'bash -s' < setupDroplet()



