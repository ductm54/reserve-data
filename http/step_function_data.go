package http

import (
	"github.com/KyberNetwork/reserve-data/http/httputil"
	"github.com/gin-gonic/gin"
)

//GetStepFunctionData return step function data of a token
func (h *HTTPServer) GetStepFunctionData(c *gin.Context) {
	data, err := h.app.GetStepFunctionData()
	if err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	httputil.ResponseSuccess(c, httputil.WithData(data))
}
