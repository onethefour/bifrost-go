package expand

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/JFJun/go-substrate-rpc-client/v3/scale"
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
	"github.com/onethefour/bifrost-go/expand/bifrost"
	"github.com/onethefour/bifrost-go/expand/crab"
	"github.com/onethefour/bifrost-go/expand/kusama"
	"github.com/onethefour/bifrost-go/expand/polkadot"
)

type IEventRecords interface {
	GetBalancesTransfer() []types.EventBalancesTransfer
	GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess
	GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed
}

/*
扩展： 解析event
*/
func DecodeEventRecords(meta *types.Metadata, rawData string, chainName string) (IEventRecords, error) {
	e := types.EventRecordsRaw(types.MustHexDecodeString(rawData))
	var ier IEventRecords
	//log.Println(chainName)
	//log.Println(xutils.String(len(meta.AsMetadataV13.Modules)))
	//for _,v:= range meta.AsMetadataV13.Modules{
	//	//if v.Index == 53 || v.Index == 0{
	//		log.Println(v.Name,v.Index,xutils.String(v))
	//	//}
	//}
	switch strings.ToLower(chainName) {
	case "polkadot":
		var events polkadot.PolkadotEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "kusama":
		var events kusama.KusamaEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "bifrost-parachain":
		var events bifrost.BifrostEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "crab":
		var events crab.CrabEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	default:
		return nil, fmt.Errorf("do not support this chain (%s) event decord", strings.ToLower(chainName))
	}
	return ier, nil
}

/*
todo
*/
func DecodeEventRecordsNew(m *types.Metadata, e types.EventRecordsRaw) (IEventRecords, error) {

	//log.Debug(fmt.Sprintf("will decode event records from raw hex: %#x", e))
	decoder := scale.NewDecoder(bytes.NewReader(e))
	// determine number of events
	n, err := decoder.DecodeUintCompact()
	if err != nil {
		return nil, err
	}

	//log.Debug(fmt.Sprintf("found %v events", n))

	var edf EventRecordFuck
	// iterate over events
	for i := uint64(0); i < n.Uint64(); i++ {
		//log.Debug(fmt.Sprintf("decoding event #%v", i))
		// decode Phase
		phase := types.Phase{}
		err := decoder.Decode(&phase)
		if err != nil {
			return nil, fmt.Errorf("unable to decode Phase for event #%v: %v", i, err)
		}
		// decode EventID
		id := types.EventID{}
		err = decoder.Decode(&id)
		if err != nil {
			return nil, fmt.Errorf("unable to decode EventID for event #%v: %v", i, err)
		}

		//log.Debug(fmt.Sprintf("event #%v has EventID %v", i, id))

		// ask metadata for method & event name for event
		moduleName, eventName, err := m.FindEventNamesForEventID(id)
		// moduleName, eventName, err := "System", "ExtrinsicSuccess", nil
		if err != nil {
			return nil, fmt.Errorf("unable to find event with EventID %v in metadata for event #%v: %s", id, i, err)
		}
		en := fmt.Sprintf("%v_%v", moduleName, eventName)
		fmt.Println(en)
		// todo get Args
		var args []types.Type // test

		if len(args) >= 0 {
			for _, a := range args {
				// todo 根据arg去找到对应的结构体，然后去做解码
				// DecodeEventTypeByArg(decoder,a)
				fmt.Println(a)
			}
		}
		// 解析 []types.Hash
		var hashes []types.Hash
		err = decoder.Decode(&hashes)
		if err != nil {
			return nil, err
		}
	}

	return &edf, nil
}

type EventRecordFuck struct {
	BalanceTransfer        []types.EventBalancesTransfer
	SystemExtrinsicSuccess []types.EventSystemExtrinsicSuccess
	SystemExtrinsicFailed  []types.EventSystemExtrinsicFailed
}

// todo
func (e *EventRecordFuck) GetBalancesTransfer() []types.EventBalancesTransfer {
	return nil
}
func (e *EventRecordFuck) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return nil
}
func (e *EventRecordFuck) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return nil
}
