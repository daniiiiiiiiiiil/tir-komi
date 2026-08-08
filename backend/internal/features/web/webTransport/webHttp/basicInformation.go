package webHttp

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

func (c *WebController) GetBasicInfoPage(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	html, err := c.webService.GetHTMLFile("/basic_information.html")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed get basic info html file")
		return
	}

	responseHandler.HTMLResponse(html)
}
