package server

import (
	"context"
	"demo/internal/sample/server/midlleware"
	"demo/internal/sample/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	store store.BookStore
	srv   *http.Server
}

func (s *Server) ListenAndServe() (chan error, error) {
	var err error
	errChan := make(chan error)
	go func() {
		err = s.srv.ListenAndServe()
		errChan <- err
	}()

	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func NewServer(addr string, store store.BookStore) *Server {
	srv := &Server{
		store,
		&http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBook).Methods("POST")
	router.HandleFunc("/book/{id}", srv.getBook).Methods("GET")
	srv.srv.Handler = midlleware.Logging(midlleware.Validating(router))
	return srv
}

func (s *Server) createBook(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.store.Save(&book)
	return
}

func (s *Server) getBook(w http.ResponseWriter, req *http.Request) {
	strId, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(w, "no id found in request", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	book := s.store.GetById(id)
	response(w, book)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func response(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
