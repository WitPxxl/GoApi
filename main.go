package main

import (
	"github.com/Witpxxl/GoApi/config"
	"github.com/Witpxxl/GoApi/hello"
	"github.com/Witpxxl/GoApi/server"
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
