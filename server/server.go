package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Witpxxl/GoApi/config"
	"github.com/Witpxxl/GoApi/error"
	"github.com/gorilla/mux"
)

type server struct {
	Config         *config.Configuration
	HandleFunction config.HandleFunction
	Router         *mux.Router
	Port           int
}

func NewServer(port int) server {
	serv := server{
		Config: config.NewConfiguration(""),
		HandleFunction: config.NewHandleFunction(),
		Port: port,
	}
	return serv
}

func (s server) Launch() {
	s.Router = mux.NewRouter()

	s.Router.NotFoundHandler = http.HandlerFunc(error404)

	for _, route := range s.Config.Routes {
		s.addRoute(route)
	}

	http.Handle("/", s.Router)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.Port), nil))
}

func (s server) AddFunction(list map[string]config.Handled) {
	for key, function := range list {
		s.HandleFunction.AddFunction(key, function)
	}
}

func (s server) LoadConfiguration() {
	s.Config.LoadConfig()
}

func (s server) addRoute(route config.Route) {

	s.Router.HandleFunc(route.Uri, func(w http.ResponseWriter, r *http.Request) {
		function := s.HandleFunction.GetFunction(route.Function)

		if function == nil {
			error404(w, r)
			return
		}

		result := function(r)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result.ToJson())
	}).Methods(route.Method)
}

func error404(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("Not Found: %s", r.RequestURI)
	errHTTP := error.ErrorHTTP{
		Code:404,
		Message: message,
	}

	display, err := json.Marshal(errHTTP)
	error.CheckErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	w.Write(display)
}
