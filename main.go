package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/ghmeier/bloodlines/config"
	"github.com/ghmeier/bloodlines/router"
)

func main() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	config, err := config.Init(path.Join(dir, "config.json"))
	if err != nil {
		fmt.Printf("ERROR: config initialization error.\n%s\n", err.Error())
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
