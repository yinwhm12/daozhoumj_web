package models

import (
	"gopkg.in/mgo.v2/bson"
	"daozhoumj/models/mongodb"
)

//售卡 记录
type SoldRecords struct {
	Id	bson.ObjectId `bson:"id_" json:"id,omitempty"`
	FromPlayerId	string	`bson:"from_player_id" json:"from_player_id,omitempty"` //售卡人的ID
	ToPlayerId	string	`bson:"to_player_id" json:"to_player_id,omitempty"` //买卡人ID 即显示中 玩家ID
	PlayerNickName	string	`bson:"player_nick_name" json:"player_nick_name,omitempty"`//玩家的昵称
	PlayerProxy	int	`bson:"player_proxy" json:"player_proxy,omitempty"` //玩家的代理级别
	SoldType	int	`bson:"sold_type" json:"sold_type,omitempty"` //交易类型 1为出售 0为买入

	SoldCounts	int	`bson:"sold_counts" json:"sold_counts,omitempty"` //售卡数量
	SoldState	int	`bson:"sold_state" json:"sold_state,omitempty"` //售卡 状态 1位 成功 0为失败
	SoldTime	int	`bson:"sold_time" json:"sold_time,omitempty"` //售卡时间
}

func AddSoldRecord(s *SoldRecords)error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("soldCards")
	err := c.Insert(s)
	return err
}

func GetSoldRecordsByPage(limit, offset int, fromPlayerId string)(total int, s []SoldRecords, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("soldCards")
	total, err = c.Find(bson.M{"from_player_id":fromPlayerId}).Count()
	if err != nil{
		return -1, nil, err
	}
	err = c.Find(bson.M{"from_player_id":fromPlayerId}).Skip(offset).Limit(limit).All(&s)
	return
}

func GetSoldRecordPageByTime(limit, offset int,fromPlayerId string,t []int)(total int, s []SoldRecords, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("soldCards")
	total, err = c.Find(bson.M{"from_player_id":fromPlayerId,"sold_time":bson.M{"$gte":t[0],"$lte":t[1]}}).Count()
	if err != nil{
		return -1, nil, err
	}
	err = c.Find(bson.M{"from_player_id":fromPlayerId,"sold_time":bson.M{"$gte":t[0],"$lte":t[1]}}).Skip(offset).Limit(limit).All(&s)
	return
	//total, err = c.Find(bson.M{"from_player_id":fromPlayerId,bson.M{"sold_time":bson.M{"$gte":t[0]}},bson.M{"sold_time":bson.M{"$lte":t[0]}}}).Skip(offset).Limit(limit)
}

func GetSoldRecordPageByToPlayerId(limit, offset int, toPlayerId string)(total int, s []SoldRecords,err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("soldCards")
	total, err = c.Find(bson.M{"to_player_id":toPlayerId}).Count()
	if err != nil{
		return  -1, nil, err
	}
	err = c.Find(bson.M{"to_player_id":toPlayerId}).Skip(offset).Limit(limit).All(&s)
	return
}

func GetSoldRecordPageByToPlayerIdTime(limit, offset int, toPlayerId string, t []int)(total int, s []SoldRecords, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("soldCards")
	total, err = c.Find(bson.M{"to_player_id":toPlayerId,"sold_time":bson.M{"$gte":t[0],"$lte":t[1]}}).Count()
	if err != nil{
		return  -1, nil, err
	}
	err = c.Find(bson.M{"to_player_id":toPlayerId,"sold_time":bson.M{"$gte":t[0],"$lte":t[1]}}).Skip(offset).Limit(limit).All(&s)
	return

}
