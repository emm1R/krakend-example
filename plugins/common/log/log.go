package log

import "fmt"

// PluginLogger holds the logger object and the plugin name
type PluginLogger struct {
	Log  Logger
	Name string
}

// Logger is an interface for leveraging KrakenD's internal logging
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}

// NewLogger returns a new named logger
func NewLogger(logger Logger, name string) *PluginLogger {
	return &PluginLogger{
		Log:  logger,
		Name: fmt.Sprintf("[PLUGIN: %s]", name),
	}
}

// plain loggers

func (l *PluginLogger) Debug(v ...interface{}) {
	l.Log.Debug(l.Name, v)
}

func (l *PluginLogger) Info(v ...interface{}) {
	l.Log.Info(l.Name, v)
}

func (l *PluginLogger) Warning(v ...interface{}) {
	l.Log.Warning(l.Name, v)
}

func (l *PluginLogger) Error(v ...interface{}) {
	l.Log.Error(l.Name, v)
}

func (l *PluginLogger) Critical(v ...interface{}) {
	l.Log.Critical(l.Name, v)
}

func (l *PluginLogger) Fatal(v ...interface{}) {
	l.Log.Fatal(l.Name, v)
}

// formatted loggers

func (l *PluginLogger) Debugf(format string, v ...interface{}) {
	l.Log.Debug(l.Name, fmt.Sprintf(format, v...))
}

func (l *PluginLogger) Infof(format string, v ...interface{}) {
	l.Log.Info(l.Name, fmt.Sprintf(format, v...))
}

func (l *PluginLogger) Warningf(format string, v ...interface{}) {
	l.Log.Warning(l.Name, fmt.Sprintf(format, v...))
}

func (l *PluginLogger) Errorf(format string, v ...interface{}) {
	l.Log.Error(l.Name, fmt.Sprintf(format, v...))
}

func (l *PluginLogger) Criticalf(format string, v ...interface{}) {
	l.Log.Critical(l.Name, fmt.Sprintf(format, v...))
}

func (l *PluginLogger) Fatalf(format string, v ...interface{}) {
	l.Log.Fatal(l.Name, fmt.Sprintf(format, v...))
}
