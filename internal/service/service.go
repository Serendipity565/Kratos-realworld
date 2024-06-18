package service

import (
	v1 "Kratos-realworld/api/realworld/v1"
	"Kratos-realworld/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(uc *biz.UserUsecase) *RealWorldService {

	return &RealWorldService{uc: uc, log: log.NewHelper(log.GetLogger())}
}
