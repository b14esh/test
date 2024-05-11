package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	port string
)

/* Не забудь создать файл .env
app_port = 8080
*/

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .enc file", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()
	log.Println("Router initilization successfully. Redy to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
