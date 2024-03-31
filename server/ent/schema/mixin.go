package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TimeMixin は、created_at と updated_at フィールドを定義します。
type TimeMixin struct {
	mixin.Schema
}

// Fields メソッドは、Mixinによって提供されるフィールドを返します。
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
