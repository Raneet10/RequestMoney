package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/raneet10/requestmoney/x/requestmoney/types"
)

// CreateRequest creates a request
func (k Keeper) CreateRequest(ctx sdk.Context, request types.Request) error {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.RequestPrefix + request.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(request)
	store.Set(key, value)
	err := k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, request.Creator, request.Value)
	if err != nil {
		return err
	}
	return nil
}

// GetRequest returns the request information
func (k Keeper) GetRequest(ctx sdk.Context, key string) (types.Request, error) {
	store := ctx.KVStore(k.storeKey)
	var request types.Request
	byteKey := []byte(types.RequestPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &request)
	if err != nil {
		return request, err
	}
	return request, nil
}

// SetRequest sets a request
/*func (k Keeper) SetRequest(ctx sdk.Context, request types.Request) {
	requestKey := request.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(request)
	key := []byte(types.RequestPrefix + requestKey)
	store.Set(key, bz)
}

// DeleteRequest deletes a request
func (k Keeper) DeleteRequest(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.RequestPrefix + key))
}*/

//
// Functions used by querier
//

func listRequest(ctx sdk.Context, k Keeper) ([]byte, error) {
	var requestList []types.Request
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.RequestPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var request types.Request
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &request)
		requestList = append(requestList, request)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, requestList)
	return res, nil
}

func getRequest(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	request, err := k.GetRequest(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, request)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetRequestOwner(ctx sdk.Context, key string) sdk.AccAddress {
	request, err := k.GetRequest(ctx, key)
	if err != nil {
		return nil
	}
	return request.Creator
}

// Check if the key exists in the store
func (k Keeper) RequestExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.RequestPrefix + key))
}
