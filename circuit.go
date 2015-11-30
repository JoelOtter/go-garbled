package garbled

// Circuit struct for whole circuit
type Circuit struct {
	Name      string            // a user-readable circuit name
	Inputs    map[string]*Input // a map of input names to input 'gates'
	Outputs   map[string]*Wire  // a map of output names to wires
	Encryptor CryptoFunc        // the function used for encryption. Takes a number and a key
	Decryptor CryptoFunc        // function for decryption. Ciphertext and key.
}

// NewCircuit creates a new Circuit with its
// name as a string argument.
// Returns a pointer to a new Circuit.
func NewCircuit(name string) *Circuit {
	c := Circuit{
		Name:    name,
		Inputs:  make(map[string]*Input),
		Outputs: make(map[string]*Wire),
		Encryptor: func(m, k uint32) uint32 {
			return m ^ k
		},
		Decryptor: func(c, k uint32) uint32 {
			return c ^ k
		},
	}
	return &c
}

// AddInput registers a new Input in the Circuit
// with the provided name and value.
// It returns a pointer to the Input's output wire.
func (c *Circuit) AddInput(name string) *Input {
	i := new(Input)
	i.Output = NewWire(i)
	i.circuit = c
	c.Inputs[name] = i
	return i
}

// AddOutput registers a new Output in the Circuit
// with the provided name.
func (c *Circuit) AddOutput(name string, g Gate) {
	c.Outputs[name] = g.GetOutput()
}

// Evaluate will evaluate a whole circuit for the inputs specified
// in the map 'inputs'. Returns a map of outputs to their values.
//
// E.g. For a circuit containing a single AND gate with
// inputs A and B, and one output O, the map:
// {"A": 0, "B": 1}
// will evaluate to:
// {"O": 0}.
func (c *Circuit) Evaluate(inputs map[string]uint32) map[string]uint32 {
	for k, v := range inputs {
		c.Inputs[k].Value = v
	}
	outputs := make(map[string]uint32)
	for k, v := range c.Outputs {
		_, outputP := v.Evaluate()
		outputs[k] = outputP ^ v.P
	}
	return outputs
}
