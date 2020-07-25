package main

import (
	"flag"
	"fmt"
	"log"
	"mongo-api/app"
	"mongo-api/controllers"
	"mongo-api/news"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	defAPI := os.Getenv("API_KEY")

	news.ApiKey = flag.String("apikey", defAPI, "Newsapi.org access key")
	flag.Parse()

	if *news.ApiKey == "" {
		// news.ApiKey = os.Getenv("API_KEY")
		log.Fatal("apiKey must be set")
	}

	router := mux.NewRouter()

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	router.HandleFunc("/search", news.SearchHandler).Methods("GET")

	router.HandleFunc("/", news.IndexHandler)

	router.HandleFunc("/account", news.CabinetHandler).Methods("GET")

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST", "GET")

	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthentication) // добавляем middleware проверки JWT-токена

	port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}
}
