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

type TokenLogin	struct {
	Name	string	`json:"name,omitempty"`
	Token	string	`json:"token,omitempty"`
}

type TestPlayer struct {
	NickName string	`json:"nick_name,omitempty"`
	
}