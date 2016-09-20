package inventory

import (
	"golang.org/x/net/context"
	"github.com/graniticio/granitic/ws"
	"github.com/graniticio/granitic/logging"
	"github.com/graniticio/granitic/types"
)

type CreateRecordLogic struct {
	Log logging.Logger
}

func (sl *CreateRecordLogic) UnmarshallTarget() interface{} {
	return new(RecordToCreate)
}

func (sl *CreateRecordLogic) Process(ctx context.Context, request *ws.WsRequest, response *ws.WsResponse) {

	r := request.RequestBody.(*RecordToCreate)

	sl.Log.LogInfof("'%s'/'%s' tracks(%d)", r.Name, r.Artist, len(r.Tracks))


}

type RecordToCreate struct {

	CatalogRef *types.NilableString
	Name *types.NilableString
	Artist *types.NilableString
	Tracks []string

}


