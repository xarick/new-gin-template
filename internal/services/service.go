package services

import (
	"net/http"

	"github.com/xarick/new-gin-template/config"
)

type Service struct {
	cfg      *config.Application
	CheckAPI *CheckService
}

func NewService(cfg *config.Application, httpClient *http.Client) *Service {
	return &Service{
		cfg:      cfg,
		CheckAPI: NewCheckService(httpClient, cfg),
	}
}
