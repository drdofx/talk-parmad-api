package request

type ReqSaveThread struct {
	ForumID string `json:"forum_id" validate:"req-numeric"`
	Title   string `json:"title" validate:"required"`
	Text    string `json:"text" validate:"required"`
}

type ReqVoteThread struct {
	ThreadID string `json:"thread_id" validate:"req-numeric"`
	Vote     bool   `json:"vote"`
}

type ReqEditThread struct {
	ThreadID string `json:"thread_id" validate:"req-numeric"`
	Title    string `json:"title"`
	Text     string `json:"text"`
}
