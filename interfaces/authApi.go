package interfaces

import (
	"fmt"
	netHttp "net/http"
	"os"
	"os/exec"
)

func AuthApi(w netHttp.ResponseWriter, r *netHttp.Request) {
	w.Header().Set("Content-Type", "application/json")

	cmd := exec.Command("k6", "run", "-")

	inputFile, err := os.Open("tests/index.js")

	if err != nil {
		netHttp.Error(w, fmt.Sprintf("Erro ao abrir o arquivo de script: %v", err), netHttp.StatusInternalServerError)
		return
	}

	defer inputFile.Close()

	cmd.Stdin = inputFile

	output, err := cmd.CombinedOutput()
	if err != nil {
		netHttp.Error(w, fmt.Sprintf("Erro ao executar o k6: %v", err), netHttp.StatusInternalServerError)
		return
	}

	w.WriteHeader(netHttp.StatusOK)
	w.Write(output)
}
