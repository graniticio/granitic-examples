package inventory

import (
	"golang.org/x/net/context"
	"github.com/graniticio/granitic/rdbms"
)

type InventoryDAO struct {
	DBClientManager rdbms.RDBMSClientManager
}

func (id *InventoryDAO) CreateRecord(ctx context.Context, record *RecordToCreate) error {

	var db *rdbms.RDBMSClient
	var err error

	if db, err = id.DBClientManager.ClientFromContext(ctx); err != nil {
		return err
	}

	db.StartTransaction()
	defer db.Rollback()

	if err := db.FlowExistingIDOrInsertTags("ARTIST_ID_SELECT", "ARTIST_INSERT", &record.ArtistId, record); err != nil {
		return err
	}

	if _, err = db.InsertIDTags("RECORD_INSERT", record); err != nil {
		return err
	}

	err = db.CommitTransaction()

	return err

}

