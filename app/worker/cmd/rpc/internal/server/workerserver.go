// Code generated by goctl. DO NOT EDIT.
// Source: worker.proto

package server

import (
	"context"

	"remarks_monitor/app/worker/cmd/rpc/internal/logic"
	"remarks_monitor/app/worker/cmd/rpc/internal/svc"
	"remarks_monitor/app/worker/cmd/rpc/pb"
)

type WorkerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedWorkerServer
}

func NewWorkerServer(svcCtx *svc.ServiceContext) *WorkerServer {
	return &WorkerServer{
		svcCtx: svcCtx,
	}
}

func (s *WorkerServer) Work(ctx context.Context, in *pb.WorkReq) (*pb.WorkResp, error) {
	l := logic.NewWorkLogic(ctx, s.svcCtx)
	return l.Work(in)
}
