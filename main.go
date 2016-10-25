package main

import (
	"net/http"
	"fmt"
	"os"
)

type BloodlinesHandler struct {

}

func (b *BloodlinesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func getServiceHandler() *BloodlinesHandler {
	return &BloodlinesHandler{}
}

func main() {
	service := getServiceHandler()

	port := os.Getenv("PORT");
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Bloodlines running on %s\n", port)
	err := http.ListenAndServe(":"+port, service)
	fmt.Println(err.Error())
}