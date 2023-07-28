// Code generated by goa v3.5.2, DO NOT EDIT.
//
// user HTTP client types
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package client

import (
	user "github.com/sakerhetsm/ssm-wargame/internal/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// JoinSchoolRequestBody is the type of the "user" service "JoinSchool"
// endpoint HTTP request body.
type JoinSchoolRequestBody struct {
	SchoolID int `form:"school_id" json:"school_id" xml:"school_id"`
}

// GetSelfResponseBody is the type of the "user" service "GetSelf" endpoint
// HTTP response body.
type GetSelfResponseBody struct {
	SchoolName *string `form:"school_name,omitempty" json:"school_name,omitempty" xml:"school_name,omitempty"`
	ID         *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Email      *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	FirstName  *string `form:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName   *string `form:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty"`
	Role       *string `form:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	SchoolID   *int    `form:"school_id,omitempty" json:"school_id,omitempty" xml:"school_id,omitempty"`
}

// SearchSchoolsResponseBody is the type of the "user" service "SearchSchools"
// endpoint HTTP response body.
type SearchSchoolsResponseBody []*SchoolResponse

// SchoolResponse is used to define fields on response body types.
type SchoolResponse struct {
	ID               *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name             *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	MunicipalityName *string `form:"municipality_name,omitempty" json:"municipality_name,omitempty" xml:"municipality_name,omitempty"`
}

// NewJoinSchoolRequestBody builds the HTTP request body from the payload of
// the "JoinSchool" endpoint of the "user" service.
func NewJoinSchoolRequestBody(p *user.JoinSchoolPayload) *JoinSchoolRequestBody {
	body := &JoinSchoolRequestBody{
		SchoolID: p.SchoolID,
	}
	return body
}

// NewGetSelfResultOK builds a "user" service "GetSelf" endpoint result from a
// HTTP "OK" response.
func NewGetSelfResultOK(body *GetSelfResponseBody) *user.GetSelfResult {
	v := &user.GetSelfResult{
		SchoolName: body.SchoolName,
		ID:         *body.ID,
		Email:      *body.Email,
		FirstName:  *body.FirstName,
		LastName:   *body.LastName,
		Role:       *body.Role,
		SchoolID:   body.SchoolID,
	}

	return v
}

// NewSearchSchoolsSchoolOK builds a "user" service "SearchSchools" endpoint
// result from a HTTP "OK" response.
func NewSearchSchoolsSchoolOK(body []*SchoolResponse) []*user.School {
	v := make([]*user.School, len(body))
	for i, val := range body {
		v[i] = unmarshalSchoolResponseToUserSchool(val)
	}

	return v
}

// ValidateGetSelfResponseBody runs the validations defined on
// GetSelfResponseBody
func ValidateGetSelfResponseBody(body *GetSelfResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("first_name", "body"))
	}
	if body.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last_name", "body"))
	}
	if body.Role == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("role", "body"))
	}
	return
}

// ValidateSchoolResponse runs the validations defined on SchoolResponse
func ValidateSchoolResponse(body *SchoolResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.MunicipalityName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("municipality_name", "body"))
	}
	return
}
