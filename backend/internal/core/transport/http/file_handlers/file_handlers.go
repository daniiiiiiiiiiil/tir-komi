package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
)

type FileHandler struct {
	advertisementService AdvertisementService
	materialService      MaterialService
}

type AdvertisementService interface {
	GetAdvertisement(ctx context.Context, id int) (domain.Advertisement, error)
}

type MaterialService interface {
	GetMethodologicalMaterial(ctx context.Context, id int) (domain.MethodologicalMaterial, error)
}

func NewFileHandler(adService AdvertisementService, matService MaterialService) *FileHandler {
	return &FileHandler{
		advertisementService: adService,
		materialService:      matService,
	}
}

func (h *FileHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:  "GET",
			Path:    "/advertisements/{id}/image",
			Handler: h.GetAdvertisementImage,
		},
		{
			Method:  "GET",
			Path:    "/advertisements/{id}/pdf",
			Handler: h.GetAdvertisementPDF,
		},
		{
			Method:  "GET",
			Path:    "/materials/{id}/pdf",
			Handler: h.GetMaterialPDF,
		},
	}
}

func (h *FileHandler) GetAdvertisementImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	id, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ad, err := h.advertisementService.GetAdvertisement(ctx, id)
	if err != nil {
		log.Error("Failed to get advertisement")
		http.NotFound(w, r)
		return
	}

	if len(ad.Image) == 0 {
		http.NotFound(w, r)
		return
	}

	contentType := detectContentType(ad.Image)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "public, max-age=86400")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(ad.Image)
}

func (h *FileHandler) GetAdvertisementPDF(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	id, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ad, err := h.advertisementService.GetAdvertisement(ctx, id)
	if err != nil {
		log.Error("Failed to get advertisement")
		http.NotFound(w, r)
		return
	}

	if len(ad.Pdf) == 0 {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=ad_"+strconv.Itoa(ad.ID)+".pdf")
	w.Header().Set("Cache-Control", "public, max-age=86400")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(ad.Pdf)
}

func (h *FileHandler) GetMaterialPDF(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)

	id, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	material, err := h.materialService.GetMethodologicalMaterial(ctx, id)
	if err != nil {
		log.Error("Failed to get material")
		http.NotFound(w, r)
		return
	}

	if len(material.Pdf) == 0 {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=material_"+strconv.Itoa(material.ID)+".pdf")
	w.Header().Set("Cache-Control", "public, max-age=86400")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(material.Pdf)
}

func detectContentType(data []byte) string {
	if len(data) < 4 {
		return "application/octet-stream"
	}

	if data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 {
		return "image/png"
	}
	if data[0] == 0xFF && data[1] == 0xD8 {
		return "image/jpeg"
	}
	if data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46 {
		return "image/gif"
	}
	if len(data) > 12 && data[0] == 0x52 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x46 {
		return "image/webp"
	}
	if data[0] == 0x25 && data[1] == 0x50 && data[2] == 0x44 && data[3] == 0x46 {
		return "application/pdf"
	}

	return "application/octet-stream"
}
