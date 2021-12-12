// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Project is an object representing the database table.
type Project struct {
	ID          int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Department  null.String `boil:"department" json:"department,omitempty" toml:"department" yaml:"department,omitempty"`
	Admin       null.String `boil:"admin" json:"admin,omitempty" toml:"admin" yaml:"admin,omitempty"`
	ProjectUser string      `boil:"project_user" json:"project_user" toml:"project_user" yaml:"project_user"`
	DevelopUser null.String `boil:"develop_user" json:"develop_user,omitempty" toml:"develop_user" yaml:"develop_user,omitempty"`
	CreateTime  time.Time   `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime  time.Time   `boil:"update_time" json:"update_time" toml:"update_time" yaml:"update_time"`

	R *projectR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L projectL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProjectColumns = struct {
	ID          string
	Name        string
	Description string
	Department  string
	Admin       string
	ProjectUser string
	DevelopUser string
	CreateTime  string
	UpdateTime  string
}{
	ID:          "id",
	Name:        "name",
	Description: "description",
	Department:  "department",
	Admin:       "admin",
	ProjectUser: "project_user",
	DevelopUser: "develop_user",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

var ProjectTableColumns = struct {
	ID          string
	Name        string
	Description string
	Department  string
	Admin       string
	ProjectUser string
	DevelopUser string
	CreateTime  string
	UpdateTime  string
}{
	ID:          "project.id",
	Name:        "project.name",
	Description: "project.description",
	Department:  "project.department",
	Admin:       "project.admin",
	ProjectUser: "project.project_user",
	DevelopUser: "project.develop_user",
	CreateTime:  "project.create_time",
	UpdateTime:  "project.update_time",
}

// Generated where

var ProjectWhere = struct {
	ID          whereHelperint
	Name        whereHelperstring
	Description whereHelpernull_String
	Department  whereHelpernull_String
	Admin       whereHelpernull_String
	ProjectUser whereHelperstring
	DevelopUser whereHelpernull_String
	CreateTime  whereHelpertime_Time
	UpdateTime  whereHelpertime_Time
}{
	ID:          whereHelperint{field: "`project`.`id`"},
	Name:        whereHelperstring{field: "`project`.`name`"},
	Description: whereHelpernull_String{field: "`project`.`description`"},
	Department:  whereHelpernull_String{field: "`project`.`department`"},
	Admin:       whereHelpernull_String{field: "`project`.`admin`"},
	ProjectUser: whereHelperstring{field: "`project`.`project_user`"},
	DevelopUser: whereHelpernull_String{field: "`project`.`develop_user`"},
	CreateTime:  whereHelpertime_Time{field: "`project`.`create_time`"},
	UpdateTime:  whereHelpertime_Time{field: "`project`.`update_time`"},
}

// ProjectRels is where relationship names are stored.
var ProjectRels = struct {
}{}

// projectR is where relationships are stored.
type projectR struct {
}

// NewStruct creates a new relationship struct
func (*projectR) NewStruct() *projectR {
	return &projectR{}
}

// projectL is where Load methods for each relationship are stored.
type projectL struct{}

var (
	projectAllColumns            = []string{"id", "name", "description", "department", "admin", "project_user", "develop_user", "create_time", "update_time"}
	projectColumnsWithoutDefault = []string{"name", "description", "department", "admin", "project_user", "develop_user"}
	projectColumnsWithDefault    = []string{"id", "create_time", "update_time"}
	projectPrimaryKeyColumns     = []string{"id"}
)

type (
	// ProjectSlice is an alias for a slice of pointers to Project.
	// This should almost always be used instead of []Project.
	ProjectSlice []*Project
	// ProjectHook is the signature for custom Project hook methods
	ProjectHook func(context.Context, boil.ContextExecutor, *Project) error

	projectQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	projectType                 = reflect.TypeOf(&Project{})
	projectMapping              = queries.MakeStructMapping(projectType)
	projectPrimaryKeyMapping, _ = queries.BindMapping(projectType, projectMapping, projectPrimaryKeyColumns)
	projectInsertCacheMut       sync.RWMutex
	projectInsertCache          = make(map[string]insertCache)
	projectUpdateCacheMut       sync.RWMutex
	projectUpdateCache          = make(map[string]updateCache)
	projectUpsertCacheMut       sync.RWMutex
	projectUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var projectBeforeInsertHooks []ProjectHook
var projectBeforeUpdateHooks []ProjectHook
var projectBeforeDeleteHooks []ProjectHook
var projectBeforeUpsertHooks []ProjectHook

var projectAfterInsertHooks []ProjectHook
var projectAfterSelectHooks []ProjectHook
var projectAfterUpdateHooks []ProjectHook
var projectAfterDeleteHooks []ProjectHook
var projectAfterUpsertHooks []ProjectHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Project) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Project) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Project) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Project) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Project) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Project) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Project) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Project) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Project) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProjectHook registers your hook function for all future operations.
func AddProjectHook(hookPoint boil.HookPoint, projectHook ProjectHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		projectBeforeInsertHooks = append(projectBeforeInsertHooks, projectHook)
	case boil.BeforeUpdateHook:
		projectBeforeUpdateHooks = append(projectBeforeUpdateHooks, projectHook)
	case boil.BeforeDeleteHook:
		projectBeforeDeleteHooks = append(projectBeforeDeleteHooks, projectHook)
	case boil.BeforeUpsertHook:
		projectBeforeUpsertHooks = append(projectBeforeUpsertHooks, projectHook)
	case boil.AfterInsertHook:
		projectAfterInsertHooks = append(projectAfterInsertHooks, projectHook)
	case boil.AfterSelectHook:
		projectAfterSelectHooks = append(projectAfterSelectHooks, projectHook)
	case boil.AfterUpdateHook:
		projectAfterUpdateHooks = append(projectAfterUpdateHooks, projectHook)
	case boil.AfterDeleteHook:
		projectAfterDeleteHooks = append(projectAfterDeleteHooks, projectHook)
	case boil.AfterUpsertHook:
		projectAfterUpsertHooks = append(projectAfterUpsertHooks, projectHook)
	}
}

// One returns a single project record from the query.
func (q projectQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Project, error) {
	o := &Project{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for project")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Project records from the query.
func (q projectQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProjectSlice, error) {
	var o []*Project

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Project slice")
	}

	if len(projectAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Project records in the query.
func (q projectQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count project rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q projectQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if project exists")
	}

	return count > 0, nil
}

// Projects retrieves all the records using an executor.
func Projects(mods ...qm.QueryMod) projectQuery {
	mods = append(mods, qm.From("`project`"))
	return projectQuery{NewQuery(mods...)}
}

// FindProject retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProject(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Project, error) {
	projectObj := &Project{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `project` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, projectObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from project")
	}

	if err = projectObj.doAfterSelectHooks(ctx, exec); err != nil {
		return projectObj, err
	}

	return projectObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Project) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no project provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	projectInsertCacheMut.RLock()
	cache, cached := projectInsertCache[key]
	projectInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			projectAllColumns,
			projectColumnsWithDefault,
			projectColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(projectType, projectMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(projectType, projectMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `project` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `project` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `project` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, projectPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into project")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == projectMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for project")
	}

CacheNoHooks:
	if !cached {
		projectInsertCacheMut.Lock()
		projectInsertCache[key] = cache
		projectInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Project.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Project) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	projectUpdateCacheMut.RLock()
	cache, cached := projectUpdateCache[key]
	projectUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			projectAllColumns,
			projectPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update project, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `project` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, projectPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(projectType, projectMapping, append(wl, projectPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update project row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for project")
	}

	if !cached {
		projectUpdateCacheMut.Lock()
		projectUpdateCache[key] = cache
		projectUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q projectQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for project")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for project")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProjectSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `project` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, projectPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in project slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all project")
	}
	return rowsAff, nil
}

var mySQLProjectUniqueColumns = []string{
	"id",
	"name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Project) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no project provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLProjectUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	projectUpsertCacheMut.RLock()
	cache, cached := projectUpsertCache[key]
	projectUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			projectAllColumns,
			projectColumnsWithDefault,
			projectColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			projectAllColumns,
			projectPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert project, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`project`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `project` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(projectType, projectMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(projectType, projectMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for project")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == projectMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(projectType, projectMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for project")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for project")
	}

CacheNoHooks:
	if !cached {
		projectUpsertCacheMut.Lock()
		projectUpsertCache[key] = cache
		projectUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Project record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Project) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Project provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), projectPrimaryKeyMapping)
	sql := "DELETE FROM `project` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from project")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for project")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q projectQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no projectQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from project")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for project")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProjectSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(projectBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `project` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, projectPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from project slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for project")
	}

	if len(projectAfterDeleteHooks) != 0 {
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
func (o *Project) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProject(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProjectSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProjectSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `project`.* FROM `project` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, projectPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProjectSlice")
	}

	*o = slice

	return nil
}

// ProjectExists checks if the Project row exists.
func ProjectExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `project` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if project exists")
	}

	return exists, nil
}
