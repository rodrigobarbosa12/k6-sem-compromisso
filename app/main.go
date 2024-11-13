package main

import (
	"fmt"
	"k6/interfaces"
	"log"
	netHttp "net/http"
)

func main() {
	netHttp.HandleFunc("/", interfaces.HelloWorld)
	netHttp.HandleFunc("/run-test", interfaces.AuthApi)

	fmt.Println("Servidor iniciado na porta 8080")
	log.Fatal(netHttp.ListenAndServe(":8080", nil))
}
