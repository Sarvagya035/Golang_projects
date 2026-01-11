package routers

import (
	controller "blogAPI_GORM/Controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", controller.ServeHome)
	r.HandleFunc("/posts", controller.CreateBlogPost).Methods("POST")
	r.HandleFunc("/posts", controller.GetAllPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", controller.GetPostByID).Methods("GET")
	r.HandleFunc("/posts/{id}", controller.UpdateABlogPost).Methods("PUT")
	r.HandleFunc("/posts/{id}", controller.DeleteById).Methods("DELETE")

	return r
}
