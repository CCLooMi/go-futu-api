package futuapi

import (
	"context"

	"github.com/CCLooMi/go-futu-api/protocol"
	"github.com/futuopen/ftapi4go/pb/qotupdatebroker"
)

const ProtoIDQotUpdateBroker = 3015 //Qot_UpdateBroker	推送经纪队列

func init() {
	workers[ProtoIDQotUpdateBroker] = protocol.NewUpdater()
}

// 实时经纪队列回调
func (api *FutuAPI) UpdateBroker(ctx context.Context) (<-chan *qotupdatebroker.Response, error) {
	ch := make(chan *qotupdatebroker.Response)
	if err := api.proto.RegisterUpdate(ProtoIDQotUpdateBroker, protocol.NewProtobufChan(ch)); err != nil {
		return nil, err
	}
	return ch, nil
}
