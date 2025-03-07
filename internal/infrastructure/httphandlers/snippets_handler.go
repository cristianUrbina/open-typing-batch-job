package httphandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"cristianUrbina/open-typing-batch-job/internal/app"

	"github.com/gorilla/mux"
)

type SnippetHandler struct {
	Service *app.SnippetService
}

func (h *SnippetHandler) GetSnippetByLanguage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang, ok := vars["lang"]
	if !ok {
		http.Error(w, "Missing language parameter", http.StatusBadRequest)
		return
	}

	snippet, err := h.Service.GetRandomSnippetByLanguage(lang)
	if err != nil {
		log.Printf("error getting snippet: %v", err)
		http.Error(w, "Error fetching snippet", http.StatusInternalServerError)
		return
	}

	if snippet == nil {
		http.Error(w, "No snippet found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippet)
}
