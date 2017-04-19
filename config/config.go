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
	Bloodlines Bloodlines `json:"bloodlines"`
	Covenant   Covenant   `json:"covenant"`
	Warehouse  Warehouse  `json:"warehouse"`
	Coinage    Coinage    `json:"coinage"`
	Rabbit     Rabbit     `json:"rabbit"`
	Statsd     Statsd     `json:"statsd"`
	Stripe     Stripe     `json:"stripe"`
	S3         S3         `json:"s3"`
	Port       string     `json:"port"`
	Shippo     Shippo     `json:"shippo"`
	JWT        JWT        //`json:"jwt"`
}

/*MySQL contains information for connecting to a MySQL instance */
type MySQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

/*S3 conatains information for connecting to amazon s3*/
type S3 struct {
	Region    string `json:"region"`
	AccessKey string `json:"accessKey"`
	Secret    string `json:"secret"`
	Bucket    string `json:"bucket"`
}

/*Sendgrid has connection information for the sendgrid gateway */
type Sendgrid struct {
	APIKey    string `json:"api_key"`
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name"`
	Host      string `json:"host"`
}

/*Stripe has connection information for the stripe gateway */
type Stripe struct {
	Secret           string  `json:"secret"`
	Public           string  `json:"public"`
	ApplicationFee   float64 `json:"applicationFee"`
	StripeFeePercent float64 `json:"stripeFeePercent"`
}

/*TownCenter has connection information for the town center service */
type TownCenter struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

/*Bloodlines has connection information for the bloodlines service */
type Bloodlines struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

/*Coinage has connection information for the coinage service */
type Coinage struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

/*Covenant has connection information for the covenant service */
type Covenant struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

/*Warehouse has connection information for the inventory service*/
type Warehouse struct {
	Host string `json:"host"`
	Port string `json:"port"`
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

/*JWT has the jwt token*/
type JWT struct {
	Token string `json:"token"`
}

/*Shippo connection information*/
type Shippo struct {
	Token string `json:"token"`
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

	root.JWT = JWT{Token: os.Getenv("JWT_TOKEN")}
	return root, nil
}
