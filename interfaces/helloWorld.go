package interfaces

import (
	json "encoding/json"
	netHttp "net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HelloWorld(w netHttp.ResponseWriter, r *netHttp.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "Olá, mundo!",
		Status:  200,
	}

	w.WriteHeader(netHttp.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
