package storage

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/KyberNetwork/reserve-data/boltutil"
	"github.com/KyberNetwork/reserve-data/common"
	"github.com/boltdb/bolt"
)

// StorePendingPWIEquationV2 stores the given PWIs equation data for later approval.
// Return error if occur or there is no pending PWIEquation
func (self *BoltStorage) StorePendingPWIEquationV2(data []byte) error {
	timepoint := common.GetTimepoint()
	err := self.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(pendingPWIEquationV2))
		return b.Put(boltutil.Uint64ToBytes(timepoint), data)
	})
	return err
}

// GetPendingPWIEquationV2 returns the stored PWIEquationRequestV2 in database.
func (self *BoltStorage) GetPendingPWIEquationV2() (common.PWIEquationRequestV2, error) {
	var (
		err    error
		result common.PWIEquationRequestV2
	)

	err = self.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(pendingPWIEquationV2))
		c := b.Cursor()
		_, v := c.First()
		if v == nil {
			return boltutil.ErrorNoPending
		}
		return json.Unmarshal(v, &result)
	})
	return result, err
}

// RemovePendingPWIEquationV2 deletes the pending equation request.
func (self *BoltStorage) RemovePendingPWIEquationV2() error {
	err := self.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(pendingPWIEquationV2))
		c := b.Cursor()
		k, _ := c.First()
		if k == nil {
			return boltutil.ErrorNoPending
		}
		return b.Delete(k)
	})
	return err
}

// StorePWIEquationV2 moved the pending equation request to
// pwiEquationV2 bucket and remove it from pending bucket if the
// given data matched what stored.
func (self *BoltStorage) StorePWIEquationV2(data string) error {
	err := self.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(pendingPWIEquationV2))
		c := b.Cursor()
		k, v := c.First()
		if v == nil {
			return boltutil.ErrorNoPending
		}
		confirmData := common.PWIEquationRequestV2{}
		if err := json.Unmarshal([]byte(data), &confirmData); err != nil {
			return err
		}
		currentData := common.PWIEquationRequestV2{}
		if err := json.Unmarshal(v, &currentData); err != nil {
			return err
		}
		if eq := reflect.DeepEqual(currentData, confirmData); !eq {
			return errors.New("Confirm data does not match pending data")
		}
		id := boltutil.Uint64ToBytes(common.GetTimepoint())
		if uErr := tx.Bucket([]byte(pwiEquationV2)).Put(id, v); uErr != nil {
			return uErr
		}
		// remove pending PWI equations request
		return b.Delete(k)
	})
	return err
}

func convertPWIEquationV1toV2(data string) (common.PWIEquationRequestV2, error) {
	result := common.PWIEquationRequestV2{}
	for _, dataConfig := range strings.Split(data, "|") {
		dataParts := strings.Split(dataConfig, "_")
		if len(dataParts) != 4 {
			return nil, errors.New("malform data")
		}

		a, err := strconv.ParseFloat(dataParts[1], 64)
		if err != nil {
			return nil, err
		}
		b, err := strconv.ParseFloat(dataParts[2], 64)
		if err != nil {
			return nil, err
		}
		c, err := strconv.ParseFloat(dataParts[3], 64)
		if err != nil {
			return nil, err
		}
		eq := common.PWIEquationV2{
			A: a,
			B: b,
			C: c,
		}
		result[dataParts[0]] = common.PWIEquationTokenV2{
			"bid": eq,
			"ask": eq,
		}
	}
	return result, nil
}

func pwiEquationV1toV2(tx *bolt.Tx) (common.PWIEquationRequestV2, error) {
	var eqv1 common.PWIEquation
	b := tx.Bucket([]byte(pwiEquation))
	c := b.Cursor()
	_, v := c.Last()
	if v == nil {
		return nil, errors.New("There is no equation")
	}
	if err := json.Unmarshal(v, &eqv1); err != nil {
		return nil, err
	}
	return convertPWIEquationV1toV2(eqv1.Data)
}

// GetPWIEquationV2 returns the current PWI equations from database.
func (self *BoltStorage) GetPWIEquationV2() (common.PWIEquationRequestV2, error) {
	var (
		err    error
		result common.PWIEquationRequestV2
	)
	err = self.db.View(func(tx *bolt.Tx) error {
		var vErr error // convert pwi v1 to v2 error
		b := tx.Bucket([]byte(pwiEquationV2))
		c := b.Cursor()
		_, v := c.Last()
		if v == nil {
			log.Println("there no equation in pwiEquationV2, getting from pwiEquation")
			result, vErr = pwiEquationV1toV2(tx)
			return vErr
		}
		return json.Unmarshal(v, &result)
	})
	return result, err
}
