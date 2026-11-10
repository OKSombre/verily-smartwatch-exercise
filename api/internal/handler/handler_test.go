package handler_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/verily/calendar/api/internal/builder"
	"github.com/verily/calendar/api/internal/handler"
)

const AsOfDate = "2026-10-21"

func logRequests(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                log.Printf("%s %s", r.Method, r.URL.Path)
                next.ServeHTTP(w, r)
        })
}

func newRouter(store *builder.RecordStore) http.Handler {
	h := &handler.Handler{Store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/studies/alpha/participants/{id}/calendar", h.GetParticipantRecordsByMonth)
	return logRequests(mux)
}

func TestGetParticipantRecordsByMonth(t *testing.T) {
	tests := []struct {
		name     	string
		endpoint	string
		expected 	int
	}{
		{
			name:     "missing month parameter",
			endpoint: "/api/studies/alpha/participants/1/calendar?month=",
			expected: http.StatusBadRequest,
		},
		{
			name:     "invalid month parameter, valid participantID",
			endpoint: "/api/studies/alpha/participants/A1000/calendar?month=invalid",
			expected: http.StatusBadRequest,
		},
		{
			name:     "invalid month format, valid participantID",
			endpoint: "/api/studies/alpha/participants/A1000/calendar?month=2026-1",
			expected: http.StatusBadRequest,
		},
		{
			name:     "valid month parameter, invalid participantID",
			endpoint: "/api/studies/alpha/participants/1/calendar?month=2026-10",
			expected: http.StatusNotFound,
		},
		{
			name:     "valid month parameter, valid participantID",
			endpoint: "/api/studies/alpha/participants/A1001/calendar?month=2026-10",
			expected: http.StatusOK,
		},
		{
			name:     "valid month parameter, valid participantID with no data",
			endpoint: "/api/studies/alpha/participants/A1000/calendar?month=2026-10",
			expected: http.StatusOK,
		},
	}
	router := newRouter(builder.NewRecordStoreWithSeed(AsOfDate))

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", test.endpoint, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != test.expected {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expected)
			}
		})
	}
	
}