// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/match"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
	"github.com/open-boardgame-stats/backend/internal/ent/statistic"
)

// StatisticQuery is the builder for querying Statistic entities.
type StatisticQuery struct {
	config
	limit               *int
	offset              *int
	unique              *bool
	order               []OrderFunc
	fields              []string
	predicates          []predicate.Statistic
	withMatch           *MatchQuery
	withStatDescription *StatDescriptionQuery
	withPlayer          *PlayerQuery
	withFKs             bool
	modifiers           []func(*sql.Selector)
	loadTotal           []func(context.Context, []*Statistic) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StatisticQuery builder.
func (sq *StatisticQuery) Where(ps ...predicate.Statistic) *StatisticQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *StatisticQuery) Limit(limit int) *StatisticQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *StatisticQuery) Offset(offset int) *StatisticQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *StatisticQuery) Unique(unique bool) *StatisticQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *StatisticQuery) Order(o ...OrderFunc) *StatisticQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryMatch chains the current query on the "match" edge.
func (sq *StatisticQuery) QueryMatch() *MatchQuery {
	query := &MatchQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statistic.Table, statistic.FieldID, selector),
			sqlgraph.To(match.Table, match.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, statistic.MatchTable, statistic.MatchColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStatDescription chains the current query on the "stat_description" edge.
func (sq *StatisticQuery) QueryStatDescription() *StatDescriptionQuery {
	query := &StatDescriptionQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statistic.Table, statistic.FieldID, selector),
			sqlgraph.To(statdescription.Table, statdescription.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, statistic.StatDescriptionTable, statistic.StatDescriptionColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlayer chains the current query on the "player" edge.
func (sq *StatisticQuery) QueryPlayer() *PlayerQuery {
	query := &PlayerQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statistic.Table, statistic.FieldID, selector),
			sqlgraph.To(player.Table, player.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, statistic.PlayerTable, statistic.PlayerColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Statistic entity from the query.
// Returns a *NotFoundError when no Statistic was found.
func (sq *StatisticQuery) First(ctx context.Context) (*Statistic, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{statistic.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *StatisticQuery) FirstX(ctx context.Context) *Statistic {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Statistic ID from the query.
// Returns a *NotFoundError when no Statistic ID was found.
func (sq *StatisticQuery) FirstID(ctx context.Context) (id guidgql.GUID, err error) {
	var ids []guidgql.GUID
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{statistic.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *StatisticQuery) FirstIDX(ctx context.Context) guidgql.GUID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Statistic entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Statistic entity is found.
// Returns a *NotFoundError when no Statistic entities are found.
func (sq *StatisticQuery) Only(ctx context.Context) (*Statistic, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{statistic.Label}
	default:
		return nil, &NotSingularError{statistic.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *StatisticQuery) OnlyX(ctx context.Context) *Statistic {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Statistic ID in the query.
// Returns a *NotSingularError when more than one Statistic ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *StatisticQuery) OnlyID(ctx context.Context) (id guidgql.GUID, err error) {
	var ids []guidgql.GUID
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{statistic.Label}
	default:
		err = &NotSingularError{statistic.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *StatisticQuery) OnlyIDX(ctx context.Context) guidgql.GUID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Statistics.
func (sq *StatisticQuery) All(ctx context.Context) ([]*Statistic, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *StatisticQuery) AllX(ctx context.Context) []*Statistic {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Statistic IDs.
func (sq *StatisticQuery) IDs(ctx context.Context) ([]guidgql.GUID, error) {
	var ids []guidgql.GUID
	if err := sq.Select(statistic.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *StatisticQuery) IDsX(ctx context.Context) []guidgql.GUID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *StatisticQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *StatisticQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *StatisticQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *StatisticQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StatisticQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *StatisticQuery) Clone() *StatisticQuery {
	if sq == nil {
		return nil
	}
	return &StatisticQuery{
		config:              sq.config,
		limit:               sq.limit,
		offset:              sq.offset,
		order:               append([]OrderFunc{}, sq.order...),
		predicates:          append([]predicate.Statistic{}, sq.predicates...),
		withMatch:           sq.withMatch.Clone(),
		withStatDescription: sq.withStatDescription.Clone(),
		withPlayer:          sq.withPlayer.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithMatch tells the query-builder to eager-load the nodes that are connected to
// the "match" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StatisticQuery) WithMatch(opts ...func(*MatchQuery)) *StatisticQuery {
	query := &MatchQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withMatch = query
	return sq
}

// WithStatDescription tells the query-builder to eager-load the nodes that are connected to
// the "stat_description" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StatisticQuery) WithStatDescription(opts ...func(*StatDescriptionQuery)) *StatisticQuery {
	query := &StatDescriptionQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withStatDescription = query
	return sq
}

// WithPlayer tells the query-builder to eager-load the nodes that are connected to
// the "player" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StatisticQuery) WithPlayer(opts ...func(*PlayerQuery)) *StatisticQuery {
	query := &PlayerQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withPlayer = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Statistic.Query().
//		GroupBy(statistic.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *StatisticQuery) GroupBy(field string, fields ...string) *StatisticGroupBy {
	grbuild := &StatisticGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = statistic.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//	}
//
//	client.Statistic.Query().
//		Select(statistic.FieldValue).
//		Scan(ctx, &v)
func (sq *StatisticQuery) Select(fields ...string) *StatisticSelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &StatisticSelect{StatisticQuery: sq}
	selbuild.label = statistic.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a StatisticSelect configured with the given aggregations.
func (sq *StatisticQuery) Aggregate(fns ...AggregateFunc) *StatisticSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *StatisticQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !statistic.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *StatisticQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Statistic, error) {
	var (
		nodes       = []*Statistic{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [3]bool{
			sq.withMatch != nil,
			sq.withStatDescription != nil,
			sq.withPlayer != nil,
		}
	)
	if sq.withMatch != nil || sq.withStatDescription != nil || sq.withPlayer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, statistic.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Statistic).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Statistic{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withMatch; query != nil {
		if err := sq.loadMatch(ctx, query, nodes, nil,
			func(n *Statistic, e *Match) { n.Edges.Match = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withStatDescription; query != nil {
		if err := sq.loadStatDescription(ctx, query, nodes, nil,
			func(n *Statistic, e *StatDescription) { n.Edges.StatDescription = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withPlayer; query != nil {
		if err := sq.loadPlayer(ctx, query, nodes, nil,
			func(n *Statistic, e *Player) { n.Edges.Player = e }); err != nil {
			return nil, err
		}
	}
	for i := range sq.loadTotal {
		if err := sq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *StatisticQuery) loadMatch(ctx context.Context, query *MatchQuery, nodes []*Statistic, init func(*Statistic), assign func(*Statistic, *Match)) error {
	ids := make([]guidgql.GUID, 0, len(nodes))
	nodeids := make(map[guidgql.GUID][]*Statistic)
	for i := range nodes {
		if nodes[i].match_stats == nil {
			continue
		}
		fk := *nodes[i].match_stats
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(match.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "match_stats" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *StatisticQuery) loadStatDescription(ctx context.Context, query *StatDescriptionQuery, nodes []*Statistic, init func(*Statistic), assign func(*Statistic, *StatDescription)) error {
	ids := make([]guidgql.GUID, 0, len(nodes))
	nodeids := make(map[guidgql.GUID][]*Statistic)
	for i := range nodes {
		if nodes[i].stat_description_stats == nil {
			continue
		}
		fk := *nodes[i].stat_description_stats
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(statdescription.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "stat_description_stats" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *StatisticQuery) loadPlayer(ctx context.Context, query *PlayerQuery, nodes []*Statistic, init func(*Statistic), assign func(*Statistic, *Player)) error {
	ids := make([]guidgql.GUID, 0, len(nodes))
	nodeids := make(map[guidgql.GUID][]*Statistic)
	for i := range nodes {
		if nodes[i].player_stats == nil {
			continue
		}
		fk := *nodes[i].player_stats
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(player.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "player_stats" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *StatisticQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *StatisticQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sq *StatisticQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statistic.Table,
			Columns: statistic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: statistic.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, statistic.FieldID)
		for i := range fields {
			if fields[i] != statistic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *StatisticQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(statistic.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = statistic.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StatisticGroupBy is the group-by builder for Statistic entities.
type StatisticGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *StatisticGroupBy) Aggregate(fns ...AggregateFunc) *StatisticGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *StatisticGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *StatisticGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sgb.fields {
		if !statistic.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *StatisticGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// StatisticSelect is the builder for selecting fields of Statistic entities.
type StatisticSelect struct {
	*StatisticQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *StatisticSelect) Aggregate(fns ...AggregateFunc) *StatisticSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *StatisticSelect) Scan(ctx context.Context, v any) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.StatisticQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *StatisticSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(ss.sql))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ss.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ss.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}