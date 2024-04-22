package main

import (
	"fmt"
	"net/http"
	"log"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Привет!")
}

//Генерируем сертификат с приватным ключом 
//openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem

func main() {
	http.HandleFunc("/", sayhello) // Устанавливаем роутер
	//err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера

	// Если хотите использовать https, то вместо ListenAndServe используйте ListenAndServeTLS
	 err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
