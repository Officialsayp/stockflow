package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}
	if idInt < 1 {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	qp := r.URL.Query().Get("details")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "order id: %d", idInt)
	return
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /orders/{id}", getOrderHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
