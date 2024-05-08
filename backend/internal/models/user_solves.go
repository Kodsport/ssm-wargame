// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserSolf is an object representing the database table.
type UserSolf struct {
	UserID      string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	ChallengeID string    `boil:"challenge_id" json:"challenge_id" toml:"challenge_id" yaml:"challenge_id"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *userSolfR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userSolfL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserSolfColumns = struct {
	UserID      string
	ChallengeID string
	CreatedAt   string
}{
	UserID:      "user_id",
	ChallengeID: "challenge_id",
	CreatedAt:   "created_at",
}

var UserSolfTableColumns = struct {
	UserID      string
	ChallengeID string
	CreatedAt   string
}{
	UserID:      "user_solves.user_id",
	ChallengeID: "user_solves.challenge_id",
	CreatedAt:   "user_solves.created_at",
}

// Generated where

var UserSolfWhere = struct {
	UserID      whereHelperstring
	ChallengeID whereHelperstring
	CreatedAt   whereHelpertime_Time
}{
	UserID:      whereHelperstring{field: "\"user_solves\".\"user_id\""},
	ChallengeID: whereHelperstring{field: "\"user_solves\".\"challenge_id\""},
	CreatedAt:   whereHelpertime_Time{field: "\"user_solves\".\"created_at\""},
}

// UserSolfRels is where relationship names are stored.
var UserSolfRels = struct {
	Challenge string
	User      string
}{
	Challenge: "Challenge",
	User:      "User",
}

// userSolfR is where relationships are stored.
type userSolfR struct {
	Challenge *Challenge `boil:"Challenge" json:"Challenge" toml:"Challenge" yaml:"Challenge"`
	User      *User      `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userSolfR) NewStruct() *userSolfR {
	return &userSolfR{}
}

func (r *userSolfR) GetChallenge() *Challenge {
	if r == nil {
		return nil
	}
	return r.Challenge
}

func (r *userSolfR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// userSolfL is where Load methods for each relationship are stored.
type userSolfL struct{}

var (
	userSolfAllColumns            = []string{"user_id", "challenge_id", "created_at"}
	userSolfColumnsWithoutDefault = []string{"user_id", "challenge_id"}
	userSolfColumnsWithDefault    = []string{"created_at"}
	userSolfPrimaryKeyColumns     = []string{"user_id", "challenge_id"}
	userSolfGeneratedColumns      = []string{}
)

type (
	// UserSolfSlice is an alias for a slice of pointers to UserSolf.
	// This should almost always be used instead of []UserSolf.
	UserSolfSlice []*UserSolf
	// UserSolfHook is the signature for custom UserSolf hook methods
	UserSolfHook func(context.Context, boil.ContextExecutor, *UserSolf) error

	userSolfQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userSolfType                 = reflect.TypeOf(&UserSolf{})
	userSolfMapping              = queries.MakeStructMapping(userSolfType)
	userSolfPrimaryKeyMapping, _ = queries.BindMapping(userSolfType, userSolfMapping, userSolfPrimaryKeyColumns)
	userSolfInsertCacheMut       sync.RWMutex
	userSolfInsertCache          = make(map[string]insertCache)
	userSolfUpdateCacheMut       sync.RWMutex
	userSolfUpdateCache          = make(map[string]updateCache)
	userSolfUpsertCacheMut       sync.RWMutex
	userSolfUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userSolfAfterSelectMu sync.Mutex
var userSolfAfterSelectHooks []UserSolfHook

var userSolfBeforeInsertMu sync.Mutex
var userSolfBeforeInsertHooks []UserSolfHook
var userSolfAfterInsertMu sync.Mutex
var userSolfAfterInsertHooks []UserSolfHook

var userSolfBeforeUpdateMu sync.Mutex
var userSolfBeforeUpdateHooks []UserSolfHook
var userSolfAfterUpdateMu sync.Mutex
var userSolfAfterUpdateHooks []UserSolfHook

var userSolfBeforeDeleteMu sync.Mutex
var userSolfBeforeDeleteHooks []UserSolfHook
var userSolfAfterDeleteMu sync.Mutex
var userSolfAfterDeleteHooks []UserSolfHook

var userSolfBeforeUpsertMu sync.Mutex
var userSolfBeforeUpsertHooks []UserSolfHook
var userSolfAfterUpsertMu sync.Mutex
var userSolfAfterUpsertHooks []UserSolfHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserSolf) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userSolfAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserSolf) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userSolfBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserSolf) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userSolfAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserSolf) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userSolfBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserSolf) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userSolfAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserSolf) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userSolfBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserSolf) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userSolfAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserSolf) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userSolfBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserSolf) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userSolfAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserSolfHook registers your hook function for all future operations.
func AddUserSolfHook(hookPoint boil.HookPoint, userSolfHook UserSolfHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userSolfAfterSelectMu.Lock()
		userSolfAfterSelectHooks = append(userSolfAfterSelectHooks, userSolfHook)
		userSolfAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		userSolfBeforeInsertMu.Lock()
		userSolfBeforeInsertHooks = append(userSolfBeforeInsertHooks, userSolfHook)
		userSolfBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		userSolfAfterInsertMu.Lock()
		userSolfAfterInsertHooks = append(userSolfAfterInsertHooks, userSolfHook)
		userSolfAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		userSolfBeforeUpdateMu.Lock()
		userSolfBeforeUpdateHooks = append(userSolfBeforeUpdateHooks, userSolfHook)
		userSolfBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		userSolfAfterUpdateMu.Lock()
		userSolfAfterUpdateHooks = append(userSolfAfterUpdateHooks, userSolfHook)
		userSolfAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		userSolfBeforeDeleteMu.Lock()
		userSolfBeforeDeleteHooks = append(userSolfBeforeDeleteHooks, userSolfHook)
		userSolfBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		userSolfAfterDeleteMu.Lock()
		userSolfAfterDeleteHooks = append(userSolfAfterDeleteHooks, userSolfHook)
		userSolfAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		userSolfBeforeUpsertMu.Lock()
		userSolfBeforeUpsertHooks = append(userSolfBeforeUpsertHooks, userSolfHook)
		userSolfBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		userSolfAfterUpsertMu.Lock()
		userSolfAfterUpsertHooks = append(userSolfAfterUpsertHooks, userSolfHook)
		userSolfAfterUpsertMu.Unlock()
	}
}

// One returns a single userSolf record from the query.
func (q userSolfQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserSolf, error) {
	o := &UserSolf{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_solves")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserSolf records from the query.
func (q userSolfQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserSolfSlice, error) {
	var o []*UserSolf

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserSolf slice")
	}

	if len(userSolfAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserSolf records in the query.
func (q userSolfQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_solves rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userSolfQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_solves exists")
	}

	return count > 0, nil
}

// Challenge pointed to by the foreign key.
func (o *UserSolf) Challenge(mods ...qm.QueryMod) challengeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChallengeID),
	}

	queryMods = append(queryMods, mods...)

	return Challenges(queryMods...)
}

// User pointed to by the foreign key.
func (o *UserSolf) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadChallenge allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userSolfL) LoadChallenge(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserSolf interface{}, mods queries.Applicator) error {
	var slice []*UserSolf
	var object *UserSolf

	if singular {
		var ok bool
		object, ok = maybeUserSolf.(*UserSolf)
		if !ok {
			object = new(UserSolf)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserSolf)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserSolf))
			}
		}
	} else {
		s, ok := maybeUserSolf.(*[]*UserSolf)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserSolf)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserSolf))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &userSolfR{}
		}
		args[object.ChallengeID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userSolfR{}
			}

			args[obj.ChallengeID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`challenges`),
		qm.WhereIn(`challenges.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Challenge")
	}

	var resultSlice []*Challenge
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Challenge")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for challenges")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for challenges")
	}

	if len(challengeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Challenge = foreign
		if foreign.R == nil {
			foreign.R = &challengeR{}
		}
		foreign.R.UserSolves = append(foreign.R.UserSolves, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChallengeID == foreign.ID {
				local.R.Challenge = foreign
				if foreign.R == nil {
					foreign.R = &challengeR{}
				}
				foreign.R.UserSolves = append(foreign.R.UserSolves, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userSolfL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserSolf interface{}, mods queries.Applicator) error {
	var slice []*UserSolf
	var object *UserSolf

	if singular {
		var ok bool
		object, ok = maybeUserSolf.(*UserSolf)
		if !ok {
			object = new(UserSolf)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserSolf)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserSolf))
			}
		}
	} else {
		s, ok := maybeUserSolf.(*[]*UserSolf)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserSolf)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserSolf))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &userSolfR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userSolfR{}
			}

			args[obj.UserID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserSolves = append(foreign.R.UserSolves, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserSolves = append(foreign.R.UserSolves, local)
				break
			}
		}
	}

	return nil
}

// SetChallenge of the userSolf to the related item.
// Sets o.R.Challenge to related.
// Adds o to related.R.UserSolves.
func (o *UserSolf) SetChallenge(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Challenge) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_solves\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"challenge_id"}),
		strmangle.WhereClause("\"", "\"", 2, userSolfPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.ChallengeID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ChallengeID = related.ID
	if o.R == nil {
		o.R = &userSolfR{
			Challenge: related,
		}
	} else {
		o.R.Challenge = related
	}

	if related.R == nil {
		related.R = &challengeR{
			UserSolves: UserSolfSlice{o},
		}
	} else {
		related.R.UserSolves = append(related.R.UserSolves, o)
	}

	return nil
}

// SetUser of the userSolf to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserSolves.
func (o *UserSolf) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_solves\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userSolfPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.ChallengeID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userSolfR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserSolves: UserSolfSlice{o},
		}
	} else {
		related.R.UserSolves = append(related.R.UserSolves, o)
	}

	return nil
}

// UserSolves retrieves all the records using an executor.
func UserSolves(mods ...qm.QueryMod) userSolfQuery {
	mods = append(mods, qm.From("\"user_solves\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_solves\".*"})
	}

	return userSolfQuery{q}
}

// FindUserSolf retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserSolf(ctx context.Context, exec boil.ContextExecutor, userID string, challengeID string, selectCols ...string) (*UserSolf, error) {
	userSolfObj := &UserSolf{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_solves\" where \"user_id\"=$1 AND \"challenge_id\"=$2", sel,
	)

	q := queries.Raw(query, userID, challengeID)

	err := q.Bind(ctx, exec, userSolfObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_solves")
	}

	if err = userSolfObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userSolfObj, err
	}

	return userSolfObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserSolf) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_solves provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userSolfColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userSolfInsertCacheMut.RLock()
	cache, cached := userSolfInsertCache[key]
	userSolfInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userSolfAllColumns,
			userSolfColumnsWithDefault,
			userSolfColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userSolfType, userSolfMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userSolfType, userSolfMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_solves\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_solves\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into user_solves")
	}

	if !cached {
		userSolfInsertCacheMut.Lock()
		userSolfInsertCache[key] = cache
		userSolfInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserSolf.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserSolf) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userSolfUpdateCacheMut.RLock()
	cache, cached := userSolfUpdateCache[key]
	userSolfUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userSolfAllColumns,
			userSolfPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_solves, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_solves\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userSolfPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userSolfType, userSolfMapping, append(wl, userSolfPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update user_solves row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_solves")
	}

	if !cached {
		userSolfUpdateCacheMut.Lock()
		userSolfUpdateCache[key] = cache
		userSolfUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userSolfQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_solves")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_solves")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserSolfSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userSolfPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_solves\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userSolfPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userSolf slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userSolf")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserSolf) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no user_solves provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userSolfColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userSolfUpsertCacheMut.RLock()
	cache, cached := userSolfUpsertCache[key]
	userSolfUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			userSolfAllColumns,
			userSolfColumnsWithDefault,
			userSolfColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userSolfAllColumns,
			userSolfPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_solves, could not build update column list")
		}

		ret := strmangle.SetComplement(userSolfAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(userSolfPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert user_solves, could not build conflict column list")
			}

			conflict = make([]string, len(userSolfPrimaryKeyColumns))
			copy(conflict, userSolfPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_solves\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(userSolfType, userSolfMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userSolfType, userSolfMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert user_solves")
	}

	if !cached {
		userSolfUpsertCacheMut.Lock()
		userSolfUpsertCache[key] = cache
		userSolfUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserSolf record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserSolf) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserSolf provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userSolfPrimaryKeyMapping)
	sql := "DELETE FROM \"user_solves\" WHERE \"user_id\"=$1 AND \"challenge_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_solves")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_solves")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userSolfQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userSolfQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_solves")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_solves")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserSolfSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userSolfBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userSolfPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_solves\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userSolfPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userSolf slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_solves")
	}

	if len(userSolfAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserSolf) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserSolf(ctx, exec, o.UserID, o.ChallengeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserSolfSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserSolfSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userSolfPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_solves\".* FROM \"user_solves\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userSolfPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserSolfSlice")
	}

	*o = slice

	return nil
}

// UserSolfExists checks if the UserSolf row exists.
func UserSolfExists(ctx context.Context, exec boil.ContextExecutor, userID string, challengeID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_solves\" where \"user_id\"=$1 AND \"challenge_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID, challengeID)
	}
	row := exec.QueryRowContext(ctx, sql, userID, challengeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_solves exists")
	}

	return exists, nil
}

// Exists checks if the UserSolf row exists.
func (o *UserSolf) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserSolfExists(ctx, exec, o.UserID, o.ChallengeID)
}
