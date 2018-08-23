package huobi

import "fmt"

const huobiAPIEndpoint = "https://api.huobi.pro"

// Interface is Huobi exchange API endpoints interface.
type Interface interface {
	// PublicEndpoint returns the endpoint that does not requires authentication.
	PublicEndpoint() string
	// AuthenticatedEndpoint returns the endpoint that requires authentication.
	// In simulation mode, authenticated endpoint is the Huobi mock server.
	AuthenticatedEndpoint() string
}

// getSimulationURL returns url of the simulated Huobi endpoint.
// It returns the local default endpoint if given URL empty.
func getSimulationURL(baseURL string) string {
	const port = "5200"
	if len(baseURL) == 0 {
		baseURL = "http://127.0.0.1"
	}

	return fmt.Sprintf("%s:%s", baseURL, port)
}

type RealInterface struct{}

func (self *RealInterface) PublicEndpoint() string {
	return huobiAPIEndpoint
}

func (self *RealInterface) AuthenticatedEndpoint() string {
	return huobiAPIEndpoint
}

func NewRealInterface() *RealInterface {
	return &RealInterface{}
}

type SimulatedInterface struct {
	baseURL string
}

func (self *SimulatedInterface) PublicEndpoint() string {
	return getSimulationURL(self.baseURL)
}

func (self *SimulatedInterface) AuthenticatedEndpoint() string {
	return getSimulationURL(self.baseURL)
}

func NewSimulatedInterface(flagVariable string) *SimulatedInterface {
	return &SimulatedInterface{baseURL: flagVariable}
}

type RopstenInterface struct {
	baseURL string
}

func (self *RopstenInterface) PublicEndpoint() string {
	return huobiAPIEndpoint
}

func (self *RopstenInterface) AuthenticatedEndpoint() string {
	return getSimulationURL(self.baseURL)
}

func NewRopstenInterface(flagVariable string) *RopstenInterface {
	return &RopstenInterface{baseURL: flagVariable}
}

type KovanInterface struct {
	baseURL string
}

func (self *KovanInterface) PublicEndpoint() string {
	return huobiAPIEndpoint
}

func (self *KovanInterface) AuthenticatedEndpoint() string {
	return getSimulationURL(self.baseURL)
}

func NewKovanInterface(flagVariable string) *KovanInterface {
	return &KovanInterface{baseURL: flagVariable}
}

type DevInterface struct{}

func (self *DevInterface) PublicEndpoint() string {
	return huobiAPIEndpoint
}

func (self *DevInterface) AuthenticatedEndpoint() string {
	return huobiAPIEndpoint
}

func NewDevInterface() *DevInterface {
	return &DevInterface{}
}
