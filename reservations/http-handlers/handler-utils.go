package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rodsil01/solution/reservations/exceptions"
)

func createResponse(w http.ResponseWriter, payload interface{}, err *exceptions.AppError) {
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(payload)
	}
}
