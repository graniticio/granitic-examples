package endpoint

import (
	"context"
	"github.com/graniticio/granitic/ws"
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
