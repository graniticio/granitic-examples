package endpoint

import (
	"github.com/graniticio/granitic/ws"
	"golang.org/x/net/context"
)

type ArtistLogic struct {
}

func (al *ArtistLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {

	a := new(ArtistDetail)
	a.Name = "Hello, World!"

	res.Body = a
}

type ArtistDetail struct {
	Name string
}