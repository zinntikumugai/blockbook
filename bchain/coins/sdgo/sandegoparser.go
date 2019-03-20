package sdgo

import (
	"blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xe5777746
	TestnetMagic wire.BitcoinNet = 0x82547825
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{63}
	MainNetParams.ScriptHashAddrID = []byte{85}

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{111}
	TestNetParams.ScriptHashAddrID = []byte{196}
}

// SanDeGoParser handle
type SanDeGoParser struct {
	*btc.BitcoinParser
}

// NewSanDeGoParser returns new VertcoinParser instance
func NewSanDeGoParser(params *chaincfg.Params, c *btc.Configuration) *SanDeGoParser {
	return &SanDeGoParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main SanDeGo network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		if !chaincfg.IsRegistered(&chaincfg.MainNetParams) {
			chaincfg.RegisterBitcoinParams()
		}	
		if !chaincfg.IsRegistered(&MainNetParams) {
			err := chaincfg.Register(&MainNetParams)
			if err == nil {
				err = chaincfg.Register(&TestNetParams)
			}
			if err != nil {
				panic(err)
			}
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
