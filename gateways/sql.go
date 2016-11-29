package gateways

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ghmeier/bloodlines/config"
)

type Sql interface {
	Modify(string, ...interface{}) error
	Select(string, ...interface{}) (*sql.Rows, error)
	Destroy()
}

type Mysql struct {
	db *sql.DB
}

func NewSql(config config.MySql) (*Mysql, error) {
	db, err := sql.Open(
		"mysql",
		config.User+":"+config.Password+"@tcp("+config.Host+":"+string(config.Port)+")/"+config.Database,
	)
	if err != nil {
		return nil, err
	}

	return &Mysql{db: db}, nil
}

func (s *Mysql) Modify(query string, values ...interface{}) error {
	stmt, err := s.db.Prepare(query)
	if err != nil {
		fmt.Printf("ERROR: unable to prepare query %s\n", query)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		fmt.Printf("ERROR: unable to execute query %s\n", query)
		return err
	}

	//success
	return nil
}

func (s *Mysql) Select(query string, values ...interface{}) (*sql.Rows, error) {
	if values == nil {
		values = make([]interface{}, 0)
	}
	rows, err := s.db.Query(query, values...)
	if err != nil {
		fmt.Printf("ERROR: unable to run select query %s\n", query)
		return nil, err
	}

	return rows, nil
}

func (s *Mysql) Destroy() {
	s.db.Close()
}
