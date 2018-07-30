package http

import (
	"fmt"
	"log"
	"time"

	"github.com/KyberNetwork/reserve-data/http/httputil"
	raven "github.com/getsentry/raven-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

//HTTPServer for huobi which including
//app stand for huobi exchange instance in reserve data
//host is for api calling
//r for http engine
type HTTPServer struct {
	app  Huobi
	host string
	r    *gin.Engine
}

//PendingIntermediateTxs get pending transaction
func (h *HTTPServer) PendingIntermediateTxs(c *gin.Context) {
	data, err := h.app.PendingIntermediateTxs()
	if err != nil {
		httputil.ResponseFailure(c, httputil.WithReason(err.Error()))
	} else {
		httputil.ResponseSuccess(c, httputil.WithData(data))
	}

}

//Run run http server for huobi
func (h *HTTPServer) Run() {
	if h.app != nil {
		h.r.GET("/pending_intermediate_tx", h.PendingIntermediateTxs)
	}

	if err := h.r.Run(h.host); err != nil {
		log.Fatalf("Http server run error: %s", err.Error())
	}
}

//NewHuobiHTTPServer return new http instance
func NewHuobiHTTPServer(app Huobi) *HTTPServer {
	huobihost := fmt.Sprintf(":12221")
	r := gin.Default()
	sentryCli, err := raven.NewWithTags(
		"https://bf15053001464a5195a81bc41b644751:eff41ac715114b20b940010208271b13@sentry.io/228067",
		map[string]string{
			"env": "huobi",
		},
	)
	if err != nil {
		panic(err)
	}
	r.Use(sentry.Recovery(
		sentryCli,
		false,
	))
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("signed")
	corsConfig.AllowAllOrigins = true
	corsConfig.MaxAge = 5 * time.Minute
	r.Use(cors.New(corsConfig))

	return &HTTPServer{
		app, huobihost, r,
	}
}
