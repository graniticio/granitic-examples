package endpoint

import (
	"github.com/graniticio/granitic/ws"
	"context"
)

type ArtistLogic struct {
	EnvLabel string
}

func (al *ArtistLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {

	a := new(ArtistDetail)
	a.Name = "Hello, World from " + al.EnvLabel

	res.Body = a
}

type ArtistDetail struct {
	Name string
}