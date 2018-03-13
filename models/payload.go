package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
)

//玩家贡献数据

type PayLoad struct {
	Id int `json:"_id" bson:"_id"`
	Asset int `json:"asset" bson:"asset"`
	GameId string `json:"game_id" bson:"game_id"`
	CreateTime int `json:"create_time" bson:"create_time"`
}

func GetPayLoadsByGameId(gameId string,t1,t2 int)(ps []PayLoad, err error)  {

	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("payload")
	err = c.Find(bson.M{"game_id":gameId,"create_time":bson.M{"$gte":t1,"$lte":t2}}).All(&ps)
	return ps, err
}

func GetPayLoadsByTime(t1, t2 int)(ps []PayLoad, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("payload")
	err = c.Find(bson.M{"create_time":bson.M{"$gte":t1,"$lte":t2}}).All(&ps)
	return ps, err
}

func DeletePayLoadsLowDate(t int)(error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("payload")
	_,err := c.RemoveAll(bson.M{"create_time":bson.M{"$lt":t}})
	return err
}

