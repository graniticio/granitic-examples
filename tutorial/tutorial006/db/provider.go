package db
import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/graniticio/granitic/logging"
	"github.com/graniticio/granitic/rdbms"
)
type MySqlProvider struct {
	Config *mysql.Config
	Log logging.Logger
}
func (p *MySqlProvider) Database() (*sql.DB, error) {
	dsn := p.Config.FormatDSN()
	if db, err := sql.Open("mysql", dsn); err == nil {
		return db, nil
	} else {
		p.Log.LogErrorf("Unable to open connection to MySQL database: %v", err)
		return nil, err
	}
}
func (p *MySqlProvider) InsertIDFunc() rdbms.InsertWithReturnedID {
	return rdbms.DefaultInsertWithReturnedID
}