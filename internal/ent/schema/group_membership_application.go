package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type GroupMembershipApplication struct {
	ent.Schema
}

func (GroupMembershipApplication) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("message").Default(""),
	}
}

func (GroupMembershipApplication) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("group_membership_applications").Required(),
		edge.From("group", Group.Type).Ref("applications").Required(),
	}
}

func (GroupMembershipApplication) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipWhereInput),
	}
}
