package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		// many to one
		edge.From("user", User.Type).
			Ref("comments").
			Unique().
			Annotations(entsql.OnDelete(entsql.SetNull)),

		edge.From("post", Post.Type).
			Ref("comments").
			Unique().
			Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}
