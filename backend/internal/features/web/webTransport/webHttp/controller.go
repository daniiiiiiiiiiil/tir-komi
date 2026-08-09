package webHttp

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
)

type WebController struct {
	webService WebService
}

type WebService interface {
	GetMainPage() ([]byte, error)
	GetAdminPage() ([]byte, error)
	GetHTMLFile(filename string) ([]byte, error)
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
			Path:    "/basic_information",
			Handler: c.GetBasicInfoPage,
		},
		{
			Method:  "GET",
			Path:    "/admin",
			Handler: c.GetAdminPage,
		},
		{
			Method:  "GET",
			Path:    "/assets/",
			Handler: c.serveAssets, // Выносим в отдельную функцию
		},
		{
			Method:  "GET",
			Path:    "/uploads/",
			Handler: c.serveUploads, // Выносим в отдельную функцию
		},
	}
}

// Обработчик для assets с правильными заголовками
func (c *WebController) serveAssets(w http.ResponseWriter, r *http.Request) {
	projectRoot := os.Getenv("PROJECT_ROOT")
	if projectRoot == "" {
		projectRoot = "."
	}

	// Получаем путь к файлу
	relativePath := strings.TrimPrefix(r.URL.Path, "/assets/")
	fullPath := filepath.Join(projectRoot, "public", "assets", relativePath)

	// Проверяем существование файла
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	// Устанавливаем правильные заголовки в зависимости от расширения
	ext := filepath.Ext(fullPath)
	switch ext {
	case ".pdf":
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "inline; filename="+filepath.Base(fullPath))
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	}

	http.ServeFile(w, r, fullPath)
}

// Обработчик для uploads
func (c *WebController) serveUploads(w http.ResponseWriter, r *http.Request) {
	projectRoot := os.Getenv("PROJECT_ROOT")
	if projectRoot == "" {
		projectRoot = "."
	}

	relativePath := strings.TrimPrefix(r.URL.Path, "/uploads/")
	fullPath := filepath.Join(projectRoot, "uploads", relativePath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	// Для загруженных файлов тоже устанавливаем заголовки
	ext := filepath.Ext(fullPath)
	if ext == ".pdf" {
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "inline; filename="+filepath.Base(fullPath))
	}

	http.ServeFile(w, r, fullPath)
}

// Вспомогательная функция для определения MIME типа
func getContentType(filename string) string {
	ext := filepath.Ext(filename)
	switch ext {
	case ".pdf":
		return "application/pdf"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".html", ".htm":
		return "text/html"
	default:
		return "application/octet-stream"
	}
}
