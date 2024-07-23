module allas-header-client

go 1.22.5

replace common => ../common

require (
	common v1.0.0
	github.com/neicnordic/crypt4gh v1.12.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/dchest/bcrypt_pbkdf v0.0.0-20150205184540-83f37f9c154a // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
)
