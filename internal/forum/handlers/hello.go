package handlers

import (
	"net/http"
 )


type ForumHandler struct {
}


func NewForumHandler() *ForumHandler {
	return &ForumHandler{}
}


func (h *ForumHandler) HelloWorld(w http.ResponseWriter, r *http.Request ) {
	
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello World! Welcome to the protected Forum!"))

}