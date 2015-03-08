#!/usr/bin/env bats

load helpers

@test "certs: valid generation" {
    run certgen $CERT_ORG $CERT_BIT_SIZE -s $CERT_SERVER_DNS -s $CERT_SERVER_IP
    [ "$status" -eq 0  ]
}

@test "certs: ca.pem should exist" {
    run test -e $CERT_TEST_DIR/ca.pem
    [ "$status" -eq 0  ]
}

@test "certs: ca-key.pem should exist" {
    run test -e $CERT_TEST_DIR/ca-key.pem
    [ "$status" -eq 0  ]
}

@test "certs: client.pem should exist" {
    run test -e $CERT_TEST_DIR/client.pem
    [ "$status" -eq 0  ]
}

@test "certs: client-key.pem should exist" {
    run test -e $CERT_TEST_DIR/client-key.pem
    [ "$status" -eq 0  ]
}

@test "certs: server.pem should exist" {
    run test -e $CERT_TEST_DIR/server.pem
    [ "$status" -eq 0  ]
}

@test "certs: server-key.pem should exist" {
    run test -e $CERT_TEST_DIR/server-key.pem
    [ "$status" -eq 0  ]
}

@test "certs: valid ca organization" {
    run openssl x509 -in $CERT_TEST_DIR/ca.pem -noout -text
    echo "$output" | grep "Issuer: O=$CERT_ORG"
    [ "$status" -eq 0  ]
}

@test "certs: valid ca bit size" {
    run openssl x509 -in $CERT_TEST_DIR/ca.pem -noout -text
    echo "$output" | grep "($CERT_BIT_SIZE bit)"
    [ "$status" -eq 0  ]
}

@test "certs: valid CA" {
    run openssl x509 -in $CERT_TEST_DIR/ca.pem -noout -text
    echo "$output" | grep "CA:TRUE"
    [ "$status" -eq 0  ]
}

@test "certs: valid client organization" {
    run openssl x509 -in $CERT_TEST_DIR/client.pem -noout -text
    echo "$output" | grep "Issuer: O=$CERT_ORG"
    [ "$status" -eq 0  ]
}

@test "certs: valid client bit size" {
    run openssl x509 -in $CERT_TEST_DIR/client.pem -noout -text
    echo "$output" | grep "($CERT_BIT_SIZE bit)"
    [ "$status" -eq 0  ]
}

@test "certs: client cert is not a CA" {
    run openssl x509 -in $CERT_TEST_DIR/client.pem -noout -text
    echo "$output" | grep "CA:FALSE"
    [ "$status" -eq 0  ]
}

@test "certs: client cert has proper extended usage" {
    run openssl x509 -in $CERT_TEST_DIR/client.pem -noout -text
    echo "$output" | grep "TLS Web Client Authentication"
    [ "$status" -eq 0  ]
}

@test "certs: valid server organization" {
    run openssl x509 -in $CERT_TEST_DIR/server.pem -noout -text
    echo "$output" | grep "Issuer: O=$CERT_ORG"
    [ "$status" -eq 0  ]
}

@test "certs: valid server bit size" {
    run openssl x509 -in $CERT_TEST_DIR/server.pem -noout -text
    echo "$output" | grep "($CERT_BIT_SIZE bit)"
    [ "$status" -eq 0  ]
}

@test "certs: server cert is not a CA" {
    run openssl x509 -in $CERT_TEST_DIR/server.pem -noout -text
    echo "$output" | grep "CA:FALSE"
    [ "$status" -eq 0  ]
}

@test "certs: server cert has proper extended usage" {
    run openssl x509 -in $CERT_TEST_DIR/server.pem -noout -text
    echo "$output" | grep "TLS Web Server Authentication"
    [ "$status" -eq 0  ]
}

@test "certs: server has correct DNS SAN" {
    run openssl x509 -in $CERT_TEST_DIR/server.pem -noout -text
    echo "$output" | grep "DNS:$CERT_SERVER_DNS"
    [ "$status" -eq 0  ]
}

@test "certs: server has correct IP SAN" {
    run openssl x509 -in $CERT_TEST_DIR/server.pem -noout -text
    echo "$output" | grep "IP Address:$CERT_SERVER_IP"
    [ "$status" -eq 0  ]
}

@test "certs: custom org/bits generation" {
    run certgen $CERT_CUSTOM_ORG $CERT_CUSTOM_BIT_SIZE
    [ "$status" -eq 0  ]
}

@test "certs: custom org valid" {
    run openssl x509 -in $CERT_TEST_DIR/ca.pem -noout -text
    echo "$output" | grep "Issuer: O=$CERT_CUSTOM_ORG"
    [ "$status" -eq 0  ]
}

@test "certs: custom ca bit size" {
    run openssl x509 -in $CERT_TEST_DIR/ca.pem -noout -text
    echo "$output" | grep "($CERT_CUSTOM_BIT_SIZE bit)"
    [ "$status" -eq 0  ]
}
