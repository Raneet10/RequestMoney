package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateRequest{}

type MsgCreateRequest struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Value   sdk.Coins      `json:"value" yaml:"value"`
}

func NewMsgCreateRequest(creator sdk.AccAddress, value sdk.Coins) MsgCreateRequest {
	return MsgCreateRequest{
		ID:      uuid.New().String(),
		Creator: creator,
		Value:   value,
	}
}

func (msg MsgCreateRequest) Route() string {
	return RouterKey
}

func (msg MsgCreateRequest) Type() string {
	return "CreateRequest"
}

func (msg MsgCreateRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateRequest) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}

	if msg.Value.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid amount")
	}
	return nil
}
