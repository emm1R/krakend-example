package main

import (
	clog "common/log"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/neicnordic/crypt4gh/keys"
)

type configStruct struct {
	PrivateKey string
	PassPhrase string
}

// Config values
var s3Config configStruct

// allasHeaderClient is a krakend http client plugin which takes a file from allas
// and returns its' header (re-encrypted with the given public key) and the offset of the original header
func (r registerer) allasHeaderClient(_ context.Context, extra map[string]interface{}) (http.Handler, error) {
	// Load config values
	pluginConfig := extra[allasHeaderPlugin].(map[string]interface{})
	conf := configStruct{}
	conf.PrivateKey = pluginConfig["c4gh_private_key_path"].(string)
	conf.PassPhrase = pluginConfig["c4gh_private_key_passphrase"].(string)
	s3Config = conf

	_, err := getPrivateKey()
	if err != nil {
		log.Fatal("failed to get private key, reason: ", err)
	}

	log.Info("PLUGIN INJECTED")

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	}), nil
}

func getPrivateKey() (privateKey [32]byte, err error) {
	// Read private key for decrypt and encrypt operations
	passPhraseBytes := []byte(s3Config.PassPhrase)
	log.Debug(fmt.Sprintf("attempting to open private key path %s", s3Config.PrivateKey))
	file, err := os.Open(s3Config.PrivateKey)
	if err != nil {
		log.Error("failed to open private key file, error: ", err)

		return
	}
	defer file.Close()

	// Read private key for decrypt and encrypt operations
	privateKey, err = keys.ReadPrivateKey(file, passPhraseBytes)
	if err != nil {
		log.Error("failed to read private key from a file, error: ", err)

		return
	}

	return privateKey, nil
}

// KrakenD boilerplate

var PluginName = "allas-header-client"

type registerer string

var ClientRegisterer = registerer(PluginName)
var allasHeaderPlugin = string(ClientRegisterer)

var log *clog.PluginLogger

// RegisterLogger connects the plugin's logger with krakend's logging mechanism
func (r registerer) RegisterLogger(v interface{}) {
	logger, ok := v.(clog.Logger)
	if !ok {
		return
	}
	log = clog.NewLogger(logger, PluginName)
	log.Debug("LOGGER LOADED")
}

// RegisterClients loads client plugins
func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(allasHeaderPlugin, r.allasHeaderClient)

	log.Debug("PLUGIN REGISTERED!")
}
