package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id int `json:"_id" bson:"_id"`
	OpenId string `json:"open_id" bson:"open_id"`
	NickName string `json:"nick_name" bson:"nick_name"`
	HeadImgUrl string `json:"head_img_url" bson:"head_img_url"`
	GameId string `json:"game_id" bson:"game_id"`
	Gold int `json:"gold" bson:"gold"`
	Diamond int `json:"diamond" bson:"diamond"`
	Sex int `json:"sex" bson:"sex"`
	AccessToken string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	UnionId string `json:"union_id" bson:"union_id"`
}

func GetUserByGameId(gameId string)( *User, error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	//c := conn.DB("").C("user")
	c := conn.DB("ddzhu").C("user")
	user := new(User)
	err := c.Find(bson.M{"game_id":gameId}).One(user)
	return user, err
}

func GetUsersByGameIds(gameIds []string)(users []User,err error)  {
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("user")
	err = c.Find(bson.M{"game_id":bson.M{"$in":gameIds}}).All(&users)
	return
}

//每次读取一百个玩家
func GetLimitMoreUser(offset, limit int)(users []User, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("user")
	err = c.Find(nil).Sort("-_id").Skip(offset).Limit(limit).All(&users)
	return
}

// gold 为最终的数量
func EditUserGoldByGameId(gameId string,gold int)error  {
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("user")
	err := c.Update(bson.M{"game_id":gameId},bson.M{"$set":bson.M{"gold":gold}})
	return err
}

// diamonds 为最终的数量
func EditUserDiamondsByGameId(gameId string,diamonds int)error  {
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("user")
	err := c.Update(bson.M{"game_id":gameId},bson.M{"$set":bson.M{"diamond":diamonds}})
	return err
}
