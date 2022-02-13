package server

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"microblog/pkg"
	"microblog/server/handlers"
	"net/http"
)

func StartServer(port string, handl *handlers.Handler) {
	router := mux.NewRouter()

	NewRouter(router, handl)

	pkg.LogInfo("Server Start")
	err := http.ListenAndServe(port, router)
	if err != nil {
		pkg.FatalError(errors.Wrap(err, "err with http.ListenAndServe"))
	}
}
