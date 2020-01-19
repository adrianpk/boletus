<img src="docs/img/boletus_iso.png" width="480">

### Etymology 1

Borrowed from Latin bōlētus, from Ancient Greek βωλίτης (bōlítēs).
Noun

#### boleto m (plural boletos)

  * porcini (Boletus edulis, an edible mushroom)
    (in general) any bolete mushroom

### Etymology 2

Borrowed from Spanish boleta or Italian bolletta, from Latin bulla.
Noun

#### boleto m (plural boletos)

  * ticket

## Features

  * List events
  * Ticket summary by events
  * Pre booking
  * Ticket purchasing

## Screenshots

<img src="docs/img/new_event.png" width="480">

[More](docs/screenshots.md)


## Dev. Env. Setup

### Clone app

```shell
$ git clone https://gitlab.com/adrianpk/boletus
```

### Create database user

If it does not exist yet.

```shell
$ psql
psql (11.5 (Ubuntu 11.5-1))
Type "help" for help.

user=# CREATE ROLE boletus;
user=# ALTER USER boletus WITH PASSWORD 'boletus';
```

Replace rolename by the database user owner.
Replace password by prefered password.

### Create database

```shell
user=# CREATE DATABASE boletus_dev OWNER boletus;
user=# CREATE DATABASE boletus_test OWNER boletus;
```

### Update run.sh script

Edit `scripts/run.sh`

Update values according to the preferred ones and / or those of your system.

### Run app

```shell
$ make clean-and-run
```

**You should see something like this**

```shell
2:56PM INF Cookie store key value=iVuOOv4PNBnqTk2o13JsBMOPcPAe4p18
2:56PM INF Reading template path=event/_ctxbar.tmpl
2:56PM INF Reading template path=event/_flash.tmpl
2:56PM INF Reading template path=event/_form.tmpl
2:56PM INF Reading template path=event/_header.tmpl
2:56PM INF Reading template path=event/_item.tmpl
2:56PM INF Reading template path=event/_list.tmpl
2:56PM INF Reading template path=event/edit.tmpl
2:56PM INF Reading template path=event/index.tmpl
2:56PM INF Reading template path=event/initdel.tmpl
2:56PM INF Reading template path=event/new.tmpl
2:56PM INF Reading template path=event/show.tmpl
2:56PM INF Reading template path=layout/base.tmpl
2:56PM INF Reading template path=user/_ctxbar.tmpl
2:56PM INF Reading template path=user/_flash.tmpl
2:56PM INF Reading template path=user/_form.tmpl
2:56PM INF Reading template path=user/_header.tmpl
2:56PM INF Reading template path=user/_item.tmpl
2:56PM INF Reading template path=user/_list.tmpl
2:56PM INF Reading template path=user/_signin.tmpl
2:56PM INF Reading template path=user/_signup.tmpl
2:56PM INF Reading template path=user/edit.tmpl
2:56PM INF Reading template path=user/index.tmpl
2:56PM INF Reading template path=user/initdel.tmpl
2:56PM INF Reading template path=user/new.tmpl
2:56PM INF Reading template path=user/show.tmpl
2:56PM INF Reading template path=user/signin.tmpl
2:56PM INF Reading template path=user/signup.tmpl
2:56PM INF Parsed template set path=event/edit.tmpl
2:56PM INF Parsed template set path=event/index.tmpl
2:56PM INF Parsed template set path=event/new.tmpl
2:56PM INF Parsed template set path=event/show.tmpl
2:56PM INF Parsed template set path=event/initdel.tmpl
2:56PM INF Parsed template set path=user/new.tmpl
2:56PM INF Parsed template set path=user/index.tmpl
2:56PM INF Parsed template set path=user/signup.tmpl
2:56PM INF Parsed template set path=user/edit.tmpl
2:56PM INF Parsed template set path=user/initdel.tmpl
2:56PM INF Parsed template set path=user/show.tmpl
2:56PM INF Parsed template set path=user/signin.tmpl
2:56PM INF Dialing to Postgres host="host=localhost port=5432 user=boletus password=boletus dbname=boletus_dev sslmode=disable"
2:56PM INF Postgres connection established
2:56PM INF New migrator name=migrator
2:56PM INF New seeder name=seeder
2:56PM INF New handler name=mailer
2020/01/19 14:56:14 CreateUsersTable
2020/01/19 14:56:14 Migration executed: CreateUsersTable
2020/01/19 14:56:14 CreateEventsTable
2020/01/19 14:56:14 Migration executed: CreateEventsTable
2020/01/19 14:56:14 CreateTicketsTable
2020/01/19 14:56:14 Migration executed: CreateTicketsTable
2020/01/19 14:56:14 Seed step executed: Users
2020/01/19 14:56:14 Seed step executed: EventAndTickets
2:56PM INF Scheduler started
2:56PM INF gRPC server initializing port=:8082
2:56PM INF Web server initializing port=:8080
2:56PM INF Currency rates updated base=EUR date=2020-01-17
2:57PM INF Expire tickets process init.
2:57PM INF Ticket reservation expired count=0
```

## Make commands

A brief summary of the most used commands

**make build**

Builds the application

**make run**

Run the application hrough a shell script that previously sets the environment variables with required values.
The shell script can be found in `scripts/run.sh`.

**package-resources**

It generates a binary representation html templates, translations and other resources that allows to embedd them within the compiled file. `clean-and-run` runs this make task as subtask before starting the applicacion.

**build-stg**

Build a staging Docker image of this application and pushes it to Docker Hub.

**build-prod**

Same as `make build-stage` but for production images.

**build-client**

Build the gRPC client. Source files location is `cmd/client/client.go` and compiled target is `bin/client`.

**install-stg**

Deploys app to Googke GKE usando [HELM](https://helm.sh/).
I haven't created helm .yaml files yet, so this command is not functional yet.

**install-prod**

Same as `make install-stage` but for production images.

**gen-proto**

Updates go files with changes in API made in the pprotocol buffers definition file at `pkg/grpc/api/v1`.

**make client**

Compiles the client and launches it through a shell script `scripts/client.sh` who sets some environment values required to use it. Now it is mostly used for manual testin but some flag based configuration will be implemented to make it more ergonomic to use it.

## Client

## Notes

* [To do list](docs/gtd/gtd.md)

