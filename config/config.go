package config

import (
	"encoding/json"
	"os"
)

type Root struct {
	Sql  MySql  `json:"mysql"`
	Port string `json:"port"`
}

type MySql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func Init() (*Root, error) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)

	root := &Root{}
	err := decoder.Decode(root)
	if err != nil {
		return nil, err
	}
	return root, nil
}
