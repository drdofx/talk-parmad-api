package response

type ResDetailThread struct {
	ThreadData     ResThreadField  `json:"thread"`
	ReplyData      []ResReplyField `json:"reply"`
	TotalReplies   int             `json:"total_replies"`
	TotalUpvotes   int64           `json:"total_upvotes"`
	TotalDownvotes int64           `json:"total_downvotes"`
	CreatedBy      string          `json:"created_by"`
}

type ResThreadField struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type ResReplyField struct {
	ID             uint   `json:"id"`
	Text           string `json:"text"`
	CreatedBy      string `json:"created_by"`
	CreatedAt      string `json:"created_at"`
	TotalUpvotes   int64  `json:"total_upvotes"`
	TotalDownvotes int64  `json:"total_downvotes"`
}
