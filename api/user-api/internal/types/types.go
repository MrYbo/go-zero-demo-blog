// Code generated by goctl. DO NOT EDIT.
package types

type UserId struct {
	Id int64 `path:"id"`
}

type UserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Birthday string `json:"birthday"`
}

type BaseUser struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Birthday  string `json:"birthday"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UpdateUser struct {
	Id       int    `path:"id"`
	Avatar   string `json:"avatar, optional"`
	Phone    string `json:"phone, optional"`
	Name     string `json:"name, optional"`
	Address  string `json:"address, optional"`
	Birthday string `json:"birthday, optional"`
}

type SelectParameters struct {
	Page     int `form:"page, defalut=1, optional"`
	PageSize int `form:"pageSize, defalut=10, optional"`
}

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RespLogin struct {
	Token string `json:"token"`
}

type CommonOK struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
