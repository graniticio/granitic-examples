package db

import (
	"github.com/graniticio/granitic/rdbms"
	"github.com/graniticio/granitic/logging"
)

type ArtistExistsChecker struct{
	DbClientManager rdbms.RdbmsClientManager
	Log logging.Logger
}

func (aec *ArtistExistsChecker) ValidInt64(id int64) (bool, error) {

	dbc, _ := aec.DbClientManager.Client()

	var count int64

	if _, err := dbc.SelectBindSingleQIdParam("CHECK_ARTIST", "Id", id, &count); err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}