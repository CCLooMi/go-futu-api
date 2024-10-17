package futuapi

import (
	"context"

	"github.com/CCLooMi/go-futu-api/protocol"
	"github.com/futuopen/ftapi4go/pb/qotcommon"
	"github.com/futuopen/ftapi4go/pb/qotgetfutureinfo"
)

const ProtoIDQotGetFutureInfo = 3218 //Qot_GetFutureInfo	获取期货合约资料

func init() {
	workers[ProtoIDQotGetFutureInfo] = protocol.NewGetter()
}

// 获取期货合约资料
func (api *FutuAPI) GetFutureInfo(ctx context.Context, securities []*qotcommon.Security) ([]*qotgetfutureinfo.FutureInfo, error) {

	if len(securities) == 0 {
		return nil, ErrParameters
	}
	req := &qotgetfutureinfo.Request{
		C2S: &qotgetfutureinfo.C2S{
			SecurityList: securities,
		},
	}

	ch := make(chan *qotgetfutureinfo.Response)
	if err := api.proto.RegisterGet(ProtoIDQotGetFutureInfo, req, protocol.NewProtobufChan(ch)); err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return nil, ErrChannelClosed
		}
		return resp.GetS2C().GetFutureInfoList(), protocol.Error(resp)
	}
}
