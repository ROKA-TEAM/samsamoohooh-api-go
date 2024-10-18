package presenter

type CommentCreateRequest struct {
	Content string `json:"content" validate:"min=1,max=24"`
	PostID  int    `json:"postID"`
}

func (r CommentCreateRequest) ToDomain() *CommentCreateRequest {
	return &CommentCreateRequest{
		Content: r.Content,
		PostID:  r.PostID,
	}
}

type CommentUpdateRequest struct {
	Content string `json:"content" validate:"min=1,max=24"`
}

func (r CommentUpdateRequest) ToDomain() *CommentUpdateRequest {
	return &CommentUpdateRequest{
		Content: r.Content,
	}
}
