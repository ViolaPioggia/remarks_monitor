// Code generated by goctl. DO NOT EDIT.
// Source: map1.proto

package mapwork

import (
	"context"

	"remarks_monitor/app/map1/cmd/rpc/map1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetMapWorkReq  = map1.GetMapWorkReq
	GetMapWorkResp = map1.GetMapWorkResp

	MapWork interface {
		MapWork(ctx context.Context, in *GetMapWorkReq, opts ...grpc.CallOption) (*GetMapWorkResp, error)
	}

	defaultMapWork struct {
		cli zrpc.Client
	}
)

func NewMapWork(cli zrpc.Client) MapWork {
	return &defaultMapWork{
		cli: cli,
	}
}

func (m *defaultMapWork) MapWork(ctx context.Context, in *GetMapWorkReq, opts ...grpc.CallOption) (*GetMapWorkResp, error) {
	client := map1.NewMapWorkClient(m.cli.Conn())
	return client.MapWork(ctx, in, opts...)
}