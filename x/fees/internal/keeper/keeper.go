package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
)

type Keeper struct {
	cdc        *codec.Codec
	storeKey   sdk.StoreKey
	paramSpace params.Subspace
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, paramSpace params.Subspace) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace.WithKeyTable(types.ParamKeyTable()),
	}
}

// GetParams returns the total set of fees parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of fees parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

func (k Keeper) GetFeeIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.FeeKeyPrefix)
}

func (k Keeper) MustGetFeeByKey(ctx sdk.Context, key []byte) types.Fee {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("fee not found")
	}

	bz := store.Get(key)
	var fee types.Fee
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &fee)

	return fee
}

func (k Keeper) FeeExists(ctx sdk.Context, feeId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFeePrefixKey(feeId))
}

func (k Keeper) GetFee(ctx sdk.Context, feeId uint64) (types.Fee, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeePrefixKey(feeId)

	bz := store.Get(key)
	if bz == nil {
		return types.Fee{}, sdk.ErrInternal("invalid fee")
	}

	var fee types.Fee
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &fee)

	return fee, nil
}

func (k Keeper) SubmitFee(ctx sdk.Context, content types.FeeContent) (types.Fee, sdk.Error) {
	feeId, err := k.GetFeeID(ctx)
	if err != nil {
		return types.Fee{}, err
	}

	fee := types.NewFee(feeId, content)

	k.SetFee(ctx, fee)
	k.SetFeeID(ctx, feeId+1)

	return fee, nil
}

func (k Keeper) SetFee(ctx sdk.Context, fee types.Fee) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeePrefixKey(fee.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(fee))
}

func (k Keeper) GetFeeContractIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.FeeContractKeyPrefix)
}

func (k Keeper) MustGetFeeContractByKey(ctx sdk.Context, key []byte) types.FeeContract {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("fee contract not found")
	}

	bz := store.Get(key)
	var feeContract types.FeeContract
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeContract)

	return feeContract
}

func (k Keeper) FeeContractExists(ctx sdk.Context, feeContractId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFeeContractPrefixKey(feeContractId))
}

func (k Keeper) GetFeeContract(ctx sdk.Context, feeContractId uint64) (types.FeeContract, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeContractPrefixKey(feeContractId)

	bz := store.Get(key)
	if bz == nil {
		return types.FeeContract{}, sdk.ErrInternal("invalid fee contract")
	}

	var feeContract types.FeeContract
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeContract)

	return feeContract, nil
}

func (k Keeper) SubmitFeeContract(ctx sdk.Context, content types.FeeContractContent) (types.FeeContract, sdk.Error) {
	feeContractId, err := k.GetFeeContractID(ctx)
	if err != nil {
		return types.FeeContract{}, err
	}

	feeContract := types.NewFeeContract(feeContractId, content)

	k.SetFeeContract(ctx, feeContract)
	k.SetFeeContractID(ctx, feeContractId+1)

	return feeContract, nil
}

func (k Keeper) SetFeeContract(ctx sdk.Context, feeContract types.FeeContract) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeContractPrefixKey(feeContract.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(feeContract))
}

// GetFeeID gets the highest fee ID
func (k Keeper) GetFeeID(ctx sdk.Context) (feeId uint64, err sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FeeIdKey)
	if bz == nil {
		return 0, types.ErrInvalidGenesis(types.DefaultCodespace, "initial fee ID hasn't been set")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeId)
	return feeId, nil
}

// Set the fee ID
func (k Keeper) SetFeeID(ctx sdk.Context, feeId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(feeId)
	store.Set(types.FeeIdKey, bz)
}

// GetFeeContractID gets the highest fee contract ID
func (k Keeper) GetFeeContractID(ctx sdk.Context) (feeContractId uint64, err sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FeeContractIdKey)
	if bz == nil {
		return 0, types.ErrInvalidGenesis(types.DefaultCodespace, "initial fee contract ID hasn't been set")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeContractId)
	return feeContractId, nil
}

// Set the fee contract ID
func (k Keeper) SetFeeContractID(ctx sdk.Context, feeContractId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(feeContractId)
	store.Set(types.FeeContractIdKey, bz)
}
