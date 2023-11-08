package template

import (
	"encoding/json"
	"net/http"
)

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"ko": "rocks"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}