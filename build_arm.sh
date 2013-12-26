#!/bin/bash

PKG=$1

GOARM=5 GOARCH=arm go build $PKG
