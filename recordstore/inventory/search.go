package inventory

import (
	"golang.org/x/net/context"
	"github.com/graniticio/granitic/ws"
	"github.com/graniticio/granitic/types"
)

type SearchLogic struct {

}

func (sl *SearchLogic) Process(ctx context.Context, request *ws.WsRequest, response *ws.WsResponse) {

}


func (sl *SearchLogic) UnmarshallTarget() interface{} {
	return new(SearchParams)
}


func (sl *SearchLogic) Validate(ctx context.Context, errors *ws.ServiceErrors, request *ws.WsRequest){
	sp := request.RequestBody.(*SearchParams)


	if sp.Mode.String() == "artist" {

		if sp.Format.IsSet() {
			errors.AddPredefinedError("ARTIST_FORMAT")
		}

	}


}

type SearchParams struct {
	Mode *types.NilableString
	Format *types.NilableString
}