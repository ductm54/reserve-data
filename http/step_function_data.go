package http

import (
	"github.com/KyberNetwork/reserve-data/http/httputil"
	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

//GetStepFunctionData return step function data of a token
func (h *HTTPServer) GetStepFunctionData(c *gin.Context) {
	tokenAddr := c.Param("tokenAddress")
	token := ethereum.HexToAddress(tokenAddr)
	data, err := h.core.GetStepFunctionData(token)
	if err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	httputil.ResponseSuccess(c, httputil.WithData(data))
}
