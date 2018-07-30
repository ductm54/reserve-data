package binance

import "fmt"

const binanceAPIEndpoint = "https://api.binance.com"

// Interface is Binance exchange API endpoints interface.
type Interface interface {
	// PublicEndpoint returns the endpoint that does not requires authentication.
	PublicEndpoint() string
	// AuthenticatedEndpoint returns the endpoint that requires authentication.
	// In simulation mode, authenticated endpoint is the Binance mock server.
	AuthenticatedEndpoint() string
}

type RealInterface struct{}

// getSimulationURL returns url of the simulated Binance endpoint.
// It returns the local default endpoint if given URL empty.
func getSimulationURL(baseURL string) string {
	const port = "5100"
	if len(baseURL) == 0 {
		baseURL = "http://127.0.0.1"
	}
	return fmt.Sprintf("%s:%s", baseURL, port)
}

func (ri *RealInterface) PublicEndpoint() string {
	return binanceAPIEndpoint
}

func (ri *RealInterface) AuthenticatedEndpoint() string {
	return binanceAPIEndpoint
}

func NewRealInterface() *RealInterface {
	return &RealInterface{}
}

type SimulatedInterface struct {
	baseURL string
}

func (si *SimulatedInterface) PublicEndpoint() string {
	return getSimulationURL(si.baseURL)
}

func (si *SimulatedInterface) AuthenticatedEndpoint() string {
	return getSimulationURL(si.baseURL)
}

func NewSimulatedInterface(flagVariable string) *SimulatedInterface {
	return &SimulatedInterface{baseURL: flagVariable}
}

type RopstenInterface struct {
	baseURL string
}

func (roi *RopstenInterface) PublicEndpoint() string {
	return binanceAPIEndpoint
}

func (roi *RopstenInterface) AuthenticatedEndpoint() string {
	return getSimulationURL(roi.baseURL)
}

func NewRopstenInterface(flagVariable string) *RopstenInterface {
	return &RopstenInterface{baseURL: flagVariable}
}

type KovanInterface struct {
	baseURL string
}

func (ki *KovanInterface) PublicEndpoint() string {
	return binanceAPIEndpoint
}

func (ki *KovanInterface) AuthenticatedEndpoint() string {
	return getSimulationURL(ki.baseURL)
}

func NewKovanInterface(flagVariable string) *KovanInterface {
	return &KovanInterface{baseURL: flagVariable}
}

type DevInterface struct{}

func (di *DevInterface) PublicEndpoint() string {
	return binanceAPIEndpoint
}

func (di *DevInterface) AuthenticatedEndpoint() string {
	return binanceAPIEndpoint
}

func NewDevInterface() *DevInterface {
	return &DevInterface{}
}
