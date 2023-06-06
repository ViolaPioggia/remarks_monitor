// Code generated by goctl. DO NOT EDIT.
// Source: reduce.proto

package server

import (
	"context"

	"remarks_monitor/app/reduce/cmd/rpc/internal/logic"
	"remarks_monitor/app/reduce/cmd/rpc/internal/svc"
	"remarks_monitor/app/reduce/cmd/rpc/pb"
)

type ReduceWorkServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedReduceWorkServer
}

func NewReduceWorkServer(svcCtx *svc.ServiceContext) *ReduceWorkServer {
	return &ReduceWorkServer{
		svcCtx: svcCtx,
	}
}

func (s *ReduceWorkServer) ReduceWork(ctx context.Context, in *pb.ReduceWorkReq) (*pb.ReduceWorkResp, error) {
	l := logic.NewReduceWorkLogic(ctx, s.svcCtx)
	return l.ReduceWork(in)
}
