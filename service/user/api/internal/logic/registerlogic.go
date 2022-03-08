package logic

import (
	"context"
	"fmt"

	"github.com/realjf/go-zero-demo/service/user/api/internal/svc"
	"github.com/realjf/go-zero-demo/service/user/api/internal/types"
	"github.com/realjf/go-zero-demo/service/user/api/internal/utils"
	"github.com/realjf/go-zero-demo/service/user/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (resp *types.RegisterResp, err error) {
	if !utils.VerifyEmailFormat(req.Email) {
		return &types.RegisterResp{}, fmt.Errorf("email format error: %v", req.Email)
	}
	_, err = l.svcCtx.Transformer.Register(l.ctx, &transformer.RegisterReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return &types.RegisterResp{}, err
	}

	return &types.RegisterResp{}, nil
}
