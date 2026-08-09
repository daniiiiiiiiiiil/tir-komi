package webHttp

import (
	"log"
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
			Method: "GET",
			Path:   "/assets/{filename}", // <-- без слеша в конце
			Handler: func(w http.ResponseWriter, r *http.Request) {
				// Получаем имя файла из URL
				filename := strings.TrimPrefix(r.URL.Path, "/assets/")

				// Путь к файлу
				filePath := filepath.Join("./public/assets", filename)

				// Проверяем существование
				if _, err := os.Stat(filePath); err == nil {
					// Устанавливаем Content-Type для изображений
					ext := filepath.Ext(filename)
					switch ext {
					case ".jpg", ".jpeg":
						w.Header().Set("Content-Type", "image/jpeg")
					case ".png":
						w.Header().Set("Content-Type", "image/png")
					case ".pdf":
						w.Header().Set("Content-Type", "application/pdf")
					}
					http.ServeFile(w, r, filePath)
				} else {
					log.Printf("File not found: %s", filePath)
					http.NotFound(w, r)
				}
			},
		},
	}
}

// Обработчик для assets с правильными заголовками
func (c *WebController) serveAssets(w http.ResponseWriter, r *http.Request) {
	projectRoot := os.Getenv("PROJECT_ROOT")
	if projectRoot == "" {
		projectRoot = "."
	}

	relativePath := strings.TrimPrefix(r.URL.Path, "/assets/")
	fullPath := filepath.Join(projectRoot, "public", "assets", relativePath)

	// DEBUG: логируем путь
	log.Printf("🔍 Looking for file: %s", fullPath)

	// Проверяем существование файла
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		// DEBUG: выводим все файлы в папке assets
		assetsDir := filepath.Join(projectRoot, "public", "assets")
		if files, err := os.ReadDir(assetsDir); err == nil {
			log.Printf("📁 Files in assets: %v", files)
		} else {
			log.Printf("❌ Cannot read assets dir: %v", err)
		}
		http.NotFound(w, r)
		return
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

func (c *WebController) GetImageProxy(w http.ResponseWriter, r *http.Request) {
	filename := strings.TrimPrefix(r.URL.Path, "/images/")

	if strings.HasPrefix(filename, "ad_") {
		id := strings.TrimPrefix(filename, "ad_")
		id = strings.TrimSuffix(id, ".jpg")
		http.Redirect(w, r, "/api/v1/advertisements/"+id+"/image", http.StatusSeeOther)
		return
	}

	http.NotFound(w, r)
}
