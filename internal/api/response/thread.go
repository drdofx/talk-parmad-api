package response

import "github.com/drdofx/talk-parmad/internal/api/models"

type ResDetailThread struct {
	ThreadData     models.Thread  `json:"thread"`
	ReplyData      []models.Reply `json:"replies"`
	TotalReplies   int64          `json:"total_replies"`
	TotalUpvotes   int64          `json:"total_upvotes"`
	TotalDownvotes int64          `json:"total_downvotes"`
}
