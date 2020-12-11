package types

import (
	"github.com/evdatsion/cosmos-sdk/codec"
	cryptocodec "github.com/evdatsion/cosmos-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
