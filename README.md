# Simple example of grpc

This is just a simple example for [grpc](https://grpc.io/) with Go.

## Usage

Install protobuf (3.0) and grpc. Then:

1. `./make-certs.sh` creates a self-signed wildcard cart for `*.go`
2. Open a second terminal.
3. Run `make serve` on the first terminal.
4. Run `male cli` on the second terminal.

This really doesn't do anything besides a simple RPC call, including TLS
configuration and server-name-based transport security.

# License

MIT
