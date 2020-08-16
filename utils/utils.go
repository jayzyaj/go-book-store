package utils

import (
	"encoding/json"
	"github.com/jayzyaj/go-book-store/models"
	"net/http"
)

// SendError ...
func SendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

// SendSuccess ...
func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
