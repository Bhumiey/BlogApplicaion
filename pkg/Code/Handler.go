package Code

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var (
	blog Blog
)

func getHandler(s PostgressService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		blogList, err := s.GetAllBlog()
		if err != nil {
			json.NewEncoder(w).Encode("Something went wrong while getting blogs...")
		}

		json.NewEncoder(w).Encode(&blogList)
	}

}

func postHandler(s PostgressService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		json.NewDecoder(r.Body).Decode(&blog)

		if err := s.AddNewBlog(blog); err != nil {
			errorMsg := "Something went wrong..." + err.Error()
			json.NewEncoder(w).Encode(errorMsg)
			return
		}
		println("Successfully added to DB...")
		json.NewEncoder(w).Encode("You added blog to Postgress DB")
	}
}
func deleteHandler(s PostgressService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		id, _ := strconv.Atoi(r.FormValue("id"))
		if err := s.DeleteBlog(id); err != nil {
			json.NewEncoder(w).Encode("Something went wrong...")
			return
		}
		json.NewEncoder(w).Encode("You have deleted blog")
	}
}
