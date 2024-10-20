package services

import (
	"context"
	"portfolio-api/graph/db"
	"portfolio-api/graph/model"

	"github.com/volatiletech/sqlboiler/boil"
)

type appService struct {
	exec boil.ContextExecutor
}

func convertApps(apps db.AppSlice) []*model.App {
	var result []*model.App

	for _, app := range apps {
		result = append(result, &model.App{
			ID:          app.ID,
			Title:       app.Title,
			Description: app.Description,
			Link:         app.URL,
			URLType:     app.UrlType,
		})
	}

	return result
}

func (u *appService) GetApps(ctx context.Context) ([]*model.App, error) {
	apps, err := db.Apps().All(ctx, u.exec)

	if err != nil {
		return nil, err
	}


	return convertApps(apps), nil
}