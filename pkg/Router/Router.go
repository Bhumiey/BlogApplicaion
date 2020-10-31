package Router

import (
	"BlogApplication/pkg/Code"
	database "BlogApplication/pkg/Database"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func StartServer() *chi.Mux {
	r, err := database.SetUpPostgreSQL()
	if err != nil {
		fmt.Println(err)
	}
	blg := Code.PostGressNewService(r)
	router := chi.NewRouter()
	router.Mount("/blog", Code.BlogRoutes(blg))
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

	return router
}
