package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GetRouter Set up a mux router
func GetRouter() *mux.Router {
	muxRouter := mux.NewRouter()
	postRouter := muxRouter.Methods(http.MethodPost).Subrouter()

	postRouter.HandleFunc("/article", CreateArticleHandler)

	return muxRouter
}
