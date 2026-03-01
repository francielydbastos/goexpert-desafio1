package http_utils

import (
	"Client-Server-API/server/service"
	"encoding/json"
	"net/http"
)

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido para este endpoint.", http.StatusMethodNotAllowed)
		return
	}

	cotacao, err := service.ProcessarRequest(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"bid": cotacao,
	})

}
