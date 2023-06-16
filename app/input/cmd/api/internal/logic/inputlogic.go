package logic

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"remarks_monitor/app/input/cmd/rpc/input"
	"remarks_monitor/app/usercenter/cmd/rpc/usercenter"
	"remarks_monitor/common/ctxdata"
	"remarks_monitor/common/tool"
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
	// 打开文件
	filePath := tool.GetWD() + "/data/remarks_monitor/words.txt" // 文件路径
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("无法打开文件: %s\n", err.Error())
		return
	}
	defer file.Close()

	// 逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// 比较字符串
		if line == req.Info.Content {
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
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("读取文件时出错: %s\n", err.Error())
		return nil, err
	}
	return &types.InputResponse{Message: "input failed, incorrect words"}, nil
}
