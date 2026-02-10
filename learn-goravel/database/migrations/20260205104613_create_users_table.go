package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"learn-goravel/app/facades"
)

type M20260205104613CreateUsersTable struct{}

// Signature The unique signature for the migration.
func (r *M20260205104613CreateUsersTable) Signature() string {
	return "20260205104613_create_users_table"
}

// Up Run the migrations.
func (r *M20260205104613CreateUsersTable) Up() error {
	if !facades.Schema().HasTable("users") {
		return facades.Schema().Create("users", func(table schema.Blueprint) {
			table.ID()                           // id (primary key)
			table.String("name")                 // name (varchar 255)
			table.String("email")                // email
			table.String("password")             // password
			table.Integer("age").Nullable()      // age (can be null)
			table.String("city", 100).Nullable() // city (varchar 100, nullable)
			table.Boolean("is_active").Default("1")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20260205104613CreateUsersTable) Down() error {
	return facades.Schema().DropIfExists("users")
}
