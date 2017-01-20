package config

import (
	"encoding/json"
	"os"
)

/*Root is the base struct containing bloodlines configs */
type Root struct {
	SQL        MySQL      `json:"mysql"`
	Sendgrid   Sendgrid   `json:"sendgrid"`
	TownCenter TownCenter `json:"towncenter"`
	Rabbit     Rabbit     `json:"rabbit"`
	Statsd     Statsd     `json:"statsd"`
	Port       string     `json:"port"`
}

/*MySQL contains information for connecting to a MySQL instance */
type MySQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

/*Sendgrid has connection information for the sendgrid gateway */
type Sendgrid struct {
	APIKey    string `json:"api_key"`
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name"`
	Host      string `json:"host"`
}

/*TownCenter has connection information for the town center service */
type TownCenter struct {
}

/*Rabbit has connection info for RabbitMQ*/
type Rabbit struct {
	Host string `json:"host"`
	Port string `json:"port"`
	PubQ string `json:"pubq"`
}

/*Statsd contains connection information for graphite stats*/
type Statsd struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Prefix string `json:"prefix"`
}

/*Init returns a populated Root struct from config.json */
func Init(filename string) (*Root, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)

	root := &Root{}
	err = decoder.Decode(root)
	if err != nil {
		return nil, err
	}
	return root, nil
}
