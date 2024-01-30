package services

import (
	"context"
	"net/http"

	"github.com/xarick/new-gin-template/config"
)

type CheckService struct {
	client *http.Client
	cfg    *config.Application
}

func NewCheckService(client *http.Client, cfg *config.Application) *CheckService {
	return &CheckService{client: client, cfg: cfg}
}

func (ch *CheckService) CheckFunc(ctx context.Context, token string) (string, error) {
	return "Success", nil
}
