package web

import (
	"fmt"
	"net/http"

	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
)

const (
	eventRes = "event"
)

const (
	// Defined in 'assets/web/embed/i18n/xx.json'
	EventCreatedInfoMsg = "event_created_info_msg"
	EventUpdatedInfoMsg = "event_updated_info_msg"
	EventDeletedInfoMsg = "event_deleted_info_msg"
	// Error
	CreateEventErrMsg = "create_event_err_msg"
	IndexEventsErrMsg = "index_events_err_msg"
	GetEventErrMsg    = "get_event_err_msg"
	GetEventsErrMsg   = "get_events_err_msg"
	UpdateEventErrMsg = "update_event_err_msg"
	DeleteEventErrMsg = "delete_event_err_msg"
)

// IndexEvents web endpoint.
func (ep *Endpoint) IndexEvents(w http.ResponseWriter, r *http.Request) {
	// Get events list from registered service
	events, err := ep.Service.IndexEvents()
	if err != nil {
		ep.ErrorRedirect(w, r, "/", IndexEventsErrMsg, err)
		return
	}

	// Convert result list into a form list
	// Models use sql null types but templates looks
	// clearer if we use plain Go type.
	// i.e.: $event.Eventname instead of $event.Eventname.String
	l := model.ToEventFormList(events)
	wr := ep.WrapRes(w, r, l, nil)

	// Get template to render from cache.
	ts, err := ep.TemplateFor(eventRes, fnd.IndexTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, "/", CannotProcErrMsg, err)
		return
	}

	// Execute it and redirect if error.
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, "/", CannotProcErrMsg, err)
		return
	}
}

func (ep *Endpoint) NewEvent(w http.ResponseWriter, r *http.Request) {
	eventForm := model.EventForm{IsNew: true}

	// Wrap response
	wr := ep.WrapRes(w, r, &eventForm, nil)
	wr.SetAction(eventCreateAction())

	// Get template to render from cache.
	ts, err := ep.TemplateFor(eventRes, fnd.NewTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	// Write response
	// Execute it and redirect if error.
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}
}

func (ep *Endpoint) CreateEvent(w http.ResponseWriter, r *http.Request) {
	// Decode request data into a form.
	eventForm := model.EventForm{}
	err := ep.FormToModel(r, &eventForm)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	// Create a model using form values.
	event := eventForm.ToModel()

	// Use registered service to do everything related
	// to event creation.
	ves, err := ep.Service.CreateEvent(&event)

	// First take care of service validation errors.
	if !ves.IsEmpty() {
		ep.rerenderEventForm(w, r, event.ToForm(), ves, fnd.NewTmpl, eventCreateAction())
		return
	}

	// Then take care of other kind of possible errors
	// that service can generate.
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	// Localize Ok info message, put it into a flash message
	// and redirect to index.
	m := ep.Localize(r, EventCreatedInfoMsg)
	ep.RedirectWithFlash(w, r, EventPath(), m, fnd.InfoMT)
}

// ShowEvent web endpoint.
func (ep *Endpoint) ShowEvent(w http.ResponseWriter, r *http.Request) {
	// Get slug from request context.
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	// Use registered service to do everything related
	// to event creation.
	event, err := ep.Service.GetEvent(s)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), GetEventErrMsg, err)
		return
	}

	// Wrap response
	wr := ep.WrapRes(w, r, event.ToForm(), nil)

	// Template
	ts, err := ep.TemplateFor(eventRes, fnd.ShowTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	// Write response
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}
}

// EditEvent web endpoint.
func (ep *Endpoint) EditEvent(w http.ResponseWriter, r *http.Request) {
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	// Use registerd service to get the event from repo.
	event, err := ep.Service.GetEvent(s)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), GetEventErrMsg, err)
		return
	}

	// Wrap response
	eventForm := event.ToForm()
	wr := ep.WrapRes(w, r, &eventForm, nil)
	wr.SetAction(eventUpdateAction(&eventForm))

	// Template
	ts, err := ep.TemplateFor(eventRes, fnd.EditTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	// Write response
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}
}

// UpdateEvent web endpoint.
func (ep *Endpoint) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), GetEventErrMsg, err)
		return
	}

	// Decode request data into a form.
	eventForm := model.EventForm{}
	err = ep.FormToModel(r, &eventForm)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	// Create a model using form values.
	event := eventForm.ToModel()

	// Use registered service to do everything related
	// to event update.
	ves, err := ep.Service.UpdateEvent(s, &event)

	// First take care of service validation errors.
	if !ves.IsEmpty() {
		ep.Log.Debug("Validation errors", "dump", fmt.Sprintf("%+v", ves.FieldErrors))
		ep.rerenderEventForm(w, r, event.ToForm(), ves, fnd.NewTmpl, eventCreateAction())
		return
	}

	// Non validation errors
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), UpdateEventErrMsg, err)
		return
	}

	m := ep.Localize(r, EventUpdatedInfoMsg)
	ep.RedirectWithFlash(w, r, EventPath(), m, fnd.InfoMT)
}

// InitDeleteEvent web endpoint.
func (ep *Endpoint) InitDeleteEvent(w http.ResponseWriter, r *http.Request) {
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	// Use registerd service to get the event from repo.
	event, err := ep.Service.GetEvent(s)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), GetEventsErrMsg, err)
		return
	}

	// Wrap response
	eventForm := event.ToForm()
	wr := ep.WrapRes(w, r, &eventForm, nil)
	wr.SetAction(eventDeleteAction(&eventForm))

	// Template
	ts, err := ep.TemplateFor(eventRes, fnd.InitDelTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	// Write response
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}
}

// DeleteEvent web endpoint.
func (ep *Endpoint) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), DeleteEventErrMsg, err)
		return
	}

	// Service
	err = ep.Service.DeleteEvent(s)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), DeleteEventErrMsg, err)
		return
	}

	m := ep.Localize(r, EventDeletedInfoMsg)
	ep.RedirectWithFlash(w, r, EventPath(), m, fnd.InfoMT)
}

func (ep *Endpoint) rerenderEventForm(w http.ResponseWriter, r *http.Request, data interface{}, valErrors fnd.ValErrorSet, template string, action fnd.FormAction) {
	wr := ep.WrapRes(w, r, data, valErrors)
	wr.AddErrorFlash(InputValuesErrMsg)
	wr.SetAction(action)

	ts, err := ep.TemplateFor(eventRes, template)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), InputValuesErrMsg, err)
		return
	}

	// Write response
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, EventPath(), CannotProcErrMsg, err)
		return
	}

	return
}

// Misc
// eventCreateAction
func eventCreateAction() fnd.FormAction {
	return fnd.FormAction{Target: fmt.Sprintf("%s", EventAdminPath()), Method: "POST"}
}

// eventUpdateAction
func eventUpdateAction(model fnd.Identifiable) fnd.FormAction {
	return fnd.FormAction{Target: EventAdminPathSlug(model), Method: "PUT"}
}

// eventDeleteAction
func eventDeleteAction(model fnd.Identifiable) fnd.FormAction {
	return fnd.FormAction{Target: EventAdminPathSlug(model), Method: "DELETE"}
}
