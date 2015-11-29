# go-garbled

![[Travis CI](https://travis-ci.org/JoelOtter/go-garbled)](https://travis-ci.org/JoelOtter/go-garbled.svg?branch=master) [![GoDoc](https://godoc.org/github.com/JoelOtter/go-garbled?status.svg)](http://godoc.org/github.com/JoelOtter/go-garbled)

A library for building and evaluating Yao's Garbled Circuits in Go.

Supports gates:
- AND
- OR
- NOT
- XOR
- NAND
- NOR
- XNOR

The aim is for ease of use and extensibility.

### Keys

Keys are formatted as 32-bit unsigned integers. The least significant 16 bits are the key itself, and the most significant 16 bits are a SHA-1 checksum.

The default encryption method is simply XOR between plaintext and key. Users can modify this themselves by adjusting the `Encryptor` field on the `Circuit` struct. An encryption/decryption function should maintain the checksum/key format.
