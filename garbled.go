package garbled

type Gate interface {
	Evaluate() bool
}

// Circuit struct for whole circuit
type Circuit struct {
	Name    string            // a user-readable circuit name
	Inputs  map[string]*Input // a map of input names to input 'gates'
	Outputs map[string]Gate   // a map of output names to gates
}

// NewCircuit creates a new Circuit with its
// name as a string argument.
// Returns a pointer to a new Circuit.
func NewCircuit(name string) *Circuit {
	c := Circuit{
		Name:    name,
		Inputs:  make(map[string]*Input),
		Outputs: make(map[string]Gate),
	}
	return &c
}

// AddInput registers a new Input in the Circuit
// with the provided name and value.
// It returns a pointer to the Input.
func (c *Circuit) AddInput(name string) *Input {
	i := &Input{}
	c.Inputs[name] = i
	return i
}

// AddOutput registers a new Output in the Circuit
// with the provided name.
func (c *Circuit) AddOutput(name string, g Gate) {
	c.Outputs[name] = g
}

// Evaluate will evaluate a whole circuit for the inputs specified
// in the map 'inputs'. Returns a map of outputs to their values.
//
// E.g. For a circuit containing a single AND gate with
// inputs A and B, and one output O, the map:
// {"A": false, "B": true}
// will evaluate to:
// {"O": false}.
func (c *Circuit) Evaluate(inputs map[string]bool) map[string]bool {
	for k, v := range inputs {
		c.Inputs[k].Value = v
	}
	outputs := make(map[string]bool)
	for k, v := range c.Outputs {
		outputs[k] = v.Evaluate()
	}
	return outputs
}

// Input 'gate', used to supply inputs to the circuit.
type Input struct {
	Value bool
}

// Evaluate returns the Input's value.
func (i *Input) Evaluate() bool {
	return i.Value
}

// Binary gate
type BinaryGate struct {
	Name     string                // a user-readable name
	Left     Gate                  // the gate on the 'left' input
	Right    Gate                  // the gate on the 'right' input
	EvalFunc func(bool, bool) bool // the function to evaluate the inputs
}

// Evaluate will use the left and right inputs to produce
// the appropriate output value.
func (b *BinaryGate) Evaluate() bool {
	return b.EvalFunc(b.Left.Evaluate(), b.Right.Evaluate())
}

// Unary gate
type UnaryGate struct {
	Name     string          // a user-readable name
	Input    Gate            // the gate used for input
	EvalFunc func(bool) bool // the function to evaulate the input
}

// Evaluate will use the input to produce
// the appropriate output value.
func (g *UnaryGate) Evaluate() bool {
	return g.EvalFunc(g.Input.Evaluate())
}
