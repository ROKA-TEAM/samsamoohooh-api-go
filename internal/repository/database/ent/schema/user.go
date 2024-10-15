package schema

import (
	"entgo.io/ent"
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
				"ADMIN", "GUEST",
			),
		field.Enum("social").
			Values(
				"KAKAO", "APPLE", "GOOGLE",
			),
		field.String("socialSub").
			Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{

		// many to many
		edge.To("groups", Group.Type),

		// one to many
		edge.To("comments", Comment.Type),
		edge.To("posts", Post.Type),
		edge.To("topics", Topic.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		SoftDeleteMixin{},
	}
}
