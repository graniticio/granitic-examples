package endpoint

import (
	"context"
	"github.com/graniticio/granitic/logging"
	"github.com/graniticio/granitic/types"
	"github.com/graniticio/granitic/ws"
	"github.com/graniticio/granitic/rdbms"
	"net/http"

)

type ArtistLogic struct {
	EnvLabel string
	Log      logging.Logger
	DbClientManager rdbms.RdbmsClientManager
}

func (al *ArtistLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {
	ar := req.RequestBody.(*ArtistRequest)

	// Obtain an RdmsClient from the rdbms.RdbmsClientManager injected into this component
	dbc, _ := al.DbClientManager.Client()

	// Create a new object to store the results of our database call
	result := new(ArtistDetail)

	// Call the database and populate our object
	if found, err := dbc.SelectBindSingleQIdParams("ARTIST_BY_ID", result, ar); found {
		// Make our result object the body of the HTTP response we'll send
		res.Body = result

	} else if err != nil{
		// Something went wrong when communicating with the database - return HTTP 500
		al.Log.LogErrorf(err.Error())
		res.HttpStatus = http.StatusInternalServerError

	} else {
		// No results were returned by the database call - return HTTP 404
		res.HttpStatus = http.StatusNotFound
	}
}

func (al *ArtistLogic) UnmarshallTarget() interface{} {
	return new(ArtistRequest)
}

type ArtistDetail struct {
	Name string
}

type ArtistRequest struct {
	Id            int
	NormaliseName *types.NilableBool
}

type SubmitArtistLogic struct {
	Log logging.Logger
	DbClientManager rdbms.RdbmsClientManager
}

func (sal *SubmitArtistLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {

	sar := req.RequestBody.(*SubmittedArtistRequest)

	// Obtain an RdmsClient from the rdbms.RdbmsClientManager injected into this component
	dbc, _ := sal.DbClientManager.Client()
	defer dbc.Rollback()

	// Start a database transaction
	dbc.StartTransaction()

	// Declare a variable to capture the ID of the newly inserted artist
	var id int64

	// Execute the insert, storing the generated ID in our variable
	if err := dbc.InsertCaptureQIdParams("CREATE_ARTIST", &id, sar); err != nil {
		// Something went wrong when communicating with the database - return HTTP 500
		sal.Log.LogErrorf(err.Error())
		res.HttpStatus = http.StatusInternalServerError

		return

	}

	// Insert a row for each related artist
	params := make(map[string]interface{})
	params["ArtistId"] = id

	for _, raId := range sar.RelatedArtists {
		params["RelatedArtistId"] = raId

		if _, err := dbc.InsertQIdParams("RELATE_ARTIST", params); err != nil {
			// Something went wrong inserting the relationship
			sal.Log.LogErrorf(err.Error())
			res.HttpStatus = http.StatusInternalServerError

			return
		}

	}

	// Commit the transaction
	dbc.CommitTransaction()

	// Use the new ID as the HTTP response, wrapped in a struct
	res.Body = struct {
		Id int64
	}{id}

}

func (sal *SubmitArtistLogic) UnmarshallTarget() interface{} {
	return new(SubmittedArtistRequest)
}

type SubmittedArtistRequest struct {
	Name            *types.NilableString
	FirstYearActive *types.NilableInt64
	RelatedArtists []int64
}
