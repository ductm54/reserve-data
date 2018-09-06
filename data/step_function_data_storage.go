package data

import "github.com/KyberNetwork/reserve-data/common"

//StepFunctionDataStorage represent a storage for step function data
type StepFunctionDataStorage interface {
	StoreStepFunctionData(data common.StepFunctionData) error
	GetStepFunctionData() (common.StepFunctionData, error)
}
