#!/bin/bash


DockerSetup() {
    curl -fsSL https://get.docker.com -o get-docker.sh \
        && sh get-docker.sh \
        && apt intsall docker-compose
    if (($?)); then apt-get update && apt-get install -y docker-compose;fi
}

ssh root@$1 'bash -s' < DockerSetup()
