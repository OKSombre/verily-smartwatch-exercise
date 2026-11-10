package main

import (
	"log"
	"net/http"

	"github.com/verily/calendar/api/internal/builder"
	"github.com/verily/calendar/api/internal/handler"
)

const AsOfDate = "2026-10-21"

// logRequests is a small middleware that prints every request.
func logRequests(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                log.Printf("%s %s", r.Method, r.URL.Path)
                next.ServeHTTP(w, r)
        })
}

// newRouter creates a router for the API.
func newRouter(store *builder.RecordStore) http.Handler {
	h := &handler.Handler{Store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/studies/alpha/participants/{id}/calendar", h.GetParticipantRecordsByMonth)
	return logRequests(mux)
}

func main() {
	store := builder.NewRecordStoreWithSeed(AsOfDate)
	log.Fatal(http.ListenAndServe(":8080", newRouter(store)))
}