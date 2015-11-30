package garbled

// Binary gate
type BinaryGate struct {
	Name         string         // a user-readable name
	Left         *Wire          // the wire on the 'left' input
	Right        *Wire          // the wire on the 'right' input
	Output       *Wire          // the output wire
	GarbledTable [2][2]uint32   // garbled table
	Evaluator    BinaryEvalFunc // the function to evaluate the gate
}

// Evaluate will use the left and right inputs to produce
// the appropriate output key and p-value.
func (b *BinaryGate) Evaluate() (uint32, uint32) {
	keyL, pL := b.Left.Evaluate()
	keyR, pR := b.Right.Evaluate()
	encrypted := b.GarbledTable[pL][pR]

	// Decrypt
	c := b.Circuit()
	decrypted := c.Decryptor(
		c.Decryptor(encrypted, keyR),
		keyL,
	)
	pOut := decrypted >> 31              // MSB
	keyOut := decrypted & ^uint32(1<<31) // all bits except MSB
	return keyOut, pOut
}

// GetOutput returns a pointer to the gate's output wire.
func (b *BinaryGate) GetOutput() *Wire {
	return b.Output
}

// Circuit returns a pointer to the Circuit this gate is part of.
func (b *BinaryGate) Circuit() *Circuit {
	return b.Left.Circuit()
}

func (b *BinaryGate) generateGarbledTable() {

	var table [2][2]uint32

	inputs := [4][2]uint32{
		[2]uint32{0, 0},
		[2]uint32{0, 1},
		[2]uint32{1, 0},
		[2]uint32{1, 1},
	}

	for _, ins := range inputs {
		x := ins[0] ^ b.Left.P
		y := ins[1] ^ b.Right.P
		z := b.Evaluator(x, y)
		t := z ^ b.Output.P
		keyOut := b.Output.Keys[z]
		toEncrypt := keyOut | (t << 31) // MSB of key is t
		keyR := b.Right.Keys[y]
		keyL := b.Left.Keys[x]

		// Encrypt
		c := b.Circuit()
		table[ins[0]][ins[1]] = c.Encryptor(
			c.Encryptor(toEncrypt, keyR),
			keyL,
		)
	}

	b.GarbledTable = table
}
