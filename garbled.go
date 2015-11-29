package garbled

import (
	"crypto/sha1"
	"encoding/binary"
	"math/rand"
)

// A Gate represents any binary or unary gate.
type Gate interface {
	Evaluate() uint32
	Circuit() *Circuit
	GetOutput() *Wire
}

// A CryptoFunc is any function that takes a value and a key and
// returns an encrypted or decrypted value.
type CryptoFunc func(uint32, uint32) uint32

func generateKey() uint32 {
	r := uint16(rand.Uint32())
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, r)
	sum := sha1.Sum(b)
	bytes := []byte{sum[0], sum[1], b[0], b[1]}
	return binary.BigEndian.Uint32(bytes)
}

func decryptionValid(key uint32) bool {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, key)
	sum := sha1.Sum(b[2:])
	if sum[0] == b[0] && sum[1] == b[1] {
		return true
	}
	return false
}

// A Wire represents the wire between two gates.
// It has a key for 0 and a key for 1.
type Wire struct {
	Input  Gate
	Output Gate
	Keys   [2]uint32
}

// NewWire returns a pointer to a new wire with randomised keys.
func NewWire(input Gate) *Wire {
	return &Wire{
		Input: input,
		Keys:  [2]uint32{generateKey(), generateKey()},
	}
}

// Evaluate uses a Wire's input gate to return a value.
func (w *Wire) Evaluate() uint32 {
	return w.Input.Evaluate()
}

// Circuit returns a pointer to the Circuit the Wire is in.
func (w *Wire) Circuit() *Circuit {
	return w.Input.Circuit()
}

// Input 'gate', used to supply inputs to the circuit.
type Input struct {
	Value   uint32
	circuit *Circuit // a pointer back to the Circuit
	Output  *Wire    // the output wire
}

// Evaluate returns the Input's value.
func (i *Input) Evaluate() uint32 {
	return i.Output.Keys[i.Value]
}

// Circuit returns a pointer to the Circuit the Input is in.
func (i *Input) Circuit() *Circuit {
	return i.circuit
}

// GetOutput returns a pointer to the Input's output wire.
func (i *Input) GetOutput() *Wire {
	return i.Output
}
