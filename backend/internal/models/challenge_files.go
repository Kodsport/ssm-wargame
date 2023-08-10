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

// ChallengeFile is an object representing the database table.
type ChallengeFile struct {
	ID           string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ChallengeID  null.String `boil:"challenge_id" json:"challenge_id,omitempty" toml:"challenge_id" yaml:"challenge_id,omitempty"`
	FriendlyName string      `boil:"friendly_name" json:"friendly_name" toml:"friendly_name" yaml:"friendly_name"`
	URL          string      `boil:"url" json:"url" toml:"url" yaml:"url"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *challengeFileR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L challengeFileL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChallengeFileColumns = struct {
	ID           string
	ChallengeID  string
	FriendlyName string
	URL          string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	ChallengeID:  "challenge_id",
	FriendlyName: "friendly_name",
	URL:          "url",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

var ChallengeFileTableColumns = struct {
	ID           string
	ChallengeID  string
	FriendlyName string
	URL          string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "challenge_files.id",
	ChallengeID:  "challenge_files.challenge_id",
	FriendlyName: "challenge_files.friendly_name",
	URL:          "challenge_files.url",
	CreatedAt:    "challenge_files.created_at",
	UpdatedAt:    "challenge_files.updated_at",
}

// Generated where

var ChallengeFileWhere = struct {
	ID           whereHelperstring
	ChallengeID  whereHelpernull_String
	FriendlyName whereHelperstring
	URL          whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
}{
	ID:           whereHelperstring{field: "\"challenge_files\".\"id\""},
	ChallengeID:  whereHelpernull_String{field: "\"challenge_files\".\"challenge_id\""},
	FriendlyName: whereHelperstring{field: "\"challenge_files\".\"friendly_name\""},
	URL:          whereHelperstring{field: "\"challenge_files\".\"url\""},
	CreatedAt:    whereHelpertime_Time{field: "\"challenge_files\".\"created_at\""},
	UpdatedAt:    whereHelpernull_Time{field: "\"challenge_files\".\"updated_at\""},
}

// ChallengeFileRels is where relationship names are stored.
var ChallengeFileRels = struct {
	Challenge string
}{
	Challenge: "Challenge",
}

// challengeFileR is where relationships are stored.
type challengeFileR struct {
	Challenge *Challenge `boil:"Challenge" json:"Challenge" toml:"Challenge" yaml:"Challenge"`
}

// NewStruct creates a new relationship struct
func (*challengeFileR) NewStruct() *challengeFileR {
	return &challengeFileR{}
}

func (r *challengeFileR) GetChallenge() *Challenge {
	if r == nil {
		return nil
	}
	return r.Challenge
}

// challengeFileL is where Load methods for each relationship are stored.
type challengeFileL struct{}

var (
	challengeFileAllColumns            = []string{"id", "challenge_id", "friendly_name", "url", "created_at", "updated_at"}
	challengeFileColumnsWithoutDefault = []string{"id", "friendly_name", "url"}
	challengeFileColumnsWithDefault    = []string{"challenge_id", "created_at", "updated_at"}
	challengeFilePrimaryKeyColumns     = []string{"id"}
	challengeFileGeneratedColumns      = []string{}
)

type (
	// ChallengeFileSlice is an alias for a slice of pointers to ChallengeFile.
	// This should almost always be used instead of []ChallengeFile.
	ChallengeFileSlice []*ChallengeFile
	// ChallengeFileHook is the signature for custom ChallengeFile hook methods
	ChallengeFileHook func(context.Context, boil.ContextExecutor, *ChallengeFile) error

	challengeFileQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	challengeFileType                 = reflect.TypeOf(&ChallengeFile{})
	challengeFileMapping              = queries.MakeStructMapping(challengeFileType)
	challengeFilePrimaryKeyMapping, _ = queries.BindMapping(challengeFileType, challengeFileMapping, challengeFilePrimaryKeyColumns)
	challengeFileInsertCacheMut       sync.RWMutex
	challengeFileInsertCache          = make(map[string]insertCache)
	challengeFileUpdateCacheMut       sync.RWMutex
	challengeFileUpdateCache          = make(map[string]updateCache)
	challengeFileUpsertCacheMut       sync.RWMutex
	challengeFileUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var challengeFileAfterSelectHooks []ChallengeFileHook

var challengeFileBeforeInsertHooks []ChallengeFileHook
var challengeFileAfterInsertHooks []ChallengeFileHook

var challengeFileBeforeUpdateHooks []ChallengeFileHook
var challengeFileAfterUpdateHooks []ChallengeFileHook

var challengeFileBeforeDeleteHooks []ChallengeFileHook
var challengeFileAfterDeleteHooks []ChallengeFileHook

var challengeFileBeforeUpsertHooks []ChallengeFileHook
var challengeFileAfterUpsertHooks []ChallengeFileHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ChallengeFile) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeFileAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ChallengeFile) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeFileBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ChallengeFile) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeFileAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ChallengeFile) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeFileBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ChallengeFile) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeFileAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ChallengeFile) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeFileBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ChallengeFile) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeFileAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ChallengeFile) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeFileBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ChallengeFile) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range challengeFileAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChallengeFileHook registers your hook function for all future operations.
func AddChallengeFileHook(hookPoint boil.HookPoint, challengeFileHook ChallengeFileHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		challengeFileAfterSelectHooks = append(challengeFileAfterSelectHooks, challengeFileHook)
	case boil.BeforeInsertHook:
		challengeFileBeforeInsertHooks = append(challengeFileBeforeInsertHooks, challengeFileHook)
	case boil.AfterInsertHook:
		challengeFileAfterInsertHooks = append(challengeFileAfterInsertHooks, challengeFileHook)
	case boil.BeforeUpdateHook:
		challengeFileBeforeUpdateHooks = append(challengeFileBeforeUpdateHooks, challengeFileHook)
	case boil.AfterUpdateHook:
		challengeFileAfterUpdateHooks = append(challengeFileAfterUpdateHooks, challengeFileHook)
	case boil.BeforeDeleteHook:
		challengeFileBeforeDeleteHooks = append(challengeFileBeforeDeleteHooks, challengeFileHook)
	case boil.AfterDeleteHook:
		challengeFileAfterDeleteHooks = append(challengeFileAfterDeleteHooks, challengeFileHook)
	case boil.BeforeUpsertHook:
		challengeFileBeforeUpsertHooks = append(challengeFileBeforeUpsertHooks, challengeFileHook)
	case boil.AfterUpsertHook:
		challengeFileAfterUpsertHooks = append(challengeFileAfterUpsertHooks, challengeFileHook)
	}
}

// One returns a single challengeFile record from the query.
func (q challengeFileQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ChallengeFile, error) {
	o := &ChallengeFile{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for challenge_files")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ChallengeFile records from the query.
func (q challengeFileQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChallengeFileSlice, error) {
	var o []*ChallengeFile

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ChallengeFile slice")
	}

	if len(challengeFileAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ChallengeFile records in the query.
func (q challengeFileQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count challenge_files rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q challengeFileQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if challenge_files exists")
	}

	return count > 0, nil
}

// Challenge pointed to by the foreign key.
func (o *ChallengeFile) Challenge(mods ...qm.QueryMod) challengeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ChallengeID),
	}

	queryMods = append(queryMods, mods...)

	return Challenges(queryMods...)
}

// LoadChallenge allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (challengeFileL) LoadChallenge(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChallengeFile interface{}, mods queries.Applicator) error {
	var slice []*ChallengeFile
	var object *ChallengeFile

	if singular {
		var ok bool
		object, ok = maybeChallengeFile.(*ChallengeFile)
		if !ok {
			object = new(ChallengeFile)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeChallengeFile)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeChallengeFile))
			}
		}
	} else {
		s, ok := maybeChallengeFile.(*[]*ChallengeFile)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeChallengeFile)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeChallengeFile))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &challengeFileR{}
		}
		if !queries.IsNil(object.ChallengeID) {
			args = append(args, object.ChallengeID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &challengeFileR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ChallengeID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ChallengeID) {
				args = append(args, obj.ChallengeID)
			}

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
		foreign.R.ChallengeFiles = append(foreign.R.ChallengeFiles, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ChallengeID, foreign.ID) {
				local.R.Challenge = foreign
				if foreign.R == nil {
					foreign.R = &challengeR{}
				}
				foreign.R.ChallengeFiles = append(foreign.R.ChallengeFiles, local)
				break
			}
		}
	}

	return nil
}

// SetChallenge of the challengeFile to the related item.
// Sets o.R.Challenge to related.
// Adds o to related.R.ChallengeFiles.
func (o *ChallengeFile) SetChallenge(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Challenge) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"challenge_files\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"challenge_id"}),
		strmangle.WhereClause("\"", "\"", 2, challengeFilePrimaryKeyColumns),
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

	queries.Assign(&o.ChallengeID, related.ID)
	if o.R == nil {
		o.R = &challengeFileR{
			Challenge: related,
		}
	} else {
		o.R.Challenge = related
	}

	if related.R == nil {
		related.R = &challengeR{
			ChallengeFiles: ChallengeFileSlice{o},
		}
	} else {
		related.R.ChallengeFiles = append(related.R.ChallengeFiles, o)
	}

	return nil
}

// RemoveChallenge relationship.
// Sets o.R.Challenge to nil.
// Removes o from all passed in related items' relationships struct.
func (o *ChallengeFile) RemoveChallenge(ctx context.Context, exec boil.ContextExecutor, related *Challenge) error {
	var err error

	queries.SetScanner(&o.ChallengeID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("challenge_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Challenge = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ChallengeFiles {
		if queries.Equal(o.ChallengeID, ri.ChallengeID) {
			continue
		}

		ln := len(related.R.ChallengeFiles)
		if ln > 1 && i < ln-1 {
			related.R.ChallengeFiles[i] = related.R.ChallengeFiles[ln-1]
		}
		related.R.ChallengeFiles = related.R.ChallengeFiles[:ln-1]
		break
	}
	return nil
}

// ChallengeFiles retrieves all the records using an executor.
func ChallengeFiles(mods ...qm.QueryMod) challengeFileQuery {
	mods = append(mods, qm.From("\"challenge_files\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"challenge_files\".*"})
	}

	return challengeFileQuery{q}
}

// FindChallengeFile retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChallengeFile(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ChallengeFile, error) {
	challengeFileObj := &ChallengeFile{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"challenge_files\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, challengeFileObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from challenge_files")
	}

	if err = challengeFileObj.doAfterSelectHooks(ctx, exec); err != nil {
		return challengeFileObj, err
	}

	return challengeFileObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ChallengeFile) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no challenge_files provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(challengeFileColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	challengeFileInsertCacheMut.RLock()
	cache, cached := challengeFileInsertCache[key]
	challengeFileInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			challengeFileAllColumns,
			challengeFileColumnsWithDefault,
			challengeFileColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(challengeFileType, challengeFileMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(challengeFileType, challengeFileMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"challenge_files\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"challenge_files\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into challenge_files")
	}

	if !cached {
		challengeFileInsertCacheMut.Lock()
		challengeFileInsertCache[key] = cache
		challengeFileInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ChallengeFile.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ChallengeFile) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	challengeFileUpdateCacheMut.RLock()
	cache, cached := challengeFileUpdateCache[key]
	challengeFileUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			challengeFileAllColumns,
			challengeFilePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update challenge_files, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"challenge_files\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, challengeFilePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(challengeFileType, challengeFileMapping, append(wl, challengeFilePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update challenge_files row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for challenge_files")
	}

	if !cached {
		challengeFileUpdateCacheMut.Lock()
		challengeFileUpdateCache[key] = cache
		challengeFileUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q challengeFileQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for challenge_files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for challenge_files")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChallengeFileSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), challengeFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"challenge_files\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, challengeFilePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in challengeFile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all challengeFile")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ChallengeFile) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no challenge_files provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(challengeFileColumnsWithDefault, o)

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

	challengeFileUpsertCacheMut.RLock()
	cache, cached := challengeFileUpsertCache[key]
	challengeFileUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			challengeFileAllColumns,
			challengeFileColumnsWithDefault,
			challengeFileColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			challengeFileAllColumns,
			challengeFilePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert challenge_files, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(challengeFilePrimaryKeyColumns))
			copy(conflict, challengeFilePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"challenge_files\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(challengeFileType, challengeFileMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(challengeFileType, challengeFileMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert challenge_files")
	}

	if !cached {
		challengeFileUpsertCacheMut.Lock()
		challengeFileUpsertCache[key] = cache
		challengeFileUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ChallengeFile record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ChallengeFile) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ChallengeFile provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), challengeFilePrimaryKeyMapping)
	sql := "DELETE FROM \"challenge_files\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from challenge_files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for challenge_files")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q challengeFileQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no challengeFileQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from challenge_files")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for challenge_files")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChallengeFileSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(challengeFileBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), challengeFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"challenge_files\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, challengeFilePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from challengeFile slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for challenge_files")
	}

	if len(challengeFileAfterDeleteHooks) != 0 {
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
func (o *ChallengeFile) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChallengeFile(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChallengeFileSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChallengeFileSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), challengeFilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"challenge_files\".* FROM \"challenge_files\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, challengeFilePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChallengeFileSlice")
	}

	*o = slice

	return nil
}

// ChallengeFileExists checks if the ChallengeFile row exists.
func ChallengeFileExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"challenge_files\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if challenge_files exists")
	}

	return exists, nil
}

// Exists checks if the ChallengeFile row exists.
func (o *ChallengeFile) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ChallengeFileExists(ctx, exec, o.ID)
}
