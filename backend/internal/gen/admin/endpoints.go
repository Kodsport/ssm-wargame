// Code generated by goa v3.5.2, DO NOT EDIT.
//
// admin endpoints
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package admin

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "admin" service endpoints.
type Endpoints struct {
	ListChallenges         goa.Endpoint
	GetChallengeMeta       goa.Endpoint
	CreateChallenge        goa.Endpoint
	UpdateChallenge        goa.Endpoint
	PresignChallFileUpload goa.Endpoint
	ListMonthlyChallenges  goa.Endpoint
	DeleteMonthlyChallenge goa.Endpoint
	DeleteFile             goa.Endpoint
	CreateMonthlyChallenge goa.Endpoint
	ListUsers              goa.Endpoint
	AddFlag                goa.Endpoint
	DeleteFlag             goa.Endpoint
	ListCategories         goa.Endpoint
	ChalltoolsImport       goa.Endpoint
}

// NewEndpoints wraps the methods of the "admin" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		ListChallenges:         NewListChallengesEndpoint(s, a.JWTAuth),
		GetChallengeMeta:       NewGetChallengeMetaEndpoint(s, a.JWTAuth),
		CreateChallenge:        NewCreateChallengeEndpoint(s, a.JWTAuth),
		UpdateChallenge:        NewUpdateChallengeEndpoint(s, a.JWTAuth),
		PresignChallFileUpload: NewPresignChallFileUploadEndpoint(s, a.JWTAuth),
		ListMonthlyChallenges:  NewListMonthlyChallengesEndpoint(s, a.JWTAuth),
		DeleteMonthlyChallenge: NewDeleteMonthlyChallengeEndpoint(s, a.JWTAuth),
		DeleteFile:             NewDeleteFileEndpoint(s, a.JWTAuth),
		CreateMonthlyChallenge: NewCreateMonthlyChallengeEndpoint(s, a.JWTAuth),
		ListUsers:              NewListUsersEndpoint(s, a.JWTAuth),
		AddFlag:                NewAddFlagEndpoint(s, a.JWTAuth),
		DeleteFlag:             NewDeleteFlagEndpoint(s, a.JWTAuth),
		ListCategories:         NewListCategoriesEndpoint(s, a.JWTAuth),
		ChalltoolsImport:       NewChalltoolsImportEndpoint(s, a.APIKeyAuth),
	}
}

// Use applies the given middleware to all the "admin" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ListChallenges = m(e.ListChallenges)
	e.GetChallengeMeta = m(e.GetChallengeMeta)
	e.CreateChallenge = m(e.CreateChallenge)
	e.UpdateChallenge = m(e.UpdateChallenge)
	e.PresignChallFileUpload = m(e.PresignChallFileUpload)
	e.ListMonthlyChallenges = m(e.ListMonthlyChallenges)
	e.DeleteMonthlyChallenge = m(e.DeleteMonthlyChallenge)
	e.DeleteFile = m(e.DeleteFile)
	e.CreateMonthlyChallenge = m(e.CreateMonthlyChallenge)
	e.ListUsers = m(e.ListUsers)
	e.AddFlag = m(e.AddFlag)
	e.DeleteFlag = m(e.DeleteFlag)
	e.ListCategories = m(e.ListCategories)
	e.ChalltoolsImport = m(e.ChalltoolsImport)
}

// NewListChallengesEndpoint returns an endpoint function that calls the method
// "ListChallenges" of service "admin".
func NewListChallengesEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListChallengesPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.ListChallenges(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedSsmAdminChallengeCollection(res, "default")
		return vres, nil
	}
}

// NewGetChallengeMetaEndpoint returns an endpoint function that calls the
// method "GetChallengeMeta" of service "admin".
func NewGetChallengeMetaEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetChallengeMetaPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.GetChallengeMeta(ctx, p)
	}
}

// NewCreateChallengeEndpoint returns an endpoint function that calls the
// method "CreateChallenge" of service "admin".
func NewCreateChallengeEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateChallengePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.CreateChallenge(ctx, p)
	}
}

// NewUpdateChallengeEndpoint returns an endpoint function that calls the
// method "UpdateChallenge" of service "admin".
func NewUpdateChallengeEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateChallengePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.UpdateChallenge(ctx, p)
	}
}

// NewPresignChallFileUploadEndpoint returns an endpoint function that calls
// the method "PresignChallFileUpload" of service "admin".
func NewPresignChallFileUploadEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PresignChallFileUploadPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.PresignChallFileUpload(ctx, p)
	}
}

// NewListMonthlyChallengesEndpoint returns an endpoint function that calls the
// method "ListMonthlyChallenges" of service "admin".
func NewListMonthlyChallengesEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListMonthlyChallengesPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ListMonthlyChallenges(ctx, p)
	}
}

// NewDeleteMonthlyChallengeEndpoint returns an endpoint function that calls
// the method "DeleteMonthlyChallenge" of service "admin".
func NewDeleteMonthlyChallengeEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteMonthlyChallengePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteMonthlyChallenge(ctx, p)
	}
}

// NewDeleteFileEndpoint returns an endpoint function that calls the method
// "DeleteFile" of service "admin".
func NewDeleteFileEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteFilePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteFile(ctx, p)
	}
}

// NewCreateMonthlyChallengeEndpoint returns an endpoint function that calls
// the method "CreateMonthlyChallenge" of service "admin".
func NewCreateMonthlyChallengeEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateMonthlyChallengePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.CreateMonthlyChallenge(ctx, p)
	}
}

// NewListUsersEndpoint returns an endpoint function that calls the method
// "ListUsers" of service "admin".
func NewListUsersEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListUsersPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ListUsers(ctx, p)
	}
}

// NewAddFlagEndpoint returns an endpoint function that calls the method
// "AddFlag" of service "admin".
func NewAddFlagEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddFlagPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.AddFlag(ctx, p)
	}
}

// NewDeleteFlagEndpoint returns an endpoint function that calls the method
// "DeleteFlag" of service "admin".
func NewDeleteFlagEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteFlagPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteFlag(ctx, p)
	}
}

// NewListCategoriesEndpoint returns an endpoint function that calls the method
// "ListCategories" of service "admin".
func NewListCategoriesEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListCategoriesPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.ListCategories(ctx, p)
	}
}

// NewChalltoolsImportEndpoint returns an endpoint function that calls the
// method "ChalltoolsImport" of service "admin".
func NewChalltoolsImportEndpoint(s Service, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ChalltoolsImportPayload)
		var err error
		sc := security.APIKeyScheme{
			Name:           "import_token",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authAPIKeyFn(ctx, p.ImportToken, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.ChalltoolsImport(ctx, p)
	}
}
