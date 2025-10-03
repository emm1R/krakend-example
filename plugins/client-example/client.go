package main

import (
	"context"
	"fmt"
	"net/http"
)

var (
	pluginName = "client-example"
	// ClientRegisterer is the symbol the plugin loader will try to load. It must implement the RegisterClient interface
	ClientRegisterer = registerer(pluginName)
)

type registerer string

func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), r.registerClients)
}

func (r registerer) registerClients(_ context.Context, _ map[string]interface{}) (http.Handler, error) {
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Client injected", r))

	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		blob := `
		<animals>
			<animal>gopher</animal>
			<animal>armadillo</animal>
			<animal>zebra</animal>
			<animal>unknown</animal>
			<animal>gopher</animal>
			<animal>bee</animal>
			<animal>gopher</animal>
			<animal>zebra</animal>
		</animals>`

		w.Header().Set("Content-Type", "application/xml")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(blob)))

		if _, err := w.Write([]byte(blob)); err != nil {
			logger.Error("Failed to write response: ", err)
			http.Error(w, "Request failed", http.StatusInternalServerError)

			return
		}
		w.WriteHeader(http.StatusOK)
	}), nil
}

func main() {}

var logger Logger = nil

func (registerer) RegisterLogger(v interface{}) {
	l, ok := v.(Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] client plugin loaded", pluginName))
}

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}
