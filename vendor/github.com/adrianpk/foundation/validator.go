package kabestan

import (
	"strings"
	"unicode/utf8"
)

// NOTE: Experiment, preferring to avoid reflection
// at the expense of greater verbosity.
// The cli generator will have to do the heavy lifting.

type (
	//Changes map[string][]string

	ValErrorSet map[string][]string

	Validator struct {
		Errors ValErrorSet
	}
)

var ValMsg = newValMsg()

func NewValidator() Validator {
	return Validator{
		ValErrorSet(map[string][]string{}),
	}
}

func newValMsg() *valMsg {
	return &valMsg{
		RequiredErrMsg:   "required",
		MinLengthErrMsg:  "too short",
		MaxLengthErrMsg:  "too long",
		NotAllowedErrMsg: "not in allowed list",
		NotEmailErrMsg:   "not an email address",
		NoMatchErrMsg:    "confirmation does not match",
	}
}

type valMsg struct {
	RequiredErrMsg   string
	MinLengthErrMsg  string
	MaxLengthErrMsg  string
	NotAllowedErrMsg string
	NotEmailErrMsg   string
	NoMatchErrMsg    string
}

// ValidateRequired value.
func (v *Validator) ValidateRequired(val string, errMsg ...string) (ok bool) {
	val = strings.Trim(val, " ")
	return utf8.RuneCountInString(val) > 0
}

// ValidateMinLength value.
func (v *Validator) ValidateMinLength(val string, min int, errMsg ...string) (ok bool) {
	return utf8.RuneCountInString(val) >= min
}

// ValidateMaxLength value.
func (v *Validator) ValidateMaxLength(val string, max int) (ok bool) {
	return utf8.RuneCountInString(val) <= max
}

// ValidateEmail value.
func (v *Validator) ValidateEmail(val string) (ok bool) {
	return len(val) < 254 && emailRegex.MatchString(val)
}

// ValidateConfirmation value.
func (v *Validator) ValidateConfirmation(val, confirmation string) (ok bool) {
	return val == confirmation
}

func (v *Validator) HasErrors() bool {
	return !v.IsValid()
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (es ValErrorSet) Add(field, msg string) {
	es[field] = append(es[field], msg)
}

func (es ValErrorSet) FieldErrors(field string) []string {
	return es[field]
}

func (es ValErrorSet) IsEmpty() bool {
	return len(es) < 1
}
