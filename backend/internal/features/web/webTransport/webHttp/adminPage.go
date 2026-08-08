package webHttp

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

func (c *WebController) GetAdminPage(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	html, err := c.webService.GetAdminPage()
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get admin html file")
		return
	}

	responseHandler.HTMLResponse(html)
}
