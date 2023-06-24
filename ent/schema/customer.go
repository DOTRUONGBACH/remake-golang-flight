package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Customer holds the schema definition for the Customer2 entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID")),
		field.String("fullname").Annotations(entgql.OrderField("FULLNAME")),
		field.String("phone").Annotations(entgql.OrderField("PHONE")),
		field.String("address").Annotations(entgql.OrderField("ADDRESS")),
		field.Enum("gender").Values("Male", "Female", "Other").Annotations(entgql.OrderField("GENDER")),
		field.String("citizen_id").MinLen(9).MaxLen(12).Unique().Annotations(entgql.OrderField("CITIZEN_ID")),
		field.Time("date_of_birth").Annotations(entgql.OrderField("DOB")),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Edges of the Customer2.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type),
	}
}
func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
