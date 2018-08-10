package statpruner

import (
	"time"
)

// ControllerRunner contains the tickers that control the pruner services.
type ControllerRunner interface {
	GetAnalyticStorageControlTicker() <-chan time.Time
	GetRateStorageControlTicker() <-chan time.Time
	Start() error
	Stop() error
}

type ControllerTickerRunner struct {
	ascduration  time.Duration
	rateDuration time.Duration
	ascclock     *time.Ticker
	rateClock    *time.Ticker
	signal       chan bool
}

func (self *ControllerTickerRunner) GetAnalyticStorageControlTicker() <-chan time.Time {
	if self.ascclock == nil {
		<-self.signal
	}
	return self.ascclock.C
}

func (self *ControllerTickerRunner) GetRateStorageControlTicker() <-chan time.Time {
	if self.rateClock == nil {
		<-self.signal
	}
	return self.rateClock.C
}

func (self *ControllerTickerRunner) Start() error {
	self.ascclock = time.NewTicker(self.ascduration)
	self.signal <- true
	self.rateClock = time.NewTicker(self.rateDuration)
	self.signal <- true
	return nil
}

func (self *ControllerTickerRunner) Stop() error {
	self.ascclock.Stop()
	return nil
}

func NewControllerTickerRunner(
	ascduration time.Duration,
	rateDuration time.Duration) *ControllerTickerRunner {
	return &ControllerTickerRunner{
		ascduration:  ascduration,
		rateDuration: rateDuration,
		ascclock:     nil,
		rateClock:    nil,
		signal:       make(chan bool, 2),
	}
}
