package main

import (
	"api-getaway/api"
	"api-getaway/cluster"
	"api-getaway/cluster/authservice"
	"api-getaway/cluster/storageservice"
	"api-getaway/cluster/userservice"
	"api-getaway/settings"
	"context"

	"github.com/GOAT-prod/goatlogger"
)

type App struct {
	mainCtx context.Context
	logger  goatlogger.Logger
	cfg     settings.Config

	server *api.Server

	storageServiceClient *storageservice.Client
	userServiceClient    *userservice.Client
	authServiceClient    *authservice.Client
}

func NewApp(ctx context.Context, cfg settings.Config, logger goatlogger.Logger) *App {
	return &App{
		mainCtx: ctx,
		logger:  logger,
		cfg:     cfg,
	}
}

func (a *App) Start() {
	go a.server.Start(a.cfg.Port)
}

func (a *App) Stop(_ context.Context) {}

func (a *App) initClusterClients() {
	a.storageServiceClient = storageservice.NewClient(cluster.NewBaseClient(a.cfg.Cluster.StorageServiceUrl))
	a.userServiceClient = userservice.NewClient(cluster.NewBaseClient(a.cfg.Cluster.UserServiceUrl))
	a.authServiceClient = authservice.NewClient(cluster.NewBaseClient(a.cfg.Cluster.AuthServiceUrl))
}

func (a *App) initServer() {
	a.server = api.NewServer()
}

func (a *App) initHandlers() {
	a.server.StorageServiceHandlers(a.storageServiceClient)
	a.server.UserServiceHandlers(a.userServiceClient)
	a.server.AuthServiceHandlers(a.authServiceClient)
	a.server.Swagger()
}
