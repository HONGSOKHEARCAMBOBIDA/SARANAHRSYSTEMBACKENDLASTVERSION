package auth

type LoginReq struct {
	Contact  string `json:"contact" binding:"required"`
	Password string `json:"password" binding:"required"`
}
