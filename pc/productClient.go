package pc

import (
	"context"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"time"
)

//启用全局链接
type Client struct {
	Conn *grpc.ClientConn
	Ctx  context.Context
	Cf   context.CancelFunc
	RPC  DcProductClient
}

//全局grpc链接
var client = new(Client)

// 获取product的链接
func GetDcProductGrpcClient(url string, c ...echo.Context) *Client {
	var err error
	if len(url) <= 0 {
		url = "127.0.0.1:8802"
	}
	if client.Conn, err = grpc.Dial(url, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(50000000))); err != nil {
		glog.Error(err)
		return nil
	} else {
		client.RPC = NewDcProductClient(client.Conn)
		//添加用户信息
		//client.Ctx = AppendToOutgoingContextLoginUserInfo(context.Background(), c...)
		client.Ctx = context.Background()
		client.Ctx, client.Cf = context.WithTimeout(client.Ctx, time.Minute*30)
		return client
	}
}

//全局的不需要关闭链接
func (d *Client) Close() {
	//d.Conn.Close()
	//d.Cf()
}
