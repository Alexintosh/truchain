package trustory

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/assert"
	amino "github.com/tendermint/go-amino"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

func TestAddGetStory(t *testing.T) {
	ms, storyKey := setupMultiStore()
	cdc := makeCodec()

	keeper := NewStoryKeeper(storyKey, cdc)

	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	body := "Body of story."
	creator := sdk.AccAddress([]byte{1, 2})

	storyID, err := keeper.AddStory(ctx, body, creator)
	assert.Nil(t, err)

	savedStory, err := keeper.GetStory(ctx, storyID)
	assert.Nil(t, err)

	story := Story{
		ID:      storyID,
		Body:    body,
		Creator: creator,
	}

	assert.Equal(t, savedStory, story, "Story received from store does not match expected value")
}

// ============================================================================

func setupMultiStore() (sdk.MultiStore, *sdk.KVStoreKey) {
	db := dbm.NewMemDB()
	storyKey := sdk.NewKVStoreKey("StoryKey")
	// coinKey := sdk.NewKVStoreKey("CoinKey")
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storyKey, sdk.StoreTypeIAVL, db)
	// ms.MountStoreWithDB(coinKey, sdk.StoreTypeIAVL, db)
	ms.LoadLatestVersion()
	return ms, storyKey
}

func makeCodec() *amino.Codec {
	cdc := amino.NewCodec()
	RegisterAmino(cdc)
	crypto.RegisterAmino(cdc)
	cdc.RegisterInterface((*auth.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "cosmos-sdk/BaseAccount", nil)
	return cdc
}
