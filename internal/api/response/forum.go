package response

import "github.com/drdofx/talk-parmad/internal/api/models"

type ResDetailForum struct {
	ForumData       models.Forum    `json:"forum"`
	ThreadData      []models.Thread `json:"threads"`
	TotalThreads    int             `json:"total_threads"`
	NumberOfMembers int64           `json:"number_of_members"`
}

type ResThreadForum struct {
	ForumData  models.Forum  `json:"forum"`
	ThreadData models.Thread `json:"thread"`
}

type ResThreadForumHome struct {
	UserID     uint   `json:"user_id"`
	ForumID    uint   `json:"forum_id"`
	ForumName  string `json:"forum_name"`
	ForumImage string `json:"forum_image"`
	ThreadID   uint   `json:"thread_id"`
	Title      string `json:"title"`
	Text       string `json:"text"`
}
