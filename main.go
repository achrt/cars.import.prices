package main

import (
	"fmt"
	"os"

	// "github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("APPLICATION_PORT")

	Run(port)
}

func Run(port string) {
	//r := mux.NewRouter()
	fmt.Println("PRICES_WAS_CONNECTED")
	//fmt.Println(r)
}
