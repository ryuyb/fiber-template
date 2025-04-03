package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"live-pilot/pkg/utils/entgo"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			SchemaType(map[string]string{
				dialect.MySQL:    "int",
				dialect.Postgres: "serial",
			}).
			Positive().
			Immutable().
			Unique(),
		field.String("username").
			Unique(),
		field.String("password"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entgo.Time{},
	}
}
