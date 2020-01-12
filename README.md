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

## Routes

### Auth

| Method | Path          | Handler    |
|--------|---------------|------------|
| GET    | /auth/signup  | InitSignUp |
| POST   | /auth/signup  | SignUp     |
| GET    | /auth/signin  | InitSignIn |
| POST   | /auth/signin  | SignIn     |
| GET    | /auth/signout | SignOut    |

### User

| Method | Path                          | Handler    |
|--------|-------------------------------|------------|
| GET    | /users                        | Index      |
| GET    | /users/new                    | New        |
| POST   | /users                        | Create     |
| GET    | /users/{slug}                 | Show       |
| GET    | /users/{slug}/edit            | Edit       |
| PUT    | /users/{slug}                 | Update     |
| PATCH  | /users/{slug}                 | Update     |
| POST   | /users/{slug}/init-delete     | InitDelete |
| DELETE | /users/{slug}                 | Delete     |
| GET    | /users/{slug}/{token}/confirm | Confirm    |


## Dev. Env. Setup

After stabilizing this base app a (still not published) generator will be updated to automate most of these steps.

### Clone app

```shell
$ git clone https://gitlab.com/foundation/repo/boletus appname
```

Replace appname by the name of your app.

## Create database user

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
user=# CREATE DATABASE boletus OWNER boletus;
user=# CREATE DATABASE dbname_test OWNER boletus;
```

### Update run.sh script

Edit `scripts/run.sh`

Config system uses envar prefixes to set app configuration values.
By default this value is `BLT` but you can replace it with any other.

```shell
# Service
export BLT_SVC_NAME="boletus"
export BLT_SVC_REVISION=$REV
export BLT_(...)
```

Edit other values according to the preferred ones and / or those of your system.

### Edit main

Edit `cmd/appname`

If you change this envvar prefix from "BLT" to, lets say, "APP"

```go
  (...)
	cfg := fnd.LoadConfig("blt") // <- change this
	// cfg := fnd.LoadConfig("app") // <- to something like this
  (...)
```

