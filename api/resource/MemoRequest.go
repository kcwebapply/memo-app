package memoRequestResource

type MemoRequest struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}
