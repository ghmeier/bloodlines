package main

import (
	"fmt"
	"os"

	"github.com/ghmeier/bloodlines/router"
)

func main() {
	b := router.New()

	port := os.Getenv("PORT");
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Bloodlines running on %s\n", port)
	b.Start(":"+port)
}
