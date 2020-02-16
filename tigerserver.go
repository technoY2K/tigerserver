package tigerserver

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore for server methods
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// TigerServer main server struct.
type TigerServer struct {
	Store PlayerStore
}

func (t *TigerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := t.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
