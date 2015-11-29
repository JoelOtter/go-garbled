package garbled

func makeBinaryGate(left, right *Wire) *BinaryGate {
	g := &BinaryGate{
		Left:  left,
		Right: right,
	}
	left.Output = g
	right.Output = g
	g.Output = NewWire(g)
	return g
}

func makeUnaryGate(input *Wire) *UnaryGate {
	g := &UnaryGate{
		Input: input,
	}
	input.Output = g
	g.Output = NewWire(g)
	return g
}

// AndGate will return a pointer to a BinaryGate
// with inputs left and right and with AND
// as its evaluation function.
func AndGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.generateGarbledTable([4]uint32{0, 0, 0, 1})
	g.Name = "AND"
	return g
}

// OrGate will return a pointer to a BinaryGate
// with inputs left and right and with OR
// as its evaluation function.
func OrGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.generateGarbledTable([4]uint32{0, 1, 1, 1})
	g.Name = "OR"
	return g
}

// XorGate will return a pointer to a BinaryGate
// with inputs left and right and with XOR
// as its evaluation function.
func XorGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.generateGarbledTable([4]uint32{0, 1, 1, 0})
	g.Name = "XOR"
	return g
}

// NandGate will return a pointer to a BinaryGate
// with inputs left and right and with NAND
// as its evaluation function.
func NandGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.generateGarbledTable([4]uint32{1, 1, 1, 0})
	g.Name = "NAND"
	return g
}

// NorGate will return a pointer to a BinaryGate
// with inputs left and right and with NOR
// as its evaluation function.
func NorGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.generateGarbledTable([4]uint32{1, 0, 0, 0})
	g.Name = "NOR"
	return g
}

// XnorGate will return a pointer to a BinaryGate
// with inputs left and right and with XNOR
// as its evaluation function.
func XnorGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.generateGarbledTable([4]uint32{1, 0, 0, 1})
	g.Name = "XNOR"
	return g
}

// NotGate will return a pointer to a UnaryGate
// with input 'input' and with NOT as its
// evaluation function.
func NotGate(input Gate) *UnaryGate {
	g := makeUnaryGate(input.GetOutput())
	g.generateGarbledTable([2]uint32{1, 0})
	g.Name = "NOT"
	return g
}
