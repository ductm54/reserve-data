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

func (ri *RealInterface) PublicEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/public"
}

func (ri *RealInterface) MarketEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/market"
}

func (ri *RealInterface) AccountEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/account"
}

func NewRealInterface() *RealInterface {
	return &RealInterface{}
}

type SimulatedInterface struct {
	baseURL string
}

func (si *SimulatedInterface) PublicEndpoint() string {
	return fmt.Sprintf("%s/api/%s/public", getSimulationURL(si.baseURL), apiVersion)
}

func (si *SimulatedInterface) MarketEndpoint() string {
	return fmt.Sprintf("%s/api/%s/market", getSimulationURL(si.baseURL), apiVersion)
}

func (si *SimulatedInterface) AccountEndpoint() string {
	return fmt.Sprintf("%s/api/%s/account", getSimulationURL(si.baseURL), apiVersion)
}

func NewSimulatedInterface(flagVariable string) *SimulatedInterface {
	return &SimulatedInterface{baseURL: flagVariable}
}

type DevInterface struct{}

func (di *DevInterface) PublicEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/public"
}

func (di *DevInterface) MarketEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/market"
}

func (di *DevInterface) AccountEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/account"
}

func NewDevInterface() *DevInterface {
	return &DevInterface{}
}

type RopstenInterface struct {
	baseURL string
}

func (ri *RopstenInterface) PublicEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/public"
}

func (ri *RopstenInterface) MarketEndpoint() string {
	return fmt.Sprintf("%s/api/%s/market", getSimulationURL(ri.baseURL), apiVersion)
}

func (ri *RopstenInterface) AccountEndpoint() string {
	return fmt.Sprintf("%s/api/%s/account", getSimulationURL(ri.baseURL), apiVersion)
}

func NewRopstenInterface(flagVariable string) *RopstenInterface {
	return &RopstenInterface{baseURL: flagVariable}
}

type KovanInterface struct {
	baseURL string
}

func (ki *KovanInterface) PublicEndpoint() string {
	return bittrexAPIEndpoint + apiVersion + "/public"
}

func (ki *KovanInterface) MarketEndpoint() string {
	return fmt.Sprintf("%s/api/%s/market", getSimulationURL(ki.baseURL), apiVersion)
}

func (ki *KovanInterface) AccountEndpoint() string {
	return fmt.Sprintf("%s/api/%s/account", getSimulationURL(ki.baseURL), apiVersion)
}

func NewKovanInterface(flagVariable string) *KovanInterface {
	return &KovanInterface{baseURL: flagVariable}
}
