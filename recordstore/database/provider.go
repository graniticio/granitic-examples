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

}

func (p *DBProvider) Database() (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", p.User, p.Password, p.Host, p.Port, p.DBName)
	//"api:apipass@tcp(deadpixel:3306)/recordstore"

	return sql.Open("mysql", dsn)
}

func (p *DBProvider) DatabaseFromContext(ctx context.Context) (*sql.DB, error) {

	return p.Database()

}

func (p *DBProvider) InsertIDFunc() rdbms.InsertWithReturnedID {
	return MySQLInsertWithId
}

func MySQLInsertWithId(query string, client *rdbms.RDBMSClient) (int64, error) {

	if r, err := client.Exec(query); err != nil {
		return 0, err
	} else {
		if id, err := r.LastInsertId(); err != nil {

			fmt.Printf("Last ID err %s\n", err)

			return 0, err
		} else {

			fmt.Printf("Last ID found %d\n", id)

			return id, nil
		}
	}


}



