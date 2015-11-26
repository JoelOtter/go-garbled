package garbled

import "testing"

func testGate(t *testing.T, c *Circuit, expd [4]uint) {
	inputs := [4][2]uint{
		[2]uint{0, 0},
		[2]uint{0, 1},
		[2]uint{1, 0},
		[2]uint{1, 1},
	}
	for idx, inp := range inputs {
		out := c.Evaluate(map[string]uint{"A": inp[0], "B": inp[1]})
		if out["O"] != expd[idx] {
			t.Errorf("Circuit %v with inputs:\nA: %v\nB: %v\nReturned %v, expd %v",
				c.Name, inp[0], inp[1], out["O"], expd[idx])
		}
	}
}

func testUnaryGate(t *testing.T, c *Circuit, expd [2]uint) {
	inputs := [2]uint{0, 1}

	for idx, inp := range inputs {
		out := c.Evaluate(map[string]uint{"A": inp})
		if out["O"] != expd[idx] {
			t.Errorf("Circuit %v with input:\nA: %v\nReturned %v, expd %v",
				c.Name, inp, out["O"], expd[idx])
		}
	}
}

func TestAndGate(t *testing.T) {
	c := NewCircuit("AND")
	a := c.AddInput("A")
	b := c.AddInput("B")
	and := AndGate(a, b)
	c.AddOutput("O", and)
	testGate(t, c, [4]uint{0, 0, 0, 1})
}

func TestOrGate(t *testing.T) {
	c := NewCircuit("OR")
	a := c.AddInput("A")
	b := c.AddInput("B")
	or := OrGate(a, b)
	c.AddOutput("O", or)
	testGate(t, c, [4]uint{0, 1, 1, 1})
}

func TestXorGate(t *testing.T) {
	c := NewCircuit("XOR")
	a := c.AddInput("A")
	b := c.AddInput("B")
	xor := XorGate(a, b)
	c.AddOutput("O", xor)
	testGate(t, c, [4]uint{0, 1, 1, 0})
}

func TestNandGate(t *testing.T) {
	c := NewCircuit("NAND")
	a := c.AddInput("A")
	b := c.AddInput("B")
	nand := NandGate(a, b)
	c.AddOutput("O", nand)
	testGate(t, c, [4]uint{1, 1, 1, 0})
}

func TestNorGate(t *testing.T) {
	c := NewCircuit("NOR")
	a := c.AddInput("A")
	b := c.AddInput("B")
	nor := NorGate(a, b)
	c.AddOutput("O", nor)
	testGate(t, c, [4]uint{1, 0, 0, 0})
}

func TestXnorGate(t *testing.T) {
	c := NewCircuit("XNOR")
	a := c.AddInput("A")
	b := c.AddInput("B")
	xnor := XnorGate(a, b)
	c.AddOutput("O", xnor)
	testGate(t, c, [4]uint{1, 0, 0, 1})
}

func TestNotGate(t *testing.T) {
	c := NewCircuit("NOT")
	a := c.AddInput("A")
	not := NotGate(a)
	c.AddOutput("O", not)
	testUnaryGate(t, c, [2]uint{1, 0})
}
