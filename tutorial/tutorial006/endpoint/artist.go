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
	l := al.Log
	l.LogTracef("Request for artist with ID %d", ar.Id)
	result := new(ArtistDetail)
	dbc, _ := al.DbClientManager.Client()
	if found, err := dbc.SelectBindSingleQIDParams("ARTIST_BY_ID", result, ar); found {
		res.Body = result
	} else if err != nil{
		l.LogErrorf(err.Error())
		res.HttpStatus = http.StatusInternalServerError
	} else {
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
}

func (sal *SubmitArtistLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {

	sar := req.RequestBody.(*SubmittedArtistRequest)

	sal.Log.LogInfof("New artist %s", sar.Name)

	//Hardcoded 'ID' of newly created artist - just a placeholder
	res.Body = struct {
		Id int
	}{0}

}

func (sal *SubmitArtistLogic) UnmarshallTarget() interface{} {
	return new(SubmittedArtistRequest)
}

type SubmittedArtistRequest struct {
	Name            *types.NilableString
	FirstYearActive *types.NilableInt64
}
