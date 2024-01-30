package bootstrap

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/new-gin-template/config"
	v1 "github.com/xarick/new-gin-template/internal/gateways/rest/v1"
	"github.com/xarick/new-gin-template/internal/services"
)

type App struct {
	cfg    config.Application
	engine *gin.Engine
}

func New(cfg config.Application) *App {
	service := services.NewService(&cfg, CreateHTTPSClient())
	ctrl := v1.NewController(service, &cfg)

	r := gin.Default()
	engine := v1.NewRouter(r, ctrl)

	app := App{
		cfg:    cfg,
		engine: engine,
	}

	return &app
}

func (app *App) Run(ctx context.Context, cfg config.Application) {
	go func() {
		err := app.engine.Run(cfg.ServerRunPort)
		if err != nil {
			log.Panic(err)
		}
	}()

	<-ctx.Done()
}

func CreateHTTPSClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			},
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{Transport: tr}
}
