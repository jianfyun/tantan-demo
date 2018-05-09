package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"tantan-demo/config"
	"tantan-demo/lib/log"
	"tantan-demo/lib/uuid"
	"time"

	"github.com/gorilla/mux"
)

// RESTful API Resource Type
const (
	ResTypeUser         = "user"
	ResTypeRelationship = "relationship"
)

// Handler handles all the HTTP requests.
type Handler struct {
	Handle func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query().Get("u")
	if len(u) == 0 {
		u = uuid.New()
	}
	ctx := context.WithValue(context.Background(), log.TraceKey, u)
	log.Infofx(ctx, "method: %s, path: %s", r.Method, r.URL.Path)
	h.Handle(ctx, w, r)
}

// Handle starts to handle HTTP requests.
func Handle() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.Handle("/", Handler{Handle: GetIndex}).Methods("GET")
	s.Handle("/users", Handler{Handle: CreateUser}).Methods("POST")
	s.Handle("/users", Handler{Handle: GetAllUsers}).Methods("GET")
	s.Handle("/users/{user_id:[0-9]+}/relationships/{other_user_id:[0-9]+}", Handler{Handle: PutUserRelation}).Methods("PUT")
	s.Handle("/users/{user_id:[0-9]+}/relationships", Handler{Handle: GetUserRelations}).Methods("GET")
	srv := &http.Server{
		Handler:      s,
		Addr:         config.Config["rest-api.addr"].Str,
		WriteTimeout: time.Duration(config.Config["rest-api.timeout"].Int) * time.Second,
		ReadTimeout:  time.Duration(config.Config["rest-api.timeout"].Int) * time.Second,
	}
	log.Infof("start HTTP server on addr %s", srv.Addr)
	go func() {
		log.Infof("HTTP server is shutting down, err=%v", srv.ListenAndServe())
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	if err := srv.Shutdown(nil); err != nil {
		log.Fatalf("HTTP server shutdown err(%v)", err)
	}
	log.Info("HTTP server is stopped")
}

// ReadJSON reads the HTTP request body and decodes it from JSON to v.
func ReadJSON(ctx context.Context, r *http.Request, v interface{}) error {
	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		log.Errorfx(ctx, "invalid Content-Type: %s", r.Header.Get("Content-Type"))
		return fmt.Errorf("invalid Content-Type")
	}
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		log.Errorfx(ctx, "decode JSON err: %v", err)
		return err
	}
	log.Infofx(ctx, "request params: %v", r.Method, r.URL.Path, v)
	return nil
}

// WriteJSON encodes HTTP response body to JSON from code and data, then writes HTTP response status and body.
func WriteJSON(ctx context.Context, w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Errorfx(ctx, "encode JSON err: %v", err)
			return err
		}
	}
	log.Infofx(ctx, "response status: %d, result: %v", status, data)
	return nil
}

// GetIndex handles GET requests to the default root.
func GetIndex(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(ctx, w, http.StatusOK, nil)
}
