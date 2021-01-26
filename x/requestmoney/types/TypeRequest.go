package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Request struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Value   sdk.Coins      `json:"value" yaml:"value"`
}
