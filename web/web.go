package web

import (
	"distrikv/db"
	"fmt"
	"net/http"
)

type Server struct {
	db *db.Database
}

func NewServer(db *db.Database) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")

	value, err := s.db.GetKey(key)
	fmt.Fprintf(w, "Value = %q, error = %v", value, err)

}

func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value := r.Form.Get("value")

	err := s.db.SetKey(key, []byte(value))
	fmt.Fprintf(w, "Error = %v", err)
}
