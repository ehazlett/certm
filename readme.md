# CertM
CertM is a simple tool to generate TLS certificates and keys.

# Usage
## Show Help
`docker run --rm ehazlett/certm -h`

## Generate CA
`docker run --rm -v $(pwd)/certs:/certs ehazlett/certm -d /certs ca generate -o=local`

This will generate a CA with the organization "local".

## Generate server certificate
`docker run --rm -v $(pwd)/certs:/certs ehazlett/certm -d /certs server generate --host localhost --host 127.0.0.1 -o=local`

This will generate a server certificate with a SAN of "localhost" and an
IP SAN of "127.0.0.1" with the organization "local".

## Generate client certificate
`docker run --rm -v $(pwd)/certs:/certs ehazlett/certm -d /certs client generate --common-name=ehazlett -o=local`

This will generate a client certificate with the common name of "ehazlett".

## Generate CA, server and client certificates / keys
`docker run --rm -v $(pwd)/certs:/certs ehazlett/certm -d /certs bundle generate --host 127.0.0.1 -o=local`

This will generate a CA using the org "local", a server certificate with an
IP SAN of "127.0.0.1" and a client certificate.

## Generate CA, client and server certificates/keys
`docker run --rm -v $(pwd)/certs:/certs ehazlett/certm -d /certs bundle generate -o=local -s localhost -s 127.0.0.1 -s foo.local`

This will generate a CA using the org "local", a client cert, and a server
certificate that is valid using the DNS names "localhost" and "foo.local" as
well as the IP "127.0.0.1"

Server cert can be used for swarm and has cert extensions for both docker server and client.

# Integration Tests
To run integration tests, use `./script/test`.  This will run the integration
tests in a container to validate proper usage.
