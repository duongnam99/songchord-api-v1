package routes

import (
	"songchord-api/api/grapqh/songResolver"

	"github.com/gorilla/mux"
	"github.com/graphql-go/handler"
)

func RegisterRoutes() *mux.Router {
	graphRoutes := mux.NewRouter()
	r := registerGraphRoutes(graphRoutes)

	return r

}

func registerGraphRoutes(r *mux.Router) *mux.Router {
	graphQL := handler.New(&handler.Config{
		Schema:   &songResolver.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	r.Handle("/song", graphQL)
	return r
}
