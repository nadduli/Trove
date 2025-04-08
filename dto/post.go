package dto

type CreatePostDto struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"Content" binding:"required"`
}

type UpdatePostDto struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"Content" binding:"required"`
}
