package requestmoney

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/raneet10/requestmoney/x/requestmoney/keeper"
	"github.com/raneet10/requestmoney/x/requestmoney/types"
)

func handleMsgCreateRequest(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateRequest) (*sdk.Result, error) {
	var request = types.Request{
		Creator: msg.Creator,
		ID:      msg.ID,
		Value:   msg.Value,
	}
	k.CreateRequest(ctx, request)
	err := k.MintKeeper.MintCoins(ctx, msg.Value)
	if err != nil {
		return nil, err
	}
	/*err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msg.Creator, msg.Value)
	if err != nil {
		return nil, err
	}*/
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
