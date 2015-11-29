package garbled

import "math/rand"
import "fmt"

// Binary gate
type BinaryGate struct {
	Name         string    // a user-readable name
	Left         *Wire     // the wire on the 'left' input
	Right        *Wire     // the wire on the 'right' input
	Output       *Wire     // the output wire
	GarbledTable [4]uint32 // garbled table
}

// Evaluate will use the left and right inputs to produce
// the appropriate output value.
func (b *BinaryGate) Evaluate() uint32 {
	keyL := b.Left.Evaluate()
	keyR := b.Right.Evaluate()
	c := b.Circuit()
	for _, k := range b.GarbledTable {
		res := c.Decryptor(c.Decryptor(k, keyL), keyR)
		if decryptionValid(res) {
			return res
		}
	}
	fmt.Printf("Decryption error in gate %v\n", b.Name)
	return 0
}

// GetOutput returns a pointer to the gate's output wire.
func (b *BinaryGate) GetOutput() *Wire {
	return b.Output
}

// Circuit returns a pointer to the Circuit this gate is part of.
func (b *BinaryGate) Circuit() *Circuit {
	return b.Left.Circuit()
}

func (b *BinaryGate) generateGarbledTable(inputs [4]uint32) {
	table := [4]uint32{}
	table[0] = b.Circuit().Encryptor( // input 0, 0
		b.Circuit().Encryptor(b.Output.Keys[inputs[0]], b.Left.Keys[0]),
		b.Right.Keys[0],
	)
	table[1] = b.Circuit().Encryptor( // input 0, 1
		b.Circuit().Encryptor(b.Output.Keys[inputs[1]], b.Left.Keys[0]),
		b.Right.Keys[1],
	)
	table[2] = b.Circuit().Encryptor( // input 1, 0
		b.Circuit().Encryptor(b.Output.Keys[inputs[2]], b.Left.Keys[1]),
		b.Right.Keys[0],
	)
	table[3] = b.Circuit().Encryptor( // input 1, 1
		b.Circuit().Encryptor(b.Output.Keys[inputs[3]], b.Left.Keys[1]),
		b.Right.Keys[1],
	)

	// Shuffle the table
	for i := range inputs {
		r := rand.Intn(4)
		swap := table[r]
		table[r] = table[i]
		table[i] = swap
	}

	b.GarbledTable = table
}
