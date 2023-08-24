package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/d8x/sgw/providers"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	p providers.Provider
}

func NewHandler(p providers.Provider) *Handler {
	return &Handler{
		p: p,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("request: %s", r.URL.Path)

	if r.URL.Path == "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	uriNoSlash := strings.TrimPrefix(r.URL.Path, "/")
	stat, err := h.p.GetObjectStat(r.Context(), uriNoSlash)
	if err != nil {
		logrus.Warnf("failed to stat object: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	object, err := h.p.GetObject(r.Context(), uriNoSlash)
	if err != nil {
		logrus.Warnf("failed to get object: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer object.Close()

	w.Header().Set("Content-Type", stat.ContentType)
	w.Header().Set("Content-Length", strconv.FormatInt(stat.Size, 10))

	_, err = io.Copy(w, object)
	if err != nil {
		logrus.Warnf("failed to copy object: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
