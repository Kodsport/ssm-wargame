// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge HTTP client CLI support package
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	challenge "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge"
	goa "goa.design/goa/v3/pkg"
)

// BuildListChallengesPayload builds the payload for the challenge
// ListChallenges endpoint from CLI flags.
func BuildListChallengesPayload(challengeListChallengesToken string) (*challenge.ListChallengesPayload, error) {
	var token *string
	{
		if challengeListChallengesToken != "" {
			token = &challengeListChallengesToken
		}
	}
	v := &challenge.ListChallengesPayload{}
	v.Token = token

	return v, nil
}

// BuildListEventsPayload builds the payload for the challenge ListEvents
// endpoint from CLI flags.
func BuildListEventsPayload(challengeListEventsToken string) (*challenge.ListEventsPayload, error) {
	var token *string
	{
		if challengeListEventsToken != "" {
			token = &challengeListEventsToken
		}
	}
	v := &challenge.ListEventsPayload{}
	v.Token = token

	return v, nil
}

// BuildGetCurrentMonthlyChallengePayload builds the payload for the challenge
// GetCurrentMonthlyChallenge endpoint from CLI flags.
func BuildGetCurrentMonthlyChallengePayload(challengeGetCurrentMonthlyChallengeToken string) (*challenge.GetCurrentMonthlyChallengePayload, error) {
	var token *string
	{
		if challengeGetCurrentMonthlyChallengeToken != "" {
			token = &challengeGetCurrentMonthlyChallengeToken
		}
	}
	v := &challenge.GetCurrentMonthlyChallengePayload{}
	v.Token = token

	return v, nil
}

// BuildListMonthlyChallengesPayload builds the payload for the challenge
// ListMonthlyChallenges endpoint from CLI flags.
func BuildListMonthlyChallengesPayload(challengeListMonthlyChallengesToken string) (*challenge.ListMonthlyChallengesPayload, error) {
	var token *string
	{
		if challengeListMonthlyChallengesToken != "" {
			token = &challengeListMonthlyChallengesToken
		}
	}
	v := &challenge.ListMonthlyChallengesPayload{}
	v.Token = token

	return v, nil
}

// BuildSubmitFlagPayload builds the payload for the challenge SubmitFlag
// endpoint from CLI flags.
func BuildSubmitFlagPayload(challengeSubmitFlagBody string, challengeSubmitFlagChallengeID string, challengeSubmitFlagToken string) (*challenge.SubmitFlagPayload, error) {
	var err error
	var body SubmitFlagRequestBody
	{
		err = json.Unmarshal([]byte(challengeSubmitFlagBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"flag\": \"SSM{flag}\"\n   }'")
		}
		if utf8.RuneCountInString(body.Flag) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.flag", body.Flag, utf8.RuneCountInString(body.Flag), 200, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var challengeID string
	{
		challengeID = challengeSubmitFlagChallengeID
		err = goa.MergeErrors(err, goa.ValidateFormat("challengeID", challengeID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = challengeSubmitFlagToken
	}
	v := &challenge.SubmitFlagPayload{
		Flag: body.Flag,
	}
	v.ChallengeID = challengeID
	v.Token = token

	return v, nil
}

// BuildSchoolScoreboardPayload builds the payload for the challenge
// SchoolScoreboard endpoint from CLI flags.
func BuildSchoolScoreboardPayload(challengeSchoolScoreboardToken string) (*challenge.SchoolScoreboardPayload, error) {
	var token *string
	{
		if challengeSchoolScoreboardToken != "" {
			token = &challengeSchoolScoreboardToken
		}
	}
	v := &challenge.SchoolScoreboardPayload{}
	v.Token = token

	return v, nil
}

// BuildUserScoreboardPayload builds the payload for the challenge
// UserScoreboard endpoint from CLI flags.
func BuildUserScoreboardPayload(challengeUserScoreboardToken string) (*challenge.UserScoreboardPayload, error) {
	var token *string
	{
		if challengeUserScoreboardToken != "" {
			token = &challengeUserScoreboardToken
		}
	}
	v := &challenge.UserScoreboardPayload{}
	v.Token = token

	return v, nil
}
