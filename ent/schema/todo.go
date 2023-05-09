package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ToDo holds the schema definition for the ToDo entity.
type ToDo struct {
	ent.Schema
}

// Fields of the ToDo.
func (ToDo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("todo_uuid", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_uuid", uuid.UUID{}),
		field.Bool("is_done").
			Default(false),
		field.String("context").
			Default("null"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the ToDo.
func (ToDo) Edges() []ent.Edge {
	return nil
}
