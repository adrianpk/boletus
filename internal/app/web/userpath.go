package web

import (
	fnd "github.com/adrianpk/foundation"
)

// AuthRoot - Auth root path.
var AuthRoot = "auth"

// UserRoot - User resource root path.
var UserRoot = "users"

func AuthPath() string {
	return fnd.ResPath(AuthRoot)
}

// AuthPathSignUp
func AuthPathSignUp() string {
	return fnd.ResPath(AuthRoot) + "/signup"
}

// AuthPathSignIn
func AuthPathSignIn() string {
	return fnd.ResPath(AuthRoot) + "/signin"
}

// UserPath
func UserPath() string {
	return fnd.ResPath(UserRoot)
}

// UserPathEdit
func UserPathEdit(res fnd.Identifiable) string {
	// TODO: Analize if in a multi-tenant setup this could be
	// a problem.
	return fnd.ResPathEdit(UserRoot, res)
	//return fmt.Sprintf("/%s/%s/edit", UserRoot, res.U)
}

// UserPathNew
func UserPathNew() string {
	return fnd.ResPathNew(UserRoot)
}

// UserPathInitDelete
func UserPathInitDelete(res fnd.Identifiable) string {
	return fnd.ResPathInitDelete(UserRoot, res)
}

// UserPathSlug
func UserPathSlug(res fnd.Identifiable) string {
	return fnd.ResPathSlug(UserRoot, res)
}
