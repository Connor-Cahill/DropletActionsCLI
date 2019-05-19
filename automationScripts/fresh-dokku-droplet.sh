#!/bin/bash

wget https://raw.githubusercontent.com/dokku/dokku/v0.16.1/bootstrap.sh \
    && sudo DOKKU_TAG=v0.16.1 bash bootstrap.sh
