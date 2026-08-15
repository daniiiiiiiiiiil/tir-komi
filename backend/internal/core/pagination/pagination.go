package pagination

import (
	"net/http"
	"strconv"

	core_errors "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

type Pagination struct {
	Total       int  `json:"total"`
	Limit       int  `json:"limit"`
	Offset      int  `json:"offset"`
	CurrentPage int  `json:"current_page"`
	TotalPages  int  `json:"total_pages"`
	HasNext     bool `json:"has_next"`
	HasPrev     bool `json:"has_prev"`
}

func NewPagination(total, limit, offset int) Pagination {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	currentPage := (offset / limit) + 1
	totalPages := 0
	if total > 0 {
		totalPages = (total + limit - 1) / limit
	}

	return Pagination{
		Total:       total,
		Limit:       limit,
		Offset:      offset,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		HasNext:     currentPage < totalPages,
		HasPrev:     currentPage > 1,
	}
}

func FromRequest(limit, offset int) (int, int) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return limit, offset
}

func LimitOffset(limit, offset int) (int, int) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return limit, offset
}

func GetPagination(r *http.Request) (int, int, error) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil {
			return 0, 0, core_errors.ErrInvalidArgument
		}
		if l > 0 {
			limit = l
		}
	}

	offset := 0
	if offsetStr != "" {
		o, err := strconv.Atoi(offsetStr)
		if err != nil {
			return 0, 0, core_errors.ErrInvalidArgument
		}
		if o >= 0 {
			offset = o
		}
	}

	if limit > 100 {
		limit = 100
	}

	return limit, offset, nil
}
