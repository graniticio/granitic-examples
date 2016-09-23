package inventory

import (
	"golang.org/x/net/context"
	"github.com/graniticio/granitic/ws"
	"github.com/graniticio/granitic/types"
	"github.com/graniticio/granitic/logging"
)

type SearchLogic struct {
	Log logging.Logger
	DAO *InventoryDAO
}

func (sl *SearchLogic) Process(ctx context.Context, request *ws.WsRequest, response *ws.WsResponse) {
	sp := request.RequestBody.(*SearchParams)

	if results, err := sl.DAO.ArtistSearch(ctx, sp); err != nil {
		sl.Log.LogErrorfCtx(ctx, err.Error())
	} else {

		response.Body = results

	}
}


func (sl *SearchLogic) UnmarshallTarget() interface{} {
	return new(SearchParams)
}


func (sl *SearchLogic) Validate(ctx context.Context, errors *ws.ServiceErrors, request *ws.WsRequest){
	sp := request.RequestBody.(*SearchParams)


	if sp.Mode.String() == "artist" {

		if sp.Format != nil && sp.Format.IsSet() {
			errors.AddPredefinedError("ARTIST_FORMAT")
		}

	}


}

type SearchParams struct {
	Mode *types.NilableString
	Format *types.NilableString
}

type ArtistSearchResult struct {

	Name *types.NilableString
	ID *types.NilableInt64
	Active *types.NilableBool
	Weighting float64

}