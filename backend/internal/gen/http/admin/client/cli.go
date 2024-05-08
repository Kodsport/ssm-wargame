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

// BuildGetChallengeMetaPayload builds the payload for the admin
// GetChallengeMeta endpoint from CLI flags.
func BuildGetChallengeMetaPayload(adminGetChallengeMetaChallengeID string, adminGetChallengeMetaToken string) (*admin.GetChallengeMetaPayload, error) {
	var err error
	var challengeID string
	{
		challengeID = adminGetChallengeMetaChallengeID
		err = goa.MergeErrors(err, goa.ValidateFormat("challengeID", challengeID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminGetChallengeMetaToken
	}
	v := &admin.GetChallengeMetaPayload{}
	v.ChallengeID = challengeID
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"authors\": [\n         \"1b678293-6737-4cc7-8eae-aa821100293f\"\n      ],\n      \"category_id\": \"1b678292-6737-4cc7-8eae-aa821100293f\",\n      \"ctf_event_id\": \"c397efb2-b171-4d77-9166-d105cf4f521a\",\n      \"description\": \"A heap overflow challenge\",\n      \"hide\": false,\n      \"publish_at\": 1638384718,\n      \"slug\": \"pwnme\",\n      \"static_score\": 50,\n      \"title\": \"pwnme\"\n   }'")
		}
		if body.CtfEventID != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.ctf_event_id", *body.CtfEventID, goa.FormatUUID))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.category_id", body.CategoryID, goa.FormatUUID))

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
		PublishAt:   body.PublishAt,
		CtfEventID:  body.CtfEventID,
		Hide:        body.Hide,
		StaticScore: body.StaticScore,
		CategoryID:  body.CategoryID,
	}
	if body.Authors != nil {
		v.Authors = make([]string, len(body.Authors))
		for i, val := range body.Authors {
			v.Authors[i] = val
		}
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"authors\": [\n         \"1b678293-6737-4cc7-8eae-aa821100293f\"\n      ],\n      \"category_id\": \"1b678292-6737-4cc7-8eae-aa821100293f\",\n      \"ctf_event_id\": \"c397efb2-b171-4d77-9166-d105cf4f521a\",\n      \"description\": \"A heap overflow challenge\",\n      \"hide\": false,\n      \"publish_at\": 1638384718,\n      \"slug\": \"pwnme\",\n      \"static_score\": 50,\n      \"title\": \"pwnme\"\n   }'")
		}
		if body.CtfEventID != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.ctf_event_id", *body.CtfEventID, goa.FormatUUID))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.category_id", body.CategoryID, goa.FormatUUID))

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
		PublishAt:   body.PublishAt,
		CtfEventID:  body.CtfEventID,
		Hide:        body.Hide,
		StaticScore: body.StaticScore,
		CategoryID:  body.CategoryID,
	}
	if body.Authors != nil {
		v.Authors = make([]string, len(body.Authors))
		for i, val := range body.Authors {
			v.Authors[i] = val
		}
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"challenge_id\": \"85163218-8735-42ed-a7a6-42a9de2294df\",\n      \"display_month\": \"Januari/Februari\",\n      \"end_date\": 1690884841,\n      \"start_date\": 1690874841\n   }'")
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

// BuildListAuthorsPayload builds the payload for the admin ListAuthors
// endpoint from CLI flags.
func BuildListAuthorsPayload(adminListAuthorsToken string) (*admin.ListAuthorsPayload, error) {
	var token string
	{
		token = adminListAuthorsToken
	}
	v := &admin.ListAuthorsPayload{}
	v.Token = token

	return v, nil
}

// BuildUpdateAuthorPayload builds the payload for the admin UpdateAuthor
// endpoint from CLI flags.
func BuildUpdateAuthorPayload(adminUpdateAuthorBody string, adminUpdateAuthorID string, adminUpdateAuthorToken string) (*admin.UpdateAuthorPayload, error) {
	var err error
	var body UpdateAuthorRequestBody
	{
		err = json.Unmarshal([]byte(adminUpdateAuthorBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Movitz gör saker\",\n      \"full_name\": \"Movitz Sunar\",\n      \"image_url\": \"movitz\",\n      \"publish\": true,\n      \"slug\": \"movitz\",\n      \"sponsor\": true\n   }'")
		}
	}
	var id string
	{
		id = adminUpdateAuthorID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminUpdateAuthorToken
	}
	v := &admin.UpdateAuthorPayload{
		FullName:    body.FullName,
		Description: body.Description,
		Sponsor:     body.Sponsor,
		Slug:        body.Slug,
		ImageURL:    body.ImageURL,
		Publish:     body.Publish,
	}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildCreateAuthorPayload builds the payload for the admin CreateAuthor
// endpoint from CLI flags.
func BuildCreateAuthorPayload(adminCreateAuthorBody string, adminCreateAuthorToken string) (*admin.CreateAuthorPayload, error) {
	var err error
	var body CreateAuthorRequestBody
	{
		err = json.Unmarshal([]byte(adminCreateAuthorBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Movitz gör saker\",\n      \"full_name\": \"Movitz Sunar\",\n      \"image_url\": \"movitz\",\n      \"publish\": true,\n      \"slug\": \"movitz\",\n      \"sponsor\": true\n   }'")
		}
	}
	var token string
	{
		token = adminCreateAuthorToken
	}
	v := &admin.CreateAuthorPayload{
		FullName:    body.FullName,
		Description: body.Description,
		Sponsor:     body.Sponsor,
		Slug:        body.Slug,
		ImageURL:    body.ImageURL,
		Publish:     body.Publish,
	}
	v.Token = token

	return v, nil
}

// BuildDeleteAuthorPayload builds the payload for the admin DeleteAuthor
// endpoint from CLI flags.
func BuildDeleteAuthorPayload(adminDeleteAuthorID string, adminDeleteAuthorToken string) (*admin.DeleteAuthorPayload, error) {
	var id string
	{
		id = adminDeleteAuthorID
	}
	var token string
	{
		token = adminDeleteAuthorToken
	}
	v := &admin.DeleteAuthorPayload{}
	v.ID = id
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

// BuildListCategoriesPayload builds the payload for the admin ListCategories
// endpoint from CLI flags.
func BuildListCategoriesPayload(adminListCategoriesToken string) (*admin.ListCategoriesPayload, error) {
	var token string
	{
		token = adminListCategoriesToken
	}
	v := &admin.ListCategoriesPayload{}
	v.Token = token

	return v, nil
}

// BuildChalltoolsImportPayload builds the payload for the admin
// ChalltoolsImport endpoint from CLI flags.
func BuildChalltoolsImportPayload(adminChalltoolsImportBody string, adminChalltoolsImportImportToken string) (*admin.ChalltoolsImportPayload, error) {
	var err error
	var body ChalltoolsImportRequestBody
	{
		err = json.Unmarshal([]byte(adminChalltoolsImportBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"authors\": [\n         \"Movitz Sunar\"\n      ],\n      \"categories\": [\n         \"web\"\n      ],\n      \"challenge_id\": \"225ada44-3fde-460d-84a4-2f16ff579618\",\n      \"description\": \"how to dns\",\n      \"file_urls\": [\n         \"https://bucket/key\"\n      ],\n      \"flag_format_prefix\": \"SSM{\",\n      \"flag_format_suffix\": \"}\",\n      \"flags\": [\n         {\n            \"flag\": \"fl4g_l0l\",\n            \"type\": \"regex\"\n         },\n         {\n            \"flag\": \"fl4g_l0l\",\n            \"type\": \"regex\"\n         }\n      ],\n      \"human_metadata\": {\n         \"event_name\": \"Architecto a est est labore eligendi.\"\n      },\n      \"order\": 5,\n      \"score\": 100,\n      \"services\": [\n         {\n            \"hyperlink\": true,\n            \"user_display\": \"nc 0.0.0.0 1234\"\n         },\n         {\n            \"hyperlink\": true,\n            \"user_display\": \"nc 0.0.0.0 1234\"\n         },\n         {\n            \"hyperlink\": true,\n            \"user_display\": \"nc 0.0.0.0 1234\"\n         },\n         {\n            \"hyperlink\": true,\n            \"user_display\": \"nc 0.0.0.0 1234\"\n         }\n      ],\n      \"title\": \"DNS 101\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.challenge_id", body.ChallengeID, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var importToken string
	{
		importToken = adminChalltoolsImportImportToken
	}
	v := &admin.ChalltoolsImportPayload{
		Title:            body.Title,
		Description:      body.Description,
		Score:            body.Score,
		ChallengeID:      body.ChallengeID,
		FlagFormatPrefix: body.FlagFormatPrefix,
		FlagFormatSuffix: body.FlagFormatSuffix,
		Order:            body.Order,
	}
	if body.Authors != nil {
		v.Authors = make([]string, len(body.Authors))
		for i, val := range body.Authors {
			v.Authors[i] = val
		}
	}
	if body.Categories != nil {
		v.Categories = make([]string, len(body.Categories))
		for i, val := range body.Categories {
			v.Categories[i] = val
		}
	}
	if body.FileUrls != nil {
		v.FileUrls = make([]string, len(body.FileUrls))
		for i, val := range body.FileUrls {
			v.FileUrls[i] = val
		}
	}
	if body.Flags != nil {
		v.Flags = make([]*admin.ImportChallFlag, len(body.Flags))
		for i, val := range body.Flags {
			v.Flags[i] = marshalImportChallFlagRequestBodyToAdminImportChallFlag(val)
		}
	}
	if body.Services != nil {
		v.Services = make([]*admin.ImportChallService, len(body.Services))
		for i, val := range body.Services {
			v.Services[i] = marshalImportChallServiceRequestBodyToAdminImportChallService(val)
		}
	}
	if body.HumanMetadata != nil {
		v.HumanMetadata = marshalImportChallHumanMetadataRequestBodyToAdminImportChallHumanMetadata(body.HumanMetadata)
	}
	v.ImportToken = importToken

	return v, nil
}

// BuildListCTFEventsPayload builds the payload for the admin ListCTFEvents
// endpoint from CLI flags.
func BuildListCTFEventsPayload(adminListCTFEventsToken string) (*admin.ListCTFEventsPayload, error) {
	var token string
	{
		token = adminListCTFEventsToken
	}
	v := &admin.ListCTFEventsPayload{}
	v.Token = token

	return v, nil
}

// BuildCreateCTFEventPayload builds the payload for the admin CreateCTFEvent
// endpoint from CLI flags.
func BuildCreateCTFEventPayload(adminCreateCTFEventBody string, adminCreateCTFEventToken string) (*admin.CreateCTFEventPayload, error) {
	var err error
	var body CreateCTFEventRequestBody
	{
		err = json.Unmarshal([]byte(adminCreateCTFEventBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Säkerhet-SM 2023\"\n   }'")
		}
	}
	var token string
	{
		token = adminCreateCTFEventToken
	}
	v := &admin.CreateCTFEventPayload{
		Name: body.Name,
	}
	v.Token = token

	return v, nil
}

// BuildDeleteCTFEventPayload builds the payload for the admin DeleteCTFEvent
// endpoint from CLI flags.
func BuildDeleteCTFEventPayload(adminDeleteCTFEventID string, adminDeleteCTFEventToken string) (*admin.DeleteCTFEventPayload, error) {
	var err error
	var id string
	{
		id = adminDeleteCTFEventID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminDeleteCTFEventToken
	}
	v := &admin.DeleteCTFEventPayload{}
	v.ID = id
	v.Token = token

	return v, nil
}

// BuildCreateCTFEventImportTokenPayload builds the payload for the admin
// CreateCTFEventImportToken endpoint from CLI flags.
func BuildCreateCTFEventImportTokenPayload(adminCreateCTFEventImportTokenBody string, adminCreateCTFEventImportTokenToken string) (*admin.CreateCTFEventImportTokenPayload, error) {
	var err error
	var body CreateCTFEventImportTokenRequestBody
	{
		err = json.Unmarshal([]byte(adminCreateCTFEventImportTokenBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"expires_in\": \"year\",\n      \"name\": \"e3bb4dc5-9479-42ce-aed3-b41e8139fccb\"\n   }'")
		}
		if !(body.ExpiresIn == "hour" || body.ExpiresIn == "week" || body.ExpiresIn == "year") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.expires_in", body.ExpiresIn, []interface{}{"hour", "week", "year"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminCreateCTFEventImportTokenToken
	}
	v := &admin.CreateCTFEventImportTokenPayload{
		Name:      body.Name,
		ExpiresIn: body.ExpiresIn,
	}
	v.Token = token

	return v, nil
}

// BuildListCoursesPayload builds the payload for the admin ListCourses
// endpoint from CLI flags.
func BuildListCoursesPayload(adminListCoursesToken string) (*admin.ListCoursesPayload, error) {
	var token string
	{
		token = adminListCoursesToken
	}
	v := &admin.ListCoursesPayload{}
	v.Token = token

	return v, nil
}

// BuildCreateCoursePayload builds the payload for the admin CreateCourse
// endpoint from CLI flags.
func BuildCreateCoursePayload(adminCreateCourseBody string, adminCreateCourseToken string) (*admin.CreateCoursePayload, error) {
	var err error
	var body CreateCourseRequestBody
	{
		err = json.Unmarshal([]byte(adminCreateCourseBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"author_ids\": [\n         \"b29869ec-6eef-451c-a389-249a6d6ac47b\"\n      ],\n      \"category\": \"web\",\n      \"description\": \"markdown text\",\n      \"difficulty\": \"advanced\",\n      \"publish\": true,\n      \"slug\": \"sqli\",\n      \"title\": \"SQL-Injektioner\"\n   }'")
		}
	}
	var token string
	{
		token = adminCreateCourseToken
	}
	v := &admin.CreateCoursePayload{
		Title:       body.Title,
		Slug:        body.Slug,
		Category:    body.Category,
		Difficulty:  body.Difficulty,
		Description: body.Description,
		Publish:     body.Publish,
	}
	if body.AuthorIds != nil {
		v.AuthorIds = make([]string, len(body.AuthorIds))
		for i, val := range body.AuthorIds {
			v.AuthorIds[i] = val
		}
	}
	v.Token = token

	return v, nil
}

// BuildUpdateCoursePayload builds the payload for the admin UpdateCourse
// endpoint from CLI flags.
func BuildUpdateCoursePayload(adminUpdateCourseBody string, adminUpdateCourseID string, adminUpdateCourseToken string) (*admin.UpdateCoursePayload, error) {
	var err error
	var body UpdateCourseRequestBody
	{
		err = json.Unmarshal([]byte(adminUpdateCourseBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"author_ids\": [\n         \"b29869ec-6eef-451c-a389-249a6d6ac47b\"\n      ],\n      \"category\": \"web\",\n      \"description\": \"markdown text\",\n      \"difficulty\": \"advanced\",\n      \"publish\": true,\n      \"slug\": \"sqli\",\n      \"title\": \"SQL-Injektioner\"\n   }'")
		}
	}
	var id string
	{
		id = adminUpdateCourseID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var token string
	{
		token = adminUpdateCourseToken
	}
	v := &admin.UpdateCoursePayload{
		Title:       body.Title,
		Slug:        body.Slug,
		Category:    body.Category,
		Difficulty:  body.Difficulty,
		Description: body.Description,
		Publish:     body.Publish,
	}
	if body.AuthorIds != nil {
		v.AuthorIds = make([]string, len(body.AuthorIds))
		for i, val := range body.AuthorIds {
			v.AuthorIds[i] = val
		}
	}
	v.ID = id
	v.Token = token

	return v, nil
}
