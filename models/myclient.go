package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
)

//自己拉的人  直属人
type MyClient struct {
	Id int `json:"_id" bson:"_id"`
	GameId string `json:"game_id" bson:"game_id"`
	Sons []string `json:"sons" bson:"sons"`
	NewCount int `json:"new_count" bson:"new_count"`
	LastWeekTime int `json:"last_week_time" bson:"last_week_time"`
}

func GetMyClientByGameId(gameId string)(m *MyClient,err error)  {

	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("myclient")
	err = c.Find(bson.M{"game_id":gameId}).One(&m)
	return
}



