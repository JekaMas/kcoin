package mining

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/common/hexutil"
	"github.com/kowala-tech/kcoin/client/log"
)

//go:generate gencodec -type Config -field-override configMarshaling -formats toml -out gen_config.go

// DefaultConfig contains default settings for use on the Kowala main net.
var DefaultConfig = Config{
	GasPrice: big.NewInt(1),
}

type Config struct {
	Coinbase common.Address `toml:",omitempty"`

	Deposit *big.Int `toml:",omitempty"`

	ExtraData []byte `toml:",omitempty"`

	// Logger is a custom logger to use with the p2p.Server.
	Logger log.Logger `toml:",omitempty"`

	GasPrice *big.Int
}

type configMarshaling struct {
	ExtraData hexutil.Bytes
}