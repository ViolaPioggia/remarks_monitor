package logic

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"remarks_monitor/app/input/cmd/rpc/internal/svc"
	"remarks_monitor/app/input/cmd/rpc/pb"
	"remarks_monitor/app/input/model"
	"remarks_monitor/common/tool"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type InputLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInputLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InputLogic {
	return &InputLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InputLogic) Input(in *pb.InputReq) (*pb.InputResp, error) {
	dirPath := tool.GetWD() + "/data/remarks_monitor/input" // 指定文件夹路径
	fileCount := 0

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing file %s: %v\n", path, err)
			return nil
		}

		if info.Size() >= 67108864 {
			fileCount++
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		return nil, err
	}

	content := in.Username + " " + in.Domain + " " + in.Content + " " + in.Time
	fmt.Println(content)
	file, err := os.OpenFile("C:/Users/ViolaPioggia/GolandProjects/remarks_monitor/data/remarks_monitor/input/input"+strconv.Itoa(fileCount), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644) // 打开文件
	if err != nil {
		logx.Error("open file failed")
		return nil, err
	}
	defer file.Close() // 使用完文件后关闭
	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	logx.Info("write into file success")
	r := &model.Remarks{
		Username: in.Username,
		Domain:   in.Domain,
		Content:  in.Content,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	err = l.svcCtx.RemarksModel.Insert(l.ctx, r)
	if err != nil {
		logx.Error("write remark into mongo failed ")
		return nil, err
	}
	return &pb.InputResp{}, nil
}
