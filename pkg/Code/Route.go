package Code

import "github.com/go-chi/chi"

func BlogRoutes(ps PostgressService) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/getBlog", getHandler(ps))
	r.Post("/addBlog", postHandler(ps))
	r.Delete("/deleteBlog", deleteHandler(ps))
	return r
}
