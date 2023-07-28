// Code generated by goa v3.5.2, DO NOT EDIT.
//
// admin HTTP client CLI support package
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package client

import (
	"encoding/json"
	"fmt"

	admin "github.com/sakerhetsm/ssm-wargame/internal/gen/admin"
	goa "goa.design/goa/v3/pkg"
)

// BuildListChallengesPayload builds the payload for the admin ListChallenges
// endpoint from CLI flags.
func BuildListChallengesPayload(adminListChallengesToken string) (*admin.ListChallengesPayload, error) {
	var token string
	{
		token = adminListChallengesToken
	}
	v := &admin.ListChallengesPayload{}
	v.Token = token

	return v, nil
}

// BuildCreateChallengePayload builds the payload for the admin CreateChallenge
// endpoint from CLI flags.
func BuildCreateChallengePayload(adminCreateChallengeBody string, adminCreateChallengeToken string) (*admin.CreateChallengePayload, error) {
	var err error
	var body CreateChallengeRequestBody
	{
		err = json.Unmarshal([]byte(adminCreateChallengeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"ctf_event_id\": \"c397efb2-b171-4d77-9166-d105cf4f521a\",\n      \"description\": \"A heap overflow challenge\",\n      \"publish_at\": 1638384718,\n      \"score\": 50,\n      \"slug\": \"pwnme\",\n      \"title\": \"pwnme\"\n   }'")
		}
		if body.CtfEventID != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.ctf_event_id", *body.CtfEventID, goa.FormatUUID))
		}
		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminCreateChallengeToken
	}
	v := &admin.CreateChallengePayload{
		Slug:        body.Slug,
		Title:       body.Title,
		Description: body.Description,
		Score:       body.Score,
		PublishAt:   body.PublishAt,
		CtfEventID:  body.CtfEventID,
	}
	v.Token = token

	return v, nil
}

// BuildUpdateChallengePayload builds the payload for the admin UpdateChallenge
// endpoint from CLI flags.
func BuildUpdateChallengePayload(adminUpdateChallengeBody string, adminUpdateChallengeChallengeID string, adminUpdateChallengeToken string) (*admin.UpdateChallengePayload, error) {
	var err error
	var body UpdateChallengeRequestBody
	{
		err = json.Unmarshal([]byte(adminUpdateChallengeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"ctf_event_id\": \"c397efb2-b171-4d77-9166-d105cf4f521a\",\n      \"description\": \"A heap overflow challenge\",\n      \"publish_at\": 1638384718,\n      \"score\": 50,\n      \"slug\": \"pwnme\",\n      \"title\": \"pwnme\"\n   }'")
		}
		if body.CtfEventID != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.ctf_event_id", *body.CtfEventID, goa.FormatUUID))
		}
		if err != nil {
			return nil, err
		}
	}
	var challengeID string
	{
		challengeID = adminUpdateChallengeChallengeID
		err = goa.MergeErrors(err, goa.ValidateFormat("challengeID", challengeID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminUpdateChallengeToken
	}
	v := &admin.UpdateChallengePayload{
		Slug:        body.Slug,
		Title:       body.Title,
		Description: body.Description,
		Score:       body.Score,
		PublishAt:   body.PublishAt,
		CtfEventID:  body.CtfEventID,
	}
	v.ChallengeID = challengeID
	v.Token = token

	return v, nil
}

// BuildPresignChallFileUploadPayload builds the payload for the admin
// PresignChallFileUpload endpoint from CLI flags.
func BuildPresignChallFileUploadPayload(adminPresignChallFileUploadBody string, adminPresignChallFileUploadChallengeID string, adminPresignChallFileUploadToken string) (*admin.PresignChallFileUploadPayload, error) {
	var err error
	var body PresignChallFileUploadRequestBody
	{
		err = json.Unmarshal([]byte(adminPresignChallFileUploadBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"filename\": \"decryptor.exe\",\n      \"md5\": \"cq02dBbcuugBHM1oKyvMlQ==\",\n      \"size\": 41239\n   }'")
		}
	}
	var challengeID string
	{
		challengeID = adminPresignChallFileUploadChallengeID
		err = goa.MergeErrors(err, goa.ValidateFormat("challengeID", challengeID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminPresignChallFileUploadToken
	}
	v := &admin.PresignChallFileUploadPayload{
		Md5:      body.Md5,
		Filename: body.Filename,
		Size:     body.Size,
	}
	v.ChallengeID = challengeID
	v.Token = token

	return v, nil
}

// BuildListMonthlyChallengesPayload builds the payload for the admin
// ListMonthlyChallenges endpoint from CLI flags.
func BuildListMonthlyChallengesPayload(adminListMonthlyChallengesToken string) (*admin.ListMonthlyChallengesPayload, error) {
	var token string
	{
		token = adminListMonthlyChallengesToken
	}
	v := &admin.ListMonthlyChallengesPayload{}
	v.Token = token

	return v, nil
}

// BuildDeleteMonthlyChallengePayload builds the payload for the admin
// DeleteMonthlyChallenge endpoint from CLI flags.
func BuildDeleteMonthlyChallengePayload(adminDeleteMonthlyChallengeChallengeID string, adminDeleteMonthlyChallengeToken string) (*admin.DeleteMonthlyChallengePayload, error) {
	var err error
	var challengeID string
	{
		challengeID = adminDeleteMonthlyChallengeChallengeID
		err = goa.MergeErrors(err, goa.ValidateFormat("challengeID", challengeID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminDeleteMonthlyChallengeToken
	}
	v := &admin.DeleteMonthlyChallengePayload{}
	v.ChallengeID = challengeID
	v.Token = token

	return v, nil
}

// BuildDeleteFilePayload builds the payload for the admin DeleteFile endpoint
// from CLI flags.
func BuildDeleteFilePayload(adminDeleteFileFileID string, adminDeleteFileToken string) (*admin.DeleteFilePayload, error) {
	var err error
	var fileID string
	{
		fileID = adminDeleteFileFileID
		err = goa.MergeErrors(err, goa.ValidateFormat("fileID", fileID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminDeleteFileToken
	}
	v := &admin.DeleteFilePayload{}
	v.FileID = fileID
	v.Token = token

	return v, nil
}

// BuildCreateMonthlyChallengePayload builds the payload for the admin
// CreateMonthlyChallenge endpoint from CLI flags.
func BuildCreateMonthlyChallengePayload(adminCreateMonthlyChallengeBody string, adminCreateMonthlyChallengeToken string) (*admin.CreateMonthlyChallengePayload, error) {
	var err error
	var body CreateMonthlyChallengeRequestBody
	{
		err = json.Unmarshal([]byte(adminCreateMonthlyChallengeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"challenge_id\": \"85163218-8735-42ed-a7a6-42a9de2294df\",\n      \"description\": \"A heap overflow challenge\",\n      \"display_month\": \"Januari/Februari\",\n      \"end_date\": \"2006-02-01\",\n      \"slug\": \"pwnme\",\n      \"start_date\": \"2006-02-01\",\n      \"title\": \"pwnme\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.challenge_id", body.ChallengeID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminCreateMonthlyChallengeToken
	}
	v := &admin.CreateMonthlyChallengePayload{
		Slug:         body.Slug,
		Title:        body.Title,
		Description:  body.Description,
		ChallengeID:  body.ChallengeID,
		DisplayMonth: body.DisplayMonth,
		StartDate:    body.StartDate,
		EndDate:      body.EndDate,
	}
	v.Token = token

	return v, nil
}

// BuildListUsersPayload builds the payload for the admin ListUsers endpoint
// from CLI flags.
func BuildListUsersPayload(adminListUsersToken string) (*admin.ListUsersPayload, error) {
	var token string
	{
		token = adminListUsersToken
	}
	v := &admin.ListUsersPayload{}
	v.Token = token

	return v, nil
}

// BuildAddFlagPayload builds the payload for the admin AddFlag endpoint from
// CLI flags.
func BuildAddFlagPayload(adminAddFlagBody string, adminAddFlagChallengeID string, adminAddFlagToken string) (*admin.AddFlagPayload, error) {
	var err error
	var body AddFlagRequestBody
	{
		err = json.Unmarshal([]byte(adminAddFlagBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"flag\": \"SSM{...}\"\n   }'")
		}
	}
	var challengeID string
	{
		challengeID = adminAddFlagChallengeID
		err = goa.MergeErrors(err, goa.ValidateFormat("challengeID", challengeID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminAddFlagToken
	}
	v := &admin.AddFlagPayload{
		Flag: body.Flag,
	}
	v.ChallengeID = challengeID
	v.Token = token

	return v, nil
}

// BuildDeleteFlagPayload builds the payload for the admin DeleteFlag endpoint
// from CLI flags.
func BuildDeleteFlagPayload(adminDeleteFlagChallengeID string, adminDeleteFlagFlagID string, adminDeleteFlagToken string) (*admin.DeleteFlagPayload, error) {
	var err error
	var challengeID string
	{
		challengeID = adminDeleteFlagChallengeID
		err = goa.MergeErrors(err, goa.ValidateFormat("challengeID", challengeID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var flagID string
	{
		flagID = adminDeleteFlagFlagID
		err = goa.MergeErrors(err, goa.ValidateFormat("flagID", flagID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminDeleteFlagToken
	}
	v := &admin.DeleteFlagPayload{}
	v.ChallengeID = challengeID
	v.FlagID = flagID
	v.Token = token

	return v, nil
}
