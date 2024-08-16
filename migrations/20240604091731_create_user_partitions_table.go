package migrations

import "gofr.dev/pkg/gofr/migration"

const createUserPartitions = `
CREATE TABLE user_p0 PARTITION OF "user"
    FOR VALUES WITH (MODULUS 5, REMAINDER 0);

CREATE TABLE user_p1 PARTITION OF "user"
    FOR VALUES WITH (MODULUS 5, REMAINDER 1);

CREATE TABLE user_p2 PARTITION OF "user"
    FOR VALUES WITH (MODULUS 5, REMAINDER 2);

CREATE TABLE user_p3 PARTITION OF "user"
    FOR VALUES WITH (MODULUS 5, REMAINDER 3);

CREATE TABLE user_p4 PARTITION OF "user"
    FOR VALUES WITH (MODULUS 5, REMAINDER 4);
`

func createPartitionsTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createUserPartitions)
			return err
		},
	}
}
