package database

import (
	"database/sql"
	"golang.org/x/net/context"
	 _ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/graniticio/granitic/rdbms"
)

type DBProvider struct {
	User string
	Password string
	Host string
	Port int
	DBName string
	openDB *sql.DB

}

func (p *DBProvider) Database() (*sql.DB, error) {

	if p.openDB != nil {
		return p.openDB, nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", p.User, p.Password, p.Host, p.Port, p.DBName)

	if db ,err :=  sql.Open("mysql", dsn); err != nil {
		return nil, err
	} else {
		p.openDB = db

		p.openDB.SetConnMaxLifetime(-1)
		p.openDB.SetMaxIdleConns(10)
		p.openDB.SetMaxOpenConns(10)

		return p.openDB, nil
	}
}

func (p *DBProvider) DatabaseFromContext(ctx context.Context) (*sql.DB, error) {

	return p.Database()

}

func (p *DBProvider) InsertIDFunc() rdbms.InsertWithReturnedID {
	return rdbms.DefaultInsertWithReturnedID
}





