package inventory

import (
	"golang.org/x/net/context"
	"github.com/graniticio/granitic/rdbms"
	"database/sql"
)

type InventoryDAO struct {
	DBClientManager rdbms.RDBMSClientManager
}

func (id *InventoryDAO) ArtistExists(ctx context.Context, name string, client ...*rdbms.RDBMSClient) (bool, int64, error) {

	var db *rdbms.RDBMSClient
	var err error
	var rows *sql.Rows

	if len(client) > 0 {
		db = client[0]
	}else if db, err = id.DBClientManager.ClientFromContext(ctx); err != nil {
		return false, 0,  err
	}

	rows, err = db.SelectIDParam("ARTIST_ID_SELECT", "artistName", name)

	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return false, 0, err
	} else {

		if rows.Next() {

			var id int64
			rows.Scan(&id)

			return true, id, nil
		} else {
			return false, 0, nil
		}

	}
}

func (id *InventoryDAO) CreateRecord(ctx context.Context, record *RecordToCreate) error {

	var db *rdbms.RDBMSClient
	var err error

	if db, err = id.DBClientManager.ClientFromContext(ctx); err != nil {
		return err
	}

	db.StartTransaction()
	defer db.Rollback()

	if exists, id, err := id.ArtistExists(ctx, record.Artist.String(), db); err != nil {
		return err
	} else if exists {
		record.ArtistId = id
	} else {
		if id, err = db.InsertIDTagsAssigned("ARTIST_INSERT", record); err != nil {
			return err
		} else {
			record.ArtistId = id
		}

	}

	if _, err = db.InsertIDTags("RECORD_INSERT", record); err != nil {
		return err
	}

	err = db.CommitTransaction()

	return err

}

