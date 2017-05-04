package main

import (
	"./config"
	"./hello"
	"./server"
)

var list = map[string]config.Handled{
	"SayHello": hello.SayHello,
}

func main() {
	serv := server.NewServer(3000)
	serv.LoadConfiguration()
	serv.AddFunction(list)
	serv.Launch()
}
