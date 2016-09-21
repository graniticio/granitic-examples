package database

import (
	"database/sql"
	"golang.org/x/net/context"
)

type DBProvider struct {

}

func (p *DBProvider) Database() (*sql.DB, error) {
	return nil,nil
}

func (p *DBProvider) DatabaseFromContext(ctx context.Context) (*sql.DB, error) {
	return nil,nil
}


