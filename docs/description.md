# Description

**Application** basically integrates nnn componentes

<img src="img/app_model.svg" width="480">

## Migrator

Manage creation and update of new tables, views and indexes.
The formal way to create a new migration file under `internal/mig` directory.
Currently therer three migrations files that you can use as a reference to create a new ones.

```shell
├── 00001createuserstable.go
├── 00002createeventstable.go
├── 00003createticketstable.go
├── mig.go
└── step.go
```

Basically each of these files defines two functions with the following structure.

**Sample 00004createxxxxxtable.go**

```go`
package mig

func (s *step) CreateXXXXXTable() error {
	tx := s.GetTx()

	st := `CREATE TABLE xxxxx
	(
		id UUID PRIMARY KEY,
		field1 VARCHAR(36) UNIQUE,
		field2 VARCHAR(16),
		field_etc VARCHAR(32) UNIQUE,
	);`

	_, err := tx.Exec(st)
	if err != nil {
		return err
	}

	return nil
}

// DropXXXXXTable rollback
func (s *step) DropXXXXXTable() error {
	tx := s.GetTx()

	st := `DROP TABLE xxxxx;`

	_, err := tx.Exec(st)
	if err != nil {
		return err
	}

	return nil
}
```

Later you need to append these functions as migrations steps in `internal/mig/mig.go`

```go
// GetMigrator configured.
func (m *Migrator) addSteps() {
	// Migrations

  // Intentionally omitted for clarity

	// Create xxxxx table
	s = &step{}
	s.Config(s.CreateXXXXXTable, s.DropXXXXXTable)
	m.AddMigration(s)
}
```

Next time you run the application it will be executed and register in migrations table.


## Seeder

Works in the same way as the migrator.
First you create a file under `internal/seed`

```shell
├── 00001users.go
├── 00002tickets.go
├── 00002xxxxxs.go
├── seed.go
└── step.go
```

And you add these steps in `internal/seed/seed.go`

```go
// GetSeeder configured.
func (s *Seeder) addSteps() {
	// Seeds

  // Intentionally omitted for clarity

  // Events and tickets
	st = &step{}
	st.Config(st.XXXXX)
	s.AddSeed(st)
}
```

Unlike the migrator, seeding process currently g is only executed the first time the application runs. In following restarts the ap verifies that the app base users have been created and if it finds them it does not try to execute the operation again.


It is planned to alter its operation so that it also keeps a record of applied seedings in a way that  allow it to detect those that are still pending.
