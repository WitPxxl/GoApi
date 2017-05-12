package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfiguration(t *testing.T) {
	fixtureFile := "config_test.json"

	res := NewConfiguration(fixtureFile)
	assert.IsType(t, &Configuration{}, res)

	assert.Equal(t, fixtureFile, res.File)
}

func TestSaveConfig(t *testing.T) {
	f, err := ioutil.TempFile(".", "test")

	if err != nil {
		t.Fatal(err)
	}

	filename := f.Name()
	config := NewConfiguration(filename)

	route := Route{
		Name:     "name",
		Uri:      "uri",
		Method:   "GET",
		Function: "Test",
	}

	config.Routes = append(config.Routes, route)

	config.SaveConfig()

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	expectedConfig := Configuration{
		File: filename,
	}
	err = json.Unmarshal(bytes, &expectedConfig)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, config, &expectedConfig)
	os.Remove(f.Name())
}
