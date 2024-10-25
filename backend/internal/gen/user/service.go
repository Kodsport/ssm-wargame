// Code generated by goa v3.5.2, DO NOT EDIT.
//
// user service
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the user service interface.
type Service interface {
	// GetSelf implements GetSelf.
	GetSelf(context.Context, *GetSelfPayload) (res *GetSelfResult, err error)
	// UpdateSelf implements UpdateSelf.
	UpdateSelf(context.Context, *UpdateSelfPayload) (err error)
	// CompleteOnboarding implements CompleteOnboarding.
	CompleteOnboarding(context.Context, *CompleteOnboardingPayload) (err error)
	// JoinSchool implements JoinSchool.
	JoinSchool(context.Context, *JoinSchoolPayload) (err error)
	// LeaveSchool implements LeaveSchool.
	LeaveSchool(context.Context, *LeaveSchoolPayload) (err error)
	// SearchSchools implements SearchSchools.
	SearchSchools(context.Context, *SearchSchoolsPayload) (res []*School, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "user"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"GetSelf", "UpdateSelf", "CompleteOnboarding", "JoinSchool", "LeaveSchool", "SearchSchools"}

// GetSelfPayload is the payload type of the user service GetSelf method.
type GetSelfPayload struct {
	Token string
}

// GetSelfResult is the result type of the user service GetSelf method.
type GetSelfResult struct {
	SchoolName     *string
	OnboardingDone bool
	ID             string
	Email          string
	FullName       string
	Role           string
	SchoolID       *string
}

// UpdateSelfPayload is the payload type of the user service UpdateSelf method.
type UpdateSelfPayload struct {
	FullName string
	Token    string
}

// CompleteOnboardingPayload is the payload type of the user service
// CompleteOnboarding method.
type CompleteOnboardingPayload struct {
	Token string
}

// JoinSchoolPayload is the payload type of the user service JoinSchool method.
type JoinSchoolPayload struct {
	SchoolID string
	Token    string
}

// LeaveSchoolPayload is the payload type of the user service LeaveSchool
// method.
type LeaveSchoolPayload struct {
	Token string
}

// SearchSchoolsPayload is the payload type of the user service SearchSchools
// method.
type SearchSchoolsPayload struct {
	Q          string
	University *bool
	Token      string
}

type School struct {
	Name             string
	MunicipalityName string
	IsUniversity     bool
	// ID of a file
	ID string
}

// MakeInvalidUsername builds a goa.ServiceError from an error.
func MakeInvalidUsername(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "InvalidUsername",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
