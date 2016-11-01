package gateways

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Sql struct{
	db *sql.DB
}

func NewSql() (*Sql, error) {
	db, err := sql.Open("mysql", "root:bloodlines@tcp(172.17.0.2:3306)/bloodlines")
	if err != nil {
		return nil, err
	}


	return &Sql{db: db}, nil
}

func (s *Sql) Modify(query string, values ...interface{}) error {
	stmt, err := s.db.Prepare(query)
	if err != nil {
		fmt.Printf("ERROR: unable to prepare query %s\n",query)
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
	rows, err := s.db.Query(query, values)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (s *Sql) Destroy() {
	s.db.Close()
}