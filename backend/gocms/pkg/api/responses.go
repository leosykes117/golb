package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseAPI struct {
	Success bool        `json:"success"`
	Status  int         `json:"status,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func (r *ResponseAPI) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	err := json.NewEncoder(w).Encode(r)

	if err != nil {
		log.Printf("Error al convertir la estructura usuario en json: %v", err)
	}
}

func Success(result interface{}, status int) *ResponseAPI {
	return &ResponseAPI{
		Success: true,
		Status:  status,
		Result:  result,
	}
}
