package bugsnag

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type reportJSON struct {
	APIKey   string       `json:"apiKey"`
	Events   []eventJSON  `json:"events"`
	Notifier notifierJSON `json:"notifier"`
}

type notifierJSON struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Version string `json:"version"`
}

type eventJSON struct {
	App            *appJSON            `json:"app"`
	Context        string              `json:"context,omitempty"`
	Device         *deviceJSON         `json:"device,omitempty"`
	Exceptions     []exceptionJSON     `json:"exceptions"`
	GroupingHash   string              `json:"groupingHash,omitempty"`
	Metadata       interface{}         `json:"metaData"`
	PayloadVersion string              `json:"payloadVersion"`
	Session        *sessionJSON        `json:"session,omitempty"`
	Severity       string              `json:"severity"`
	SeverityReason *severityReasonJSON `json:"severityReason,omitempty"`
	Unhandled      bool                `json:"unhandled"`
	User           *User               `json:"user,omitempty"`
}

type sessionJSON struct {
	StartedAt time.Time       `json:"startedAt"`
	ID        uuid.UUID       `json:"id"`
	Events    eventCountsJSON `json:"events"`
}

type eventCountsJSON struct {
	Handled   int `json:"handled"`
	Unhandled int `json:"unhandled"`
}

type appJSON struct {
	ReleaseStage string `json:"releaseStage"`
	Type         string `json:"type,omitempty"`
	Version      string `json:"version,omitempty"`
}

type exceptionJSON struct {
	ErrorClass string       `json:"errorClass"`
	Message    string       `json:"message"`
	Stacktrace []stackFrame `json:"stacktrace"`
}

type severityAttributesJSON struct {
	Framework string `json:"framework,omitempty"`
}

type severityReasonJSON struct {
	Attributes *severityAttributesJSON `json:"attributes,omitempty"`
	Type       SeverityReason          `json:"type,omitempty"`
}

type deviceJSON struct {
	Hostname string `json:"hostname,omitempty"`
}
