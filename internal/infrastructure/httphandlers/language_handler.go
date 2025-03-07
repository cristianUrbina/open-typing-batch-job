package httphandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"cristianUrbina/open-typing-batch-job/internal/app"

	"github.com/gorilla/mux"
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

func (h *LanguageHandler) GetLanguageByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang, ok := vars["lang"]
	if !ok {
		http.Error(w, "Missing language parameter", http.StatusBadRequest)
		return
	}
	language, err := h.Service.GetLanguageByName(lang)
	if err != nil {
		log.Printf("error: %v", err)
		http.Error(w, "Failed to fetch language", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(language); err != nil {
		log.Printf("error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
