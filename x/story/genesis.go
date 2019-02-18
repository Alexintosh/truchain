package story

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all story state that must be provided at genesis
type GenesisState struct {
	Params Params `json:"params"`
}

// DefaultGenesisState for tests
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params: DefaultParams(),
	}
}

// InitGenesis initializes story state from genesis file
func InitGenesis(ctx sdk.Context, storyKeeper WriteKeeper, data GenesisState) {
	storyKeeper.SetParams(ctx, data.Params)
}