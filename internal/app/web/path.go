package web

import "html/template"

var pathFxs = template.FuncMap{
	// User
	"userPath":           UserPath,
	"userPathEdit":       UserPathEdit,
	"userPathSlug":       UserPathSlug,
	"userPathInitDelete": UserPathInitDelete,
	"userPathNew":        UserPathNew,

	// Event
	"eventPath":           EventPath,
	"eventPathEdit":       EventPathEdit,
	"eventPathSlug":       EventPathSlug,
	"eventPathInitDelete": EventPathInitDelete,
	"eventPathNew":        EventPathNew,
}
