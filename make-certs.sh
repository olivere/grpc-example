#!/usr/bin/env bash
DIR="./etc/"
DOMAIN="*.go"
PASSPHRASE=""
SUBJ="
C=DE
ST=
O=
localityName=
commonName=$DOMAIN
organizationalUnitName=
emailAddress=
"

mkdir -p "$DIR"

# Generate our Private Key, CSR and Certificate
openssl genrsa -out "$DIR/star.go.key" 2048
openssl req -new -subj "$(echo -n "$SUBJ" | tr "\n" "/")" -key "$DIR/star.go.key" -out "$DIR/star.go.csr" -passin pass:$PASSPHRASE
openssl x509 -req -days 365 -in "$DIR/star.go.csr" -signkey "$DIR/star.go.key" -out "$DIR/star.go.crt"
