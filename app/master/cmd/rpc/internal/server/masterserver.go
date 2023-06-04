// Code generated by goctl. DO NOT EDIT.
// Source: master.proto

package server

import (
	"context"

	"remarks_monitor/app/master/cmd/rpc/internal/logic"
	"remarks_monitor/app/master/cmd/rpc/internal/svc"
	"remarks_monitor/app/master/cmd/rpc/pb"
)

type MasterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedMasterServer
}

func NewMasterServer(svcCtx *svc.ServiceContext) *MasterServer {
	return &MasterServer{
		svcCtx: svcCtx,
	}
}

func (s *MasterServer) Master(ctx context.Context, in *pb.WorkReq) (*pb.WorkResp, error) {
	l := logic.NewMasterLogic(ctx, s.svcCtx)
	return l.Master(in)
}