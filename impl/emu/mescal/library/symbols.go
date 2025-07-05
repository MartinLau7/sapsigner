package library

import (
	"fmt"
)

type Symbol string

const (
	SymbolFairPlayUnknown0 = Symbol("_WIn9UJ86JKdV4dM")
	SymbolFairPlayUnknown1 = Symbol("_X46O5IeS")
	SymbolFairPlayUnknown2 = Symbol("_YlCJ3lg")
	SymbolFairPlayUnknown3 = Symbol("_dku592fbFAj")
	SymbolFairPlayUnknown4 = Symbol("_fdjkDSAFjklaf2s")
	SymbolFairPlayUnknown5 = Symbol("_lxpgvVMLd0S7uRl")

	SymbolFairPlayGetMacAddress = Symbol("_get_mac_address")

	SymbolFairPlayDisposeStorage = Symbol("_jEHf8Xzsv8K")
	SymbolFairPlaySAPExchange    = Symbol("_Mib5yocT")
	SymbolFairPlaySAPInit        = Symbol("_cp2g1b9ro")
	SymbolFairPlaySAPPrime       = Symbol("_jfkdDAjba3jd")
	SymbolFairPlaySAPSign        = Symbol("_Fc3vhtJDvr")
	SymbolFairPlaySAPTeardown    = Symbol("_IPaI1oem5iL")
	SymbolFairPlaySAPVerify      = Symbol("_gLg1CWr7p")

	SymbolFairPlayGlobalContextInit  = Symbol("$0") // anonymous function called by -[FairPlayHelper init]
	SymbolFairPlayKBSyncDataWithDSID = Symbol("$1") // anonymous function called by -[ADIProvisioningOperation _kbSyncDataWithDSID:], -[LoadDownloadQueueOperation _kbSyncDataWithDSID:], and -[PurchaseOperation _kbSyncDataWithDSID:]
)

var _ fmt.Stringer = Symbol("")

func (s Symbol) String() string {
	return string(s)
}
