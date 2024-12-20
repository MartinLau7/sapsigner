package library

import (
	"fmt"
)

type Symbol string

const (
	SymbolFairPlayDisposeStorage = Symbol("_jEHf8Xzsv8K")
	SymbolFairPlaySAPExchange    = Symbol("_Mib5yocT")
	SymbolFairPlaySAPInit        = Symbol("_cp2g1b9ro")
	SymbolFairPlaySAPPrime       = Symbol("_jfkdDAjba3jd")
	SymbolFairPlaySAPSign        = Symbol("_Fc3vhtJDvr")
	SymbolFairPlaySAPTeardown    = Symbol("_IPaI1oem5iL")
	SymbolFairPlaySAPVerify      = Symbol("_gLg1CWr7p")
)

var _ fmt.Stringer = Symbol("")

func (s Symbol) String() string {
	return string(s)
}
