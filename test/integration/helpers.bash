#!/bin/bash

# Root directory of the repository.
ROOT=${BATS_TEST_DIRNAME}/../..
CERT_TEST_DIR=/tmp/certs
CERT_ORG=bats-test
CERT_BIT_SIZE=2048
CERT_SERVER_DNS=bats.test.local
CERT_SERVER_IP=1.2.3.4
CERT_CUSTOM_ORG=foo-org
CERT_CUSTOM_BIT_SIZE=1024

build() {
    pushd $ROOT >/dev/null
    godep go build
    popd >/dev/null
}

# build machine binary if needed
if [ ! -e $MACHINE_ROOT/certm ]; then
    build
fi

certm() {
    ${ROOT}/certm "$@"
}

certgen() {
    if [ -e $CERT_TEST_DIR ]; then
        rm -rf $CERT_TEST_DIR
    fi
    certtool -d $CERT_TEST_DIR -o=$1 -b $2 ${*:3}
}
