package garbled

import "testing"

func testGate(t *testing.T, c *Circuit, expd [4]bool) {
	inputs := [4][2]bool{
		[2]bool{false, false},
		[2]bool{false, true},
		[2]bool{true, false},
		[2]bool{true, true},
	}
	for idx, inp := range inputs {
		out := c.Evaluate(map[string]bool{"A": inp[0], "B": inp[1]})
		if out["O"] != expd[idx] {
			t.Errorf("Circuit %v with inputs:\nA: %t\nB: %t\nReturned %t, expd %t",
				c.Name, inp[0], inp[1], out["O"], expd[idx])
		}
	}
}

func TestAndGate(t *testing.T) {
	c := NewCircuit("AND")
	a := c.AddInput("A")
	b := c.AddInput("B")
	and := AndGate(a, b)
	c.AddOutput("O", and)
	testGate(t, c, [4]bool{false, false, false, true})
}

func TestOrGate(t *testing.T) {
	c := NewCircuit("OR")
	a := c.AddInput("A")
	b := c.AddInput("B")
	or := OrGate(a, b)
	c.AddOutput("O", or)
	testGate(t, c, [4]bool{false, true, true, true})
}

func TestXorGate(t *testing.T) {
	c := NewCircuit("XOR")
	a := c.AddInput("A")
	b := c.AddInput("B")
	xor := XorGate(a, b)
	c.AddOutput("O", xor)
	testGate(t, c, [4]bool{false, true, true, false})
}

func TestNandGate(t *testing.T) {
	c := NewCircuit("NAND")
	a := c.AddInput("A")
	b := c.AddInput("B")
	nand := NandGate(a, b)
	c.AddOutput("O", nand)
	testGate(t, c, [4]bool{true, true, true, false})
}

func TestNorGate(t *testing.T) {
	c := NewCircuit("NOR")
	a := c.AddInput("A")
	b := c.AddInput("B")
	nor := NorGate(a, b)
	c.AddOutput("O", nor)
	testGate(t, c, [4]bool{true, false, false, false})
}

func TestXnorGate(t *testing.T) {
	c := NewCircuit("XNOR")
	a := c.AddInput("A")
	b := c.AddInput("B")
	xnor := XnorGate(a, b)
	c.AddOutput("O", xnor)
	testGate(t, c, [4]bool{true, false, false, true})
}
