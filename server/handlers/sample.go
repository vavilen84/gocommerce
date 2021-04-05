package handlers

import (
	"github.com/vavilen84/gocommerce/store"
	"net/http"
)

func Sample(w http.ResponseWriter, r *http.Request) {
	conn, _ := store.GetNewDBConn()
	defer conn.Close()

	w.WriteHeader(http.StatusCreated)
}
