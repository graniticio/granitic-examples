package endpoint

import (
	"context"
	"github.com/graniticio/granitic/logging"
	"github.com/graniticio/granitic/ws"
)

type ArtistLogic struct {
	EnvLabel string
	Log      logging.Logger
}

func (al *ArtistLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {

	a := new(ArtistDetail)
	a.Name = "Hello, World from " + al.EnvLabel

	res.Body = a

	l := al.Log
	l.LogInfof("Environment is set to '%s'", al.EnvLabel)
	l.LogTracef("Request served")

}

type ArtistDetail struct {
	Name string
}
