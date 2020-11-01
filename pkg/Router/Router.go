package Router

import (
	"BlogApplication/pkg/Code"
	database "BlogApplication/pkg/Database"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func StartServer() *chi.Mux {
	r, err := database.SetUpPostgreSQL()
	if err != nil {
		fmt.Println(err)
	}
	blg := Code.PostGressNewService(r)
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Mount("/blog", Code.BlogRoutes(blg))
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

	return router
}
