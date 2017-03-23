#!/usr/bin/env bash
APP_NAME=$1
if [ -z ${APP_NAME} ]; then
    echo "Missing AppName arg"
    exit 1
fi
SLOG_DEBUG=1
go install ${APP_NAME} && bin/${APP_NAME} "${@:2}"
