package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return nil
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("messages").Unique().Comment("ユーザー"),
	}
}
