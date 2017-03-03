#!/usr/bin/env bash

#Builds a binary with attached information
go build -ldflags "-X main.version=0.1.0-beta -X main.sha1=`git rev-parse HEAD` -X main.date=`date +'%d-%m-%Y'`"