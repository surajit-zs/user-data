package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

func All() map[int64]migration.Migrate {
	return map[int64]migration.Migrate{
		20240604091730: createUserTable(),
		20240604091731: createPartitionsTable(),
		20240604091732: createIdxUserName(),
	}
}
