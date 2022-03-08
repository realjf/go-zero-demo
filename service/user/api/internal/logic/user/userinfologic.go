package user

import (
	"context"
	"encoding/json"

	"github.com/realjf/go-zero-demo/service/user/api/internal/svc"
	"github.com/realjf/go-zero-demo/service/user/api/internal/types"
	"github.com/realjf/go-zero-demo/service/user/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	res, err := l.svcCtx.Transformer.GetUser(l.ctx, &transformer.UserInfoReq{
		Uid: uid,
	})

	if err != nil {
		return &types.UserInfoResp{}, err
	}

	return &types.UserInfoResp{
		Uid:      res.Uid,
		Username: res.Username,
		Email:    res.Email,
	}, nil
}
