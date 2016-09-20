package inventory

import (
	"golang.org/x/net/context"
	"github.com/graniticio/granitic/ws"
)

type CreateRecordLogic struct {

}

func (sl *CreateRecordLogic) Process(ctx context.Context, request *ws.WsRequest, response *ws.WsResponse) {

}

type RecordToCreate struct {

	Name string
	Artist string
	Tracks []string

}


