package models

//玩家贡献数据

type PayLoad struct {
	Id int `json:"_id" bson:"_id"`
	Asset int `json:"asset" bson:"asset"`
	GameId string `json:"game_id" bson:"game_id"`
	CreateTime int `json:"create_time" bson:"create_time"`
}

//func GetPayLoadByGameId(gameId string)(ps []PayLoad, err error)  {
//
//}

