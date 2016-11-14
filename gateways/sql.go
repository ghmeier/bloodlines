package gateways

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ghmeier/bloodlines/config"
)

type Sql struct {
	db *sql.DB
}

func NewSql(config config.MySql) (*Sql, error) {
	db, err := sql.Open(
		"mysql",
		config.User+":"+config.Password+"@tcp("+config.Host+":"+string(config.Port)+")/"+config.Database,
	)
	if err != nil {
		return nil, err
	}

	return &Sql{db: db}, nil
}

func (s *Sql) Modify(query string, values ...interface{}) error {
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

func (s *Sql) Select(query string, values ...interface{}) (*sql.Rows, error) {
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

func (s *Sql) Destroy() {
	s.db.Close()
}
