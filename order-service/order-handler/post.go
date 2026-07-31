package main

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	Id        int64
	FirstName string
	LastName  string
	Gender    string
	City      string
	Age       byte
}

func CreateClient(firstName string, lastName string, gender string, city string, age byte) *Client {
	return &Client{
		Id:        13, //имитируем, что здесь подтягивается из бд id
		FirstName: firstName,
		LastName:  lastName,
		Gender:    gender,
		City:      city,
		Age:       age,
	}
}

type ClientRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	City      string `json:"city"`
	Age       byte   `json:"age"`
}

// Обычная функция-обработчик
func createClientHandler(w http.ResponseWriter, r *http.Request) {
	// r — содержит данные запроса (URL, метод, заголовки, тело)
	// w — используется для записи ответа клиенту
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req ClientRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "incorrect json", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	createClient:= CreateClient(req.FirstName, req.LastName, req.Gender, req.City, req.Age)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createClient)
}

func main() {
	// Регистрируем хендлер для пути "/"
	http.HandleFunc("/clients", createClientHandler)

	// Запускаем сервер на порту 8080
	http.ListenAndServe(":8080", nil)

}
