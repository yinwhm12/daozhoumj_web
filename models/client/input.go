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