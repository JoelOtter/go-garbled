# go-garbled

[![Travis CI](https://travis-ci.org/JoelOtter/go-garbled.svg?branch=master)](https://travis-ci.org/JoelOtter/go-garbled) [![GoDoc](https://godoc.org/github.com/JoelOtter/go-garbled?status.svg)](http://godoc.org/github.com/JoelOtter/go-garbled)

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

### Examples

**max.go** - A circuit which determines the maximum of two 2-bit numbers. To evaluate for all possible inputs, navigate to the `_examples/` directory and run `go test`.

### Keys

Keys have the type uint32, but are in fact effectively 31-bit numbers. The most significant bit is a space left for the p-value which is to be encrypted along with the key. Keys are generated randomly, so it's not really necessary for a user to care about this detail.

The default encryption method is simply XOR between plaintext and key. Users can modify this themselves by adjusting the `Encryptor` field on the `Circuit` struct.
