package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("content"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		// many to one
		edge.To("comments", Comment.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),

		edge.From("user", User.Type).
			Ref("posts").
			Unique().
			Annotations(entsql.OnDelete(entsql.SetNull)),

		edge.From("group", Group.Type).
			Ref("posts").Unique().
			Unique().
			Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
