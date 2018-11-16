package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//MysqlConfig has the attributes for making the mysql connection
type MysqlConfig struct {
	MysqlAddr     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDB       string
}


type DataSource interface {
	Exec(strSQL string, args ...interface{}) error
	Query(strSQL string, args ...interface{}) (*sql.Rows, error)
	QueryRow(strSQL string, args ...interface{}) (*sql.Row, error)

}


//DataSource provides the basic database operations
type SQLDataSource struct {
	db *sql.DB
}

func NewSQLDataSource(db  *sql.DB) *SQLDataSource{
	return &SQLDataSource{db:db}
}

//Exec is a basic function for updating, deleting, inserting data into database
func (ds *SQLDataSource) Exec(strSQL string, args ...interface{}) error {
	stmt, err := ds.db.Prepare(strSQL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if args == nil {
		args = []interface{}{}
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//Query is a basic function for querying data from database
func (ds *SQLDataSource) Query(strSQL string, args ...interface{}) (*sql.Rows, error) {
	rows, err := ds.db.Query(strSQL, args...)
	if err != nil {
		fmt.Println(err)
		return rows, err
	}
	return rows, nil
}

//Query is a basic function for querying data from database
func (ds *SQLDataSource) QueryRow(strSQL string, args ...interface{}) (*sql.Row, error) {
	row := ds.db.QueryRow(strSQL, args...)
	if row == nil {
		return row, errors.New("no data")
	}
	return row, nil
}
