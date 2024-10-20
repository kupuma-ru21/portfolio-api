package services

import (
	"context"
	"portfolio-api/graph/model"

	"github.com/volatiletech/sqlboiler/boil"
)

type AppService interface {
	GetApps(ctx context.Context) ([]*model.App, error)
}

type Services interface {
	AppService
}

type services struct {
	*appService
}



func New(exec boil.ContextExecutor) Services {
	return &services{
		appService: &appService{exec: exec},
	}
}