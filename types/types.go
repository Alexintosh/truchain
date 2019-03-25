package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	// AppName is the name of the Cosmos app
	AppName = "TruChain"
	// StakeDenom is the name of the main staking currency (will be "trustake" on mainnet launch)
	StakeDenom = "trusteak"
	// Hostname is the address the app's HTTP server will bind to
	Hostname = "0.0.0.0"
	// Portname is the port the app's HTTP server will bind to
	Portname = "1337"
)

// Coin units
const (
	Preethi = 1
	Shanev  = 1000000000 * Preethi
)

// InitialCredAmount is the initial amount of cred for categories
var InitialCredAmount = sdk.NewInt(1000000000)

// InitialTruStake is an `sdk.Coins` representing the balance a new user is granted upon registration
var InitialTruStake = sdk.Coin{Amount: sdk.NewInt(1000000000000), Denom: StakeDenom}

// LikeCredAmount is the amount of cred for a like action
var LikeCredAmount = sdk.NewInt(1 * Shanev)

// RegistrationFee is an `auth.StdFee` representing the coin and gas cost of registering a new account
// TODO: Use more accurate gas estimate [notduncansmith]
var RegistrationFee = auth.StdFee{
	Amount: sdk.Coins{sdk.Coin{Amount: sdk.NewInt(1), Denom: StakeDenom}},
	Gas:    20000,
}

// Timestamp records the timestamp for a type
type Timestamp struct {
	CreatedBlock int64     `json:"created_block,omitempty"`
	CreatedTime  time.Time `json:"created_time,omitempty"`
	UpdatedBlock int64     `json:"updated_block,omitempty"`
	UpdatedTime  time.Time `json:"updated_time,omitempty"`
}

// NewTimestamp creates a new default Timestamp
func NewTimestamp(blockHeader abci.Header) Timestamp {
	return Timestamp{
		blockHeader.Height,
		blockHeader.Time,
		blockHeader.Height,
		blockHeader.Time,
	}
}

// Update updates an existing Timestamp and returns a new one
func (t Timestamp) Update(blockHeader abci.Header) Timestamp {
	t.UpdatedBlock = blockHeader.Height
	t.UpdatedTime = blockHeader.Time

	return t
}
