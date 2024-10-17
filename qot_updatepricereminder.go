package futuapi

import (
	"context"

	"github.com/CCLooMi/go-futu-api/protocol"
	"github.com/futuopen/ftapi4go/pb/qotupdatepricereminder"
)

const ProtoIDQotUpdatePriceReminder = 3019 //Qot_UpdatePriceReminder	到价提醒通知

func init() {
	workers[ProtoIDQotUpdatePriceReminder] = protocol.NewUpdater()
}

// 到价提醒回调
func (api *FutuAPI) UpdatePriceReminder(ctx context.Context) (<-chan *qotupdatepricereminder.Response, error) {
	ch := make(chan *qotupdatepricereminder.Response)
	if err := api.proto.RegisterUpdate(ProtoIDQotUpdatePriceReminder, protocol.NewProtobufChan(ch)); err != nil {
		return nil, err
	}
	return ch, nil
}
