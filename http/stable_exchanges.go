package http

import (
	"encoding/json"
	"github.com/KyberNetwork/reserve-data/common"
	"log"

	"github.com/KyberNetwork/reserve-data/http/httputil"
	"github.com/gin-gonic/gin"
)

func (self *HTTPServer) GetGoldData(c *gin.Context) {
	log.Printf("Getting gold data")

	data, err := self.app.GetGoldData(getTimePoint(c, true))
	if err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
	} else {
		httputil.ResponseSuccess(c, httputil.WithData(data))
	}
}

func (self *HTTPServer) GetBTCData(c *gin.Context) {
	data, err := self.app.GetBTCData(getTimePoint(c, true))
	if err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
	} else {
		httputil.ResponseSuccess(c, httputil.WithData(data))
	}
}

func (self *HTTPServer) UpdateFeedConfiguration(c *gin.Context) {
	const dataPostFormKey = "data"

	postForm, ok := self.Authenticated(c, []string{dataPostFormKey}, []Permission{ConfigurePermission})
	if !ok {
		return
	}

	data := []byte(postForm.Get(dataPostFormKey))
	if len(data) > maxDataSize {
		httputil.ResponseFailure(c, httputil.WithError(errDataSizeExceed))
		return
	}

	var input common.FeedConfiguration
	if err := json.Unmarshal(data, &input); err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}

	if err := self.app.UpdateFeedConfiguration(input.Name, input.Enabled); err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	httputil.ResponseSuccess(c)
}

func (self *HTTPServer) GetFeedConfiguration(c *gin.Context) {
	_, ok := self.Authenticated(c, []string{}, []Permission{ReadOnlyPermission, RebalancePermission, ConfigurePermission, ConfirmConfPermission})
	if !ok {
		return
	}

	data, err := self.app.GetFeedConfiguration()
	if err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	httputil.ResponseSuccess(c, httputil.WithData(data))
}
