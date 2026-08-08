package webHttp

import (
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
)

type WebController struct {
	webService WebService
}

type WebService interface {
	GetMainPage() ([]byte, error)
}

func NewWebController(webService WebService) *WebController {
	return &WebController{
		webService: webService,
	}
}

func (c *WebController) Routes() []server.Route {
	return []server.Route{
		{
			Path:    "/",
			Handler: c.GetMainPage,
		},
	}
}
