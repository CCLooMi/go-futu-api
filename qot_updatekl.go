package futuapi

import (
	"context"

	"github.com/CCLooMi/go-futu-api/protocol"
	"github.com/futuopen/ftapi4go/pb/qotupdatekl"
)

const ProtoIDQotUpdateKL = 3007 //Qot_UpdateKL	推送 K 线

func init() {
	workers[ProtoIDQotUpdateKL] = protocol.NewUpdater()
}

// 实时 K 线回调
func (api *FutuAPI) UpdateKL(ctx context.Context) (<-chan *qotupdatekl.Response, error) {
	ch := make(chan *qotupdatekl.Response)
	if err := api.proto.RegisterUpdate(ProtoIDQotUpdateKL, protocol.NewProtobufChan(ch)); err != nil {
		return nil, err
	}
	return ch, nil
}
