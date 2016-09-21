package inventory

import (
	"golang.org/x/net/context"
	"github.com/graniticio/granitic/ws"
	"github.com/graniticio/granitic/logging"
	"github.com/graniticio/granitic/types"
)

type CreateRecordLogic struct {
	Log logging.Logger
	DAO *InventoryDAO
}

func (sl *CreateRecordLogic) UnmarshallTarget() interface{} {
	return new(RecordToCreate)
}

func (sl *CreateRecordLogic) Process(ctx context.Context, request *ws.WsRequest, response *ws.WsResponse) {

	r := request.RequestBody.(*RecordToCreate)

	sl.Log.LogInfof("'%s'/'%s' tracks(%d)", r.Name, r.Artist, len(r.Tracks))

	if err := sl.DAO.CreateRecord(ctx, r); err != nil {

		sl.Log.LogErrorfCtx(ctx, err.Error())

	}


}

type RecordToCreate struct {

	CatalogRef *types.NilableString `dbparam:"catRef"`
	Name *types.NilableString 		`dbparam:"recordName"`
	Artist *types.NilableString		`dbparam:"artistName"`
	ArtistId int	`dbparam:"artistID"`
	Tracks []string

}
