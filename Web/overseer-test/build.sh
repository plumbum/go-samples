#!/bin/sh

go build -ldflags "-X main.buildDate=`date +%Y-%m-%d_%H:%M:%S`"

