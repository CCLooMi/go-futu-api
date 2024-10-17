package futuapi

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/futuopen/ftapi4go/pb/qotcommon"
	"google.golang.org/protobuf/proto"
)

func printJson(t *testing.T, v interface{}) {
	js, _ := json.Marshal(v)
	t.Log(string(js))
}
func TestConnect(t *testing.T) {

	api := NewFutuAPI()
	defer api.Close(context.Background())
	api.SetClientInfo("1000", 1)

	if err := api.Connect(context.Background(), ":11111"); err != nil {
		t.Error(err)
		return
	}

	api.SetRecvNotify(true)
	nCh, err := api.SysNotify(context.Background())
	if err != nil {
		t.Error(err)
	}

	if sub, err := api.QuerySubscription(context.Background(), true); err != nil {
		t.Error(err)
	} else {
		printJson(t, sub)
	}

	tCh, err := api.UpdateTicker(context.Background())
	if err != nil {
		t.Error(err)
	}
	if err := api.Subscribe(context.Background(), []*qotcommon.Security{
		{Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_CNSZ_Security)),
			Code: proto.String("002352")},
	},
		[]qotcommon.SubType{qotcommon.SubType_SubType_Ticker},
		true, true, true, true); err != nil {
		t.Error(err)
	}
	select {
	case notify := <-nCh:
		t.Log(notify)
	case ticker := <-tCh:
		printJson(t, ticker)
	}

	if sub, err := api.QuerySubscription(context.Background(), true); err != nil {
		t.Error(err)
	} else {
		printJson(t, sub)
	}

	secs, err := api.GetUserSecurity(context.Background(), "全部")
	if err != nil {
		t.Error(err)
	} else {
		for _, sec := range secs {
			//js, _ := json.Marshal(sec.Basic)
			//t.Log(string(js))
			printJson(t, sec)
		}
	}

	nw := time.Now()
	tds, err := api.RequestTradingDays(context.Background(),
		qotcommon.TradeDateMarket_TradeDateMarket_CN,
		fmt.Sprintf("%d-1-1", nw.Year()), fmt.Sprintf("%d-12-31", nw.Year()),
		&qotcommon.Security{
			Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_CNSZ_Security)),
			Code:   proto.String("002352"),
		})

	if err != nil {
		t.Error(err)
	} else {
		for _, td := range tds {
			js, _ := json.Marshal(td)
			t.Log(string(js))
		}
	}
}
