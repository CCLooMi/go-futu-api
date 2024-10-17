package futuapi

import (
	"context"

	"github.com/CCLooMi/go-futu-api/protocol"
	"github.com/futuopen/ftapi4go/pb/getoptionexpirationdate"
	"github.com/futuopen/ftapi4go/pb/qotcommon"
	"google.golang.org/protobuf/proto"
)

const ProtoIDQotGetOptionExpirationDate = 3224 //Qot_GetOptionExpirationDate 获取期权到期日

func init() {
	workers[ProtoIDQotGetOptionExpirationDate] = protocol.NewGetter()
}

func (api *FutuAPI) GetOptionExpirationDate(ctx context.Context, owner *qotcommon.Security,
	indexOptType qotcommon.IndexOptionType) ([]*getoptionexpirationdate.OptionExpirationDate, error) {

	if owner == nil {
		return nil, ErrParameters
	}
	req := &getoptionexpirationdate.Request{
		C2S: &getoptionexpirationdate.C2S{
			Owner: owner,
		},
	}
	if indexOptType != qotcommon.IndexOptionType_IndexOptionType_Unknown {
		req.C2S.IndexOptionType = proto.Int32(int32(indexOptType))
	}

	ch := make(chan *getoptionexpirationdate.Response)
	if err := api.proto.RegisterGet(ProtoIDQotGetOptionExpirationDate, req, protocol.NewProtobufChan(ch)); err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ErrInterrupted
	case resp, ok := <-ch:
		if !ok {
			return nil, ErrChannelClosed
		}
		return resp.GetS2C().GetDateList(), protocol.Error(resp)
	}
}
