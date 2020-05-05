package v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Boot(router *mux.Router) {
	r := router.PathPrefix("/api/v1").Subrouter()
	w := Wrapper()

	r.HandleFunc("/hooks/car/retail_price", w(ReplacePrice)).Methods(http.MethodPost)
}
