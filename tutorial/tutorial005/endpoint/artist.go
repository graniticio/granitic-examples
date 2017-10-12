package endpoint

import (
	"context"
	"github.com/graniticio/granitic/logging"
	"github.com/graniticio/granitic/types"
	"github.com/graniticio/granitic/ws"
	"strings"
)

type ArtistLogic struct {
	EnvLabel string
	Log      logging.Logger
}

func (al *ArtistLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {

	ar := req.RequestBody.(*ArtistRequest)

	a := new(ArtistDetail)
	a.Name = "Some Artist"

	res.Body = a

	l := al.Log
	l.LogTracef("Request for artist with ID %d", ar.Id)

	if ar.NormaliseName != nil && ar.NormaliseName.Bool() {
		a.Name = strings.ToUpper(a.Name)
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
