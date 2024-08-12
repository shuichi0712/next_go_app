package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// StudyTime holds the schema definition for the StudyTime entity.
type StudyTime struct {
	ent.Schema
}

// Fields of the StudyTime.
func (StudyTime) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Positive(),
		field.String("title").Unique().NotEmpty(),
		field.Int("hour"),
		field.Int("minite"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the StudyTime.
func (StudyTime) Edges() []ent.Edge {
	
		return []ent.Edge{
					edge.From("users", User.Type).
							Ref("studytimes").
							Field("user_id").
							Unique().
							Required(),
				}
	
}
