// Code generated by ent, DO NOT EDIT.

package game

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// ID filters vertices based on their ID field.
func ID(id guidgql.GUID) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id guidgql.GUID) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id guidgql.GUID) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...guidgql.GUID) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...guidgql.GUID) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id guidgql.GUID) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id guidgql.GUID) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id guidgql.GUID) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id guidgql.GUID) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// MinPlayers applies equality check predicate on the "min_players" field. It's identical to MinPlayersEQ.
func MinPlayers(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinPlayers), v))
	})
}

// MaxPlayers applies equality check predicate on the "max_players" field. It's identical to MaxPlayersEQ.
func MaxPlayers(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxPlayers), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// BoardgamegeekURL applies equality check predicate on the "boardgamegeek_url" field. It's identical to BoardgamegeekURLEQ.
func BoardgamegeekURL(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBoardgamegeekURL), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Game {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Game {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// MinPlayersEQ applies the EQ predicate on the "min_players" field.
func MinPlayersEQ(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinPlayers), v))
	})
}

// MinPlayersNEQ applies the NEQ predicate on the "min_players" field.
func MinPlayersNEQ(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMinPlayers), v))
	})
}

// MinPlayersIn applies the In predicate on the "min_players" field.
func MinPlayersIn(vs ...int) predicate.Game {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMinPlayers), v...))
	})
}

// MinPlayersNotIn applies the NotIn predicate on the "min_players" field.
func MinPlayersNotIn(vs ...int) predicate.Game {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMinPlayers), v...))
	})
}

// MinPlayersGT applies the GT predicate on the "min_players" field.
func MinPlayersGT(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMinPlayers), v))
	})
}

// MinPlayersGTE applies the GTE predicate on the "min_players" field.
func MinPlayersGTE(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMinPlayers), v))
	})
}

// MinPlayersLT applies the LT predicate on the "min_players" field.
func MinPlayersLT(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMinPlayers), v))
	})
}

// MinPlayersLTE applies the LTE predicate on the "min_players" field.
func MinPlayersLTE(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMinPlayers), v))
	})
}

// MaxPlayersEQ applies the EQ predicate on the "max_players" field.
func MaxPlayersEQ(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxPlayers), v))
	})
}

// MaxPlayersNEQ applies the NEQ predicate on the "max_players" field.
func MaxPlayersNEQ(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxPlayers), v))
	})
}

// MaxPlayersIn applies the In predicate on the "max_players" field.
func MaxPlayersIn(vs ...int) predicate.Game {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMaxPlayers), v...))
	})
}

// MaxPlayersNotIn applies the NotIn predicate on the "max_players" field.
func MaxPlayersNotIn(vs ...int) predicate.Game {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMaxPlayers), v...))
	})
}

// MaxPlayersGT applies the GT predicate on the "max_players" field.
func MaxPlayersGT(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxPlayers), v))
	})
}

// MaxPlayersGTE applies the GTE predicate on the "max_players" field.
func MaxPlayersGTE(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxPlayers), v))
	})
}

// MaxPlayersLT applies the LT predicate on the "max_players" field.
func MaxPlayersLT(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxPlayers), v))
	})
}

// MaxPlayersLTE applies the LTE predicate on the "max_players" field.
func MaxPlayersLTE(v int) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxPlayers), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Game {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Game {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// BoardgamegeekURLEQ applies the EQ predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLEQ(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBoardgamegeekURL), v))
	})
}

// BoardgamegeekURLNEQ applies the NEQ predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLNEQ(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBoardgamegeekURL), v))
	})
}

// BoardgamegeekURLIn applies the In predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLIn(vs ...string) predicate.Game {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBoardgamegeekURL), v...))
	})
}

// BoardgamegeekURLNotIn applies the NotIn predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLNotIn(vs ...string) predicate.Game {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBoardgamegeekURL), v...))
	})
}

// BoardgamegeekURLGT applies the GT predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLGT(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBoardgamegeekURL), v))
	})
}

// BoardgamegeekURLGTE applies the GTE predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLGTE(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBoardgamegeekURL), v))
	})
}

// BoardgamegeekURLLT applies the LT predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLLT(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBoardgamegeekURL), v))
	})
}

// BoardgamegeekURLLTE applies the LTE predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLLTE(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBoardgamegeekURL), v))
	})
}

// BoardgamegeekURLContains applies the Contains predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLContains(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBoardgamegeekURL), v))
	})
}

// BoardgamegeekURLHasPrefix applies the HasPrefix predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLHasPrefix(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBoardgamegeekURL), v))
	})
}

// BoardgamegeekURLHasSuffix applies the HasSuffix predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLHasSuffix(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBoardgamegeekURL), v))
	})
}

// BoardgamegeekURLIsNil applies the IsNil predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLIsNil() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBoardgamegeekURL)))
	})
}

// BoardgamegeekURLNotNil applies the NotNil predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLNotNil() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBoardgamegeekURL)))
	})
}

// BoardgamegeekURLEqualFold applies the EqualFold predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLEqualFold(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBoardgamegeekURL), v))
	})
}

// BoardgamegeekURLContainsFold applies the ContainsFold predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLContainsFold(v string) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBoardgamegeekURL), v))
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFavorites applies the HasEdge predicate on the "favorites" edge.
func HasFavorites() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FavoritesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FavoritesTable, FavoritesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFavoritesWith applies the HasEdge predicate on the "favorites" edge with a given conditions (other predicates).
func HasFavoritesWith(preds ...predicate.GameFavorite) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FavoritesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FavoritesTable, FavoritesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNumericalStatDescriptions applies the HasEdge predicate on the "numerical_stat_descriptions" edge.
func HasNumericalStatDescriptions() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NumericalStatDescriptionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, NumericalStatDescriptionsTable, NumericalStatDescriptionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNumericalStatDescriptionsWith applies the HasEdge predicate on the "numerical_stat_descriptions" edge with a given conditions (other predicates).
func HasNumericalStatDescriptionsWith(preds ...predicate.NumericalStatDescription) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NumericalStatDescriptionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, NumericalStatDescriptionsTable, NumericalStatDescriptionsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEnumStatDescriptions applies the HasEdge predicate on the "enum_stat_descriptions" edge.
func HasEnumStatDescriptions() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EnumStatDescriptionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, EnumStatDescriptionsTable, EnumStatDescriptionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEnumStatDescriptionsWith applies the HasEdge predicate on the "enum_stat_descriptions" edge with a given conditions (other predicates).
func HasEnumStatDescriptionsWith(preds ...predicate.EnumStatDescription) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EnumStatDescriptionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, EnumStatDescriptionsTable, EnumStatDescriptionsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMatches applies the HasEdge predicate on the "matches" edge.
func HasMatches() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MatchesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MatchesTable, MatchesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMatchesWith applies the HasEdge predicate on the "matches" edge with a given conditions (other predicates).
func HasMatchesWith(preds ...predicate.Match) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MatchesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MatchesTable, MatchesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Game) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Game) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Game) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		p(s.Not())
	})
}
