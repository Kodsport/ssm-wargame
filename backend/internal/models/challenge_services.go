// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ChallengeService is an object representing the database table.
type ChallengeService struct {
	ID          string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	ChallengeID string    `boil:"challenge_id" json:"challenge_id" toml:"challenge_id" yaml:"challenge_id"`
	UserDisplay string    `boil:"user_display" json:"user_display" toml:"user_display" yaml:"user_display"`
	Hyperlink   string    `boil:"hyperlink" json:"hyperlink" toml:"hyperlink" yaml:"hyperlink"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *challengeServiceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L challengeServiceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChallengeServiceColumns = struct {
	ID          string
	ChallengeID string
	UserDisplay string
	Hyperlink   string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	ChallengeID: "challenge_id",
	UserDisplay: "user_display",
	Hyperlink:   "hyperlink",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var ChallengeServiceTableColumns = struct {
	ID          string
	ChallengeID string
	UserDisplay string
	Hyperlink   string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "challenge_services.id",
	ChallengeID: "challenge_services.challenge_id",
	UserDisplay: "challenge_services.user_display",
	Hyperlink:   "challenge_services.hyperlink",
	CreatedAt:   "challenge_services.created_at",
	UpdatedAt:   "challenge_services.updated_at",
}

// Generated where

var ChallengeServiceWhere = struct {
	ID          whereHelperstring
	ChallengeID whereHelperstring
	UserDisplay whereHelperstring
	Hyperlink   whereHelperstring
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpernull_Time
}{
	ID:          whereHelperstring{field: "\"challenge_services\".\"id\""},
	ChallengeID: whereHelperstring{field: "\"challenge_services\".\"challenge_id\""},
	UserDisplay: whereHelperstring{field: "\"challenge_services\".\"user_display\""},
	Hyperlink:   whereHelperstring{field: "\"challenge_services\".\"hyperlink\""},
	CreatedAt:   whereHelpertime_Time{field: "\"challenge_services\".\"created_at\""},
	UpdatedAt:   whereHelpernull_Time{field: "\"challenge_services\".\"updated_at\""},
}

// ChallengeServiceRels is where relationship names are stored.
var ChallengeServiceRels = struct {
	Challenge string
}{
	Challenge: "Challenge",
}

// challengeServiceR is where relationships are stored.
type challengeServiceR struct {
	Challenge *Challenge `boil:"Challenge" json:"Challenge" toml:"Challenge" yaml:"Challenge"`
}

// NewStruct creates a new relationship struct
func (*challengeServiceR) NewStruct() *challengeServiceR {
	return &challengeServiceR{}
}

func (r *challengeServiceR) GetChallenge() *Challenge {
	if r == nil {
		return nil
	}
	return r.Challenge
}

// challengeServiceL is where Load methods for each relationship are stored.
type challengeServiceL struct{}

var (
	challengeServiceAllColumns            = []string{"id", "challenge_id", "user_display", "hyperlink", "created_at", "updated_at"}
	challengeServiceColumnsWithoutDefault = []string{"id", "challenge_id", "user_display", "hyperlink"}
	challengeServiceColumnsWithDefault    = []string{"created_at", "updated_at"}
	challengeServicePrimaryKeyColumns     = []string{"id"}
	challengeServiceGeneratedColumns      = []string{}
)

type (
	// ChallengeServiceSlice is an alias for a slice of pointers to ChallengeService.
	// This should almost always be used instead of []ChallengeService.
	ChallengeServiceSlice []*ChallengeService
	// ChallengeServiceHook is the signature for custom ChallengeService hook methods
	ChallengeServiceHook func(context.Context, boil.ContextExecutor, *ChallengeService) error

	challengeServiceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	challengeServiceType                 = reflect.TypeOf(&ChallengeService{})
	challengeServiceMapping              = queries.MakeStructMapping(challengeServiceType)
	challengeServicePrimaryKeyMapping, _ = queries.BindMapping(challengeServiceType, challengeServiceMapping, challengeServicePrimaryKeyColumns)
	challengeServiceInsertCacheMut       sync.RWMutex
	challengeServiceInsertCache          = make(map[string]insertCache)
	challengeServiceUpdateCacheMut       sync.RWMutex
	challengeServiceUpdateCache          = make(map[string]updateCache)
	challengeServiceUpsertCacheMut       sync.RWMutex
	challengeServiceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var challengeServiceAfterSelectHooks []ChallengeServiceHook

var challengeServiceBeforeInsertHooks []ChallengeServiceHook
var challengeServiceAfterInsertHooks []ChallengeServiceHook

var challengeServiceBeforeUpdateHooks []ChallengeServiceHook
var challengeServiceAfterUpdateHooks []ChallengeServiceHook

var challengeServiceBeforeDeleteHooks []ChallengeServiceHook
var challengeServiceAfterDeleteHooks []ChallengeServiceHook

var challengeServiceBeforeUpsertHooks []ChallengeServiceHook
var challengeServiceAfterUpsertHooks []ChallengeServiceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ChallengeService) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeServiceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ChallengeService) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeServiceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ChallengeService) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeServiceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ChallengeService) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeServiceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ChallengeService) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeServiceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ChallengeService) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeServiceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ChallengeService) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeServiceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ChallengeService) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeServiceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ChallengeService) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeServiceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChallengeServiceHook registers your hook function for all future operations.
func AddChallengeServiceHook(hookPoint boil.HookPoint, challengeServiceHook ChallengeServiceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		challengeServiceAfterSelectHooks = append(challengeServiceAfterSelectHooks, challengeServiceHook)
	case boil.BeforeInsertHook:
		challengeServiceBeforeInsertHooks = append(challengeServiceBeforeInsertHooks, challengeServiceHook)
	case boil.AfterInsertHook:
		challengeServiceAfterInsertHooks = append(challengeServiceAfterInsertHooks, challengeServiceHook)
	case boil.BeforeUpdateHook:
		challengeServiceBeforeUpdateHooks = append(challengeServiceBeforeUpdateHooks, challengeServiceHook)
	case boil.AfterUpdateHook:
		challengeServiceAfterUpdateHooks = append(challengeServiceAfterUpdateHooks, challengeServiceHook)
	case boil.BeforeDeleteHook:
		challengeServiceBeforeDeleteHooks = append(challengeServiceBeforeDeleteHooks, challengeServiceHook)
	case boil.AfterDeleteHook:
		challengeServiceAfterDeleteHooks = append(challengeServiceAfterDeleteHooks, challengeServiceHook)
	case boil.BeforeUpsertHook:
		challengeServiceBeforeUpsertHooks = append(challengeServiceBeforeUpsertHooks, challengeServiceHook)
	case boil.AfterUpsertHook:
		challengeServiceAfterUpsertHooks = append(challengeServiceAfterUpsertHooks, challengeServiceHook)
	}
}

// One returns a single challengeService record from the query.
func (q challengeServiceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ChallengeService, error) {
	o := &ChallengeService{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for challenge_services")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ChallengeService records from the query.
func (q challengeServiceQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChallengeServiceSlice, error) {
	var o []*ChallengeService

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ChallengeService slice")
	}

	if len(challengeServiceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ChallengeService records in the query.
func (q challengeServiceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count challenge_services rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q challengeServiceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if challenge_services exists")
	}

	return count > 0, nil
}

// Challenge pointed to by the foreign key.
func (o *ChallengeService) Challenge(mods ...qm.QueryMod) challengeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChallengeID),
	}

	queryMods = append(queryMods, mods...)

	return Challenges(queryMods...)
}

// LoadChallenge allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (challengeServiceL) LoadChallenge(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChallengeService interface{}, mods queries.Applicator) error {
	var slice []*ChallengeService
	var object *ChallengeService

	if singular {
		var ok bool
		object, ok = maybeChallengeService.(*ChallengeService)
		if !ok {
			object = new(ChallengeService)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeChallengeService)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeChallengeService))
			}
		}
	} else {
		s, ok := maybeChallengeService.(*[]*ChallengeService)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeChallengeService)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeChallengeService))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &challengeServiceR{}
		}
		args = append(args, object.ChallengeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &challengeServiceR{}
			}

			for _, a := range args {
				if a == obj.ChallengeID {
					continue Outer
				}
			}

			args = append(args, obj.ChallengeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`challenges`),
		qm.WhereIn(`challenges.id in ?`, args...),
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
		foreign.R.ChallengeServices = append(foreign.R.ChallengeServices, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChallengeID == foreign.ID {
				local.R.Challenge = foreign
				if foreign.R == nil {
					foreign.R = &challengeR{}
				}
				foreign.R.ChallengeServices = append(foreign.R.ChallengeServices, local)
				break
			}
		}
	}

	return nil
}

// SetChallenge of the challengeService to the related item.
// Sets o.R.Challenge to related.
// Adds o to related.R.ChallengeServices.
func (o *ChallengeService) SetChallenge(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Challenge) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"challenge_services\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"challenge_id"}),
		strmangle.WhereClause("\"", "\"", 2, challengeServicePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

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
		o.R = &challengeServiceR{
			Challenge: related,
		}
	} else {
		o.R.Challenge = related
	}

	if related.R == nil {
		related.R = &challengeR{
			ChallengeServices: ChallengeServiceSlice{o},
		}
	} else {
		related.R.ChallengeServices = append(related.R.ChallengeServices, o)
	}

	return nil
}

// ChallengeServices retrieves all the records using an executor.
func ChallengeServices(mods ...qm.QueryMod) challengeServiceQuery {
	mods = append(mods, qm.From("\"challenge_services\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"challenge_services\".*"})
	}

	return challengeServiceQuery{q}
}

// FindChallengeService retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChallengeService(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ChallengeService, error) {
	challengeServiceObj := &ChallengeService{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"challenge_services\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, challengeServiceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from challenge_services")
	}

	if err = challengeServiceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return challengeServiceObj, err
	}

	return challengeServiceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ChallengeService) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no challenge_services provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(challengeServiceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	challengeServiceInsertCacheMut.RLock()
	cache, cached := challengeServiceInsertCache[key]
	challengeServiceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			challengeServiceAllColumns,
			challengeServiceColumnsWithDefault,
			challengeServiceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(challengeServiceType, challengeServiceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(challengeServiceType, challengeServiceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"challenge_services\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"challenge_services\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into challenge_services")
	}

	if !cached {
		challengeServiceInsertCacheMut.Lock()
		challengeServiceInsertCache[key] = cache
		challengeServiceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ChallengeService.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ChallengeService) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	challengeServiceUpdateCacheMut.RLock()
	cache, cached := challengeServiceUpdateCache[key]
	challengeServiceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			challengeServiceAllColumns,
			challengeServicePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update challenge_services, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"challenge_services\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, challengeServicePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(challengeServiceType, challengeServiceMapping, append(wl, challengeServicePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update challenge_services row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for challenge_services")
	}

	if !cached {
		challengeServiceUpdateCacheMut.Lock()
		challengeServiceUpdateCache[key] = cache
		challengeServiceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q challengeServiceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for challenge_services")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for challenge_services")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChallengeServiceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), challengeServicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"challenge_services\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, challengeServicePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in challengeService slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all challengeService")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ChallengeService) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no challenge_services provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(challengeServiceColumnsWithDefault, o)

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

	challengeServiceUpsertCacheMut.RLock()
	cache, cached := challengeServiceUpsertCache[key]
	challengeServiceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			challengeServiceAllColumns,
			challengeServiceColumnsWithDefault,
			challengeServiceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			challengeServiceAllColumns,
			challengeServicePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert challenge_services, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(challengeServicePrimaryKeyColumns))
			copy(conflict, challengeServicePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"challenge_services\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(challengeServiceType, challengeServiceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(challengeServiceType, challengeServiceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert challenge_services")
	}

	if !cached {
		challengeServiceUpsertCacheMut.Lock()
		challengeServiceUpsertCache[key] = cache
		challengeServiceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ChallengeService record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ChallengeService) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ChallengeService provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), challengeServicePrimaryKeyMapping)
	sql := "DELETE FROM \"challenge_services\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from challenge_services")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for challenge_services")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q challengeServiceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no challengeServiceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from challenge_services")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for challenge_services")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChallengeServiceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(challengeServiceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), challengeServicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"challenge_services\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, challengeServicePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from challengeService slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for challenge_services")
	}

	if len(challengeServiceAfterDeleteHooks) != 0 {
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
func (o *ChallengeService) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChallengeService(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChallengeServiceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChallengeServiceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), challengeServicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"challenge_services\".* FROM \"challenge_services\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, challengeServicePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChallengeServiceSlice")
	}

	*o = slice

	return nil
}

// ChallengeServiceExists checks if the ChallengeService row exists.
func ChallengeServiceExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"challenge_services\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if challenge_services exists")
	}

	return exists, nil
}

// Exists checks if the ChallengeService row exists.
func (o *ChallengeService) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ChallengeServiceExists(ctx, exec, o.ID)
}
