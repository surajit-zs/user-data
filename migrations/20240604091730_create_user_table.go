package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const createUser = `
CREATE TABLE IF NOT EXISTS "user" (
    id UUID NOT NULL primary key ,
    name VARCHAR(255) NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP default NULL) PARTITION BY HASH (id);
`

func createUserTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createUser)
			return err
		},
	}
}
