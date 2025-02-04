package futuapi

import (
	"context"

	"github.com/CCLooMi/go-futu-api/protocol"
	"github.com/futuopen/ftapi4go/pb/qotupdateorderbook"
)

const ProtoIDQotUpdateOrderBook = 3013 //Qot_UpdateOrderBook	推送买卖盘

func init() {
	workers[ProtoIDQotUpdateOrderBook] = protocol.NewUpdater()
}

// 实时摆盘回调
func (api *FutuAPI) UpdateOrderBook(ctx context.Context) (<-chan *qotupdateorderbook.Response, error) {
	ch := make(chan *qotupdateorderbook.Response)
	if err := api.proto.RegisterUpdate(ProtoIDQotUpdateOrderBook, protocol.NewProtobufChan(ch)); err != nil {
		return nil, err
	}
	return ch, nil
}
