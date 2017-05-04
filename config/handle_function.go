package config

import (
	"net/http"

	"../interfaces"
)

type HandleFunction struct {
	functions map[string]Handled
}

type Handled func(r *http.Request) interfaces.Json

func NewHandleFunction() HandleFunction {
	makeFunction := make(map[string]Handled)
	return HandleFunction{
		functions: makeFunction,
	}
}

func (hf HandleFunction) AddFunction(name string, function Handled) {
	hf.functions[name] = function
}

func (hf HandleFunction) GetFunctions() map[string]Handled {
	return hf.functions
}

func (hf HandleFunction) GetFunction(name string) Handled {
	if _, ok := hf.functions[name]; ok {
		return hf.functions[name]
	}

	return nil
}
