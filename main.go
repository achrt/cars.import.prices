package main

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"gitlab.kodixauto.ru/nw/logger"

	"cars.import.prices/api"
	v1 "cars.import.prices/api/v1"
)

func main() {
	port := os.Getenv("APPLICATION_PORT")

	Run(port)
}

func Run(port string) {
	r := mux.NewRouter()
	l := logger.New("runtime")
	v1.Boot(r)
	s := api.New(r)

	l.Info(fmt.Sprintf("server started on port [%s]", port), nil)

	panic(s.Run(":" + port))
}
