# Go File encrypt

Allows running file encryption operations using Go and AES cyphers.

## How to run

To encrypt a file run the following command, where ``<decrypted_file>`` should be replaced by origin file path and
``<encrypted_file>`` should be replaced by destination file path.

``go run ./cmd/encrypt/main.go <decrypted_file> <encrypted_file>``

To decrypt a file run the following command, where ``<encrypted_file>`` should be replaced by origin file path and .
``<decrypted_file>`` should be replaced by destination file path.

``go run ./cmd/decrypt/main.go <encrypted_file> <decrypted_file>``

## Generate token

Generate a token by running the following command, where ``<length>`` should be replaced by token length.

``go run ./cmd/token/main.go <length>``
