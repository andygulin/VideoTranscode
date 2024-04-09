#!/usr/bin/env bash

#go install github.com/mitchellh/gox@latest

gox -osarch="darwin/amd64"
gox -osarch="windows/amd64"