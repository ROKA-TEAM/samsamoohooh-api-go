package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("book_title"),
		field.String("author"),
		field.Int("max_page"),
		field.String("publisher"),
		field.String("description"),
		field.Int("book_mark"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{

		// many to many
		edge.From("users", User.Type).
			Ref("groups"),

		// many to one
		edge.To("posts", Post.Type),
		edge.To("tasks", Task.Type),
	}
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}
