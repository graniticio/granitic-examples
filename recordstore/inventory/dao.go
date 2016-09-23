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

		return db.SelectBindSingleQIDParam("CAT_REF_SELECT", "catRef", catRef, &recordId)

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

	if err := db.ExistingIDOrInsertTags("ARTIST_ID_SELECT", "ARTIST_INSERT", &record.ArtistId, record); err != nil {
		return err
	}

	var recordId int64

	if err = db.InsertCaptureQIDTags("RECORD_INSERT", record, &recordId); err != nil {
		return err
	}

	for i, name := range record.Tracks {

		t := recordTrack{recordId, name, i+1}

		if _, err := db.InsertQIDTags("TRACK_INSERT", t); err != nil {
			return err
		}

	}

	err = db.CommitTransaction()

	return err

}

func (id *InventoryDAO) ArtistSearch(ctx context.Context, sp *SearchParams) ([]*ArtistSearchResult, error) {

	if db, err := id.DBClientManager.ClientFromContext(ctx); err != nil {
		return nil, err
	} else {

		ar := new(ArtistSearchResult)

		if r, err := db.SelectBindQIDParams("ARTIST_SEARCH_BASE", make(map[string]interface{}), ar); err != nil {
			return nil, err
		} else {
			return id.artistResults(r), nil
		}
	}

}

func (id *InventoryDAO) ArtistDetail(ctx context.Context, rid *resourceId) (*ArtistDetail, error){

	if db, err := id.DBClientManager.ClientFromContext(ctx); err != nil {
		return nil, err
	} else {

		ad := new(ArtistDetail)

		 if found, err := db.SelectBindSingleQIDTags("ARTIST_DETAIL", rid, ad); found {
			 return ad, err
		 } else {
			 return nil, err
		 }
	}

}

func (id *InventoryDAO) artistResults(is []interface{}) []*ArtistSearchResult {

	ar := make([]*ArtistSearchResult, len(is))

	for i, v := range is {
		ar[i] = v.(*ArtistSearchResult)
	}

	return ar
}

type recordTrack struct {

	RecordId int64 `dbparam:"recordId"`
	Name     string `dbparam:"name"`
	Number   int `dbparam:"trackNumber"`

}

