module github.com/adrianpk/boletus

go 1.13

require (
	github.com/adrianpk/foundation v0.0.0-00010101000000-000000000000
	github.com/adrianpk/foundation/db v0.0.0-20200111024303-1eea754ea2f3
	github.com/adrianpk/foundation/db/pg v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v4.0.3+incompatible
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.3.0
	github.com/markbates/pkger v0.14.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200109152110-61a87790db17
	golang.org/x/text v0.3.2
)

replace github.com/adrianpk/foundation => ../foundation

replace github.com/adrianpk/foundation/db => ../foundation/db

replace github.com/adrianpk/foundation/db/pg => ../foundation/db/pg
