package request

type ReqSaveForum struct {
	ForumName        string `json:"forum_name" validate:"required"`
	IntroductionText string `json:"introduction_text"`
	Category         string `json:"category"`
}

type ReqJoinForum struct {
	ForumID uint `json:"forum_id" validate:"required"`
}

type ReqCheckModeratorForum struct {
	ForumID uint `json:"forum_id" validate:"required"`
	UserID  uint `json:"-" validate:"required"`
}

type ReqEditForum struct {
	ForumID          uint   `json:"forum_id" validate:"required"`
	ForumName        string `json:"forum_name"`
	IntroductionText string `json:"introduction_text"`
	Category         string `json:"category"`
}

type ReqDeleteForum struct {
	ForumID uint `json:"forum_id" validate:"required"`
}

type ReqDetailForum struct {
	ForumID uint `json:"forum_id" validate:"required"`
}
