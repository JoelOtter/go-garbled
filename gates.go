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
	g.Name = "AND"
	g.Evaluator = func(left, right uint32) uint32 {
		return left & right
	}
	g.generateGarbledTable()
	return g
}

// OrGate will return a pointer to a BinaryGate
// with inputs left and right and with OR
// as its evaluation function.
func OrGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.Name = "OR"
	g.Evaluator = func(left, right uint32) uint32 {
		return left | right
	}
	g.generateGarbledTable()
	return g
}

// XorGate will return a pointer to a BinaryGate
// with inputs left and right and with XOR
// as its evaluation function.
func XorGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.Name = "XOR"
	g.Evaluator = func(left, right uint32) uint32 {
		return left ^ right
	}
	g.generateGarbledTable()
	return g
}

// NandGate will return a pointer to a BinaryGate
// with inputs left and right and with NAND
// as its evaluation function.
func NandGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.Name = "NAND"
	g.Evaluator = func(left, right uint32) uint32 {
		return 1 & ^(left & right)
	}
	g.generateGarbledTable()
	return g
}

// NorGate will return a pointer to a BinaryGate
// with inputs left and right and with NOR
// as its evaluation function.
func NorGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.Name = "NOR"
	g.Evaluator = func(left, right uint32) uint32 {
		return 1 & ^(left | right)
	}
	g.generateGarbledTable()
	return g
}

// XnorGate will return a pointer to a BinaryGate
// with inputs left and right and with XNOR
// as its evaluation function.
func XnorGate(left, right Gate) *BinaryGate {
	g := makeBinaryGate(left.GetOutput(), right.GetOutput())
	g.Name = "XNOR"
	g.Evaluator = func(left, right uint32) uint32 {
		return 1 & ^(left ^ right)
	}
	g.generateGarbledTable()
	return g
}

// NotGate will return a pointer to a UnaryGate
// with input 'input' and with NOT as its
// evaluation function.
func NotGate(input Gate) *UnaryGate {
	g := makeUnaryGate(input.GetOutput())
	g.Name = "NOT"
	g.Evaluator = func(input uint32) uint32 {
		return 1 & ^input
	}
	g.generateGarbledTable()
	return g
}
