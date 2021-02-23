#!/bin/sh
set -e
export GOBIN=`pwd`/bin
go install ./cmd/...
unset GOBIN
echo "Your files in /bin/"
./bin/goodsGateway --config-path ./configs/goodsGateway/goodsGateway.properties
