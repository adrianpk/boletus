module github.com/adrianpk/boletus

go 1.13

require (
	github.com/adrianpk/foundation v0.0.0-00010101000000-000000000000
	github.com/adrianpk/foundation/db v0.0.0-20200111024303-1eea754ea2f3
	github.com/adrianpk/foundation/db/pg v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v1.1.1
	github.com/go-chi/chi v4.0.3+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.3.0
	github.com/markbates/pkger v0.14.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200109152110-61a87790db17
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sys v0.0.0-20200113162924-86b910548bc1 // indirect
	golang.org/x/text v0.3.2
	google.golang.org/genproto v0.0.0-20200113173426-e1de0a7b01eb // indirect
	google.golang.org/grpc v1.26.0
)

replace github.com/adrianpk/foundation => ../foundation

replace github.com/adrianpk/foundation/db => ../foundation/db

replace github.com/adrianpk/foundation/db/pg => ../foundation/db/pg
