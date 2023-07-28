// Code generated by goa v3.5.2, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package server

import (
	user "github.com/sakerhetsm/ssm-wargame/internal/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// JoinSchoolRequestBody is the type of the "user" service "JoinSchool"
// endpoint HTTP request body.
type JoinSchoolRequestBody struct {
	SchoolID *int `form:"school_id,omitempty" json:"school_id,omitempty" xml:"school_id,omitempty"`
}

// GetSelfResponseBody is the type of the "user" service "GetSelf" endpoint
// HTTP response body.
type GetSelfResponseBody struct {
	SchoolName *string `form:"school_name,omitempty" json:"school_name,omitempty" xml:"school_name,omitempty"`
	ID         string  `form:"id" json:"id" xml:"id"`
	Email      string  `form:"email" json:"email" xml:"email"`
	FirstName  string  `form:"first_name" json:"first_name" xml:"first_name"`
	LastName   string  `form:"last_name" json:"last_name" xml:"last_name"`
	Role       string  `form:"role" json:"role" xml:"role"`
	SchoolID   *int    `form:"school_id,omitempty" json:"school_id,omitempty" xml:"school_id,omitempty"`
}

// SearchSchoolsResponseBody is the type of the "user" service "SearchSchools"
// endpoint HTTP response body.
type SearchSchoolsResponseBody []*SchoolResponse

// SchoolResponse is used to define fields on response body types.
type SchoolResponse struct {
	ID               int    `form:"id" json:"id" xml:"id"`
	Name             string `form:"name" json:"name" xml:"name"`
	MunicipalityName string `form:"municipality_name" json:"municipality_name" xml:"municipality_name"`
}

// NewGetSelfResponseBody builds the HTTP response body from the result of the
// "GetSelf" endpoint of the "user" service.
func NewGetSelfResponseBody(res *user.GetSelfResult) *GetSelfResponseBody {
	body := &GetSelfResponseBody{
		SchoolName: res.SchoolName,
		ID:         res.ID,
		Email:      res.Email,
		FirstName:  res.FirstName,
		LastName:   res.LastName,
		Role:       res.Role,
		SchoolID:   res.SchoolID,
	}
	return body
}

// NewSearchSchoolsResponseBody builds the HTTP response body from the result
// of the "SearchSchools" endpoint of the "user" service.
func NewSearchSchoolsResponseBody(res []*user.School) SearchSchoolsResponseBody {
	body := make([]*SchoolResponse, len(res))
	for i, val := range res {
		body[i] = marshalUserSchoolToSchoolResponse(val)
	}
	return body
}

// NewGetSelfPayload builds a user service GetSelf endpoint payload.
func NewGetSelfPayload(token string) *user.GetSelfPayload {
	v := &user.GetSelfPayload{}
	v.Token = token

	return v
}

// NewJoinSchoolPayload builds a user service JoinSchool endpoint payload.
func NewJoinSchoolPayload(body *JoinSchoolRequestBody, token string) *user.JoinSchoolPayload {
	v := &user.JoinSchoolPayload{
		SchoolID: *body.SchoolID,
	}
	v.Token = token

	return v
}

// NewLeaveSchoolPayload builds a user service LeaveSchool endpoint payload.
func NewLeaveSchoolPayload(token string) *user.LeaveSchoolPayload {
	v := &user.LeaveSchoolPayload{}
	v.Token = token

	return v
}

// NewSearchSchoolsPayload builds a user service SearchSchools endpoint payload.
func NewSearchSchoolsPayload(q string, token string) *user.SearchSchoolsPayload {
	v := &user.SearchSchoolsPayload{}
	v.Q = q
	v.Token = token

	return v
}

// ValidateJoinSchoolRequestBody runs the validations defined on
// JoinSchoolRequestBody
func ValidateJoinSchoolRequestBody(body *JoinSchoolRequestBody) (err error) {
	if body.SchoolID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("school_id", "body"))
	}
	return
}
