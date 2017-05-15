package hello

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Witpxxl/GoApi/error"
	"github.com/Witpxxl/GoApi/interfaces"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Hello struct {
	Id int `json:"id"`
}

func (h Hello) ToJson() []byte {
	bytes, err := json.Marshal(h)

	error.CheckErr(err)

	return bytes
}

func SayHello(r *http.Request) interfaces.Json {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	error.CheckErr(err)

	hello := Hello{
		Id: id,
	}

	return hello
}
