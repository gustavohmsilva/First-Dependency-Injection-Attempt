package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavohmsilva/test-dependency-injection/datasource"
	"github.com/gustavohmsilva/test-dependency-injection/handler"
	"github.com/gustavohmsilva/test-dependency-injection/service"
)

// Start the server
func Start(addr, port string) error {
	// carrega as rotas
	r := mux.NewRouter()
	ds, err := datasource.NewUserDataSource("demoDB.db")
	if err != nil {
		return err
	}
	us := service.NewUserService(&ds)
	h := handler.NewHandler(us)
	r.HandleFunc("/user", h.GetLatestUser).Methods("GET")
	r.HandleFunc("/user", h.SetLatestUser).Methods("POST")
	// inicia o servidor
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), r)
	if err != nil {
		return err
	}
	return nil
}
