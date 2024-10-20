package services

import (
	"context"
	"portfolio-api/graph/db"
	"portfolio-api/graph/model"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
			Link:         app.Link,
			LinkType:     app.LinkType,
      ImageURL:    app.ImageUrl,
		})
	}

	return result
}

func (a *appService) GetApps(ctx context.Context) ([]*model.App, error) {
	apps, err := db.Apps().All(ctx, a.exec)

	if err != nil {
		return nil, err
	}


	return convertApps(apps), nil
}

func (a *appService) CreateApp(ctx context.Context, input model.NewApp) (*model.App, error) {
	itemID := uuid.New()
	newApp := db.App{ID: itemID.String(), Title: input.Title, Description: input.Description, Link: input.Link, LinkType: input.LinkType, ImageUrl: input.ImageURL}
	err := newApp.Insert(ctx, a.exec, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &model.App{ID: newApp.ID}, nil
}