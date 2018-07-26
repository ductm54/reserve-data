package http

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/KyberNetwork/reserve-data/boltutil"
	"github.com/KyberNetwork/reserve-data/common"
	"github.com/KyberNetwork/reserve-data/http/httputil"
	"github.com/gin-gonic/gin"
)

// GetPWIEquationV2 returns the current PWI equations.
func (self *HTTPServer) GetPWIEquationV2(c *gin.Context) {
	_, ok := self.Authenticated(c, []string{}, []Permission{ReadOnlyPermission, RebalancePermission, ConfigurePermission, ConfirmConfPermission})
	if !ok {
		return
	}
	data, err := self.metric.GetPWIEquationV2()
	if err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	httputil.ResponseSuccess(c, httputil.WithData(data))
}

// SetPWIEquationV2 stores the given PWI equations to pending for later evaluation.
func (self *HTTPServer) SetPWIEquationV2(c *gin.Context) {
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

	// check if there is pending PWIEquation 2
	if _, err := self.metric.GetPendingPWIEquationV2(); err != boltutil.ErrorNoPending {
		httputil.ResponseFailure(c, httputil.WithReason(fmt.Sprintf("Expect there is no pending in PWIEquationv2, got different case (err = %v)", err)))
		return
	}

	var input common.PWIEquationRequestV2
	if err := json.Unmarshal(data, &input); err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}

	// check if the tokens in the PWIEquation is supported
	for tokenID := range input {
		if _, err := self.setting.GetInternalTokenByID(tokenID); err != nil {
			httputil.ResponseFailure(c, httputil.WithReason(fmt.Sprintf("Token %s is unsupported", tokenID)))
		}
	}

	if err := self.metric.StorePendingPWIEquationV2(data); err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	httputil.ResponseSuccess(c)
}

// GetPendingPWIEquationV2 returns the pending PWI equations.
func (self *HTTPServer) GetPendingPWIEquationV2(c *gin.Context) {
	_, ok := self.Authenticated(c, []string{}, []Permission{ReadOnlyPermission, RebalancePermission, ConfigurePermission, ConfirmConfPermission})
	if !ok {
		return
	}
	data, err := self.metric.GetPendingPWIEquationV2()
	if err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	httputil.ResponseSuccess(c, httputil.WithData(data))
}

// ConfirmPWIEquationV2 accepts the pending PWI equations and remove it from pending bucket.
func (self *HTTPServer) ConfirmPWIEquationV2(c *gin.Context) {
	const dataPostFormKey = "data"

	postForm, ok := self.Authenticated(c, []string{dataPostFormKey}, []Permission{ConfirmConfPermission})
	if !ok {
		return
	}
	postData := postForm.Get(dataPostFormKey)
	// check if the confirmData is as correct format.
	confirmData := common.PWIEquationRequestV2{}
	if err := json.Unmarshal([]byte(postData), &confirmData); err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	// check if there is pending PWIEquation and it is deep equal to confirm Data
	pending, err := self.metric.GetPendingPWIEquationV2()
	if err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	if eq := reflect.DeepEqual(pending, confirmData); !eq {
		httputil.ResponseFailure(c, httputil.WithReason("Confirm data does not match pending data"))
		return
	}

	if err := self.metric.StorePWIEquationV2(postData); err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	httputil.ResponseSuccess(c)
}

// RejectPWIEquationV2 rejects the PWI equations request and removes
// it from pending storage.
func (self *HTTPServer) RejectPWIEquationV2(c *gin.Context) {
	_, ok := self.Authenticated(c, []string{}, []Permission{ConfirmConfPermission})
	if !ok {
		return
	}

	if err := self.metric.RemovePendingPWIEquationV2(); err != nil {
		httputil.ResponseFailure(c, httputil.WithError(err))
		return
	}
	httputil.ResponseSuccess(c)
}
