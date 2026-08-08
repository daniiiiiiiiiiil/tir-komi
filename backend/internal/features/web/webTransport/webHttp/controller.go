package webHttp

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
)

type WebController struct {
	webService WebService
}

type WebService interface {
	GetMainPage() ([]byte, error)
	GetAdminPage() ([]byte, error)
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
		{
			Method:  "GET",
			Path:    "/admin",
			Handler: c.GetAdminPage,
		},
		{
			Method: "GET",
			Path:   "/assets/",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				projectRoot := os.Getenv("PROJECT_ROOT")
				if projectRoot == "" {
					projectRoot = "."
				}
				relativePath := r.URL.Path[len("/assets/"):]
				fullPath := filepath.Join(projectRoot, "public", "assets", relativePath)
				http.ServeFile(w, r, fullPath)
			},
		},
		{
			Method: "GET",
			Path:   "/uploads/",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				projectRoot := os.Getenv("PROJECT_ROOT")
				if projectRoot == "" {
					projectRoot = "."
				}
				relativePath := r.URL.Path[len("/uploads/"):]
				fullPath := filepath.Join(projectRoot, "uploads", relativePath)
				http.ServeFile(w, r, fullPath)
			},
		},
	}
}
