package garbled

// Unary gate
type UnaryGate struct {
	Name         string        // a user-readable name
	Input        *Wire         // the wire used for input
	Output       *Wire         //the output wire
	GarbledTable [2]uint32     // garbled table
	Evaluator    UnaryEvalFunc // the function to evaluate the gate
}

// Evaluate will use the input to produce
// the appropriate key and p-value.
func (g *UnaryGate) Evaluate() (uint32, uint32) {
	key, p := g.Input.Evaluate()
	encrypted := g.GarbledTable[p]

	// Decrypt
	c := g.Circuit()
	decrypted := c.Decryptor(encrypted, key)
	pOut := decrypted >> 31              // MSB
	keyOut := decrypted & ^uint32(1<<31) // all bits except MSB
	return keyOut, pOut
}

// GetOutput returns a pointer to the gate's output wire.
func (g *UnaryGate) GetOutput() *Wire {
	return g.Output
}

// Circuit returns a pointer to the Circuit this gate is part of.
func (g *UnaryGate) Circuit() *Circuit {
	return g.Input.Circuit()
}

func (g *UnaryGate) generateGarbledTable() {
	var table [2]uint32
	for i := range table {
		x := uint32(i) ^ g.Input.P
		z := g.Evaluator(x)
		t := z ^ g.Output.P
		keyOut := g.Output.Keys[z]
		toEncrypt := keyOut | (t << 31) // MSg.of key is t
		key := g.Input.Keys[x]

		// Encrypt
		c := g.Circuit()
		table[i] = c.Encryptor(toEncrypt, key)
	}
	g.GarbledTable = table
}
