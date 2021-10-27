package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").Optional(),
		field.String("uid").MaxLen(32),
		field.String("phone").MaxLen(32).Nillable().Optional().Unique(),
		field.Int32("status").Default(0),
		field.JSON("extend", make(map[string]string)).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
