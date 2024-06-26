// Code generated by ent, DO NOT EDIT.

package groupmembership

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// ID filters vertices based on their ID field.
func ID(id guidgql.GUID) predicate.GroupMembership {
	return predicate.GroupMembership(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id guidgql.GUID) predicate.GroupMembership {
	return predicate.GroupMembership(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id guidgql.GUID) predicate.GroupMembership {
	return predicate.GroupMembership(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...guidgql.GUID) predicate.GroupMembership {
	return predicate.GroupMembership(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...guidgql.GUID) predicate.GroupMembership {
	return predicate.GroupMembership(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id guidgql.GUID) predicate.GroupMembership {
	return predicate.GroupMembership(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id guidgql.GUID) predicate.GroupMembership {
	return predicate.GroupMembership(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id guidgql.GUID) predicate.GroupMembership {
	return predicate.GroupMembership(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id guidgql.GUID) predicate.GroupMembership {
	return predicate.GroupMembership(sql.FieldLTE(FieldID, id))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v enums.Role) predicate.GroupMembership {
	vc := v
	return predicate.GroupMembership(sql.FieldEQ(FieldRole, vc))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v enums.Role) predicate.GroupMembership {
	vc := v
	return predicate.GroupMembership(sql.FieldNEQ(FieldRole, vc))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...enums.Role) predicate.GroupMembership {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMembership(sql.FieldIn(FieldRole, v...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...enums.Role) predicate.GroupMembership {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMembership(sql.FieldNotIn(FieldRole, v...))
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.GroupMembership {
	return predicate.GroupMembership(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.GroupMembership {
	return predicate.GroupMembership(func(s *sql.Selector) {
		step := newGroupStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.GroupMembership {
	return predicate.GroupMembership(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.GroupMembership {
	return predicate.GroupMembership(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GroupMembership) predicate.GroupMembership {
	return predicate.GroupMembership(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GroupMembership) predicate.GroupMembership {
	return predicate.GroupMembership(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GroupMembership) predicate.GroupMembership {
	return predicate.GroupMembership(sql.NotPredicates(p))
}
