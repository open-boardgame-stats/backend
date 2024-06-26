// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequestapproval"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// PlayerSupervisionRequestApprovalQuery is the builder for querying PlayerSupervisionRequestApproval entities.
type PlayerSupervisionRequestApprovalQuery struct {
	config
	ctx                    *QueryContext
	order                  []playersupervisionrequestapproval.OrderOption
	inters                 []Interceptor
	predicates             []predicate.PlayerSupervisionRequestApproval
	withApprover           *UserQuery
	withSupervisionRequest *PlayerSupervisionRequestQuery
	withFKs                bool
	modifiers              []func(*sql.Selector)
	loadTotal              []func(context.Context, []*PlayerSupervisionRequestApproval) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlayerSupervisionRequestApprovalQuery builder.
func (psraq *PlayerSupervisionRequestApprovalQuery) Where(ps ...predicate.PlayerSupervisionRequestApproval) *PlayerSupervisionRequestApprovalQuery {
	psraq.predicates = append(psraq.predicates, ps...)
	return psraq
}

// Limit the number of records to be returned by this query.
func (psraq *PlayerSupervisionRequestApprovalQuery) Limit(limit int) *PlayerSupervisionRequestApprovalQuery {
	psraq.ctx.Limit = &limit
	return psraq
}

// Offset to start from.
func (psraq *PlayerSupervisionRequestApprovalQuery) Offset(offset int) *PlayerSupervisionRequestApprovalQuery {
	psraq.ctx.Offset = &offset
	return psraq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psraq *PlayerSupervisionRequestApprovalQuery) Unique(unique bool) *PlayerSupervisionRequestApprovalQuery {
	psraq.ctx.Unique = &unique
	return psraq
}

// Order specifies how the records should be ordered.
func (psraq *PlayerSupervisionRequestApprovalQuery) Order(o ...playersupervisionrequestapproval.OrderOption) *PlayerSupervisionRequestApprovalQuery {
	psraq.order = append(psraq.order, o...)
	return psraq
}

// QueryApprover chains the current query on the "approver" edge.
func (psraq *PlayerSupervisionRequestApprovalQuery) QueryApprover() *UserQuery {
	query := (&UserClient{config: psraq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psraq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psraq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playersupervisionrequestapproval.Table, playersupervisionrequestapproval.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, playersupervisionrequestapproval.ApproverTable, playersupervisionrequestapproval.ApproverColumn),
		)
		fromU = sqlgraph.SetNeighbors(psraq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySupervisionRequest chains the current query on the "supervision_request" edge.
func (psraq *PlayerSupervisionRequestApprovalQuery) QuerySupervisionRequest() *PlayerSupervisionRequestQuery {
	query := (&PlayerSupervisionRequestClient{config: psraq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psraq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psraq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playersupervisionrequestapproval.Table, playersupervisionrequestapproval.FieldID, selector),
			sqlgraph.To(playersupervisionrequest.Table, playersupervisionrequest.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, playersupervisionrequestapproval.SupervisionRequestTable, playersupervisionrequestapproval.SupervisionRequestColumn),
		)
		fromU = sqlgraph.SetNeighbors(psraq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PlayerSupervisionRequestApproval entity from the query.
// Returns a *NotFoundError when no PlayerSupervisionRequestApproval was found.
func (psraq *PlayerSupervisionRequestApprovalQuery) First(ctx context.Context) (*PlayerSupervisionRequestApproval, error) {
	nodes, err := psraq.Limit(1).All(setContextOp(ctx, psraq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{playersupervisionrequestapproval.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psraq *PlayerSupervisionRequestApprovalQuery) FirstX(ctx context.Context) *PlayerSupervisionRequestApproval {
	node, err := psraq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PlayerSupervisionRequestApproval ID from the query.
// Returns a *NotFoundError when no PlayerSupervisionRequestApproval ID was found.
func (psraq *PlayerSupervisionRequestApprovalQuery) FirstID(ctx context.Context) (id guidgql.GUID, err error) {
	var ids []guidgql.GUID
	if ids, err = psraq.Limit(1).IDs(setContextOp(ctx, psraq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{playersupervisionrequestapproval.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psraq *PlayerSupervisionRequestApprovalQuery) FirstIDX(ctx context.Context) guidgql.GUID {
	id, err := psraq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PlayerSupervisionRequestApproval entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PlayerSupervisionRequestApproval entity is found.
// Returns a *NotFoundError when no PlayerSupervisionRequestApproval entities are found.
func (psraq *PlayerSupervisionRequestApprovalQuery) Only(ctx context.Context) (*PlayerSupervisionRequestApproval, error) {
	nodes, err := psraq.Limit(2).All(setContextOp(ctx, psraq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{playersupervisionrequestapproval.Label}
	default:
		return nil, &NotSingularError{playersupervisionrequestapproval.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psraq *PlayerSupervisionRequestApprovalQuery) OnlyX(ctx context.Context) *PlayerSupervisionRequestApproval {
	node, err := psraq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PlayerSupervisionRequestApproval ID in the query.
// Returns a *NotSingularError when more than one PlayerSupervisionRequestApproval ID is found.
// Returns a *NotFoundError when no entities are found.
func (psraq *PlayerSupervisionRequestApprovalQuery) OnlyID(ctx context.Context) (id guidgql.GUID, err error) {
	var ids []guidgql.GUID
	if ids, err = psraq.Limit(2).IDs(setContextOp(ctx, psraq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{playersupervisionrequestapproval.Label}
	default:
		err = &NotSingularError{playersupervisionrequestapproval.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psraq *PlayerSupervisionRequestApprovalQuery) OnlyIDX(ctx context.Context) guidgql.GUID {
	id, err := psraq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PlayerSupervisionRequestApprovals.
func (psraq *PlayerSupervisionRequestApprovalQuery) All(ctx context.Context) ([]*PlayerSupervisionRequestApproval, error) {
	ctx = setContextOp(ctx, psraq.ctx, "All")
	if err := psraq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PlayerSupervisionRequestApproval, *PlayerSupervisionRequestApprovalQuery]()
	return withInterceptors[[]*PlayerSupervisionRequestApproval](ctx, psraq, qr, psraq.inters)
}

// AllX is like All, but panics if an error occurs.
func (psraq *PlayerSupervisionRequestApprovalQuery) AllX(ctx context.Context) []*PlayerSupervisionRequestApproval {
	nodes, err := psraq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PlayerSupervisionRequestApproval IDs.
func (psraq *PlayerSupervisionRequestApprovalQuery) IDs(ctx context.Context) (ids []guidgql.GUID, err error) {
	if psraq.ctx.Unique == nil && psraq.path != nil {
		psraq.Unique(true)
	}
	ctx = setContextOp(ctx, psraq.ctx, "IDs")
	if err = psraq.Select(playersupervisionrequestapproval.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psraq *PlayerSupervisionRequestApprovalQuery) IDsX(ctx context.Context) []guidgql.GUID {
	ids, err := psraq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psraq *PlayerSupervisionRequestApprovalQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, psraq.ctx, "Count")
	if err := psraq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, psraq, querierCount[*PlayerSupervisionRequestApprovalQuery](), psraq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (psraq *PlayerSupervisionRequestApprovalQuery) CountX(ctx context.Context) int {
	count, err := psraq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psraq *PlayerSupervisionRequestApprovalQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, psraq.ctx, "Exist")
	switch _, err := psraq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (psraq *PlayerSupervisionRequestApprovalQuery) ExistX(ctx context.Context) bool {
	exist, err := psraq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlayerSupervisionRequestApprovalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psraq *PlayerSupervisionRequestApprovalQuery) Clone() *PlayerSupervisionRequestApprovalQuery {
	if psraq == nil {
		return nil
	}
	return &PlayerSupervisionRequestApprovalQuery{
		config:                 psraq.config,
		ctx:                    psraq.ctx.Clone(),
		order:                  append([]playersupervisionrequestapproval.OrderOption{}, psraq.order...),
		inters:                 append([]Interceptor{}, psraq.inters...),
		predicates:             append([]predicate.PlayerSupervisionRequestApproval{}, psraq.predicates...),
		withApprover:           psraq.withApprover.Clone(),
		withSupervisionRequest: psraq.withSupervisionRequest.Clone(),
		// clone intermediate query.
		sql:  psraq.sql.Clone(),
		path: psraq.path,
	}
}

// WithApprover tells the query-builder to eager-load the nodes that are connected to
// the "approver" edge. The optional arguments are used to configure the query builder of the edge.
func (psraq *PlayerSupervisionRequestApprovalQuery) WithApprover(opts ...func(*UserQuery)) *PlayerSupervisionRequestApprovalQuery {
	query := (&UserClient{config: psraq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psraq.withApprover = query
	return psraq
}

// WithSupervisionRequest tells the query-builder to eager-load the nodes that are connected to
// the "supervision_request" edge. The optional arguments are used to configure the query builder of the edge.
func (psraq *PlayerSupervisionRequestApprovalQuery) WithSupervisionRequest(opts ...func(*PlayerSupervisionRequestQuery)) *PlayerSupervisionRequestApprovalQuery {
	query := (&PlayerSupervisionRequestClient{config: psraq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psraq.withSupervisionRequest = query
	return psraq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Approved bool `json:"approved,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PlayerSupervisionRequestApproval.Query().
//		GroupBy(playersupervisionrequestapproval.FieldApproved).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (psraq *PlayerSupervisionRequestApprovalQuery) GroupBy(field string, fields ...string) *PlayerSupervisionRequestApprovalGroupBy {
	psraq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PlayerSupervisionRequestApprovalGroupBy{build: psraq}
	grbuild.flds = &psraq.ctx.Fields
	grbuild.label = playersupervisionrequestapproval.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Approved bool `json:"approved,omitempty"`
//	}
//
//	client.PlayerSupervisionRequestApproval.Query().
//		Select(playersupervisionrequestapproval.FieldApproved).
//		Scan(ctx, &v)
func (psraq *PlayerSupervisionRequestApprovalQuery) Select(fields ...string) *PlayerSupervisionRequestApprovalSelect {
	psraq.ctx.Fields = append(psraq.ctx.Fields, fields...)
	sbuild := &PlayerSupervisionRequestApprovalSelect{PlayerSupervisionRequestApprovalQuery: psraq}
	sbuild.label = playersupervisionrequestapproval.Label
	sbuild.flds, sbuild.scan = &psraq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PlayerSupervisionRequestApprovalSelect configured with the given aggregations.
func (psraq *PlayerSupervisionRequestApprovalQuery) Aggregate(fns ...AggregateFunc) *PlayerSupervisionRequestApprovalSelect {
	return psraq.Select().Aggregate(fns...)
}

func (psraq *PlayerSupervisionRequestApprovalQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range psraq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, psraq); err != nil {
				return err
			}
		}
	}
	for _, f := range psraq.ctx.Fields {
		if !playersupervisionrequestapproval.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if psraq.path != nil {
		prev, err := psraq.path(ctx)
		if err != nil {
			return err
		}
		psraq.sql = prev
	}
	return nil
}

func (psraq *PlayerSupervisionRequestApprovalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PlayerSupervisionRequestApproval, error) {
	var (
		nodes       = []*PlayerSupervisionRequestApproval{}
		withFKs     = psraq.withFKs
		_spec       = psraq.querySpec()
		loadedTypes = [2]bool{
			psraq.withApprover != nil,
			psraq.withSupervisionRequest != nil,
		}
	)
	if psraq.withApprover != nil || psraq.withSupervisionRequest != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, playersupervisionrequestapproval.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PlayerSupervisionRequestApproval).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PlayerSupervisionRequestApproval{config: psraq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(psraq.modifiers) > 0 {
		_spec.Modifiers = psraq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, psraq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := psraq.withApprover; query != nil {
		if err := psraq.loadApprover(ctx, query, nodes, nil,
			func(n *PlayerSupervisionRequestApproval, e *User) { n.Edges.Approver = e }); err != nil {
			return nil, err
		}
	}
	if query := psraq.withSupervisionRequest; query != nil {
		if err := psraq.loadSupervisionRequest(ctx, query, nodes, nil,
			func(n *PlayerSupervisionRequestApproval, e *PlayerSupervisionRequest) { n.Edges.SupervisionRequest = e }); err != nil {
			return nil, err
		}
	}
	for i := range psraq.loadTotal {
		if err := psraq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (psraq *PlayerSupervisionRequestApprovalQuery) loadApprover(ctx context.Context, query *UserQuery, nodes []*PlayerSupervisionRequestApproval, init func(*PlayerSupervisionRequestApproval), assign func(*PlayerSupervisionRequestApproval, *User)) error {
	ids := make([]guidgql.GUID, 0, len(nodes))
	nodeids := make(map[guidgql.GUID][]*PlayerSupervisionRequestApproval)
	for i := range nodes {
		if nodes[i].user_supervision_request_approvals == nil {
			continue
		}
		fk := *nodes[i].user_supervision_request_approvals
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_supervision_request_approvals" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (psraq *PlayerSupervisionRequestApprovalQuery) loadSupervisionRequest(ctx context.Context, query *PlayerSupervisionRequestQuery, nodes []*PlayerSupervisionRequestApproval, init func(*PlayerSupervisionRequestApproval), assign func(*PlayerSupervisionRequestApproval, *PlayerSupervisionRequest)) error {
	ids := make([]guidgql.GUID, 0, len(nodes))
	nodeids := make(map[guidgql.GUID][]*PlayerSupervisionRequestApproval)
	for i := range nodes {
		if nodes[i].player_supervision_request_approvals == nil {
			continue
		}
		fk := *nodes[i].player_supervision_request_approvals
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(playersupervisionrequest.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "player_supervision_request_approvals" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (psraq *PlayerSupervisionRequestApprovalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psraq.querySpec()
	if len(psraq.modifiers) > 0 {
		_spec.Modifiers = psraq.modifiers
	}
	_spec.Node.Columns = psraq.ctx.Fields
	if len(psraq.ctx.Fields) > 0 {
		_spec.Unique = psraq.ctx.Unique != nil && *psraq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, psraq.driver, _spec)
}

func (psraq *PlayerSupervisionRequestApprovalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(playersupervisionrequestapproval.Table, playersupervisionrequestapproval.Columns, sqlgraph.NewFieldSpec(playersupervisionrequestapproval.FieldID, field.TypeString))
	_spec.From = psraq.sql
	if unique := psraq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if psraq.path != nil {
		_spec.Unique = true
	}
	if fields := psraq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playersupervisionrequestapproval.FieldID)
		for i := range fields {
			if fields[i] != playersupervisionrequestapproval.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := psraq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psraq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psraq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psraq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psraq *PlayerSupervisionRequestApprovalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psraq.driver.Dialect())
	t1 := builder.Table(playersupervisionrequestapproval.Table)
	columns := psraq.ctx.Fields
	if len(columns) == 0 {
		columns = playersupervisionrequestapproval.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psraq.sql != nil {
		selector = psraq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if psraq.ctx.Unique != nil && *psraq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range psraq.predicates {
		p(selector)
	}
	for _, p := range psraq.order {
		p(selector)
	}
	if offset := psraq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psraq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlayerSupervisionRequestApprovalGroupBy is the group-by builder for PlayerSupervisionRequestApproval entities.
type PlayerSupervisionRequestApprovalGroupBy struct {
	selector
	build *PlayerSupervisionRequestApprovalQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psragb *PlayerSupervisionRequestApprovalGroupBy) Aggregate(fns ...AggregateFunc) *PlayerSupervisionRequestApprovalGroupBy {
	psragb.fns = append(psragb.fns, fns...)
	return psragb
}

// Scan applies the selector query and scans the result into the given value.
func (psragb *PlayerSupervisionRequestApprovalGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psragb.build.ctx, "GroupBy")
	if err := psragb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlayerSupervisionRequestApprovalQuery, *PlayerSupervisionRequestApprovalGroupBy](ctx, psragb.build, psragb, psragb.build.inters, v)
}

func (psragb *PlayerSupervisionRequestApprovalGroupBy) sqlScan(ctx context.Context, root *PlayerSupervisionRequestApprovalQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(psragb.fns))
	for _, fn := range psragb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*psragb.flds)+len(psragb.fns))
		for _, f := range *psragb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*psragb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psragb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PlayerSupervisionRequestApprovalSelect is the builder for selecting fields of PlayerSupervisionRequestApproval entities.
type PlayerSupervisionRequestApprovalSelect struct {
	*PlayerSupervisionRequestApprovalQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (psras *PlayerSupervisionRequestApprovalSelect) Aggregate(fns ...AggregateFunc) *PlayerSupervisionRequestApprovalSelect {
	psras.fns = append(psras.fns, fns...)
	return psras
}

// Scan applies the selector query and scans the result into the given value.
func (psras *PlayerSupervisionRequestApprovalSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psras.ctx, "Select")
	if err := psras.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlayerSupervisionRequestApprovalQuery, *PlayerSupervisionRequestApprovalSelect](ctx, psras.PlayerSupervisionRequestApprovalQuery, psras, psras.inters, v)
}

func (psras *PlayerSupervisionRequestApprovalSelect) sqlScan(ctx context.Context, root *PlayerSupervisionRequestApprovalQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(psras.fns))
	for _, fn := range psras.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*psras.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psras.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
