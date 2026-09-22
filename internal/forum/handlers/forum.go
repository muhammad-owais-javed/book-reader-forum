package handlers

import (
	"net/http"
	"forum/internal/forum/services"
	"html/template"
	constants "forum/internal/constants"
	"encoding/json"
 )


type ForumHandler struct {
	PostService *services.PostService
	CommentService *services.CommentService
	PostReactionService *services.PostReactionService
}


func NewForumHandler(postService *services.PostService, commentService *services.CommentService, postReactionService *services.PostReactionService) *ForumHandler {
	
	return &ForumHandler{PostService: postService, CommentService: commentService, PostReactionService: postReactionService}
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

	userID, _ := ctx.Value(constants.UserIDKey).(string)
	// adding the comments to the post struct
	for i := range posts {
		comments, err := h.CommentService.GetCommentsByPostID(ctx, posts[i].ID)
		if err != nil {
			http.Error(w, "Failed to load comments", http.StatusInternalServerError)
			return
		}
		posts[i].Comments = comments

		// this part is to show the amount of like and dislikes per post
		likeCount, dislikeCount, err := h.PostReactionService.GetReactionCounts(ctx, posts[i].ID)
		if err != nil {
			http.Error(w, "Failed to load reactions", http.StatusInternalServerError)
			return
		}
		posts[i].LikeCount = likeCount
		posts[i].DislikeCount = dislikeCount

		// this part is to mark if the active user reacted to the post
		reaction, err := h.PostReactionService.GetReaction(ctx, userID, posts[i].ID)
		if err != nil {
			http.Error(w, "Failed to load user reaction", http.StatusInternalServerError)
			return
		}
		if reaction != nil {
			if reaction.IsLike {
				posts[i].UserLiked = true
			} else {
				posts[i].UserDisliked = true
			}
		}
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
	comments, err := h.CommentService.GetCommentsByPostID(ctx, postID)
	if err != nil {
		http.Error(w, "Failed to load comments", http.StatusInternalServerError)
		return
	}
	if len(comments) == 0 {
		http.Error(w, "Comment was not found", http.StatusInternalServerError)
		return
	}
	newComment := comments[0]
	response := struct {
		Username  string `json:"username"`
		Content   string `json:"content"`
		CreatedAt string `json:"createdAt"`
	}{
		Username:  newComment.Username,
		Content:   newComment.Content,
		CreatedAt: newComment.CreatedAt.Format("Jan 02, 2006 15:04"),
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ForumHandler) TogglePostReaction(w http.ResponseWriter, r *http.Request) {
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
	reaction := r.FormValue("reaction")
	ctx := r.Context()
	userID, ok := ctx.Value(constants.UserIDKey).(string)
	if !ok || userID == "" {
		http.Error(w, "Unauthorized: User ID not found in context", http.StatusUnauthorized)
		return
	}
	var isLike bool
	switch reaction {
	case "like":
		isLike = true
	case "dislike":
		isLike = false
	default:
		http.Error(w, "Invalid reaction", http.StatusBadRequest)
		return
	}
	err = h.PostReactionService.ToggleReaction(ctx, userID, postID, isLike)
	if err != nil {
		http.Error(w, "Failed to update reaction", http.StatusInternalServerError)
		return
	}
	likeCount, dislikeCount, err := h.PostReactionService.GetReactionCounts(ctx, postID)
	if err != nil {
		http.Error(w, "Failed to load reaction counts", http.StatusInternalServerError)
		return
	}
	currentReaction, err := h.PostReactionService.GetReaction(ctx, userID, postID)
	if err != nil {
		http.Error(w, "Failed to load user reaction", http.StatusInternalServerError)
		return
	}
	userReaction := ""
	if currentReaction != nil {
		if currentReaction.IsLike {
			userReaction = "like"
		} else {
			userReaction = "dislike"
		}
	}
	response := struct {
		LikeCount    int    `json:"likeCount"`
		DislikeCount int    `json:"dislikeCount"`
		UserReaction string `json:"userReaction"`
	}{
		LikeCount:    likeCount,
		DislikeCount: dislikeCount,
		UserReaction: userReaction,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// func (h *ForumHandler) HelloWorld(w http.ResponseWriter, r *http.Request ) {
	
// 	w.Header().Set("Content-Type", "text/plain")
// 	w.Write([]byte("Hello World! Welcome to the protected Forum!"))

// }