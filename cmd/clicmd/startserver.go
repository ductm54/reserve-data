package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/KyberNetwork/reserve-data"
	"github.com/KyberNetwork/reserve-data/blockchain"
	"github.com/KyberNetwork/reserve-data/common"
	"github.com/KyberNetwork/reserve-data/http"
	"github.com/spf13/cobra"
)

const (
	remoteLogPath  string = "core-log"
	defaultBaseURL        = "http://127.0.0.1"
	coreDefaultURL string = "http://127.0.0.1:8000"
)

// logDir is located at base of this repository.
var logDir = filepath.Join(filepath.Dir(filepath.Dir(common.CurrentDir())), "log")
var noAuthEnable bool
var servPort int = 8000
var endpointOW string
var base_url string
var noCore bool
var stdoutLog bool
var dryrun bool
var coreURL string

func serverStart(_ *cobra.Command, _ []string) {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	configLog(stdoutLog)
	//get configuration from ENV variable
	kyberENV := common.RunningMode()
	InitInterface()
	config := GetConfigFromENV(kyberENV)
	//backup other log daily
	backupLog(config.Archive, "@daily", "core.+\\.log")
	//backup core.log every 2 hour
	backupLog(config.Archive, "@every 2h", "core\\.log")

	var (
		rData reserve.ReserveData
		rCore reserve.ReserveCore
		bc    *blockchain.Blockchain
	)
	//Create Data and Core, run if not in dry mode
	if !noCore {
		var err error
		//create blockchain only if there is core
		bc, err = CreateBlockchain(config, kyberENV)
		if err != nil {
			log.Panicf("Can not create blockchain: (%s)", err)
		}
		rData, rCore = CreateDataCore(config, kyberENV, bc)
		if !dryrun {
			if kyberENV != common.SimulationMode {
				if err = rData.RunStorageController(); err != nil {
					log.Panic(err)
				}
			}
			if err = rData.Run(); err != nil {
				log.Panic(err)
			}
		}
		//set static field supportExchange from common...
		for _, ex := range config.Exchanges {
			common.SupportedExchanges[ex.ID()] = ex
		}
	}

	//Create Server
	servPortStr := fmt.Sprintf(":%d", servPort)
	server := http.NewHTTPServer(
		rData, rCore,
		config.MetricStorage,
		servPortStr,
		config.EnableAuthentication,
		config.AuthEngine,
		kyberENV,
		bc, config.Setting,
	)

	if !dryrun {
		server.Run()
	} else {
		log.Printf("Dry run finished. All configs are corrected")
	}
}

var startServer = &cobra.Command{
	Use:   "server ",
	Short: "initiate the server with specific config",
	Long: `Start reserve-data core server with preset Environment and
Allow overwriting some parameter`,
	Example: "KYBER_ENV=dev KYBER_EXCHANGES=bittrex ./cmd server --noauth -p 8000",
	Run:     serverStart,
}

func init() {
	// start server flags.
	startServer.Flags().BoolVarP(&noAuthEnable, "noauth", "", false, "disable authentication")
	startServer.Flags().IntVarP(&servPort, "port", "p", 8000, "server port")
	startServer.Flags().StringVar(&endpointOW, "endpoint", "", "endpoint, default to configuration file")
	startServer.PersistentFlags().StringVar(&base_url, "base_url", defaultBaseURL, "base_url for authenticated enpoint")
	startServer.Flags().BoolVarP(&stdoutLog, "log-to-stdout", "", false, "send log to both log file and stdout terminal")
	startServer.Flags().BoolVarP(&dryrun, "dryrun", "", false, "only test if all the configs are set correctly, will not actually run core")
	RootCmd.AddCommand(startServer)
}
