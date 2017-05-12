package config

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"

	"../error"
)

var configFile = "config.json"

type Configuration struct {
	Routes []Route `json:"routes"`
	File string
}

func NewConfiguration(file string) *Configuration {
	if file == "" {
		file = configFile
	}

	config := Configuration{
		File: file,
	}

	return &config
}

func (config *Configuration) SaveConfig() {
	bytes, err := json.MarshalIndent(config, "", "  ")
	error.CheckErr(err)

	err = ioutil.WriteFile(config.File, bytes, 0644)
	error.CheckErr(err)
}

func (config *Configuration) LoadConfig() {
	_, currentFile, _, _ := runtime.Caller(0)
	bytes, err := ioutil.ReadFile(path.Dir(currentFile) + "/" + config.File)
	error.CheckErr(err)

	err = json.Unmarshal(bytes, config)
	error.CheckErr(err)
}
