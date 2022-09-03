package pkg

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

//DBOperator is an interface for sqlx operations (TX, DB)
type DBOperator interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
}

func CreateDatabase(db DBOperator) error {
	var err error
	_, err = db.Exec(sqlCreateProjectTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(sqlCreateTaskTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(sqlCreateTrackingTable)
	if err != nil {
		return err
	}

	return nil
}
