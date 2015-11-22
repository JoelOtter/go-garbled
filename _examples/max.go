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
			g.AndGate(g.NotGate(c), b),
		),
	)
	cr.AddOutput("O2", o2)

	return cr
}

func main() {
	c := BuildCircuit()
	inputs := map[string]bool{
		"A": false,
		"B": false,
		"C": false,
		"D": false,
	}
	out := c.Evaluate(inputs)
	fmt.Println(out)
}
