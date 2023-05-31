package request

type ReqSaveForum struct {
	ForumName        string `json:"forum_name" validate:"required"`
	IntroductionText string `json:"introduction_text"`
	Category         string `json:"category"`
}
