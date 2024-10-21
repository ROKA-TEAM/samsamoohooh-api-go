package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("resolution"),
		field.Enum("role").
			Values(
				"ADMIN", "MANAGER", "USER",
			),
		field.Enum("social").
			Values(
				"KAKAO", "APPLE", "GOOGLE",
			),
		field.String("socialSub"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{

		// many to many
		edge.To("groups", Group.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),

		// one to many
		edge.To("comments", Comment.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("posts", Post.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("topics", Topic.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}
