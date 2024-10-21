package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deadline"),
		field.Int("range"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		// many to one
		edge.From("group", Group.Type).
			Ref("tasks").
			Unique().
			Annotations(entsql.OnDelete(entsql.SetNull)),

		edge.To("topics", Topic.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
