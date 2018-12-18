package blockchain

import (
	"errors"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	errEmptyResult = errors.New("all nodes failed: nonce Results is empty")
	errEqualCount  = errors.New("cannot determine dominant nonce since they have equal count")
)

func fetchNonceFromAllNode(op *Operator, clients map[string]*ethclient.Client, nonceChannel chan uint64) {
	var wg sync.WaitGroup
	for endpoint, client := range clients {
		wg.Add(1)
		go func(endpoint string, client *ethclient.Client) {
			defer wg.Done()
			nonce, err := op.NonceCorpus.MinedNonce(client)
			log.Printf("SET_RATE_MINED_NONCE: request for mined nonce from endpoint %s, got result %d, error %s", endpoint, nonce, err)
			if err == nil {
				nonceChannel <- nonce.Uint64()
			}
		}(endpoint, client)
	}
	wg.Wait()
	close(nonceChannel)
}

func getDominantMinedNonceFromResults(nonceResults map[uint64]uint64) (uint64, error) {
	var (
		mostPopularNonce      uint64
		mostPopularNonceCount uint64
		sameMaxCount          bool
	)
	if len(nonceResults) == 0 {
		return 0, errEmptyResult
	}
	for nonce := range nonceResults {
		if nonceResults[nonce] > mostPopularNonceCount {
			mostPopularNonceCount = nonceResults[nonce]
			mostPopularNonce = nonce
			sameMaxCount = false

		} else if nonceResults[nonce] == mostPopularNonceCount {
			sameMaxCount = true
		}
	}
	if sameMaxCount == true {
		return 0, errEqualCount
	}
	log.Printf("SET_RATE_MINED_NONCE: most popular none is %d, with number of occurrence %d on total of %d result from nodes", mostPopularNonce, mostPopularNonceCount, len(nonceResults))
	return mostPopularNonce, nil
}

//GetDominantMinedNonceFromAllNodes  return the nonce that is dominant from all nodes
func (self *BaseBlockchain) GetDominantMinedNonceFromAllNodes(operator string) (uint64, error) {
	var (
		op           = self.MustGetOperator(operator)
		nonceChannel = make(chan uint64)
		//nonceResults is the map of [nonce]Count, to count the occurrence of all nonces return from nodes.
		nonceResults map[uint64]uint64
	)
	go fetchNonceFromAllNode(op, self.broadcaster.clients, nonceChannel)
	for nonce := range nonceChannel {
		_, avail := nonceResults[nonce]
		if !avail {
			nonceResults[nonce] = 1
		} else {
			nonceResults[nonce]++
		}
	}
	return getDominantMinedNonceFromResults(nonceResults)
}

func (self *BaseBlockchain) GetMinedNonce(operator string) (uint64, error) {
	nonce, err := self.MustGetOperator(operator).NonceCorpus.MinedNonce(self.client)
	if err != nil {
		return 0, err
	} else {
		return nonce.Uint64(), err
	}
}
