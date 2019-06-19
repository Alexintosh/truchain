package claim

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState defines genesis data for the module
type GenesisState struct {
	Claims []Claim `json:"claims"`
	Params Params  `json:"params"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState() GenesisState {
	return GenesisState{
		Claims: nil,
		Params: DefaultParams(),
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState { return NewGenesisState() }

// InitGenesis initializes story state from genesis file
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	for _, c := range data.Claims {
		k.Set(ctx, c.ID, c)
		k.Push(ctx, k.storeKey, communityKey, c.ID, c.CommunityID)
	}
	k.SetLen(ctx, uint64(len(data.Claims)))
	k.SetParams(ctx, data.Params)
}

// ExportGenesis exports the genesis state
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Claims: k.Claims(ctx),
		Params: k.GetParams(ctx),
	}
}

// ValidateGenesis validates the genesis state data
func ValidateGenesis(data GenesisState) error {
	if data.Params.MinClaimLength < 1 {
		return fmt.Errorf("Param: MinClaimLength must have a positive value")
	}
	if data.Params.MaxClaimLength < 1 {
		return fmt.Errorf("Param: MaxClaimLength must have a positive value")
	}

	return nil
}
