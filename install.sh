#!/bin/bash

set -eu

PKGER_REPO_URL="github.com/markbates/pkger/cmd/pkger"
PKGER_PATH=`which pkger`

if [ $? = 1 ]; then
    go get ${PKGER_REPO_URL}
fi

make install
