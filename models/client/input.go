package client


type CreateSession	struct {
	Name string	`json:"name,omitempty"`
	Password	string	`json:"password,omitempty"`
}

type LoginSuccessOutPut struct {
	Uid string	`json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	Token	string	`json:"token,omitempty"`
}

type LoinManagerSuccessOutPut struct {
	Id int `json:"id"`
	NickName string `json:"nick_name"`
	Token string `json:"token"`
}

type TokenLogin	struct {
	Name	string	`json:"name,omitempty"`
	Token	string	`json:"token,omitempty"`
}

type TestPlayer struct {
	NickName string	`json:"nick_name,omitempty"`
	
}

type VersionJSON struct {
	Content	string	`json:"content,omitempty"`
	Address string	`json:"address,omitempty"`
}

//公告 添加新的
type PublicJSON struct {
	Message string `json:"message"`
} 

//请求 时间 翻页 数据
type PageByTime struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Before int `json:"before,omitempty"`
	After int `json:"after,omitempty"`

} 

//修改金币
type EditGoldParams struct {
	Id string `json:"id"`
	Value int `json:"value"`
	Type int `json:"type"`
}

//修改钻石的数量
type EditDiamodParams struct {
	Id string `json:"id"`
	Value int `json:"value"`
	Type int `json:"type"`
}