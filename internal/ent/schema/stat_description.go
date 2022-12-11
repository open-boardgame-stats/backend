package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
)

// StatDescription is the schema for the StatDescription entity.
type StatDescription struct {
	ent.Schema
}

// Fields of the StatDescription.
func (StatDescription) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(guidgql.GUID{}).DefaultFunc(guidgql.New(guidgql.StatDescription)),
		field.Enum("type").GoType(stat.StatType(stat.Numeric)),
		field.String("name").NotEmpty(),
		field.String("description").Optional().Default(""),
	}
}

// Edges of the StatDescription.
func (StatDescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game", Game.Type).Annotations(entgql.Skip(entgql.SkipAll)),
	}
}

// Annotations of the StatDescription.
func (StatDescription) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipWhereInput),
	}
}