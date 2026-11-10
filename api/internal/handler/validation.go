package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// writeJSON sends any value as JSON with the given status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(status)
        json.NewEncoder(w).Encode(v)
}

// writeError sends an error message as JSON with the given status code.
func writeError(w http.ResponseWriter, status int, msg string) {
        writeJSON(w, status, map[string]string{"error": msg})
}

// pathID extracts the participant ID from the request path.
func pathID(r *http.Request) (string, error) {
        id := r.PathValue("id")
        if id == "" {
                return "", errors.New("id is required")
        }
        return id, nil
}

// queryMonth extracts and validates the month query parameter.
func queryMonth(r *http.Request) (string, error) {
        month := r.URL.Query().Get("month")
        if month == "" {
                return "", errors.New("month is required")
        }
		if _, err := time.Parse("2006-01", month); err != nil {
			return "", errors.New("month must be in YYYY-MM format")
		}
        return month, nil
}