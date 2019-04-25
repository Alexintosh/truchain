package backing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all story state that must be provided at genesis
type GenesisState struct {
	Backings []Backing `json:"backings"`
}

// DefaultGenesisState for tests
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// InitGenesis initializes story state from genesis file
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, backing := range data.Backings {
		keeper.setBacking(ctx, backing)
		keeper.backingStoryList.Append(ctx, keeper, backing.StoryID(), backing.Creator(), backing.ID())
	}
	keeper.SetLen(ctx, int64(len(data.Backings)))
}

// ExportGenesis exports the genesis state
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	var backings []Backing
	err := keeper.EachPrefix(ctx, keeper.StorePrefix(), func(bz []byte) bool {
		var b Backing
		keeper.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &b)
		backings = append(backings, b)
		return true
	})
	if err != nil {
		panic(err)
	}

	return GenesisState{
		Backings: backings,
	}
}