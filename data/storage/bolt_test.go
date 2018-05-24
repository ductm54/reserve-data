package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/KyberNetwork/reserve-data/common"
)

func TestHasPendingDepositBoltStorage(t *testing.T) {
	boltFile := "test_bolt.db"
	tmpDir, err := ioutil.TempDir("", "pending_deposit")
	if err != nil {
		t.Fatal(err)
	}
	storage, err := NewBoltStorage(filepath.Join(tmpDir, boltFile))
	if err != nil {
		t.Fatalf("Couldn't init bolt storage %v", err)
	}
	token := common.NewToken("OMG", "0x1111111111111111111111111111111111111111", 18)
	exchange := common.TestExchange{}
	timepoint := common.GetTimepoint()
	out, err := storage.HasPendingDeposit(token, exchange)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if out != false {
		t.Fatalf("Expected ram storage to return true false there is no pending deposit for the same currency and exchange")
	}
	storage.Record(
		"deposit",
		common.NewActivityID(1, "1"),
		string(exchange.ID()),
		map[string]interface{}{
			"exchange":  exchange,
			"token":     token,
			"amount":    "1.0",
			"timepoint": timepoint,
		},
		map[string]interface{}{
			"tx":    "",
			"error": nil,
		},
		"",
		"submitted",
		common.GetTimepoint())
	out, err = storage.HasPendingDeposit(token, exchange)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if out != true {
		t.Fatalf("Expected ram storage to return true when there is pending deposit")
	}

	if err = os.RemoveAll(tmpDir); err != nil {
		t.Error(err)
	}
}
