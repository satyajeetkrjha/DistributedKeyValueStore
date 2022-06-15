//go run main.go -db-location=$PWD/my.db -config-file=$PWD/sharding.toml -shard=Washington
package web

import (
	"distrikv/db"
	"fmt"
	"hash/fnv"
	"net/http"
)

type Server struct {
	db         *db.Database
	shardIdx   int
	shardCount int
}

func (s *Server) getShard(key string) int {
	h := fnv.New64()
	h.Write([]byte(key))
	return int(h.Sum64() % uint64(s.shardCount))
}

func NewServer(db *db.Database, shardIdx, shardCount int) *Server {
	return &Server{
		db:         db,
		shardIdx:   shardIdx,
		shardCount: shardCount,
	}
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")

	shard := s.getShard(key)

	value, err := s.db.GetKey(key)
	fmt.Fprintf(w, "Shard = %d, current shard = %d, Value = %q, error = %v", shard, s.shardIdx, value, err)

}

func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value := r.Form.Get("value")

	fmt.Println(key)
	fmt.Println(value)

	h := fnv.New64()
	h.Write([]byte(key))
	shardIdx := h.Sum64() % uint64(s.shardCount)

	fmt.Println(shardIdx)

	err := s.db.SetKey(key, []byte(value))

	fmt.Fprintf(w, "Error = %v, shardIdx = %d", err, shardIdx)
}
