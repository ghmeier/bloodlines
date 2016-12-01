package main

import (
	"fmt"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/router"
)

func main() {
	config, err := config.Init("config.json")
	if err != nil {
		fmt.Printf("ERROR: config initialization error. %s\n", err.Error())
		return
	}
	b, err := router.New(config)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	fmt.Printf("Bloodlines running on %s\n", config.Port)
	b.Start(":" + config.Port)
}
