package inventory

import (
	"golang.org/x/net/context"
	"github.com/graniticio/granitic/rdbms"
	"fmt"
)

type InventoryDAO struct {
	DBClientManager rdbms.RDBMSClientManager
}

func (id *InventoryDAO) RecordExists(ctx context.Context, album, artist string) (bool, error) {

	var db *rdbms.RDBMSClient
	var err error

	if db, err = id.DBClientManager.ClientFromContext(ctx); err != nil {
		return false, err
	}

	db.CommitTransaction()

	return false, nil

}

func (id *InventoryDAO) CreateRecord(ctx context.Context, record *RecordToCreate) error {

	var db *rdbms.RDBMSClient
	var err error

	if db, err = id.DBClientManager.ClientFromContext(ctx); err != nil {
		return err
	}

	fmt.Println("Inserting")

	p := make(map[string]interface{})
	p["catRef"] = record.CatalogRef
	p["name"] = record.Name

	//todo
	p["$artistID"] = 0


	_, err = db.InsertQueryIdParamMap("RECORD_INSERT", p)

	return err

}

