package logic

import (
	"context"

	"github.com/realjf/go-zero-demo/service/user/rpc/transform/internal/svc"
	"github.com/realjf/go-zero-demo/service/user/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *transform.UserInfoReq) (*transform.UserInfoResp, error) {
	one, err := l.svcCtx.UserModel.FindOne(in.Uid)
	if err != nil {
		return nil, err
	}

	return &transform.UserInfoResp{
		Uid:      one.Id,
		Username: one.Username,
		Email:    one.Email,
	}, nil
}
