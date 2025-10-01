package main

import (
	"context"
	"fmt"
	"io"
	"net/url"
)

var (
	pluginName            = "modifier-example"
	ModifierRegisterer    = registerer(pluginName)
	errRegistererNotFound = fmt.Errorf("%s plugin disabled: config not found", pluginName)
)

type registerer string

// RegisterModifiers is the function the plugin loader will call to register the
// modifier(s) contained in the plugin using the function passed as argument.
// f will register the factoryFunc under the name and mark it as a request
// and/or response modifier.
func (r registerer) RegisterModifiers(f func(
	name string,
	factoryFunc func(map[string]any) func(any) (any, error),
	appliesToRequest bool,
	appliesToResponse bool,
)) {
	f(string(r), r.request, true, false)
}

func (r registerer) request(
	_ map[string]any,
) func(any) (any, error) {
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Modifier injected", r))

	return func(input any) (any, error) {
		logger.Info("modifier success!")
		return nil, nil
	}
}

// RequestWrapper is an interface for passing proxy request between the krakend pipe
// and the loaded plugins
type RequestWrapper interface {
	Context() context.Context
	Params() map[string]string
	Headers() map[string][]string
	Body() io.ReadCloser
	Method() string
	URL() *url.URL
	Query() url.Values
	Path() string
}

type requestWrapper struct {
	ctx     context.Context
	method  string
	url     *url.URL
	query   url.Values
	path    string
	body    io.ReadCloser
	params  map[string]string
	headers map[string][]string
}

func (r requestWrapper) Context() context.Context     { return r.ctx }
func (r requestWrapper) Method() string               { return r.method }
func (r requestWrapper) URL() *url.URL                { return r.url }
func (r requestWrapper) Query() url.Values            { return r.query }
func (r requestWrapper) Path() string                 { return r.path }
func (r requestWrapper) Body() io.ReadCloser          { return r.body }
func (r requestWrapper) Params() map[string]string    { return r.params }
func (r requestWrapper) Headers() map[string][]string { return r.headers }

// ResponseWrapper is an interface for passing proxy response between the krakend pipe
// and the loaded plugins
type ResponseWrapper interface {
	Context() context.Context
	Data() map[string]any
	Io() io.Reader
	IsComplete() bool
	StatusCode() int
	Headers() map[string][]string
}

type metadataWrapper struct {
	headers    map[string][]string
	statusCode int
}

func (m metadataWrapper) Headers() map[string][]string { return m.headers }
func (m metadataWrapper) StatusCode() int              { return m.statusCode }

type responseWrapper struct {
	ctx        context.Context
	request    any
	data       map[string]any
	isComplete bool
	metadata   metadataWrapper
	io         io.Reader
}

func (r responseWrapper) Context() context.Context     { return r.ctx }
func (r responseWrapper) Request() any                 { return r.request }
func (r responseWrapper) Data() map[string]any         { return r.data }
func (r responseWrapper) IsComplete() bool             { return r.isComplete }
func (r responseWrapper) Io() io.Reader                { return r.io }
func (r responseWrapper) Headers() map[string][]string { return r.metadata.headers }
func (r responseWrapper) StatusCode() int              { return r.metadata.statusCode }

// This logger is replaced by the RegisterLogger method to load the one from KrakenD
var logger Logger = noopLogger{}

func (r registerer) RegisterLogger(v any) {
	l, ok := v.(Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Logger loaded", r))
}

type Logger interface {
	Debug(v ...any)
	Info(v ...any)
	Warning(v ...any)
	Error(v ...any)
	Critical(v ...any)
	Fatal(v ...any)
}

// Empty logger implementation
type noopLogger struct{}

func (n noopLogger) Debug(_ ...any)    {}
func (n noopLogger) Info(_ ...any)     {}
func (n noopLogger) Warning(_ ...any)  {}
func (n noopLogger) Error(_ ...any)    {}
func (n noopLogger) Critical(_ ...any) {}
func (n noopLogger) Fatal(_ ...any)    {}

var ctx = context.Background()

// RegisterContext saves the KrakenD application context so KrakenD can inject
// the context like we do with the logger
func (registerer) RegisterContext(c context.Context) {
	ctx = c
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Context loaded", ModifierRegisterer))
}

func main() {}
