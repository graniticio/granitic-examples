package database

import (
	"database/sql"
	"golang.org/x/net/context"
	 _ "github.com/go-sql-driver/mysql"
	"fmt"
)

type DBProvider struct {
	User string
	Password string
	Host string
	Port int
	DBName string

}

func (p *DBProvider) Database() (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", p.User, p.Password, p.Host, p.Port, p.DBName)
	//"api:apipass@tcp(deadpixel:3306)/recordstore"

	return sql.Open("mysql", dsn)
}

func (p *DBProvider) DatabaseFromContext(ctx context.Context) (*sql.DB, error) {

	return p.Database()

}


