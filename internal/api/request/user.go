package request

type ReqSaveUser struct {
	NIM      string `json:"nim" validate:"req-numeric,max=10"`
	Email    string `json:"email" validate:"req-email"`
	Password string `json:"password" validate:"required"`
}

type ReqLoginUser struct {
	NIM      string `json:"nim" validate:"numeric"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}
