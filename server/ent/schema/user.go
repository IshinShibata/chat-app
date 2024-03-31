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

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").NotEmpty().Optional().Nillable().Unique().Comment("UID"),
		field.String("email").NotEmpty().Unique().Comment("メールアドレス"),
		field.String("name").
			Default("unknown"),
		field.Time("last_login_at").Optional().Nillable().Comment("最終ログイン日時"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type).Comment("メッセージ"),
		edge.From("chatRooms", ChatRoom.Type).Ref("users").Unique().Comment("グループ"),
	}
}
