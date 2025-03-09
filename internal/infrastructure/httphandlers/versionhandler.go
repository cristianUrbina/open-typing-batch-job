package httphandlers

import (
	"encoding/json"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	response := VersionResponse{Version: "1.0.0"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
