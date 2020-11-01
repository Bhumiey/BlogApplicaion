package Code

import (
	"encoding/json"
	"net/http"
)

var (
	blog Blog
)

func getHandler(s PostgressService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		blogList, err := s.GetAllBlog()
		if err != nil {
			json.NewEncoder(w).Encode("Something went wrong while getting blogs...")
		}

		json.NewEncoder(w).Encode(&blogList)
	}

}

func postHandler(s PostgressService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
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
		title := r.FormValue("title")
		if err := s.DeleteBlog(title); err != nil {
			json.NewEncoder(w).Encode("Something went wrong...")
			return
		}
		json.NewEncoder(w).Encode("You have deleted blog")
	}
}
