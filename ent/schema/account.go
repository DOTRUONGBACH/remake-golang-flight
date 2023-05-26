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

type Account struct {
	ent.Schema
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID")),
		field.String("email").Unique().Annotations(entgql.OrderField("EMAIL")),
		field.String("password").Sensitive().Annotations(entgql.OrderField("PASSWORD")),
		field.Enum("role").Values("admin", "subscriber", "customer").Annotations(entgql.OrderField("ROLE")),
		field.Enum("status").Values("inactive", "active").Annotations(entgql.OrderField("STATUS")),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entgql.OrderField("UPDATED_AT")),
	}
}
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("acc_owner", Customer.Type).Ref("accounts").Unique(),
	}
}
func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
