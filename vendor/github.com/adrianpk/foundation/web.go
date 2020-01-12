package kabestan

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	// "github.com/davecgh/go-spew/spew"

	"github.com/gorilla/csrf"
	"github.com/gorilla/schema"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/markbates/pkger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type (
	WebEndpoint struct {
		Cfg        *Config
		Log        Logger
		templates  TemplateSet
		templateFx template.FuncMap
		//session    *sessions.Session
		store          *sessions.CookieStore
		storeKey       string
		secCookieStore *securecookie.SecureCookie
		i18n           *i18n.Bundle
		// More
		Service ServiceIF
		Mailer  Mailer
	}

	TemplateSet    map[string]*template.Template
	TemplateGroups map[string]map[string][]string
)

type FormAction struct {
	Target string
	Method string
}

type (
	ContextKey string
)

type (
	Localizer struct {
		*i18n.Localizer
	}
)

type (
	// WrappedRes stands for wrapped response
	WrappedRes struct {
		// Data stores model data
		Data interface{}
		// Stores model detailed validation errors.
		Errors ValErrorSet
		// Action can be used to reuse form templates letting change target and method from controller.
		Action FormAction
		// Localizer can be used to show msgID in differente languages
		Loc *Localizer
		// Flash messages
		Flash FlashSet
		// Cross-site request forgery protection
		CSRF map[string]interface{}
	}

	FlashSet []FlashItem

	// Flash data to present in page
	FlashItem struct {
		Msg  string
		Type MsgType
	}

	// MsgType stands for message type
	MsgType string
)

type (
	Identifiable interface {
		GetSlug() string
	}
)

const (
	GetMethod    = "GET"
	PostMethod   = "POST"
	PutMethod    = "PUT"
	PatchMethod  = "PATCH"
	DeleteMethod = "DELETE"
)

const (
	templateDir = "/assets/web/embed/template"
	layoutDir   = "layout"
	layoutKey   = "layout"
	pageKey     = "page"
	partialKey  = "partial"
)

const (
	NewTmpl     = "new.tmpl"
	IndexTmpl   = "index.tmpl"
	EditTmpl    = "edit.tmpl"
	ShowTmpl    = "show.tmpl"
	InitDelTmpl = "initdel.tmpl"
	SignUpTmpl  = "signup.tmpl"
	SignInTmpl  = "signin.tmpl"
)

const (
	secCookieName    = "sec-cookie"
	sessionCookieKey = "session"
)

const (
	InfoMT  MsgType = "info"
	WarnMT  MsgType = "warn"
	ErrorMT MsgType = "error"
	DebugMT MsgType = "debug"
)

var (
	InfoMTColor    = []string{"green-800", "white", "green-500", "green-800"}
	WarnMTColor    = []string{"yellow-800", "white", "yellow-500", "yellow-800"}
	ErrorMTColor   = []string{"red-800", "white", "red-500", "red-800"}
	DebugMTColor   = []string{"blue-800", "white", "blue-500", "blue-800"}
	DefaultMTColor = []string{"white", "white", "white", "white"}
)

const (
	FlashStoreKey   = "flash"
	SessionKey      = "session"
	SecureCookieKey = "sec-cookie"
)

const (
	I18NorCtxKey ContextKey = "i18n"
)

func MakeWebEndpoint(cfg *Config, log Logger, templateFx template.FuncMap) (*WebEndpoint, error) {
	registerGobTypes()

	ep := WebEndpoint{
		Cfg:        cfg,
		Log:        log,
		templateFx: templateFx,
	}

	// Cookie store
	ep.newCookieStore()

	// Secure Cookie store
	ep.newSecCookieStore()

	// Load
	ts, err := ep.loadTemplates()
	if err != nil {
		return &ep, err
	}

	// Classify
	tg := ep.classifyTemplates(ts)

	// Parse
	ep.parseTemplates(ts, tg)

	return &ep, nil
}

// registerGobTypes
func registerGobTypes() {
	gob.Register(FlashSet{})
}

func (ep *WebEndpoint) Templates() TemplateSet {
	return ep.templates
}

func (ep *WebEndpoint) TemplatesFx() template.FuncMap {
	return ep.templateFx
}

//func (ep *WebEndpoint) Session() *sessions.Session {
//return ep.session
//}

func (ep *WebEndpoint) Store() *sessions.CookieStore {
	return ep.store
}

func (wr *WrappedRes) addCSRF(r *http.Request) {
	wr.CSRF = map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	}
}

func (wr *WrappedRes) AllFlashes() FlashSet {
	fs := MakeFlashSet()

	for _, fi := range wr.Flash {
		fs = append(fs, fi)
	}

	return fs
}

func (l *Localizer) Localize(textID string) string {
	if l.Localizer != nil {
		t, _, err := l.LocalizeWithTag(&i18n.LocalizeConfig{
			MessageID: textID,
		})

		if err != nil {
			return fmt.Sprintf("%s", textID) // "'%s' [untransalted]", textID
		}

		return t
	}

	return fmt.Sprintf("%s", textID) // "'%s' [untransalted]", textID
}

// loadTemplates from embedded filesystem (pkger)
// under '/assets/web/embed/template'
func (ep *WebEndpoint) loadTemplates() (TemplateSet, error) {
	tmpls := make(TemplateSet)

	err := pkger.Walk(templateDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				ep.Log.Error(err, "msg", "Cannot load template", "path", path)
				return err
			}

			if filepath.Ext(path) == ".tmpl" {
				list := filepath.SplitList(path)
				base := fmt.Sprintf("%s:%s", list[0], templateDir)
				p, _ := filepath.Rel(base, path)

				ep.Log.Info("Reading template", "path", p)

				tmpls[p] = template.New(base)
				return nil
			}

			// ep.Log.Warn("Not a valid template", "path", path)

			return nil
		})

	if err != nil {
		ep.Log.Error(err, "msg", "Cannot load templates", "path")
		return tmpls, err
	}

	return tmpls, nil
}

// classifyTemplates grouping them,
// first by type (layout, partial and page)
// and then by resource.
func (ep *WebEndpoint) classifyTemplates(ts TemplateSet) TemplateGroups {
	all := make(TemplateGroups)
	last := ""
	keys := ep.tmplsKeys(ts)

	for _, path := range keys {
		p := "./assets/web/embed/template" + "/" + path

		//e.Log.Debug("Classifying", "path", path)

		fileDir := filepath.Dir(path)
		fileName := filepath.Base(path)

		if fileDir != last {

			if _, ok := all[fileDir]; !ok {
				all[fileDir] = make(map[string][]string)
			}

			if isValidTemplateFile(path) {
				if isPartial(fileName) {
					all[fileDir][partialKey] = append(all[fileDir][partialKey], p)

				} else if isLayout(fileDir) {
					all[layoutDir][layoutKey] = append(all[layoutDir][layoutKey], p)

				} else {
					all[fileDir][pageKey] = append(all[fileDir][pageKey], p)
				}
			}
		}
	}

	return all
}

func (ep *WebEndpoint) tmplsKeys(ts TemplateSet) []string {
	keys := make([]string, 0, len(ts))
	for k, _ := range ts {
		keys = append(keys, k)
	}
	return keys
}

// parseTemplates parses template sets for each resource.
func (ep *WebEndpoint) parseTemplates(ts TemplateSet, tg TemplateGroups) {
	ep.templates = make(TemplateSet)
	layout := tg[layoutDir][layoutKey][0]

	for k, ts := range tg {
		pages := ts[pageKey]
		partials := ts[partialKey]

		for _, t := range pages {
			if k != layoutDir {
				ep.parseTemplate(t, partials, layout, ep.templateFx)
			}
		}
	}
}

func (ep *WebEndpoint) parseTemplate(page string, partials []string, layout string, funcs template.FuncMap) {
	parse := "base.tmpl"
	all := make([]string, 10)
	all = append(all, page)
	all = append(all, partials...)
	all = append(all, layout)
	trimSlice(&all)

	t, err := template.New(parse).Funcs(funcs).ParseFiles(all...)
	if err != nil {
		ep.Log.Error(err, "Error parsing template set", "page", page)
	}

	base := fmt.Sprintf(".%s", templateDir)
	p, _ := filepath.Rel(base, page)

	ep.Log.Info("Parsed template set", "path", p)

	ep.templates[page] = t
}

func trimSlice(slice *[]string) {
	newSlice := make([]string, 0, len(*slice))
	for _, val := range *slice {
		switch val {
		case "":
		default:
			newSlice = append(newSlice, val)
		}
	}
	*slice = newSlice
}

func isValidTemplateFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".tmpl") && !strings.HasPrefix(fileName, ".")
}

func isPartial(fileName string) bool {
	return strings.HasPrefix(fileName, "_")
}

func isLayout(fileDir string) bool {
	return strings.HasPrefix(fileDir, "layout")
}

// Output
func (ep *WebEndpoint) writeResponse(w http.ResponseWriter, res interface{}) {
	// Marshalling
	o, err := ep.toJSON(res)
	if err != nil {
		ep.Log.Error(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(o)
}

func (ep *WebEndpoint) toJSON(res interface{}) ([]byte, error) {
	return json.Marshal(res)
}

func formatRequestBody(r *http.Request) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	return buf.String()
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	// If this is a POST, add post data
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}

// Sessions
func (ep *WebEndpoint) newSecCookieStore() {
	// Set secure cookie hash
	h := ep.Cfg.ValOrDef("web.seccookie.hash", "")
	if h == "" {
		h = ep.genRandomString(32)
		ep.Log.Debug("New secure cookie hash", "value", h)
		csEnVar := fmt.Sprintf("%s_SECCOOKIE_HASH", "BLT")
		ep.Log.Info("Set a custom secure cookie hash using a 32 char string stored as an envar", "envvar", csEnVar)
	}

	// Set secure cookie block
	b := ep.Cfg.ValOrDef("web.seccookie.block", "")
	if b == "" {
		b = ep.genRandomString(16)
		ep.Log.Debug("New secure cookie block", "value", b)
		csEnVar := fmt.Sprintf("%s_SECCOOKIE_BLOCK", "BLT")
		ep.Log.Info("Set a custom secure cookie block using a 16 char string stored as an envar", "envvar", csEnVar)
	}

	var hashKey = []byte(h)
	var blockKey = []byte(b)
	ep.secCookieStore = securecookie.New(hashKey, blockKey)
}

func (ep *WebEndpoint) newCookieStore() {
	k := ep.Cfg.ValOrDef("web.cookiestore.key", "")
	if k == "" {
		k = ep.genAES256Key()
		ep.Log.Debug("New cookie store random key", "value", k)
		csEnVar := fmt.Sprintf("%s_COOKIESTORE_KEY", "BLT")
		ep.Log.Info("Set a custom cookie store key using a 32 char string stored as an envar", "envvar", csEnVar)
	}

	ep.storeKey = k
	ep.Log.Debug("Cookie store key", "value", k)

	ep.store = sessions.NewCookieStore([]byte(k))
}

func (ep *WebEndpoint) genAES256Key() string {
	return ep.genRandomString(32)
}

func (ep *WebEndpoint) genRandomString(length int) string {
	const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		indexBits = 6                // 6 bits to represent a letter index
		indexMask = 1<<indexBits - 1 // All 1-bits, as many as letterIdxBits
		indexMax  = 63 / indexBits   // # of letter indices fitting in 63 bits
	)
	src := rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(32)
	for i, cache, remain := length-1, src.Int63(), indexMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), indexMax
		}
		if idx := int(cache & indexMask); idx < len(allowed) {
			sb.WriteByte(allowed[idx])
			i--
		}
		cache >>= indexBits
		remain--
	}

	return sb.String()
}

// Signin / signout

// SignIn creates a user session.
func (ep *WebEndpoint) SignIn(w http.ResponseWriter, r *http.Request, userSlug string) {
	ep.SetSecCookieVal(w, r, sessionCookieKey, userSlug)
}

func (ep *WebEndpoint) SignOut(w http.ResponseWriter, r *http.Request) {
	ep.ClearSecCookie(w, r)
}

func (ep *WebEndpoint) IsAuthenticated(r *http.Request) (slug string, ok bool) {
	return ep.ReadSecCookieVal(r, sessionCookieKey)
}

// Cookies
func (ep *WebEndpoint) GetSession(r *http.Request, name ...string) *sessions.Session {
	session := SessionKey
	if len(name) > 0 {
		session = name[0]
	}
	s, err := ep.Store().Get(r, session)
	if err != nil {
		ep.Log.Warn("Cannot get sesssion from store", "reqID", "n/a")
	}
	return s
}

func (ep *WebEndpoint) SetSecCookieVal(w http.ResponseWriter, r *http.Request, key, value string) {
	ep.Log.Debug("SetCookieVal begin")
	c, err := r.Cookie(secCookieName)
	if err != nil {
		ep.Log.Warn("No secure cookie present")
		c = &http.Cookie{
			Name: secCookieName,
			Path: "/",
			// TODO: Get domain from ep.Cfg
			// Domain: "127.0.0.1",
			// TODO: Dev -> false, Prod -> true
			Secure: false,
		}
	}

	vals := make(map[string]string)

	// Decode the cookie content
	if c.Value != "" {
		err = ep.secCookieStore.Decode(secCookieName, c.Value, &vals)
		if err != nil {
			ep.Log.Warn("Cannot decode current secure cookie")
		}
	}

	// Update cookie value
	delete(vals, key)

	if value != "" {
		vals[key] = value
	}

	// Encode values again
	e, err := ep.secCookieStore.Encode(secCookieName, vals)
	if err != nil {
		ep.Log.Warn("Cannot encode secure cookie")
		return
	}

	c.Value = e

	ep.Log.Info("Storing secure cookie", "vals", vals, "encrypted", e)

	http.SetCookie(w, c)
	ep.Log.Debug("SetCookieVal end")
}

func (ep *WebEndpoint) ReadSecCookieVal(r *http.Request, key string) (val string, ok bool) {
	c, err := r.Cookie(SecureCookieKey)
	if err != nil {
		ep.Log.Warn("No secure cookie")
		ep.Log.Error(err)
		return "", false
	}

	var vals map[string]string

	// Decode the cookie content
	err = ep.secCookieStore.Decode(secCookieName, c.Value, &vals)
	if err != nil {
		ep.Log.Error(err)
		return "", false
	}

	val, ok = vals[key]
	if !ok {
		ep.Log.Debug("No value stored in secure cookie", "key", key)
		return val, false
	}

	ep.Log.Debug("Retrieved from secure cookie", "key", key, "val", val)

	return val, true
}

func (ep *WebEndpoint) ClearSecCookie(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name: secCookieName,
		Path: "/",
		// TODO: Get domain from ep.Cfg
		// Domain: "127.0.0.1",
		// TODO: Dev -> false, Prod -> true
		MaxAge: -1,
		Secure: false,
	}
	http.SetCookie(w, c)
}

func (ep *WebEndpoint) DumpCookieValues(r *http.Request, cookieName, key string) (values string) {
	c, err := r.Cookie(cookieName)
	if err != nil {
		ep.Log.Error(err)
		return "n/a"
	}

	var vals map[string]string

	// Decode the cookie content
	err = ep.secCookieStore.Decode(cookieName, c.Value, &vals)
	if err != nil {
		ep.Log.Error(err)
		return "n/a"
	}

	return fmt.Sprintf("+v", vals)
}

// Templates

func (ep *WebEndpoint) TemplateFor(res, name string) (*template.Template, error) {
	key := ep.template(res, name)

	t, ok := ep.templates[key]
	if !ok {
		err := errors.New("canot get template")
		ep.Log.Error(err, "resource", res, "template", name)
		return nil, err
	}

	return t, nil
}

func (ep *WebEndpoint) template(resource, template string) (tmplKey string) {
	return fmt.Sprintf(".%s/%s/%s", templateDir, resource, template)
}

func getI18NLocalizer(r *http.Request) (localizer *i18n.Localizer, ok bool) {
	localizer, ok = r.Context().Value(I18NorCtxKey).(*i18n.Localizer)
	return localizer, ok
}

// Wrap response data.
func (ep *WebEndpoint) WrapRes(w http.ResponseWriter, r *http.Request, data interface{}, errors ValErrorSet) WrappedRes {
	f := MakeFlashSet()

	// Add pending messages
	f = f.AddItems(ep.RestoreFlash(r))

	wr := WrappedRes{
		Data:   data,
		Errors: errors,
		Loc:    ep.Localizer(r),
		Flash:  f,
	}

	wr.AddCSRF(r)

	ep.ClearFlash(w, r)

	return wr
}

func (ep *WebEndpoint) ErrorRedirect(w http.ResponseWriter, r *http.Request, redirPath, msgID string, err error) {
	m := ep.Localize(r, msgID)
	ep.RedirectWithFlash(w, r, redirPath, m, ErrorMT)
	ep.Log.Error(err)
}

// Localization - I18N
func (ep *WebEndpoint) Localize(r *http.Request, msgID string) string {
	l := ep.Localizer(r)
	if l == nil {
		ep.Log.Warn("No localizer available")
		return msgID
	}

	t, _, err := l.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID: msgID,
	})

	if err != nil {
		ep.Log.Error(err)
		return msgID
	}

	//s.Log.Debug("Localized message", "value", t, "lang", lang)

	return t
}

func (ep *WebEndpoint) localizeMessageID(l *i18n.Localizer, messageID string) (string, error) {
	return l.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
}

func (wr *WrappedRes) SetAction(fa FormAction) {
	wr.Action = fa
}

func (wr *WrappedRes) AddInfoFlash(infoMsg string) {
	f := wr.Flash

	m := strings.Trim(infoMsg, " ")
	if m != "" {
		f = f.AddItem(MakeFlashItem(m, InfoMT))
	}

	wr.Flash = f
}

func (wr *WrappedRes) AddWarnFlash(warnMsgs ...string) {
	f := wr.Flash

	if len(warnMsgs) > 0 {
		for _, m := range warnMsgs {
			f = f.AddItem(MakeFlashItem(m, WarnMT))
		}
	}

	wr.Flash = f
}

func (wr *WrappedRes) AddErrorFlash(errorMsg string) {
	f := wr.Flash

	m := strings.Trim(errorMsg, " ")
	if m != "" {
		f = f.AddItem(MakeFlashItem(m, ErrorMT))
	}

	wr.Flash = f
}

// MakeFlash message.
func MakeFlashItem(msg string, msgType MsgType) FlashItem {
	return FlashItem{
		Msg:  msg,
		Type: msgType,
	}
}

func (ep *WebEndpoint) StoreFlash(w http.ResponseWriter, r *http.Request, message string, mt MsgType) (ok bool) {
	s := ep.GetSession(r)

	// Append to current ones
	f := ep.RestoreFlash(r)
	f = append(f, MakeFlashItem(message, mt))

	s.Values[FlashStoreKey] = f
	err := s.Save(r, w)
	if err != nil {
		ep.Log.Error(err)
		return true
	}

	return false
}

func (ep *WebEndpoint) RestoreFlash(r *http.Request) FlashSet {
	s := ep.GetSession(r)
	v := s.Values[FlashStoreKey]

	f, ok := v.(FlashSet)
	if ok {
		//ep.Log.Debug("Stored flash", "value", spew.Sdump(f))
		return f
	}

	ep.Log.Info("No stored flash", "key", FlashStoreKey)
	return MakeFlashSet()
}

func (ep *WebEndpoint) ClearFlash(w http.ResponseWriter, r *http.Request) (ok bool) {
	s := ep.GetSession(r)
	delete(s.Values, FlashStoreKey)
	err := s.Save(r, w)
	if err != nil {
		return true
	}
	return false
}

// Redirect to url.
func (ep *WebEndpoint) Redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, 302)
}

func (ep *WebEndpoint) RedirectWithFlash(w http.ResponseWriter, r *http.Request, url string, msg string, msgType MsgType) {
	ep.StoreFlash(w, r, msg, msgType)
	http.Redirect(w, r, url, 303)
}

func (ep *WebEndpoint) Localizer(r *http.Request) *Localizer {
	l, ok := getI18NLocalizer(r)
	if !ok {
		return nil
	}

	return &Localizer{l}
}

func (wr *WrappedRes) addLocalizer(l *Localizer) {
	wr.Loc = l
}

func (wr *WrappedRes) AddCSRF(r *http.Request) {
	wr.CSRF = map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	}
}

func (fi *FlashItem) Color() []string {
	if fi.Type == InfoMT {
		return InfoMTColor

	} else if fi.Type == WarnMT {
		return WarnMTColor

	} else if fi.Type == ErrorMT {
		return ErrorMTColor

	} else if fi.Type == DebugMT {
		return DebugMTColor

	}
	return DefaultMTColor
}

// Forms
func (ep *WebEndpoint) FormToModel(r *http.Request, model interface{}) error {
	d := schema.NewDecoder()
	d.IgnoreUnknownKeys(true)
	return d.Decode(model, r.Form)
}

// NewDecoder build a schema decoder
// that put values from a map[string][]string into a struct.
func NewDecoder() *schema.Decoder {
	d := schema.NewDecoder()
	d.IgnoreUnknownKeys(true)
	return d
}

// Flash
func MakeFlashSet() FlashSet {
	return make(FlashSet, 0)
}

func (f FlashSet) IsEmpty() bool {
	return len(f) == 0
}

func (f FlashSet) AddItem(fi FlashItem) FlashSet {
	return append(f, fi)
}

func (f FlashSet) AddItems(fis []FlashItem) FlashSet {
	return append(f, fis...)
}

func (fi FlashItem) IsEmpty() bool {
	return fi == FlashItem{}
}

// Resource paths

// Resource path functions
// IndexPath returns index path under resource root path.
func IndexPath() string {
	return ""
}

// EditPath returns edit path under resource root path.
func EditPath() string {
	return "/{id}/edit"
}

// NewPath returns new path under resource root path.
func NewPath() string {
	return "/new"
}

// ShowPath returns show path under resource root path.
func ShowPath() string {
	return "/{id}"
}

// CreatePath returns create path under resource root path.
func CreatePath() string {
	return ""
}

// UpdatePath returns update path under resource root path.
func UpdatePath() string {
	return "/{id}"
}

// InitDeletePath returns init delete path under resource root path.
func InitDeletePath() string {
	return "/{id}/init-delete"
}

// DeletePath returns delete path under resource root path.
func DeletePath() string {
	return "/{id}"
}

// SignupPath returns signup path.
func SignupPath() string {
	return "/signup"
}

// LoginPath returns login path.
func LoginPath() string {
	return "/login"
}

// ResPath
func ResPath(rootPath string) string {
	return "/" + rootPath + IndexPath()
}

// ResPathEdit
func ResPathEdit(rootPath string, r Identifiable) string {
	return fmt.Sprintf("/%s/%s/edit", rootPath, r.GetSlug())
}

// ResPathNew
func ResPathNew(rootPath string) string {
	return fmt.Sprintf("/%s/new", rootPath)
}

// ResPathInitDelete
func ResPathInitDelete(rootPath string, r Identifiable) string {
	return fmt.Sprintf("/%s/%s/init-delete", rootPath, r.GetSlug())
}

// ResPathSlug
func ResPathSlug(rootPath string, r Identifiable) string {
	return fmt.Sprintf("/%s/%s", rootPath, r.GetSlug())
}
