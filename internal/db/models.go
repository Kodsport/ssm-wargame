// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Category struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type Challenge struct {
	ID          uuid.UUID
	Slug        string
	Title       string
	Description string
	Score       int32
	Published   bool
	Services    pgtype.JSON
	Files       pgtype.JSON
	CtfEventID  uuid.NullUUID
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}

type ChallengeAuthor struct {
	ChallengeID uuid.UUID
	UserID      uuid.UUID
}

type ChallengeCategory struct {
	ChallengeID uuid.UUID
	CategoryID  uuid.UUID
}

type ChallengeFile struct {
	ID          uuid.UUID
	ChallengeID uuid.UUID
	Url         string
	Filename    string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}

type ChallengeService struct {
	ID          uuid.UUID
	ChallengeID uuid.UUID
	Type        string
	Url         sql.NullString
	Host        sql.NullString
	Port        sql.NullInt32
	Name        sql.NullString
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}

type CtfEvent struct {
	ID        uuid.UUID
	Name      string
	Url       string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type Flag struct {
	ID          uuid.UUID
	ChallengeID uuid.UUID
	Flag        string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}

type MonthlyChallenge struct {
	ChallengeID  uuid.UUID
	StartDate    time.Time
	EndDate      time.Time
	DisplayMonth time.Time
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
}

type School struct {
	ID                   int32
	Name                 string
	GeographicalAreaCode sql.NullInt32
	CreatedAt            time.Time
	UpdatedAt            sql.NullTime
}

type Submission struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	ChallengeID uuid.UUID
	Successful  bool
	Input       string
	CreatedAt   time.Time
}

type User struct {
	ID        uuid.UUID
	DiscordID sql.NullString
	FirstName sql.NullString
	LastName  sql.NullString
	Email     string
	SchoolID  sql.NullInt32
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserSolf struct {
	UserID      uuid.UUID
	ChallengeID uuid.UUID
	CreatedAt   time.Time
}
