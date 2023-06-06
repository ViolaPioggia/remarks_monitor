// Code generated by goctl. DO NOT EDIT.
// Source: MapReduce.proto

package server

import (
	"context"

	"remarks_monitor/app/map/cmd/rpc/internal/logic"
	"remarks_monitor/app/map/cmd/rpc/internal/svc"
	"remarks_monitor/app/map/cmd/rpc/pb"
)

type MapServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedMapWorkServer
}

func NewMapServer(svcCtx *svc.ServiceContext) *MapServer {
	return &MapServer{
		svcCtx: svcCtx,
	}
}

func (s *MapServer) MapWork(ctx context.Context, in *pb.GetMapWorkReq) (*pb.GetMapWorkResp, error) {
	l := logic.NewMapWorkLogic(ctx, s.svcCtx)
	return l.MapWork(in)
}
