// Code generated by ent, DO NOT EDIT.

package match

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// ID filters vertices based on their ID field.
func ID(id guidgql.GUID) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id guidgql.GUID) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id guidgql.GUID) predicate.Match {
	return predicate.Match(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...guidgql.GUID) predicate.Match {
	return predicate.Match(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...guidgql.GUID) predicate.Match {
	return predicate.Match(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id guidgql.GUID) predicate.Match {
	return predicate.Match(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id guidgql.GUID) predicate.Match {
	return predicate.Match(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id guidgql.GUID) predicate.Match {
	return predicate.Match(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id guidgql.GUID) predicate.Match {
	return predicate.Match(sql.FieldLTE(FieldID, id))
}

// HasGameVersion applies the HasEdge predicate on the "game_version" edge.
func HasGameVersion() predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GameVersionTable, GameVersionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGameVersionWith applies the HasEdge predicate on the "game_version" edge with a given conditions (other predicates).
func HasGameVersionWith(preds ...predicate.GameVersion) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := newGameVersionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlayers applies the HasEdge predicate on the "players" edge.
func HasPlayers() predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PlayersTable, PlayersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlayersWith applies the HasEdge predicate on the "players" edge with a given conditions (other predicates).
func HasPlayersWith(preds ...predicate.Player) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := newPlayersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStats applies the HasEdge predicate on the "stats" edge.
func HasStats() predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatsTable, StatsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatsWith applies the HasEdge predicate on the "stats" edge with a given conditions (other predicates).
func HasStatsWith(preds ...predicate.Statistic) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := newStatsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Match) predicate.Match {
	return predicate.Match(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Match) predicate.Match {
	return predicate.Match(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Match) predicate.Match {
	return predicate.Match(sql.NotPredicates(p))
}
