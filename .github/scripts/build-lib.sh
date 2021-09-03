#!/bin/bash
echo "🔷 Building Library..."
SCRIPTDIR=$(dirname "$0")
cd ${SCRIPTDIR}/../../

echo "Setting up Project"
PROJECT_DIR=$(pwd)
CORE_BIND_DIR=${PROJECT_DIR}/cmd/bind
CORE_RPC_DIR=${PROJECT_DIR}/cmd/rpc
go mod download golang.org/x/mobile
go mod download golang.org/x/mobile/cmd/gobind
mkdir -p ${PROJECT_DIR}/build

echo "Building Daemon"
cd ${CORE_RPC_DIR} && goreleaser build
mv ${PROJECT_DIR}/dist/rpc-* ${PROJECT_DIR}/build/
echo "✅  Finished Building Daemon ➡ `date`"

echo "Building for Android"
cd ${CORE_BIND_DIR} && gomobile bind -ldflags='-s -w' -target=android -o ${PROJECT_DIR}/build/io.sonr.core.aar
echo "✅  Finished Binding for Android ➡ `date`"

echo "Building for iOS"
cd ${CORE_BIND_DIR} && gomobile bind -ldflags='-s -w' -target=ios -bundleid=io.sonr.core -o ${PROJECT_DIR}/build/Core.framework
echo "✅  Finished Binding for iOS ➡ `date`"