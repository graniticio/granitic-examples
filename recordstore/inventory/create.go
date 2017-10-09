package inventory

import (
	"github.com/graniticio/granitic/logging"
	"github.com/graniticio/granitic/types"
	"github.com/graniticio/granitic/ws"
	"golang.org/x/net/context"
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

	if err := sl.DAO.CreateRecord(ctx, r); err != nil {

		sl.Log.LogErrorfCtx(ctx, err.Error())
	}
}

func (sl *CreateRecordLogic) Validate(ctx context.Context, errors *ws.ServiceErrors, request *ws.WsRequest) {
	r := request.RequestBody.(*RecordToCreate)

	if found, err := sl.DAO.CatRefInUse(ctx, r.CatalogRef.String()); err != nil {
		sl.Log.LogErrorfCtx(ctx, err.Error())
		errors.AddPredefinedError("UNEX")

	} else if found {
		errors.AddPredefinedError("CATALOG_REF_IN_USE")
	}

}

type RecordToCreate struct {
	CatalogRef *types.NilableString `dbparam:"catRef"`
	Name       *types.NilableString `dbparam:"recordName"`
	Artist     *types.NilableString `dbparam:"artistName"`
	ArtistId   int64                `dbparam:"artistID"`
	Tracks     []string
}
