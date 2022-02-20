//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/walrusyu/gocamp007/week4/internal/biz"
	"github.com/walrusyu/gocamp007/week4/internal/data"
	"github.com/walrusyu/gocamp007/week4/internal/service"
)

func InitUserServer() *service.Server {
	wire.Build(service.NewServer, biz.NewUserBiz, data.NewUserRepo)
	return &service.Server{}
}
