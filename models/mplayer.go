package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
)

//模拟玩家 信息
type Player struct {
	ID	bson.ObjectId	`bson:"_id" json:"id,omitempty"` //通过 ID = bson.NewObjectId() 插入
	NickName	string	`bson:"nick_name" json:"nick_name,omitempty"`
	Sex	int	`bson:"sex" json:"sex,omitempty"` //1为男性 0为女性
	GamePoint	int 	`bson:"game_point" json:"game_point,omitempty"` //游戏积分
	UsedRoomCards	int	`bson:"used_room_cards" json:"used_room_cards,omitempty"`//消耗的房卡
	BoughtRoomCards	int `bson:"bought_room_cards" json:"bought_room_cards,omitempty"` //购买的房卡 历史
	GameRecord	int	`bson:"game_record" json:"game_record,omitempty"` //游戏记录
	IsProxy	int	`bson:"is_proxy" json:"is_proxy,omitempty"` //是否是代理 1为一级代理 2 为二级代理 0不是代理
	LastGameTime	int	`bson:"last_game_time" json:"last_game_time,omitempty"`//最后一次登录游戏时间
	Image	string	`bson:"image" json:"image,omitempty"` //头像url
	
	IsBadPlayer	int	`bson:"is_bad_player" json:"is_bad_player,omitempty"`//是否是黑名单玩家 1为黑名单 0否
}

func GetAllPlayers() (p []Player,err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err = c.Find(nil).All(&p)
	return
}

func GetOnePlayer(id string)(p *Player, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err = c.FindId(id).One(&p)
	return
}

func GetPlayersByPage(limit,offset int)(total int, p []Player,err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	total, err = c.Find(nil).Count()
	if err != nil{
		return -1, nil, err
	}
	err = c.Find(nil).Skip(offset).Limit(limit).All(&p)
	return
}

func GetBadPlayersByPage(limit, offset int)(total int, p []Player, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	total, err = c.Find(bson.M{"is_bad_player": 1}).Count()
	if err != nil{
		return -1, nil, err
	}
	err = c.Find(bson.M{"is_bad_player":1}).Skip(offset).Limit(limit).All(&p)
	return
}

func GetOneBadPlayerById(id string)(p *Player, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err = c.FindId(id).One(&p)
	return
}


