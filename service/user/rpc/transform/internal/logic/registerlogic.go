package logic

import (
	"context"
	"errors"

	"github.com/realjf/go-zero-demo/service/user/model"
	"github.com/realjf/go-zero-demo/service/user/rpc/transform/internal/svc"
	"github.com/realjf/go-zero-demo/service/user/rpc/transform/internal/utils"
	"github.com/realjf/go-zero-demo/service/user/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *transform.RegisterReq) (*transform.RegisterResp, error) {
	one, err := l.svcCtx.UserModel.FindByEmail(in.Email)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	} else if one != nil {
		return nil, errors.New("email conflict")
	}

	// 密码长度
	if len(in.Password) < 6 {
		return nil, errors.New("password too short")
	}

	encodePwd, err := utils.EncodePassword(in.Password)
	if err != nil {
		return nil, err
	}

	// 注册用户
	data := &model.User{
		Username: in.Username,
		Email:    in.Email,
		Password: encodePwd,
	}
	res, err := l.svcCtx.UserModel.Insert(data)
	if err != nil {
		return nil, err
	}

	lastId, _ := res.LastInsertId()

	return &transform.RegisterResp{
		Uid:      lastId,
		Username: in.Username,
		Email:    in.Email,
	}, nil
}
