package blockchain

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/KyberNetwork/reserve-data/common/blockchain"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (self *Blockchain) GeneratedSetBaseRate(opts blockchain.TxOpts, tokens []ethereum.Address, baseBuy []*big.Int, baseSell []*big.Int, buy [][14]byte, sell [][14]byte, blockNumber *big.Int, indices []*big.Int) (*types.Transaction, error) {
	timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return self.BuildTx(timeout, opts, self.pricing, "setBaseRate", tokens, baseBuy, baseSell, buy, sell, blockNumber, indices)
}

func (self *Blockchain) GeneratedSetCompactData(opts blockchain.TxOpts, buy [][14]byte, sell [][14]byte, blockNumber *big.Int, indices []*big.Int) (*types.Transaction, error) {
	timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return self.BuildTx(timeout, opts, self.pricing, "setCompactData", buy, sell, blockNumber, indices)
}

func (self *Blockchain) GeneratedSetImbalanceStepFunction(opts blockchain.TxOpts, token ethereum.Address, xBuy []*big.Int, yBuy []*big.Int, xSell []*big.Int, ySell []*big.Int) (*types.Transaction, error) {
	timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return self.BuildTx(timeout, opts, self.pricing, "setImbalanceStepFunction", token, xBuy, yBuy, xSell, ySell)
}

func (self *Blockchain) GeneratedSetQtyStepFunction(opts blockchain.TxOpts, token ethereum.Address, xBuy []*big.Int, yBuy []*big.Int, xSell []*big.Int, ySell []*big.Int) (*types.Transaction, error) {
	timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return self.BuildTx(timeout, opts, self.pricing, "setQtyStepFunction", token, xBuy, yBuy, xSell, ySell)
}

func (self *Blockchain) GeneratedGetRate(opts blockchain.CallOpts, token ethereum.Address, currentBlockNumber *big.Int, buy bool, qty *big.Int) (*big.Int, error) {
	timeOut := 2 * time.Second
	out := big.NewInt(0)
	err := self.Call(timeOut, opts, self.pricing, out, "getRate", token, currentBlockNumber, buy, qty)
	return out, err
}

//GeneratedGetStepFunctionData get step function data for an option
func (bc *Blockchain) GeneratedGetStepFunctionData(token ethereum.Address, command *big.Int, param *big.Int) (*big.Int, error) {
	contractCaller := bc.BaseBlockchain.GetContractCaller()
	pricingCaller, err := NewPricingCaller(bc.pricing.Address, contractCaller.GetSingleContractCaller())
	if err != nil {
		return nil, err
	}

	opts := new(bind.CallOpts)
	log.Printf("Command: %d", command.Int64())

	result, err := pricingCaller.GetStepFunctionData(opts, token, command, param)
	if err != nil {
		log.Printf("command: %d, param: %d", command.Int64(), param.Int64())
	}
	return result, err
}
