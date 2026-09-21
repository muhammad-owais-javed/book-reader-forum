package handlers

import (
	
	"net/http"
	"forum/internal/forum/services"
	"html/template"
	constants "forum/internal/constants"
 )


type ForumHandler struct {

	PostService *services.PostService
	CommentService *services.CommentService

}


func NewForumHandler(postService *services.PostService, commentService *services.CommentService) *ForumHandler {
	
	return &ForumHandler{PostService: postService, CommentService: commentService}
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
	// adding the comments to the post struct
	for i := range posts {
		comments, err := h.CommentService.GetCommentsByPostID(ctx, posts[i].ID)
		if err != nil {
			http.Error(w, "Failed to load comments", http.StatusInternalServerError)
			return
		}
		posts[i].Comments = comments
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


func (h *ForumHandler) CreatePost(w http.ResponseWriter, r *http.Request ) {
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed )
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest )
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	// TODO: We need the REAL logged-in user ID here!
	//userID := "dummy-user-id" 

	ctx := r.Context()

	userID, ok := ctx.Value(constants.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized: User ID not found in context", http.StatusUnauthorized )
		return
	}

	err = h.PostService.CreatePost(ctx, userID, title, content)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError )
		return
	}

	http.Redirect(w, r, "/forum", http.StatusSeeOther )
}

func (h *ForumHandler) CreateComment(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	postID := r.FormValue("post_id")
	content := r.FormValue("content")

	ctx := r.Context()

	userID, ok := ctx.Value(constants.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized: User ID not found in context", http.StatusUnauthorized)
		return
	}

	err = h.CommentService.CreateComment(ctx, userID, postID, content)
	if err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}

// func (h *ForumHandler) HelloWorld(w http.ResponseWriter, r *http.Request ) {
	
// 	w.Header().Set("Content-Type", "text/plain")
// 	w.Write([]byte("Hello World! Welcome to the protected Forum!"))

// }