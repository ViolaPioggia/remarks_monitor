package logic

import (
	"context"
	"github.com/pkg/errors"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
	"remarks_monitor/app/usercenter/model"
	"remarks_monitor/common/tool"

	"remarks_monitor/app/usercenter/cmd/rpc/internal/svc"
	"remarks_monitor/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var userId int64
	var err error

	userId, err = l.loginByUsername(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &usercenter.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByUsername(username, password string) (int64, error) {

	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, username)
	if err != nil && err != model.ErrNotFound {
		return 0, err
	}
	if user == nil {
		return 0, err
	}

	if !(tool.Md5ByString(password) == user.Password) {
		return 0, errors.New("密码错误")
	}

	return user.Id, nil
}
