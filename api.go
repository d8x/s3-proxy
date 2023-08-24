package main

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type API struct {
	config *Config
}

func NewAPI(config *Config) *API {
	return &API{
		config: config,
	}
}

func (a *API) Create(m *http.ServeMux) {
	m.HandleFunc("/", a.ProviderDispatcher)
}

func (a *API) ProviderDispatcher(w http.ResponseWriter, r *http.Request) {
	provider := r.URL.Query().Get("provider")
	logrus.Debugf("query provider: %s", provider)
	if provider == "" {
		provider = r.Header.Get("X-Provider")
		logrus.Debugf("header provider: %s", provider)
	}

	if provider == "" {
		logrus.Errorf("provider not specified in query or header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storage, ok := a.config.Providers[strings.ToLower(provider)]
	if !ok {
		logrus.Errorf("provider %s not found", provider)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	NewHandler(storage.StorageProvider).ServeHTTP(w, r)
}
