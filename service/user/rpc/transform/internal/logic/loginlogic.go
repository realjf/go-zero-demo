package logic

import (
	"context"
	"errors"

	"github.com/realjf/go-zero-demo/service/user/rpc/transform/internal/svc"
	"github.com/realjf/go-zero-demo/service/user/rpc/transform/internal/utils"
	"github.com/realjf/go-zero-demo/service/user/rpc/transform/transform"

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

func (l *LoginLogic) Login(in *transform.LoginReq) (*transform.LoginResp, error) {

	one, err := l.svcCtx.UserModel.FindByEmail(in.Email)
	if err != nil {
		return nil, err
	}

	// 检查密码
	if !utils.CheckPassword(in.Password, one.Password) {
		return nil, errors.New("password check failed")
	}

	return &transform.LoginResp{
		Uid:      one.Id,
		Username: one.Username,
		Email:    one.Email,
	}, nil
}
