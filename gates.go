package garbled

// AndGate will return a pointer to a BinaryGate
// with inputs left and right and with AND
// as its evaluation function.
func AndGate(left, right Gate) *BinaryGate {
	return &BinaryGate{
		Name:  "AND",
		Left:  left,
		Right: right,
		EvalFunc: func(left, right bool) bool {
			return left && right
		},
	}
}

// OrGate will return a pointer to a BinaryGate
// with inputs left and right and with OR
// as its evaluation function.
func OrGate(left, right Gate) *BinaryGate {
	return &BinaryGate{
		Name:  "OR",
		Left:  left,
		Right: right,
		EvalFunc: func(left, right bool) bool {
			return left || right
		},
	}
}

// XorGate will return a pointer to a BinaryGate
// with inputs left and right and with XOR
// as its evaluation function.
func XorGate(left, right Gate) *BinaryGate {
	return &BinaryGate{
		Name:  "XOR",
		Left:  left,
		Right: right,
		EvalFunc: func(left, right bool) bool {
			return left != right
		},
	}
}

// NandGate will return a pointer to a BinaryGate
// with inputs left and right and with NAND
// as its evaluation function.
func NandGate(left, right Gate) *BinaryGate {
	return &BinaryGate{
		Name:  "NAND",
		Left:  left,
		Right: right,
		EvalFunc: func(left, right bool) bool {
			return !(left && right)
		},
	}
}

// NorGate will return a pointer to a BinaryGate
// with inputs left and right and with NOR
// as its evaluation function.
func NorGate(left, right Gate) *BinaryGate {
	return &BinaryGate{
		Name:  "NOR",
		Left:  left,
		Right: right,
		EvalFunc: func(left, right bool) bool {
			return !(left || right)
		},
	}
}

// XnorGate will return a pointer to a BinaryGate
// with inputs left and right and with XNOR
// as its evaluation function.
func XnorGate(left, right Gate) *BinaryGate {
	return &BinaryGate{
		Name:  "XNOR",
		Left:  left,
		Right: right,
		EvalFunc: func(left, right bool) bool {
			return left == right
		},
	}
}

// NotGate will return a pointer to a UnaryGate
// with input 'input' and with NOT as its
// evaluation function.
func NotGate(input Gate) *UnaryGate {
	return &UnaryGate{
		Name:  "NOT",
		Input: input,
		EvalFunc: func(val bool) bool {
			return !val
		},
	}
}
