package futuapi

import (
	"context"

	"github.com/CCLooMi/go-futu-api/protocol"
	"github.com/futuopen/ftapi4go/pb/qotcommon"
	"github.com/futuopen/ftapi4go/pb/qotgetplatesecurity"
	"google.golang.org/protobuf/proto"
)

const ProtoIDQotGetPlateSecurity = 3205 //Qot_GetPlateSecurity	获取板块下的股票

func init() {
	workers[ProtoIDQotGetPlateSecurity] = protocol.NewGetter()
}

// 获取板块内股票列表
func (api *FutuAPI) GetPlateStock(ctx context.Context, plate *qotcommon.Security, sortField qotcommon.SortField, ascend bool) ([]*qotcommon.SecurityStaticInfo, error) {

	if plate == nil {
		return nil, ErrParameters
	}
	req := &qotgetplatesecurity.Request{
		C2S: &qotgetplatesecurity.C2S{
			Plate:  plate,
			Ascend: proto.Bool(ascend),
		},
	}
	if sortField != qotcommon.SortField_SortField_Unknow {
		req.C2S.SortField = proto.Int32(int32(sortField))
	}

	ch := make(chan *qotgetplatesecurity.Response)
	if err := api.proto.RegisterGet(ProtoIDQotGetPlateSecurity, req, protocol.NewProtobufChan(ch)); err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		return nil, ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return nil, ErrChannelClosed
		}
		return resp.GetS2C().GetStaticInfoList(), protocol.Error(resp)
	}
}
