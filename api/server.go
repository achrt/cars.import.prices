package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"gitlab.kodixauto.ru/nw/logger"
)

type server struct {
	Router *mux.Router
}

type health struct {
	Status  bool `json:"status"`
	Version int  `json:"version"`
}

func (s *server) health(w http.ResponseWriter, r *http.Request) {
	response := &health{
		Status:  true,
		Version: 3,
	}

	j, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (s *server) init() {
	s.Router.Use(loggingMiddleware, contextMiddleware)
	s.Router.HandleFunc("/health", s.health).Methods(http.MethodGet)
}

func New(r *mux.Router) *server {
	s := &server{r}
	s.init()

	return s
}

func (s *server) Run(address string) error {
	return http.ListenAndServe(address, s.Router)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("X-Request-Id")

		if requestId == "" {
			requestId = uuid.New().String()
			r.Header.Set("X-Request-Id", requestId)
		}

		l := logger.New(requestId)

		l.Info(r.Method+" "+r.RequestURI, nil)

		next.ServeHTTP(w, r)
	})
}

func contextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceId := r.Header.Get("X-Request-Id")

		c := context.WithValue(r.Context(), "context", &Context{
			TraceId:         traceId,
			UserId:          r.Header.Get("X-Auth-User-Id"),
			InternalRequest: r.Header.Get("X-Auth-Internal-Request"),
			ClientIds:       extractClients(r),
		})

		next.ServeHTTP(w, r.WithContext(c))
	})
}

func extractClients(r *http.Request) []string {
	c := r.Header.Get("X-Auth-Clients-Id")

	if c == "" {
		return []string{}
	} else {
		return strings.Split(c, ",")
	}
}
