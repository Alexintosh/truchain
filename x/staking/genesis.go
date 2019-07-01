package staking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	app "github.com/TruStory/truchain/types"
)

type UserEarnedCoins struct {
	Address sdk.AccAddress
	Coins   sdk.Coins
}

// GenesisState defines genesis data for the module
type GenesisState struct {
	Arguments     []Argument        `json:"arguments"`
	Params        Params            `json:"params"`
	Stakes        []Stake           `json:"stakes"`
	UsersEarnings []UserEarnedCoins `json:"users_earnings"`
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(arguments []Argument, stakes []Stake, userEarnings []UserEarnedCoins, params Params) GenesisState {
	return GenesisState{
		Arguments:     arguments,
		Params:        params,
		Stakes:        stakes,
		UsersEarnings: userEarnings,
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params:        DefaultParams(),
		Stakes:        make([]Stake, 0),
		Arguments:     make([]Argument, 0),
		UsersEarnings: make([]UserEarnedCoins, 0),
	}
}

// InitGenesis initializes staking state from genesis file
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	for _, a := range data.Arguments {
		k.setArgument(ctx, a)
		k.setClaimArgument(ctx, a.ClaimID, a.ID)
		k.serUserArgument(ctx, a.Creator, a.ID)
	}
	for _, s := range data.Stakes {
		k.setStake(ctx, s)
		if !s.Expired {
			k.InsertActiveStakeQueue(ctx, s.ID, s.EndTime)
		}
		k.setArgumentStake(ctx, s.ArgumentID, s.ID)
		k.setUserStake(ctx, s.Creator, s.CreatedTime, s.ID)
	}
	k.setArgumentID(ctx, uint64(len(data.Arguments)+1))
	k.setStakeID(ctx, uint64(len(data.Stakes)+1))

	for _, e := range data.UsersEarnings {
		k.setEarnedCoins(ctx, e.Address, e.Coins)
	}
	k.SetParams(ctx, data.Params)
}

// ExportGenesis exports the genesis state
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	return GenesisState{
		Params:        keeper.GetParams(ctx),
		Arguments:     keeper.Arguments(ctx),
		Stakes:        keeper.Stakes(ctx),
		UsersEarnings: keeper.UsersEarnings(ctx),
	}
}

// ValidateGenesis validates the genesis state data
func ValidateGenesis(data GenesisState) error {
	if data.Params.ArgumentCreationStake.Denom != app.StakeDenom {
		return ErrInvalidArgumentStakeDenom
	}
	if data.Params.UpvoteStake.Denom != app.StakeDenom {
		return ErrInvalidUpvoteStakeDenom
	}
	return nil
}
