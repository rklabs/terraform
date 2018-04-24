package addrs

import (
	"fmt"
)

// InputVariable is the address of an input variable.
type InputVariable struct {
	referenceable
	Name string
}

func (v InputVariable) String() string {
	return "var." + v.Name
}

// AbsInputVariableInstance is the address of an input variable within a
// particular module instance.
type AbsInputVariableInstance struct {
	Module   ModuleInstance
	Variable InputVariable
}

func (v AbsInputVariableInstance) String() string {
	if len(v.Module) == 0 {
		return v.String()
	}

	return fmt.Sprintf("%s.%s", v.Module.String(), v.Variable.String())
}
