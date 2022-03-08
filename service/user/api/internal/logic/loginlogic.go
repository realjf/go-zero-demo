package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/realjf/go-zero-demo/service/user/api/internal/svc"
	"github.com/realjf/go-zero-demo/service/user/api/internal/types"
	"github.com/realjf/go-zero-demo/service/user/api/internal/utils"
	"github.com/realjf/go-zero-demo/service/user/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {
	if !utils.VerifyEmailFormat(req.Email) {
		return &types.LoginResp{}, fmt.Errorf("email format error: %v", req.Email)
	}
	res, err := l.svcCtx.Transformer.Login(l.ctx, &transformer.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return &types.LoginResp{}, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := utils.GenerateJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, res.Uid)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Username:     res.Username,
		Email:        res.Email,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
