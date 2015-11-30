package garbled

import "math/rand"

// A Gate represents any binary or unary gate.
type Gate interface {
	Evaluate() (uint32, uint32)
	Circuit() *Circuit
	GetOutput() *Wire
}

// A CryptoFunc is any function that takes a value and a key and
// returns an encrypted or decrypted value.
type CryptoFunc func(uint32, uint32) uint32

// A BinaryEvalFunc represents the function of a binary gate.
type BinaryEvalFunc func(uint32, uint32) uint32

// A UnaryEvalFunc represents the function of a unary gate.
type UnaryEvalFunc func(uint32) uint32

func generateKey() uint32 {
	return rand.Uint32() >> 1 // Discard upper bit; effectively 31-bit number
}

// A Wire represents the wire between two gates.
type Wire struct {
	Input  Gate      // the input gate
	Output Gate      // the output gate
	Keys   [2]uint32 // the keys:one for 0, one for 1
	P      uint32    // a randomised p-value
}

// NewWire returns a pointer to a new wire with randomised keys.
func NewWire(input Gate) *Wire {
	return &Wire{
		Input: input,
		Keys:  [2]uint32{generateKey(), generateKey()},
		P:     rand.Uint32() & 1, // 1 or 0
	}
}

// Evaluate uses a Wire's input gate to return a value.
func (w *Wire) Evaluate() (uint32, uint32) {
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

// Evaluate returns the Input's key and p-value.
func (i *Input) Evaluate() (uint32, uint32) {
	return i.Output.Keys[i.Value], i.Value ^ i.Output.P
}

// Circuit returns a pointer to the Circuit the Input is in.
func (i *Input) Circuit() *Circuit {
	return i.circuit
}

// GetOutput returns a pointer to the Input's output wire.
func (i *Input) GetOutput() *Wire {
	return i.Output
}
