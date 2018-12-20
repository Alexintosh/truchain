package truapi

import (
	"github.com/TruStory/truchain/x/backing"
	"github.com/TruStory/truchain/x/challenge"
	"github.com/TruStory/truchain/x/chttp"
	"github.com/TruStory/truchain/x/story"
)

var supported = chttp.MsgTypes{
	"SubmitStoryMsg":     story.SubmitStoryMsg{},
	"BackStoryMsg":       backing.BackStoryMsg{},
	"CreateChallengeMsg": challenge.CreateChallengeMsg{},
	"FlagStoryMsg":       story.FlagStoryMsg{},
}
