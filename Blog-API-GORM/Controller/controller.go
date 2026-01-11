package controller

import (
	database "blogAPI_GORM/Database"
	models "blogAPI_GORM/Models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome To Blog API using GORM and GORRILA Mux</h1>"))
}

func createBlogPosthelper(blog models.Blog) error {

	ctx := context.Background()
	inserted := database.DB.WithContext(ctx).Create(&blog)
	fmt.Println("Rows affected: ", inserted.RowsAffected)

	return inserted.Error
}

func CreateBlogPost(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	fmt.Println("Post request is called")

	var payload struct {
		Title    string   `json:"title"`
		Content  string   `json:"content"`
		Category string   `json:"category"`
		Tags     []string `json:"tags"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
	}

	tagsJSON, err := json.Marshal(payload.Tags)
	if err != nil {
		http.Error(w, "Failed to process tags", http.StatusInternalServerError)
		return
	}

	blog := models.Blog{
		Title:    payload.Title,
		Content:  payload.Content,
		Category: payload.Category,
		Tags:     tagsJSON,
	}

	err = createBlogPosthelper(blog)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(blog)

}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	blogs := []models.Blog{}

	result := database.DB.Find(&blogs)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error ": result.Error.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blogs)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	/*
		 ------ Tradational method of getting all elements by id -----

		var foundBlog bool

		blogs := []models.Blog{}
		result := database.DB.Find(&blogs)

		if result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error ": result.Error.Error()})
			return
		}

		id, _ := strconv.ParseUint(params["id"], 10, 32)

		var blog models.Blog

		for _, post := range blogs {

			if post.ID == uint(id) {
				blog = post
				foundBlog = true
				break
			}
		}

		if !foundBlog {
			fmt.Println("No Blog with given id exist")
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(blog)

	*/

	// -----GORM Style Method -------

	id, _ := strconv.ParseUint(params["id"], 10, 32)
	var blog models.Blog

	result := database.DB.First(&blog, uint(id))

	if result.Error != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": result.Error.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blog)

}

func UpdateABlogPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	var blog models.Blog

	result := database.DB.First(&blog, uint(id))

	if result.Error != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": result.Error.Error()})
		return
	}

	var payload struct {
		Title    string   `json:"title"`
		Content  string   `json:"content"`
		Category string   `json:"category"`
		Tags     []string `json:"tags"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
	}

	tagsJSON, err := json.Marshal(payload.Tags)
	if err != nil {
		http.Error(w, "Failed to process tags", http.StatusInternalServerError)
		return
	}

	blog.Title = payload.Title
	blog.Content = payload.Content
	blog.Category = payload.Category
	blog.Tags = tagsJSON

	result = database.DB.Save(&blog)
	if result.Error != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": result.Error.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blog)

}

func DeleteById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	var deletedBlog models.Blog

	result := database.DB.First(&deletedBlog, uint(id))

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Blog not found"})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error ": result.Error.Error()})
		return
	}

	err := database.DB.Delete(&deletedBlog).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error ": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedBlog)

}
