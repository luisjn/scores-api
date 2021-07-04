package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	chiDispatcher = chi.NewRouter()
)

type chiRouter struct{}

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(res http.ResponseWriter, rq *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(res http.ResponseWriter, rq *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Println("Chi HTTP server running on port", port)
	http.ListenAndServe(":"+port, chiDispatcher)
}
