package web

import (
	fnd "github.com/adrianpk/foundation"
)

// EventRoot - Event resource root path.
var EventRoot = "events"

// EventPath
func EventPath() string {
	return fnd.ResPath(EventRoot)
}

func EventAdminPath() string {
	return fnd.ResPath(EventRoot)
}

// EventPathEdit
func EventPathEdit(res fnd.Identifiable) string {
	return fnd.ResPathEdit(EventRoot, res)
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
