// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pl *PlayerQuery) CollectFields(ctx context.Context, satisfies ...string) (*PlayerQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pl, nil
	}
	if err := pl.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pl, nil
}

func (pl *PlayerQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "owner":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: pl.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pl.withOwner = query
		case "supervisors":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: pl.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pl.WithNamedSupervisors(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "supervisionRequests", "supervision_requests":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerSupervisionRequestQuery{config: pl.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pl.WithNamedSupervisionRequests(alias, func(wq *PlayerSupervisionRequestQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type playerPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PlayerPaginateOption
}

func newPlayerPaginateArgs(rv map[string]interface{}) *playerPaginateArgs {
	args := &playerPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PlayerWhereInput); ok {
		args.opts = append(args.opts, WithPlayerFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (psr *PlayerSupervisionRequestQuery) CollectFields(ctx context.Context, satisfies ...string) (*PlayerSupervisionRequestQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return psr, nil
	}
	if err := psr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return psr, nil
}

func (psr *PlayerSupervisionRequestQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "sender":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: psr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			psr.withSender = query
		case "player":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerQuery{config: psr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			psr.withPlayer = query
		case "approvals":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerSupervisionRequestApprovalQuery{config: psr.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			psr.WithNamedApprovals(alias, func(wq *PlayerSupervisionRequestApprovalQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type playersupervisionrequestPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PlayerSupervisionRequestPaginateOption
}

func newPlayerSupervisionRequestPaginateArgs(rv map[string]interface{}) *playersupervisionrequestPaginateArgs {
	args := &playersupervisionrequestPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PlayerSupervisionRequestWhereInput); ok {
		args.opts = append(args.opts, WithPlayerSupervisionRequestFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (psra *PlayerSupervisionRequestApprovalQuery) CollectFields(ctx context.Context, satisfies ...string) (*PlayerSupervisionRequestApprovalQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return psra, nil
	}
	if err := psra.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return psra, nil
}

func (psra *PlayerSupervisionRequestApprovalQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "approver":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: psra.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			psra.withApprover = query
		case "supervisionRequest", "supervision_request":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerSupervisionRequestQuery{config: psra.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			psra.withSupervisionRequest = query
		}
	}
	return nil
}

type playersupervisionrequestapprovalPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PlayerSupervisionRequestApprovalPaginateOption
}

func newPlayerSupervisionRequestApprovalPaginateArgs(rv map[string]interface{}) *playersupervisionrequestapprovalPaginateArgs {
	args := &playersupervisionrequestapprovalPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PlayerSupervisionRequestApprovalWhereInput); ok {
		args.opts = append(args.opts, WithPlayerSupervisionRequestApprovalFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "players":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedPlayers(alias, func(wq *PlayerQuery) {
				*wq = *query
			})
		case "mainPlayer", "main_player":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &PlayerQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.withMainPlayer = query
		}
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]interface{}) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}
