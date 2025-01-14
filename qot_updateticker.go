package futuapi

import (
	"context"

	"github.com/CCLooMi/go-futu-api/protocol"
	"github.com/futuopen/ftapi4go/pb/qotupdateticker"
)

const ProtoIDQotUpdateTicker = 3011 //Qot_UpdateTicker	推送逐笔

func init() {
	workers[ProtoIDQotUpdateTicker] = protocol.NewUpdater()
}

// 实时逐笔回调，异步处理已订阅股票的实时逐笔推送
func (api *FutuAPI) UpdateTicker(ctx context.Context) (<-chan *qotupdateticker.Response, error) {
	ch := make(chan *qotupdateticker.Response)
	if err := api.proto.RegisterUpdate(ProtoIDQotUpdateTicker, protocol.NewProtobufChan(ch)); err != nil {
		return nil, err
	}
	return ch, nil
}
