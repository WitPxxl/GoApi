package config

type Route struct {
	Name     string `json:"name"`
	Method   string `json:"method"`
	Uri      string `json:"uri"`
	Function string `json:"func"`
}
