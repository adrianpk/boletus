# Routes

## Auth

| Method | Path          | Handler    |
|--------|---------------|------------|
| GET    | /auth/signup  | InitSignUp |
| POST   | /auth/signup  | SignUp     |
| GET    | /auth/signin  | InitSignIn |
| POST   | /auth/signin  | SignIn     |
| GET    | /auth/signout | SignOut    |

## User

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


## Event (*)

| Method | Path                           | Handler    |
|--------|--------------------------------|------------|
| GET    | /events                        | Index      |
| GET    | /events/new                    | New        |
| POST   | /events                        | Create     |
| GET    | /events/{slug}                 | Show       |
| GET    | /events/{slug}/edit            | Edit       |
| PUT    | /events/{slug}                 | Update     |
| PATCH  | /events/{slug}                 | Update     |
| POST   | /events/{slug}/init-delete     | InitDelete |
| DELETE | /events/{slug}                 | Delete     |

* Partially implemented
