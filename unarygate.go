package garbled

import (
	"fmt"
	"math/rand"
)

// Unary gate
type UnaryGate struct {
	Name         string    // a user-readable name
	Input        *Wire     // the wire used for input
	Output       *Wire     //the output wire
	GarbledTable [2]uint32 // garbled table
}

// Evaluate will use the input to produce
// the appropriate output value.
func (g *UnaryGate) Evaluate() uint32 {
	key := g.Input.Evaluate()
	c := g.Circuit()
	for _, k := range g.GarbledTable {
		res := c.Decryptor(k, key)
		if decryptionValid(res) {
			return res
		}
	}
	fmt.Printf("Decryption error in gate %v.\n", g.Name)
	return 0
}

// GetOutput returns a pointer to the gate's output wire.
func (g *UnaryGate) GetOutput() *Wire {
	return g.Output
}

// Circuit returns a pointer to the Circuit this gate is part of.
func (g *UnaryGate) Circuit() *Circuit {
	return g.Input.Circuit()
}

func (g *UnaryGate) generateGarbledTable(inputs [2]uint32) {
	table := [2]uint32{}
	table[0] = g.Circuit().Encryptor(g.Output.Keys[inputs[0]], g.Input.Keys[0])
	table[1] = g.Circuit().Encryptor(g.Output.Keys[inputs[1]], g.Input.Keys[1])
	r := rand.Intn(2)
	if r == 1 {
		swap := table[0]
		table[0] = table[1]
		table[1] = swap
	}
	g.GarbledTable = table
}
