package web

import (
	fnd "github.com/adrianpk/foundation"
)

// AdminRoot - Admin root path.
var AdminRoot = "admin"

// EventRoot - Event resource root path.
var EventRoot = "events"

// EventPath
func EventPath() string {
	return fnd.ResPath(EventRoot)
}

func EventAdminPath() string {
	return fnd.ResAdmin(fnd.ResPath(EventRoot), AdminRoot)
}

// EventPathEdit
func EventPathEdit(res fnd.Identifiable) string {
	return fnd.ResAdmin(fnd.ResPathEdit(EventRoot, res), AdminRoot)
}

// EventPathNew
func EventPathNew() string {
	return fnd.ResPathNew(EventRoot)
}

// EventPathInitDelete
func EventPathInitDelete(res fnd.Identifiable) string {
	return fnd.ResPathInitDelete(EventRoot, res)
}

// EventPathSlug
func EventPathSlug(res fnd.Identifiable) string {
	return fnd.ResPathSlug(EventRoot, res)
}

func EventAdminPathSlug(res fnd.Identifiable) string {
	return fnd.ResAdmin(fnd.ResPathSlug(EventRoot, res), AdminRoot)
}
