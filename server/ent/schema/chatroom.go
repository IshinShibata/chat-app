package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChatRoom holds the schema definition for the ChatRoom entity.
type ChatRoom struct {
	ent.Schema
}

func (ChatRoom) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the ChatRoom.
func (ChatRoom) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().Comment("チャットルーム"),
	}
}

// Edges of the ChatRoom.
func (ChatRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
