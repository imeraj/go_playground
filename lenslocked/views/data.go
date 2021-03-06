package views

import (
	"github.com/imeraj/go_playground/lenslocked/models"
)

const (
	AlertLvlError   = "danger"
	AlertLvlWarning = "warning"
	AlertLvlInfo    = "info"
	AlertLvlSuccess = "success"

	// AlertMsgGeneric is displayed when any random error
	// is encountered by our backend.
	AlertMsgGeneric = "Something went wrong. Please try " +
		"again, and contact us if the problem persists."
)

// Data is the top level structure that views expect data
// to come in.
type Data struct {
	Alert *Alert
	User  *models.User
	Yield interface{}
}

func (d *Data) AlertError(msg string) {
	d.Alert = &Alert{
		Level:   AlertLvlError,
		Message: msg,
	}
}

func (d *Data) SetAlert(msg string, level string) {
	d.Alert = &Alert{
		Level:   level,
		Message: msg,
	}
}

// Alert is used to render Bootstrap Alert messages in templates
type Alert struct {
	Level   string
	Message string
}

type PublicError interface {
	error
	Public() string
}
