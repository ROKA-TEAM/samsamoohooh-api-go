package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(timeNowUTC()),

		field.Time("updated_at").
			Default(timeNowUTC()).
			UpdateDefault(time.Now),
	}
}

func timeNowUTC() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}
