package hello

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"../error"
	"../interfaces"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Hello struct {
	Id              int    `json:"id"`
	TransactionType string `json:"transactionType"`
	Price           int    `json:"price"`
}

func (h Hello) ToJson() []byte {
	bytes, err := json.Marshal(h)

	error.CheckErr(err)

	return bytes
}

func SayHello(r *http.Request) interfaces.Json {
	vars := mux.Vars(r)

	db, err := sql.Open("mysql", "athome:athome@tcp(127.0.0.1:3307)/immo")
	defer db.Close()

	error.CheckErr(err)

	stmt, err := db.Prepare("SELECT id, transaction_type, price FROM offer WHERE id=?")

	error.CheckErr(err)

	row := stmt.QueryRow(vars["id"])

	hello := Hello{}

	row.Scan(&hello.Id, &hello.TransactionType, &hello.Price)

	return hello
}
