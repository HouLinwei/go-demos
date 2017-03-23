#!/bin/bash
source ./setup-gopath.sh
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
APPS=`ls ./src`
for APP in ${APPS[@]}; do
    if [ $? -ne 0 ]; then
        echo "[ERROR] go install $APP"
        break
    else
        echo "[OK] go install $APP"
    fi
done
