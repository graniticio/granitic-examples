package inventory

import (
	"github.com/graniticio/granitic/logging"
	"github.com/graniticio/granitic/types"
	"github.com/graniticio/granitic/ws"
	"golang.org/x/net/context"
	"net/http"
)

type ArtistLogic struct {
	Log logging.Logger
	DAO *InventoryDAO
}

func (al *ArtistLogic) Process(ctx context.Context, request *ws.WsRequest, response *ws.WsResponse) {
	id := request.RequestBody.(*resourceId)

	if result, err := al.DAO.ArtistDetail(ctx, id); err != nil {
		al.Log.LogErrorfCtx(ctx, err.Error())
	} else {

		if result == nil {
			response.HttpStatus = http.StatusNotFound
		} else {
			response.Body = result
		}

	}
}

func (al *ArtistLogic) UnmarshallTarget() interface{} {
	return new(resourceId)
}

type resourceId struct {
	ID *types.NilableInt64 `dbparam:"id"`
}

type ArtistDetail struct {
	Name string `column:"name"`
}
