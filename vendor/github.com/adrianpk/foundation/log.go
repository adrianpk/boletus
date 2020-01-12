package kabestan

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// Disabled level
	DisabledLevel = -1
	// Debug level
	DebugLevel = iota
	// Info level
	InfoLevel
	// Warn level
	WarnLevel
	// Error level
	ErrorLevel
)

const (
	loggerCtxKey contextKey = "logger"
)

var (
	// Default is the package default logger.
	// It can be used right out of the box.
	// It can be replaced by a custom configured one
	// using package Set(*Logger) function
	// or using *Logger.Set() method.
	cfg LogConfig
)

// Logger interface
type Logger interface {
	Debug(meta ...interface{})
	Info(meta ...interface{})
	Warn(meta ...interface{})
	Error(err error, meta ...interface{})
}

// Log is structured leveled logger
type Log struct {
	// Level of min logging
	Level int
	// Version
	Version string
	// Revision
	Revision string
	// DebugLog logger
	StdLog zerolog.Logger
	// ErrorLog logger
	ErrLog zerolog.Logger
	// Dynamic fields
	dynafields []interface{}
}

type LogConfig struct {
	// Name
	name string
	// Level of min logging
	level int
	// Static fields
	stfields []interface{}
	// configured
	configured bool
}

type contextKey string

func NewLogger(cfg *Config) *Log {
	ll := int(cfg.ValAsInt("log.level", 1))
	sn := cfg.ValOrDef("app.name", "kabestan")
	sr := cfg.ValOrDef("app.revision", "n/a")
	//return NewLogger(ll, sn, sr)
	return NewDevLogger(ll, sn, sr)
}

// String is human readable representation of a context key.
func (c contextKey) String() string {
	return "mw-" + string(c)
}

// setup name and static fields.
// Each new instance of logger will always append these
// key-value pairs to the output and name if it is not empty.
// These values cannot be modified after they are configured.
func setup(name string, stfields []interface{}) {
	if cfg.configured {
		return
	}
	cfg.name = name
	cfg.stfields = append(cfg.stfields, stfields...)
	cfg.configured = true
}

// SetDyna fields.
// The receiver instance will always append these
// key-value pairs to the output.
func (l *Log) SetDyna(dynafields ...interface{}) {
	l.dynafields = make([]interface{}, 2)
	l.dynafields = append(l.dynafields, dynafields...)
}

// AddDyna fields.
// The receiver instance will always append these
// key-value pairs to the output.
func (l *Log) AddDyna(key, value interface{}) {
	l.dynafields = append(l.dynafields, []interface{}{key, value})
}

// ResetDyna fields.
// Remove dynamic fields.
func (l *Log) ResetDyna() {
	l.dynafields = make([]interface{}, 2)
}

// CtxLogger returns a logger stored in the context provided by the argument.
// We want to avoid runtime errors, then, to get the logger from the context
// the package uses an unexported key to store and retrive it.
func CtxLogger(ctx context.Context) (logger *Logger, ok bool) {
	l, ok := ctx.Value(loggerCtxKey).(*Logger)
	return l, ok
}

// NewLogger logger.
// If static fields are provided those values will define
// the default static fields for each new built instance
// if they were not yet configured.
func newLogger(level int, name string, stfields ...interface{}) *Log {
	if level < DisabledLevel || level > ErrorLevel {
		level = InfoLevel
	}

	stdl := zerolog.New(os.Stdout).With().Timestamp().Logger()
	errl := zerolog.New(os.Stderr).With().Timestamp().Logger()

	setLogLevel(&stdl, level)
	setLogLevel(&errl, level)

	l := &Log{
		Level:  level,
		StdLog: stdl,
		ErrLog: errl,
	}

	if len(stfields) > 1 && !cfg.configured {
		setup(name, stfields)
	}

	return l
}

// NewDevLogger logger.
// Pretty logging for development mode.
// Not recommended for production use.
// If static fields are provided those values will define
// the default static fields for each new built instance
// if they were not yet configured.
func NewDevLogger(level int, name string, stfields ...interface{}) *Log {
	if level < DisabledLevel || level > ErrorLevel {
		level = InfoLevel
	}

	stdl := log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	errl := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	setLogLevel(&stdl, level)
	setLogLevel(&errl, level)

	l := &Log{
		Level:  level,
		StdLog: stdl,
		ErrLog: errl,
	}

	if len(stfields) > 1 && !cfg.configured {
		setup(name, stfields)
	}

	return l
}

// InCtx returns a copy of context that also includes a configured logger.
func InCtx(ctx context.Context, fields ...string) context.Context {
	l, _ := FromCtx(ctx)
	if len(fields) > 0 {
		l.SetDyna(fields)
	}
	return context.WithValue(ctx, loggerCtxKey, l)
}

// FromCtx returns current logger in context.
// If there is not logger in context it returns
// a new one with current config values.
func FromCtx(ctx context.Context) (log *Log, fresh bool) {
	l, ok := ctx.Value(loggerCtxKey).(Log)
	if !ok {
		return newLogger(cfg.level, cfg.name), true
	}
	return &l, false
}

// Debug logs debug messages.
func (l Log) Debug(meta ...interface{}) {
	if len(meta) > 0 {
		l.debugf(stringify(meta[0]), meta[1:len(meta)])
	}
}

// Info logs info messages.
func (l Log) Info(meta ...interface{}) {
	if len(meta) > 0 {
		l.infof(stringify(meta[0]), meta[1:len(meta)])
	}
}

// Warn logs warning messages.
func (l Log) Warn(meta ...interface{}) {
	if len(meta) > 0 {
		l.warnf(stringify(meta[0]), meta[1:len(meta)])
	}
}

// Error logs error messages.
func (l Log) Error(err error, meta ...interface{}) {
	if len(meta) > 0 {
		l.errorf(err, stringify(meta[0]), meta[1:len(meta)])
		return
	}
	l.errorf(err, "", nil)
}

func (l Log) debugf(message string, fields []interface{}) {
	if l.Level > DebugLevel {
		return
	}
	le := l.StdLog.Info()
	appendKeyValues(le, l.dynafields, fields)
	le.Msg(message)
}

func (l Log) infof(message string, fields []interface{}) {
	if l.Level > InfoLevel {
		return
	}
	le := l.StdLog.Info()
	appendKeyValues(le, l.dynafields, fields)
	le.Msg(message)
}

func (l Log) warnf(message string, fields []interface{}) {
	if l.Level > WarnLevel {
		return
	}
	le := l.StdLog.Info()
	appendKeyValues(le, l.dynafields, fields)
	le.Msg(message)
}

func (l Log) errorf(err error, message string, fields []interface{}) {
	le := l.ErrLog.Error()
	appendKeyValues(le, l.dynafields, fields)
	le.Err(err)
	le.Msg(message)
}

// TODO: Optimize.
// Static key-value calculation shoud be cached.
// Dynamic key-value calculation shoud be cached if didn't changed.
func appendKeyValues(le *zerolog.Event, dynafields []interface{}, fields []interface{}) {
	if cfg.name != "" {
		le.Str("name", cfg.name)
	}

	fs := make(map[string]interface{})

	if len(fields) > 1 {
		for i := 0; i < len(fields)-1; i++ {
			if fields[i] != nil && fields[i+1] != nil {
				k := stringify(fields[i])
				fs[k] = fields[i+1]
				// fmt.Printf("field - (%s, %v)\n", k, fs[k])
				i++
			}
		}

		if len(dynafields) > 1 {
			// fs := make(map[string]interface{})
			for i := 0; i < len(dynafields)-1; i++ {
				if dynafields[i] != nil && dynafields[i+1] != nil {
					k := stringify(dynafields[i])
					fs[k] = dynafields[i+1]
					// fmt.Printf("dyna - (%s, %v)\n", k, fs[k])
					i++
				}
			}
		}

		if len(cfg.stfields) > 1 {
			for i := 0; i < len(cfg.stfields)-1; i++ {
				if cfg.stfields[i] != nil && cfg.stfields[i+1] != nil {
					k := stringify(cfg.stfields[i])
					fs[k] = cfg.stfields[i+1]
					// fmt.Printf("static - (%s, %v)\n", k, fs[k])
					i++
				}
			}
		}

	}
	le.Fields(fs)
}

func stringify(val interface{}) string {
	switch v := val.(type) {
	case nil:
		return fmt.Sprintf("%v", v)
	case int:
		return fmt.Sprintf("%d", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case string:
		return v
	default:
		return fmt.Sprintf("%+v", v)
	}
}

// UpdateLogLevel updates log level.
func (l Log) UpdateLogLevel(level int) {
	// Allow info level to log the update
	// But don't downgrade to it if Error is set.
	current := ErrorLevel
	l.Info("Log level updated", "", "log level", level)
	l.Level = current
	if level < DisabledLevel || level > ErrorLevel {
		l.Level = level
		setLogLevel(&l.StdLog, level)
		setLogLevel(&l.ErrLog, level)
	}
}
func setLogLevel(l *zerolog.Logger, level int) {
	switch level {
	case -1:
		l.Level(zerolog.Disabled)
	case 0:
		l.Level(zerolog.DebugLevel)
	case 1:
		l.Level(zerolog.InfoLevel)
	case 2:
		l.Level(zerolog.WarnLevel)
	case 3:
		l.Level(zerolog.ErrorLevel)
	default:
		l.Level(zerolog.DebugLevel)
	}
}
