package models

type LoginRespone struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	UserName string `json:"userName" form:"userName" xml:"userName" query:"userName"`
	PassWord string `json:"passWord" form:"passWord" xml:"passWord" query:"passWord"`
}
