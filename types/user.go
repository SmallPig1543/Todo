package types

type UserServiceReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type UserResp struct {
	ID       uint   `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name" `
	CreateAt string `json:"create_at" form:"create_at"`
}
