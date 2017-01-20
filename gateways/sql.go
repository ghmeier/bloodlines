package gateways

import (
	"database/sql"
	"fmt"

	// this is standard practice for using the mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/ghmeier/bloodlines/config"
)

/*SQL describes the implementation of any sql gateway */
type SQL interface {
	Modify(string, ...interface{}) error
	Select(string, ...interface{}) (*sql.Rows, error)
	Destroy()
}

/*MySQL implements SQL with the mysql driver */
type MySQL struct {
	DB     *sql.DB
	config *config.MySQL
}

/*NewSQL returns an instance of MySQL with the given connection configuration */
func NewSQL(config config.MySQL) (*MySQL, error) {
	sql := &MySQL{config: &config}
	err := sql.connect()
	return sql, err
}

/*Modify executes any query which changes the db and doesn't return result rows */
func (s *MySQL) Modify(query string, values ...interface{}) error {
	err := s.connect()
	if err != nil {
		return err
	}

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		fmt.Printf("ERROR: unable to prepare query %s\n", query)
		s.Destroy()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		fmt.Printf("ERROR: unable to execute query %s\n", query)
		s.Destroy()
		return err
	}

	//success
	return nil
}

/*Select gets rows from a select query*/
func (s *MySQL) Select(query string, values ...interface{}) (*sql.Rows, error) {
	err := s.connect()
	if err != nil {
		return nil, err
	}

	if values == nil {
		values = make([]interface{}, 0)
	}
	rows, err := s.DB.Query(query, values...)
	if err != nil {
		fmt.Printf("ERROR: unable to run select query %s\n", query)
		s.Destroy()
		return nil, err
	}

	return rows, nil
}

/*Destroy cleans up the MySQL instance*/
func (s *MySQL) Destroy() {
	s.DB.Close()
	s.DB = nil
}

func (s *MySQL) connect() error {
	if s.DB != nil {
		return nil
	}

	db, err := sql.Open(
		"mysql",
		s.config.User+":"+s.config.Password+"@tcp("+s.config.Host+":"+string(s.config.Port)+")/"+s.config.Database+"?parseTime=true",
	)
	if err != nil {
		return err
	}

	s.DB = db
	return nil
}
