package main

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	// Inputs - A B C D
	inputs := [16][4]uint32{
		[4]uint32{0, 0, 0, 0},
		[4]uint32{0, 0, 0, 1},
		[4]uint32{0, 0, 1, 0},
		[4]uint32{0, 0, 1, 1},
		[4]uint32{0, 1, 0, 0},
		[4]uint32{0, 1, 0, 1},
		[4]uint32{0, 1, 1, 0},
		[4]uint32{0, 1, 1, 1},
		[4]uint32{1, 0, 0, 0},
		[4]uint32{1, 0, 0, 1},
		[4]uint32{1, 0, 1, 0},
		[4]uint32{1, 0, 1, 1},
		[4]uint32{1, 1, 0, 0},
		[4]uint32{1, 1, 0, 1},
		[4]uint32{1, 1, 1, 0},
		[4]uint32{1, 1, 1, 1},
	}
	expd := [16][2]uint32{
		[2]uint32{0, 0},
		[2]uint32{0, 1},
		[2]uint32{1, 0},
		[2]uint32{1, 1},
		[2]uint32{0, 1},
		[2]uint32{0, 1},
		[2]uint32{1, 0},
		[2]uint32{1, 1},
		[2]uint32{1, 0},
		[2]uint32{1, 0},
		[2]uint32{1, 0},
		[2]uint32{1, 1},
		[2]uint32{1, 1},
		[2]uint32{1, 1},
		[2]uint32{1, 1},
		[2]uint32{1, 1},
	}
	c := BuildCircuit()
	for i, input := range inputs {
		inputs := map[string]uint32{
			"A": input[0],
			"B": input[1],
			"C": input[2],
			"D": input[3],
		}
		out := c.Evaluate(inputs)
		if out["O1"] == expd[i][0] && out["O2"] == expd[i][1] {
			fmt.Printf("Input %v ==> Output [%v %v]\n",
				input, out["O1"], out["O2"])
		} else {
			t.Errorf("Input %v returned [%v %v], expd [%v %v]",
				input, out["O1"], out["O2"], expd[i][0], expd[i][1])
		}
	}
}
