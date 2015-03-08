# Docker CertTool
CertTool is a simple tool to generate TLS certificates and keys for use with
Docker.


# Usage
## Show Help
`docker-certtool -h`: show help

## Generate CA and client certificates / keys
`docker-certtool -d=certs -o=local`

This will generate a CA using the org "local" and a client certificate.

## Generate CA, client and server certificates/keys
`docker-certtool -d=certs -o=local -s localhost -s 127.0.0.1 -s foo.local`

This will generate a CA using the org "local", a client cert, and a server
certificate that is valid using the DNS names "localhost" and "foo.local" as
well as the IP "127.0.0.1"

# Integration Tests
To run integration tests, use `./script/test`.  This will run the integration
tests in a container to validate proper usage.
