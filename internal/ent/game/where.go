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
	return predicate.Game(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id guidgql.GUID) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id guidgql.GUID) predicate.Game {
	return predicate.Game(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...guidgql.GUID) predicate.Game {
	return predicate.Game(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...guidgql.GUID) predicate.Game {
	return predicate.Game(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id guidgql.GUID) predicate.Game {
	return predicate.Game(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id guidgql.GUID) predicate.Game {
	return predicate.Game(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id guidgql.GUID) predicate.Game {
	return predicate.Game(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id guidgql.GUID) predicate.Game {
	return predicate.Game(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldName, v))
}

// MinPlayers applies equality check predicate on the "min_players" field. It's identical to MinPlayersEQ.
func MinPlayers(v int) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldMinPlayers, v))
}

// MaxPlayers applies equality check predicate on the "max_players" field. It's identical to MaxPlayersEQ.
func MaxPlayers(v int) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldMaxPlayers, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldDescription, v))
}

// BoardgamegeekURL applies equality check predicate on the "boardgamegeek_url" field. It's identical to BoardgamegeekURLEQ.
func BoardgamegeekURL(v string) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldBoardgamegeekURL, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Game {
	return predicate.Game(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Game {
	return predicate.Game(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Game {
	return predicate.Game(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Game {
	return predicate.Game(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Game {
	return predicate.Game(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Game {
	return predicate.Game(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Game {
	return predicate.Game(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Game {
	return predicate.Game(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Game {
	return predicate.Game(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Game {
	return predicate.Game(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Game {
	return predicate.Game(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Game {
	return predicate.Game(sql.FieldContainsFold(FieldName, v))
}

// MinPlayersEQ applies the EQ predicate on the "min_players" field.
func MinPlayersEQ(v int) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldMinPlayers, v))
}

// MinPlayersNEQ applies the NEQ predicate on the "min_players" field.
func MinPlayersNEQ(v int) predicate.Game {
	return predicate.Game(sql.FieldNEQ(FieldMinPlayers, v))
}

// MinPlayersIn applies the In predicate on the "min_players" field.
func MinPlayersIn(vs ...int) predicate.Game {
	return predicate.Game(sql.FieldIn(FieldMinPlayers, vs...))
}

// MinPlayersNotIn applies the NotIn predicate on the "min_players" field.
func MinPlayersNotIn(vs ...int) predicate.Game {
	return predicate.Game(sql.FieldNotIn(FieldMinPlayers, vs...))
}

// MinPlayersGT applies the GT predicate on the "min_players" field.
func MinPlayersGT(v int) predicate.Game {
	return predicate.Game(sql.FieldGT(FieldMinPlayers, v))
}

// MinPlayersGTE applies the GTE predicate on the "min_players" field.
func MinPlayersGTE(v int) predicate.Game {
	return predicate.Game(sql.FieldGTE(FieldMinPlayers, v))
}

// MinPlayersLT applies the LT predicate on the "min_players" field.
func MinPlayersLT(v int) predicate.Game {
	return predicate.Game(sql.FieldLT(FieldMinPlayers, v))
}

// MinPlayersLTE applies the LTE predicate on the "min_players" field.
func MinPlayersLTE(v int) predicate.Game {
	return predicate.Game(sql.FieldLTE(FieldMinPlayers, v))
}

// MaxPlayersEQ applies the EQ predicate on the "max_players" field.
func MaxPlayersEQ(v int) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldMaxPlayers, v))
}

// MaxPlayersNEQ applies the NEQ predicate on the "max_players" field.
func MaxPlayersNEQ(v int) predicate.Game {
	return predicate.Game(sql.FieldNEQ(FieldMaxPlayers, v))
}

// MaxPlayersIn applies the In predicate on the "max_players" field.
func MaxPlayersIn(vs ...int) predicate.Game {
	return predicate.Game(sql.FieldIn(FieldMaxPlayers, vs...))
}

// MaxPlayersNotIn applies the NotIn predicate on the "max_players" field.
func MaxPlayersNotIn(vs ...int) predicate.Game {
	return predicate.Game(sql.FieldNotIn(FieldMaxPlayers, vs...))
}

// MaxPlayersGT applies the GT predicate on the "max_players" field.
func MaxPlayersGT(v int) predicate.Game {
	return predicate.Game(sql.FieldGT(FieldMaxPlayers, v))
}

// MaxPlayersGTE applies the GTE predicate on the "max_players" field.
func MaxPlayersGTE(v int) predicate.Game {
	return predicate.Game(sql.FieldGTE(FieldMaxPlayers, v))
}

// MaxPlayersLT applies the LT predicate on the "max_players" field.
func MaxPlayersLT(v int) predicate.Game {
	return predicate.Game(sql.FieldLT(FieldMaxPlayers, v))
}

// MaxPlayersLTE applies the LTE predicate on the "max_players" field.
func MaxPlayersLTE(v int) predicate.Game {
	return predicate.Game(sql.FieldLTE(FieldMaxPlayers, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Game {
	return predicate.Game(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Game {
	return predicate.Game(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Game {
	return predicate.Game(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Game {
	return predicate.Game(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Game {
	return predicate.Game(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Game {
	return predicate.Game(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Game {
	return predicate.Game(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Game {
	return predicate.Game(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Game {
	return predicate.Game(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Game {
	return predicate.Game(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Game {
	return predicate.Game(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Game {
	return predicate.Game(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Game {
	return predicate.Game(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Game {
	return predicate.Game(sql.FieldContainsFold(FieldDescription, v))
}

// BoardgamegeekURLEQ applies the EQ predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLEQ(v string) predicate.Game {
	return predicate.Game(sql.FieldEQ(FieldBoardgamegeekURL, v))
}

// BoardgamegeekURLNEQ applies the NEQ predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLNEQ(v string) predicate.Game {
	return predicate.Game(sql.FieldNEQ(FieldBoardgamegeekURL, v))
}

// BoardgamegeekURLIn applies the In predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLIn(vs ...string) predicate.Game {
	return predicate.Game(sql.FieldIn(FieldBoardgamegeekURL, vs...))
}

// BoardgamegeekURLNotIn applies the NotIn predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLNotIn(vs ...string) predicate.Game {
	return predicate.Game(sql.FieldNotIn(FieldBoardgamegeekURL, vs...))
}

// BoardgamegeekURLGT applies the GT predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLGT(v string) predicate.Game {
	return predicate.Game(sql.FieldGT(FieldBoardgamegeekURL, v))
}

// BoardgamegeekURLGTE applies the GTE predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLGTE(v string) predicate.Game {
	return predicate.Game(sql.FieldGTE(FieldBoardgamegeekURL, v))
}

// BoardgamegeekURLLT applies the LT predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLLT(v string) predicate.Game {
	return predicate.Game(sql.FieldLT(FieldBoardgamegeekURL, v))
}

// BoardgamegeekURLLTE applies the LTE predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLLTE(v string) predicate.Game {
	return predicate.Game(sql.FieldLTE(FieldBoardgamegeekURL, v))
}

// BoardgamegeekURLContains applies the Contains predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLContains(v string) predicate.Game {
	return predicate.Game(sql.FieldContains(FieldBoardgamegeekURL, v))
}

// BoardgamegeekURLHasPrefix applies the HasPrefix predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLHasPrefix(v string) predicate.Game {
	return predicate.Game(sql.FieldHasPrefix(FieldBoardgamegeekURL, v))
}

// BoardgamegeekURLHasSuffix applies the HasSuffix predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLHasSuffix(v string) predicate.Game {
	return predicate.Game(sql.FieldHasSuffix(FieldBoardgamegeekURL, v))
}

// BoardgamegeekURLIsNil applies the IsNil predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLIsNil() predicate.Game {
	return predicate.Game(sql.FieldIsNull(FieldBoardgamegeekURL))
}

// BoardgamegeekURLNotNil applies the NotNil predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLNotNil() predicate.Game {
	return predicate.Game(sql.FieldNotNull(FieldBoardgamegeekURL))
}

// BoardgamegeekURLEqualFold applies the EqualFold predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLEqualFold(v string) predicate.Game {
	return predicate.Game(sql.FieldEqualFold(FieldBoardgamegeekURL, v))
}

// BoardgamegeekURLContainsFold applies the ContainsFold predicate on the "boardgamegeek_url" field.
func BoardgamegeekURLContainsFold(v string) predicate.Game {
	return predicate.Game(sql.FieldContainsFold(FieldBoardgamegeekURL, v))
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := newAuthorStep()
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
			sqlgraph.Edge(sqlgraph.O2M, false, FavoritesTable, FavoritesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFavoritesWith applies the HasEdge predicate on the "favorites" edge with a given conditions (other predicates).
func HasFavoritesWith(preds ...predicate.GameFavorite) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := newFavoritesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVersions applies the HasEdge predicate on the "versions" edge.
func HasVersions() predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, VersionsTable, VersionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVersionsWith applies the HasEdge predicate on the "versions" edge with a given conditions (other predicates).
func HasVersionsWith(preds ...predicate.GameVersion) predicate.Game {
	return predicate.Game(func(s *sql.Selector) {
		step := newVersionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Game) predicate.Game {
	return predicate.Game(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Game) predicate.Game {
	return predicate.Game(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Game) predicate.Game {
	return predicate.Game(sql.NotPredicates(p))
}
