package susu

import (
	"blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xf9beb9e9
	TestnetMagic wire.BitcoinNet = 0x0c120807
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{50}
	MainNetParams.ScriptHashAddrID = []byte{55}
	MainNetParams.Bech32HRPSegwit = "susu"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{111}
	TestNetParams.ScriptHashAddrID = []byte{117}
	TestNetParams.Bech32HRPSegwit = "tutu"
}

// SusucoinParser handle
type SusuParser struct {
	*btc.BitcoinParser
}

// NewSusucoinParser returns new SusucoinParser instance
func NewSusuParser(params *chaincfg.Params, c *btc.Configuration) *SusuParser {
	return &SusuParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Susucoin network,
// and the test Susucoin network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
