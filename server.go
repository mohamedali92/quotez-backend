package main

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)


// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getQuotesFromDb(w http.ResponseWriter, req *http.Request) {
	// setup up db connection
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("QUOTESDSN"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer conn.Close(ctx)
	quotes, err := getQuotes(ctx, conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	renderJSON(w, quotes)
}

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/quotes/", getQuotesFromDb)
	http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), router)

}
