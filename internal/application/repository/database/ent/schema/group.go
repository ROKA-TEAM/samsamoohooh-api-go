package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
			Ref("groups").
			// Group과 User 사이의 연결만 삭제되고, User는 그대로 유지됩니다.
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),

		// many to one
		edge.To("posts", Post.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("tasks", Task.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
