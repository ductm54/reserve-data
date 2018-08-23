package bittrex

import (
	"fmt"
)

const (
	apiVersion         = "v1.1"
	bittrexAPIEndpoint = "https://bittrex.com/api/"
)

type Interface interface {
	PublicEndpoint() string
	MarketEndpoint() string
	AccountEndpoint() string
}

// getSimulationURL returns url of the simulated Bittrex endpoint.
// It returns the local default endpoint if given URL empty.
func getSimulationURL(baseURL string) string {
	const port = "5300"
	if len(baseURL) == 0 {
		baseURL = "http://127.0.0.1"
	}
	return fmt.Sprintf("%s:%s", baseURL, port)
}

type RealInterface struct{}

func (self *RealInterface) PublicEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/public"
}

func (self *RealInterface) MarketEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/market"
}

func (self *RealInterface) AccountEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/account"
}

func NewRealInterface() *RealInterface {
	return &RealInterface{}
}

type SimulatedInterface struct {
	baseURL string
}

func (self *SimulatedInterface) PublicEndpoint() string {
	return fmt.Sprintf("%s/api/%s/public", getSimulationURL(self.baseURL), apiVersion)
}

func (self *SimulatedInterface) MarketEndpoint() string {
	return fmt.Sprintf("%s/api/%s/market", getSimulationURL(self.baseURL), apiVersion)
}

func (self *SimulatedInterface) AccountEndpoint() string {
	return fmt.Sprintf("%s/api/%s/account", getSimulationURL(self.baseURL), apiVersion)
}

func NewSimulatedInterface(flagVariable string) *SimulatedInterface {
	return &SimulatedInterface{baseURL: flagVariable}
}

type DevInterface struct{}

func (self *DevInterface) PublicEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/public"
}

func (self *DevInterface) MarketEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/market"
}

func (self *DevInterface) AccountEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/account"
}

func NewDevInterface() *DevInterface {
	return &DevInterface{}
}

type RopstenInterface struct {
	baseURL string
}

func (self *RopstenInterface) PublicEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/public"
}

func (self *RopstenInterface) MarketEndpoint() string {
	return fmt.Sprintf("%s/api/%s/market", getSimulationURL(self.baseURL), apiVersion)
}

func (self *RopstenInterface) AccountEndpoint() string {
	return fmt.Sprintf("%s/api/%s/account", getSimulationURL(self.baseURL), apiVersion)
}

func NewRopstenInterface(flagVariable string) *RopstenInterface {
	return &RopstenInterface{baseURL: flagVariable}
}

type KovanInterface struct {
	baseURL string
}

func (self *KovanInterface) PublicEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/public"
}

func (self *KovanInterface) MarketEndpoint() string {
	return fmt.Sprintf("%s/api/%s/market", getSimulationURL(self.baseURL), apiVersion)
}

func (self *KovanInterface) AccountEndpoint() string {
	return fmt.Sprintf("%s/api/%s/account", getSimulationURL(self.baseURL), apiVersion)
}

func NewKovanInterface(flagVariable string) *KovanInterface {
	return &KovanInterface{baseURL: flagVariable}
}
