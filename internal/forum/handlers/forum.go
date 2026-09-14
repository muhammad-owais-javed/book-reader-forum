package handlers

import (
	
	"fmt"
	"net/http"
	"forum/internal/forum/services"
 )


type ForumHandler struct {

	PostService *services.PostService

}


func NewForumHandler(postService *services.PostService) *ForumHandler {
	
	return &ForumHandler{PostService: postService}
	//return &ForumHandler{}
}

func (h *ForumHandler) ViewForum(w http.ResponseWriter, r *http.Request ) {
	
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed )
		return
	}

	
	ctx := r.Context()
	
	posts, err := h.PostService.GetAllPosts(ctx)
	if err != nil {
		http.Error(w, "Failed to load posts", http.StatusInternalServerError )
		return
	}

	// Parsing html
	tmpl, err := template.ParseFiles("./ui/html/forum.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError )
		return
	}

	// Just print the number of posts for testing!
	// w.Header().Set("Content-Type", "text/plain")
	// w.Write([]byte(fmt.Sprintf("Welcome to the Forum! There are %d posts.", len(posts))))

	tmpl.Execute(w, posts)

}


// func (h *ForumHandler) HelloWorld(w http.ResponseWriter, r *http.Request ) {
	
// 	w.Header().Set("Content-Type", "text/plain")
// 	w.Write([]byte("Hello World! Welcome to the protected Forum!"))

// }