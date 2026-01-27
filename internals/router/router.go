package router

import (
	"net/http"

	"github.com/Ankush263/http-server/internals/handler"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)

	return mux
}
