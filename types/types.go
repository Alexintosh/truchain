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
var InitialTruStake = sdk.Coin{Amount: sdk.NewInt(300 * Shanev), Denom: StakeDenom}

// RegistrationFee is an `auth.StdFee` representing the coin and gas cost of registering a new account
// TODO: Use more accurate gas estimate [notduncansmith]
var RegistrationFee = auth.StdFee{
	Amount: sdk.Coins{sdk.Coin{Amount: sdk.NewInt(1), Denom: StakeDenom}},
	Gas:    20000,
}

// Tags keys
const (
	KeyPushTag             = "tru.event"
	KeyCompletedStoriesTag = "tru.event.completedStories"
)

// PushTag signifies a push notification event for Tendermint
var PushTag = sdk.NewTags(KeyPushTag, []byte("Push"))

// MsgResult is the default success response for a chain request
type MsgResult struct {
	ID int64 `json:"id"`
}

// StakeNotificationResult defines data for a stake push notification
type StakeNotificationResult struct {
	MsgResult
	StoryID int64          `json:"story_id"`
	From    sdk.AccAddress `json:"from,omitempty"`
	To      sdk.AccAddress `json:"to,omitempty"`
	Amount  sdk.Coin       `json:"amount"`
	Cred    *sdk.Coin      `json:"cred,omitempty"`
}

// Staker represents a backer or challenger with the amount staked.
type Staker struct {
	Address sdk.AccAddress
	Amount  sdk.Coin
}

// CompletedStory defines a story result.
type CompletedStory struct {
	ID                          int64                       `json:"id"`
	Creator                     sdk.AccAddress              `json:"creator"`
	Backers                     []Staker                    `json:"backers"`
	Challengers                 []Staker                    `json:"challengers"`
	StakeDistributionResults    StakeDistributionResults    `json:"stake_destribution_results"`
	InterestDistributionResults InterestDistributionResults `json:"interest_destribution_results"`
}

// CompletedStoriesNotificationResult defines the notification result of
// completed stories in a new Block.
type CompletedStoriesNotificationResult struct {
	Stories []CompletedStory `json:"stories"`
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

// StakeReward represents the amount of stake earned by an user.
type StakeReward struct {
	Account sdk.AccAddress `json:"account"`
	Amount  sdk.Coin       `json:"amount"`
}

// StakeDistributionResultsType indicates who wins the pool.
type StakeDistributionResultsType int64

// Distribution result constants
const (
	DistributionMajorityNotReached StakeDistributionResultsType = iota
	DistributionBackersWin
	DistributionChallengersWin
)

// StakeDistributionResults contains how the stake was distributed after a story completes.
type StakeDistributionResults struct {
	Type        StakeDistributionResultsType `json:"type"`
	TotalAmount sdk.Coin                     `json:"total_amount"`
	Rewards     []StakeReward                `json:"rewards"`
}

// Interest represents the amount of interest earned by an user in trustake
type Interest struct {
	Account sdk.AccAddress `json:"account"`
	Amount  sdk.Coin       `json:"amount"`
	Rate    sdk.Int        `json:"rate"`
}

// InterestDistributionResults contains how the interest was applied after a story completes.
type InterestDistributionResults struct {
	TotalAmount sdk.Coin   `json:"total_amount"`
	Interests   []Interest `json:"interests"`
}
