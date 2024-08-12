package schema

import (
		"entgo.io/ent"
		"entgo.io/ent/schema/edge"
		"entgo.io/ent/schema/field"
		"time"
	)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Positive(),
		field.String("title").NotEmpty(),
		field.String("body").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
				edge.From("users", User.Type).
						Ref("articles").
						Field("user_id").
						Unique().
						Required(),
			}
}
