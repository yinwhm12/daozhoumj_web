package client


type CreateUser	struct {
	Name string	`json:"name,omitempty"`
	Password	string	`json:"password,omitempty"`
}

type LoginSuccessOutPut struct {
	Uid string	`json:"uid,omitempty"`
	Name string	`json:"name,omitempty"`
	Token	string	`json:"token,omitempty"`
}

type TokenLogin	struct {
	Name	string	`json:"name,omitempty"`
	Token	string	`json:"token,omitempty"`
}