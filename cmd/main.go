package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/KyberNetwork/reserve-data/blockchain"
	"github.com/KyberNetwork/reserve-data/blockchain/nonce"
	"github.com/KyberNetwork/reserve-data/common"
	"github.com/KyberNetwork/reserve-data/core"
	"github.com/KyberNetwork/reserve-data/data"
	"github.com/KyberNetwork/reserve-data/data/fetcher"
	"github.com/KyberNetwork/reserve-data/http"
	"github.com/ethereum/go-ethereum/ethclient"
)

func loadTimestamp(path string) []uint64 {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	timestamp := []uint64{}
	err = json.Unmarshal(raw, &timestamp)
	if err != nil {
		panic(err)
	}
	return timestamp
}

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	var config *Config
	switch os.Getenv("KYBER_ENV") {
	case "simulation":
		log.Printf("Running in simulation mode")
		config = GetConfigForSimulation()
		break
	case "kovan":
		log.Printf("Running in kovan mode")
		config = GetConfigForKovan()
		break
	default:
		log.Printf("Running in kovan mode")
		config = GetConfigForKovan()
	}

	logPath := "/go/src/github.com/KyberNetwork/reserve-data/cmd/log.log"
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Couldn't open log file: %v", err)
	}
	mw := io.MultiWriter(os.Stdout, f)
	defer f.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(mw)

	fetcher := fetcher.NewFetcher(
		config.FetcherStorage,
		config.FetcherRunner,
		config.ReserveAddress,
	)
	for _, ex := range config.Exchanges {
		common.SupportedExchanges[ex.ID()] = ex
	}
	for _, ex := range config.FetcherExchanges {
		fetcher.AddExchange(ex)
	}
	infura, err := ethclient.Dial(config.EthereumEndpoint)
	if err != nil {
		panic(err)
	}

	// nonceCorpus := nonce.NewAutoIncreasing(infura, fileSigner)
	nonceCorpus := nonce.NewTimeWindow(infura, config.BlockchainSigner)

	bc, err := blockchain.NewBlockchain(
		infura,
		config.WrapperAddress,
		config.ReserveAddress,
		config.BlockchainSigner,
		nonceCorpus,
	)
	for _, token := range config.SupportedTokens {
		bc.AddToken(token)
	}
	if err != nil {
		fmt.Printf("Can't connect to infura: %s\n", err)
	} else {
		fetcher.SetBlockchain(bc)
		app := data.NewReserveData(
			config.DataStorage,
			fetcher,
		)
		app.Run()
		core := core.NewReserveCore(bc, config.ActivityStorage, config.ReserveAddress)
		server := http.NewHTTPServer(app, core, ":8000")
		server.Run()
	}
}
