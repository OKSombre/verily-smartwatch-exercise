package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/verily/calendar/api/internal/builder"
	"github.com/verily/calendar/api/internal/handler"
)

const AsOfDateDefault = "2026-10-21"

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
	allArgs := os.Args
	AsOfDate := AsOfDateDefault
	if len(allArgs) > 1 {
		arg := allArgs[1]
		if _, err := time.Parse("2006-01-02", arg); err == nil {
			AsOfDate = arg
		} else {
			log.Printf("invalid date format. Must be YYYY-MM-DD: %s", arg)
		}
	}
	log.Printf("using date %s as today", AsOfDate)
	
	store := builder.NewRecordStoreWithSeed(AsOfDate)
	log.Fatal(http.ListenAndServe(":8080", newRouter(store)))
}