package blockchain

import (
	"testing"
)

type NonceResultTestCase struct {
	nonceResultInput         map[uint64]uint64
	expectedMostPopularNonce uint64
	expectedError            error
}

func TestGetDominantMinedNonceFromResults(t *testing.T) {
	var testCases = []NonceResultTestCase{
		{
			nonceResultInput: map[uint64]uint64{
				1: 3,
				2: 1,
				3: 1,
			},
			expectedMostPopularNonce: 1,
			expectedError:            nil,
		},
		{
			nonceResultInput: map[uint64]uint64{
				1: 2,
				2: 2,
				3: 1,
			},
			expectedMostPopularNonce: 0,
			expectedError:            errEqualCount,
		},
		{
			nonceResultInput:         map[uint64]uint64{},
			expectedMostPopularNonce: 0,
			expectedError:            errEmptyResult,
		},
	}
	for _, tc := range testCases {
		nonce, err := getDominantMinedNonceFromResults(tc.nonceResultInput)
		if nonce != tc.expectedMostPopularNonce {
			t.Logf("received different nonce %d compated to expected nonce %d", nonce, tc.expectedMostPopularNonce)
			t.Fail()
		}
		if err != tc.expectedError {
			t.Logf("received different error %s compated to expected error %s", err, tc.expectedError)
			t.Fail()
		}
	}
}
