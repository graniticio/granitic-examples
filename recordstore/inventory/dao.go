package inventory

import (
	"golang.org/x/net/context"
	"github.com/graniticio/granitic/rdbms"
)

type InventoryDAO struct {
	DBClientManager rdbms.RDBMSClientManager
}

func (id *InventoryDAO) CatRefInUse(ctx context.Context, catRef string) (bool, error) {
	if db, err := id.DBClientManager.ClientFromContext(ctx); err != nil {
		return false, err
	} else {

		var recordId int64

		return db.SelectIDParamSingleResult("CAT_REF_SELECT", "catRef", catRef, &recordId)

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

	if err := db.FlowExistingIDOrInsertTags("ARTIST_ID_SELECT", "ARTIST_INSERT", &record.ArtistId, record); err != nil {
		return err
	}

	var recordId int64

	if err = db.InsertIDTagsAssigned("RECORD_INSERT", record, &recordId); err != nil {
		return err
	}

	for i, name := range record.Tracks {

		t := recordTrack{recordId, name, i+1}

		if _, err := db.InsertIDTags("TRACK_INSERT", t); err != nil {
			return err
		}

	}

	err = db.CommitTransaction()

	return err

}

type recordTrack struct {

	RecordId int64 `dbparam:"recordId"`
	Name     string `dbparam:"name"`
	Number   int `dbparam:"trackNumber"`

}

