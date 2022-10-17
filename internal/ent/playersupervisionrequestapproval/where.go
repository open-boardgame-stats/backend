// Code generated by ent, DO NOT EDIT.

package playersupervisionrequestapproval

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Approved applies equality check predicate on the "approved" field. It's identical to ApprovedEQ.
func Approved(v bool) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApproved), v))
	})
}

// ApprovedEQ applies the EQ predicate on the "approved" field.
func ApprovedEQ(v bool) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApproved), v))
	})
}

// ApprovedNEQ applies the NEQ predicate on the "approved" field.
func ApprovedNEQ(v bool) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldApproved), v))
	})
}

// ApprovedIsNil applies the IsNil predicate on the "approved" field.
func ApprovedIsNil() predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldApproved)))
	})
}

// ApprovedNotNil applies the NotNil predicate on the "approved" field.
func ApprovedNotNil() predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldApproved)))
	})
}

// HasApprover applies the HasEdge predicate on the "approver" edge.
func HasApprover() predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ApproverTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApproverTable, ApproverColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApproverWith applies the HasEdge predicate on the "approver" edge with a given conditions (other predicates).
func HasApproverWith(preds ...predicate.User) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ApproverInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApproverTable, ApproverColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSupervisionRequest applies the HasEdge predicate on the "supervision_request" edge.
func HasSupervisionRequest() predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SupervisionRequestTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SupervisionRequestTable, SupervisionRequestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSupervisionRequestWith applies the HasEdge predicate on the "supervision_request" edge with a given conditions (other predicates).
func HasSupervisionRequestWith(preds ...predicate.PlayerSupervisionRequest) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SupervisionRequestInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SupervisionRequestTable, SupervisionRequestColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlayerSupervisionRequestApproval) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlayerSupervisionRequestApproval) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
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
func Not(p predicate.PlayerSupervisionRequestApproval) predicate.PlayerSupervisionRequestApproval {
	return predicate.PlayerSupervisionRequestApproval(func(s *sql.Selector) {
		p(s.Not())
	})
}
