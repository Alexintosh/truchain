package community

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// Keys for params
var (
	KeyMinNameLength        = []byte("minNameLength")
	KeyMaxNameLength        = []byte("maxNameLength")
	KeyMinSlugLength        = []byte("minSlugLength")
	KeyMaxSlugLength        = []byte("maxSlugLength")
	KeyMaxDescriptionLength = []byte("maxDescriptionLength")
)

// Params holds parameters for a Community
type Params struct {
	MinNameLength        int `json:"min_name_length"`
	MaxNameLength        int `json:"max_name_length"`
	MinSlugLength        int `json:"min_slug_length"`
	MaxSlugLength        int `json:"max_slug_length"`
	MaxDescriptionLength int `json:"max_description_length"`
}

// DefaultParams is the Community params for testing
func DefaultParams() Params {
	return Params{
		MinNameLength:        5,
		MaxNameLength:        25,
		MinSlugLength:        3,
		MaxSlugLength:        15,
		MaxDescriptionLength: 140,
	}
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{Key: KeyMinNameLength, Value: &p.MinNameLength},
		{Key: KeyMaxNameLength, Value: &p.MaxNameLength},
		{Key: KeyMinSlugLength, Value: &p.MinSlugLength},
		{Key: KeyMaxSlugLength, Value: &p.MaxSlugLength},
		{Key: KeyMaxDescriptionLength, Value: &p.MaxDescriptionLength},
	}
}

// ParamKeyTable for community module
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// GetParams gets the genesis params for the community
func (k Keeper) GetParams(ctx sdk.Context) Params {
	var paramSet Params
	k.paramStore.GetParamSet(ctx, &paramSet)
	return paramSet
}

// SetParams sets the params for the community
func (k Keeper) SetParams(ctx sdk.Context, params Params) {
	logger := ctx.Logger().With("module", ModuleName)
	k.paramStore.SetParamSet(ctx, &params)
	logger.Info(fmt.Sprintf("Loaded community params: %+v", params))
}
