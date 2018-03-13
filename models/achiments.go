package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

//获取玩家真实的个人信息
type Achiment struct {
	Id string `json:"_id" bson:"_id"` //gameId
	Achievements int `json:"achievements" bson:"achievements"`
	Commision int `json:"commision" bson:"commision"`
	Degree int `json:"degree" bson:"degree"`
	CreateTime int `json:"create_time" bson:"create_time"`
}

func GetAchievemetById(id string)(a *Achiment, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("achievements")
	err = c.Find(bson.M{"_id":id}).One(&a)
	return
}

func GetTenTopData()(as []Achiment,err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("achievements")
	err = c.Find(nil).Sort("-achievements").Limit(10).All(&as)
	return
}

//每一条 进行更新
func UpsertAchiment(a *Achiment)error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("achievements")
	changeInfo,err := c.Upsert(bson.M{"_id":a.Id},a)
	if err != nil{
		return err
	}
	fmt.Printf("%+v\n", changeInfo)
	return err
}
//翻页获取用户数据
func GetDataByPage(offset,limit int)(total int, as []Achiment,err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("achievements")
	total, err = c.Find(nil).Count()
	if err != nil{
		return -1, nil, err
	}
	err = c.Find(nil).Sort("-achievements").Skip(offset).Limit(limit).All(&as)
	return

}
