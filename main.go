package main

import (
	routes "MyFirstGo/routers"
	"fmt"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)
}

func print() {
	fmt.Println("Hello World")
}
