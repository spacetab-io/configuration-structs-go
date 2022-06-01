package cfgstructs

import (
	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type LogFormat string

func (f LogFormat) String() string {
	return string(f)
}

const (
	LogFormatText     = "text"
	LogFormatJSON     = "json"
	LogsFormatDefault = LogFormatText
)

type Logs struct {
	Level   string        `yaml:"level" valid:"required,in(debug|info|warn|error|dpanic|panic|fatal)"`
	Format  string        `yaml:"format" valid:"required,in(text|json)"`
	Colored bool          `yaml:"isColored"`
	Caller  CallerConfig  `yaml:"caller" valid:"required"`
	Sentry  *SentryConfig `yaml:"sentry,omitempty"`
}

func (l Logs) String() string {
	return "logs"
}

func (l Logs) SentryDebugEnabled() bool {
	if l.Sentry == nil {
		return false
	}

	return l.Sentry.Debug
}

func (l Logs) IsSentryEnabled() bool {
	if l.Sentry == nil {
		return false
	}

	return l.Sentry.Enable
}

func (l Logs) GetSentryDSN() string {
	if l.Sentry == nil {
		return ""
	}

	return l.Sentry.DSN
}

func (l *Logs) SetCaller(isDisabled bool, skipFrames int) {
	l.Caller.Show = isDisabled
	l.Caller.SkipFrames = skipFrames
}

func (l Logs) ShowCaller() bool {
	return l.Caller.Show
}

func (l Logs) GetCallerSkipFrames() int {
	return l.Caller.SkipFrames
}

func (l Logs) GetLevel() string {
	return l.Level
}

func (l Logs) GetFormat() string {
	if l.Format == "" {
		return LogsFormatDefault
	}

	return l.Format
}

func (l *Logs) SetFormat(format string) {
	l.Format = format
}

func (l Logs) IsColored() bool {
	return l.Colored
}

func (l Logs) GetSentryParams() (bool, bool, string) {
	return l.Sentry.Enable, l.Sentry.Debug, l.Sentry.DSN
}

func (l Logs) GetCallerParams() (bool, int) {
	return l.Caller.Show, l.Caller.SkipFrames
}

func (l Logs) Validate() (bool, error) {
	return contracts.ConfigValidate(l)
}

type CallerConfig struct {
	Show       bool `yaml:"show"`
	SkipFrames int  `yaml:"skip_frames"`
}

type SentryConfig struct {
	Enable bool   `yaml:"enable"`
	Debug  bool   `yaml:"debug"`
	DSN    string `yaml:"dsn"`
}
