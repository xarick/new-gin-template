package v1

import (
	"github.com/xarick/new-gin-template/config"
	"github.com/xarick/new-gin-template/internal/services"
)

type Controller struct {
	s   *services.Service
	cfg *config.Application
}

func NewController(s *services.Service, cfg *config.Application) *Controller {
	return &Controller{s: s, cfg: cfg}
}
