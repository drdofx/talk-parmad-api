package request

type ReqSaveUser struct {
	NIM      string `json:"nim" validate:"req-numeric,max=10"`
	Email    string `json:"email" validate:"req-email"`
	Password string `json:"password" validate:"required"`
}

type ReqLoginUser struct {
	User     string `json:"user" validate:"required"` // email or nim
	Password string `json:"password" validate:"required"`
}
