package httphandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"cristianUrbina/open-typing-batch-job/internal/app"
)

type LanguageHandler struct {
	Service *app.LanguageService
}

func (h *LanguageHandler) GetLanguages(w http.ResponseWriter, r *http.Request) {
	languages, err := h.Service.GetAvailableLanguages()
	if err != nil {
		log.Printf("error: %v", err)
		http.Error(w, "Failed to fetch languages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(languages); err != nil {
		log.Printf("error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
