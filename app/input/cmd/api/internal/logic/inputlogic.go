package logic

import (
	"context"
	"remarks_monitor/app/input/cmd/rpc/input"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
	"remarks_monitor/common/ctxdata"
	"time"

	"remarks_monitor/app/input/cmd/api/internal/svc"
	"remarks_monitor/app/input/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InputLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInputLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InputLogic {
	return &InputLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InputLogic) Input(req *types.InputRequest) (resp *types.InputResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	data, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{Id: userId})
	username := data.User.Username
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	_, err = l.svcCtx.InputRpc.Input(l.ctx, &input.InputReq{Username: username, Domain: req.Info.Domain, Content: req.Info.Content, Time: time.Now().String()})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &types.InputResponse{Message: "input success"}, nil
}
