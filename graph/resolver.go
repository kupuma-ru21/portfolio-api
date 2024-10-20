package graph

import (
	"portfolio-api/graph/model"
	"portfolio-api/graph/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	apps []*model.App
	Srv services.Services
}
