package crab

import (
	//"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
)
type CrabEventRecords struct {
	types.EventRecords
	Treasury_DepositRing 			   []types.EventTreasurySpending

}

func (p CrabEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return p.Balances_Transfer
}

func (p CrabEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return p.System_ExtrinsicSuccess
}

func (p CrabEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return p.System_ExtrinsicFailed
}
