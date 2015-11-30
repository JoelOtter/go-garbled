package main

import (
	"fmt"
	g "github.com/joelotter/go-garbled"
)

func BuildCircuit() *g.Circuit {
	cr := g.NewCircuit("Max")
	a := cr.AddInput("A")
	b := cr.AddInput("B")
	c := cr.AddInput("C")
	d := cr.AddInput("D")

	o1 := g.OrGate(a, c)
	cr.AddOutput("O1", o1)

	o2 := g.OrGate(
		g.OrGate(
			g.AndGate(a, b),
			g.AndGate(c, d),
		),
		g.OrGate(
			g.AndGate(g.NotGate(a), d),
			g.AndGate(b, g.NotGate(c)),
		),
	)
	cr.AddOutput("O2", o2)

	return cr
}

func main() {
	c := BuildCircuit()
	inputs := map[string]uint32{
		"A": 0,
		"B": 1,
		"C": 0,
		"D": 0,
	}
	out := c.Evaluate(inputs)
	fmt.Println(out)
}
