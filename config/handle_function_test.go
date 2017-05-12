package config

import (
	"net/http"
	"testing"

	"../interfaces"
	"github.com/stretchr/testify/assert"
)

func TestNewHandleFunction(t *testing.T) {
	res := NewHandleFunction()
	assert.IsType(t, HandleFunction{}, res)
	assert.IsType(t, make(map[string]Handled), res.functions)
}

type mockJson struct {
}

func (m mockJson) ToJson() []byte {
	res := []byte("test")
	return res
}

func mockHandledFunction(r *http.Request) interfaces.Json {
	return mockJson{}
}

func TestAddFunction(t *testing.T) {
	object := NewHandleFunction()

	object.AddFunction("testAdd", mockHandledFunction)

	assert.IsType(t, Handled(mockHandledFunction), object.functions["testAdd"])
}

func TestGetFunction(t *testing.T) {
	object := NewHandleFunction()

	object.functions["testAdd"] = mockHandledFunction

	f := object.GetFunction("testAdd")

	assert.ObjectsAreEqual(f, mockHandledFunction)

	f2 := object.GetFunction("test")

	assert.Nil(t, f2)
}

func TestGetFunctions(t *testing.T) {
	object := NewHandleFunction()

	fixtureFunctions := map[string]Handled{
		"test": mockHandledFunction,
	}

	object.functions = fixtureFunctions

	f := object.GetFunctions()

	assert.ObjectsAreEqual(f, fixtureFunctions)
}
