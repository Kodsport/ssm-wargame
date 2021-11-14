// Code generated by goa v3.5.2, DO NOT EDIT.
//
// challenge HTTP server types
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	challengeviews "github.com/sakerhetsm/ssm-wargame/internal/gen/challenge/views"
)

// SsmChallengeResponseCollection is the type of the "challenge" service
// "ListChallenges" endpoint HTTP response body.
type SsmChallengeResponseCollection []*SsmChallengeResponse

// SsmChallengeResponse is used to define fields on response body types.
type SsmChallengeResponse struct {
	ID          *string                     `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Title       *string                     `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Description *string                     `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Score       *int                        `form:"score,omitempty" json:"score,omitempty" xml:"score,omitempty"`
	Published   *bool                       `form:"published,omitempty" json:"published,omitempty" xml:"published,omitempty"`
	Services    []*ChallengeServiceResponse `form:"services,omitempty" json:"services,omitempty" xml:"services,omitempty"`
	Files       []*ChallengeFilesResponse   `form:"files,omitempty" json:"files,omitempty" xml:"files,omitempty"`
}

// ChallengeServiceResponse is used to define fields on response body types.
type ChallengeServiceResponse struct {
}

// ChallengeFilesResponse is used to define fields on response body types.
type ChallengeFilesResponse struct {
}

// NewSsmChallengeResponseCollection builds the HTTP response body from the
// result of the "ListChallenges" endpoint of the "challenge" service.
func NewSsmChallengeResponseCollection(res challengeviews.SsmChallengeCollectionView) SsmChallengeResponseCollection {
	body := make([]*SsmChallengeResponse, len(res))
	for i, val := range res {
		body[i] = marshalChallengeviewsSsmChallengeViewToSsmChallengeResponse(val)
	}
	return body
}
