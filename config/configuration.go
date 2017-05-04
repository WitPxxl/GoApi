package config

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"

	"../error"
)

type Configuration struct {
	Routes []Route `json:"routes"`
}

var configFile string = "config.json"

func (config *Configuration) SaveConfig() {
	bytes, err := json.MarshalIndent(config, "", "  ")
	error.CheckErr(err)

	err = ioutil.WriteFile(configFile, bytes, 0644)
	error.CheckErr(err)
}

func (config *Configuration) LoadConfig() {
	_, currentFile, _, _ := runtime.Caller(0)
	bytes, err := ioutil.ReadFile(path.Dir(currentFile) + "/" + configFile)
	error.CheckErr(err)

	err = json.Unmarshal(bytes, config)
	error.CheckErr(err)
}
