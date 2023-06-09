// Code generated by goctl. DO NOT EDIT.
// Source: reduce1.proto

package server

import (
	"context"

	"remarks_monitor/app/reduce1/cmd/rpc/internal/logic"
	"remarks_monitor/app/reduce1/cmd/rpc/internal/svc"
	"remarks_monitor/app/reduce1/cmd/rpc/reduce1"
)

type ReduceWork1Server struct {
	svcCtx *svc.ServiceContext
	reduce1.UnimplementedReduceWork1Server
}

func NewReduceWork1Server(svcCtx *svc.ServiceContext) *ReduceWork1Server {
	return &ReduceWork1Server{
		svcCtx: svcCtx,
	}
}

func (s *ReduceWork1Server) ReduceWork1(ctx context.Context, in *reduce1.ReduceWorkReq) (*reduce1.ReduceWorkResp, error) {
	l := logic.NewReduceWork1Logic(ctx, s.svcCtx)
	return l.ReduceWork1(in)
}
