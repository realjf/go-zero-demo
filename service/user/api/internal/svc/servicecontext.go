package svc

import (
	"github.com/realjf/go-zero-demo/service/user/api/internal/config"
	"github.com/realjf/go-zero-demo/service/user/api/internal/middleware"
	"github.com/realjf/go-zero-demo/service/user/rpc/transform/transformer"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	Transformer    transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		Transformer:    transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
	}
}
