package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

const idxUserName = `CREATE INDEX idx_user_user_name ON "user" (user_name);`

func createIdxUserName() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(idxUserName)
			return err
		},
	}
}
