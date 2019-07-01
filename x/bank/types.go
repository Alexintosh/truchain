package bank

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Defines bank module constants
const (
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	QuerierRoute      = ModuleName
	DefaultParamspace = ModuleName
)

// Association list keys
var (
	accountKey = sdk.NewKVStoreKey("account")
)

// TransactionType defines the type of transaction.
// NOTE: This needs to stay in sync with x/account/expected_keepers.go
type TransactionType int8

func (t TransactionType) String() string {
	if int(t) >= len(TransactionTypeName) {
		return "Unknown"
	}
	return TransactionTypeName[t]
}

// Types of transactions
const (
	TransactionRegistration TransactionType = iota
	TransactionBacking
	TransactionBackingReturned
	TransactionChallenge
	TransactionChallengeReturned
	TransactionUpvote
	TransactionUpvoteReturned
	TransactionInterestArgumentCreation
	TransactionInterestUpvoteReceived
	TransactionInterestUpvoteGiven
	TransactionRewardPayout
)

var TransactionTypeName = []string{
	TransactionRegistration:             "TransactionRegistration",
	TransactionBacking:                  "TransactionBacking",
	TransactionBackingReturned:          "TransactionBackingReturned",
	TransactionChallenge:                "TransactionChallenge",
	TransactionChallengeReturned:        "TransactionChallengeReturned",
	TransactionUpvote:                   "TransactionUpvote",
	TransactionUpvoteReturned:           "TransactionUpvoteReturned",
	TransactionInterestArgumentCreation: "TransactionInterestArgumentCreation",
	TransactionInterestUpvoteReceived:   "TransactionInterestUpvoteReceived",
	TransactionInterestUpvoteGiven:      "TransactionInterestUpvoteGiven",
	TransactionRewardPayout:             "TransactionRewardPayout",
}

var allowedTransactionsForAddition = []TransactionType{
	TransactionRegistration,
	TransactionBackingReturned,
	TransactionChallengeReturned,
	TransactionUpvoteReturned,
	TransactionInterestArgumentCreation,
	TransactionInterestUpvoteReceived,
	TransactionInterestUpvoteGiven,
	TransactionRewardPayout,
}

var allowedTransactionsForDeduction = []TransactionType{
	TransactionBacking,
	TransactionChallenge,
	TransactionUpvote,
}

func (t TransactionType) allowedForAddition() bool {
	return t.oneOf(allowedTransactionsForAddition)
}

func (t TransactionType) allowedForDeduction() bool {
	return t.oneOf(allowedTransactionsForDeduction)
}

func (t TransactionType) oneOf(types []TransactionType) bool {
	for _, tType := range types {
		if tType == t {
			return true
		}
	}
	return false
}

// Transaction stores data related to a transaction
type Transaction struct {
	ID                uint64
	Type              TransactionType
	AppAccountAddress sdk.AccAddress
	ReferenceID       uint64
	Amount            sdk.Coin
	CreatedTime       time.Time
}
