package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	idInt, err1 := strconv.Atoi(idStr)
	if err1 != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}
	if idInt < 1 {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	detailsStr := r.URL.Query().Get("details")
	details, err2 := strconv.ParseBool(detailsStr)
	if err2 != nil && detailsStr != "" {
		http.Error(w, "Некорректный параметр details", http.StatusBadRequest)
		return
	}

	switch details {
	case true:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "order id: %d, details: %v", idInt, details)
		return
	default:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "order id: %d, details:", idInt)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /orders/{id}", getOrderHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
