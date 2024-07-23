package main

import (
	"context"
	"net/http"

	clog "common/log"
)

var PluginName = "profile-client"

type registerer string

var ClientRegisterer = registerer(PluginName)
var profilePlugin = string(ClientRegisterer)

var log *clog.PluginLogger

func (registerer) RegisterLogger(v interface{}) {
	logger, ok := v.(clog.Logger)
	if !ok {
		return
	}
	log = clog.NewLogger(logger, PluginName)
	log.Debug("LOGGER LOADED")
}

func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(profilePlugin, r.registerClients)

	log.Debug("PLUGIN REGISTERED")
}

func (r registerer) registerClients(_ context.Context, _ map[string]interface{}) (http.Handler, error) {
	log.Info("PLUGIN INJECTED")

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	}), nil
}

func main() {}
