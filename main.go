package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Структура для животных
type Animal struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Age  int    `json:"age"`
}

// Данные животных
var animals = map[int]Animal{
	1: {ID: 1, Name: "Лев", Type: "Хищник", Age: 5},
	2: {ID: 2, Name: "Тигр", Type: "Хищник", Age: 4},
	3: {ID: 3, Name: "Слон", Type: "Травоядное", Age: 10},
	4: {ID: 4, Name: "Жираф", Type: "Травоядное", Age: 7},
	5: {ID: 5, Name: "Котик", Type: "Домашний", Age: 2},
}

// Обработчик для главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

// Обработчик для получения списка всех животных
func getAnimals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	animalList := []Animal{}
	for _, animal := range animals {
		animalList = append(animalList, animal)
	}
	json.NewEncoder(w).Encode(animalList)
}

// Обработчик для создания нового животного
func createAnimal(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		animalType := r.FormValue("type")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			http.Error(w, "Возраст должен быть числом", http.StatusBadRequest)
			return
		}

		newAnimal := Animal{
			ID:   len(animals) + 1,
			Name: name,
			Type: animalType,
			Age:  age,
		}
		animals[newAnimal.ID] = newAnimal

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// Главная функция
func main() {
	r := mux.NewRouter()

	// Статические файлы
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Маршруты
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/animals", getAnimals).Methods("GET")
	r.HandleFunc("/animals", createAnimal).Methods("POST")

	fmt.Println("Сервер запущен на порту :8080")
	http.ListenAndServe(":8080", r)
}
