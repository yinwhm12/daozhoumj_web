package models

import (
	"gopkg.in/mgo.v2/bson"
	"daozhoumj/models/mongodb"
)

//充值记录
type RechargeRecord struct {
	Id	bson.ObjectId `bson:"id_" json:"id,omitempty"`
	PlayerId	string	`bson:"player_id" json:"player_id,omitempty"` //充卡玩家ID
	RechargeCounts	int	`bson:"recharge_counts" json:"recharge_counts,omitempty"`//充卡数量
	LeftCounts	int	`bson:"left_counts" json:"left_counts,omitempty"`//剩余卡数量
	RechargeState	int	`bson:"recharge_state" json:"recharge_state,omitempty"` //充卡状态 1为成功 0为失败
	CreatedTime	int	`bson:"created_time" json:"created_time,omitempty"` //充卡时间
}

func AddOneRechargeRecord(r *RechargeRecord) error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("rechargeRecord")
	err := c.Insert(r)
	return err
}

func GetAllRechargeRecordByPage(limit, offset int, playerId string)(total int,r []RechargeRecord, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("rechargeRecord")
	total, err = c.Find(bson.M{"player_id": playerId}).Count()
	if err != nil{
		return -1, nil, err
	}
	err = c.Find(bson.M{"player_id":playerId}).Skip(offset).Limit(limit).All(&r)
	return
}

func GetRechargeRecordPageByTime(limit, offset int, playerId string, t []int)(total int, r []RechargeRecord, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("rechargeRecord")
	total, err = c.Find(bson.M{"player_id":playerId,"created_time":bson.M{"$gte":t[0],"$lte":t[1]}}).Count()
	if err != nil{
		return  -1, nil , err
	}
	err = c.Find(bson.M{"player_id":playerId,"created_time":bson.M{"$gte":t[0],"$lte":t[1]}}).Skip(offset).Limit(limit).All(&r)
	return

}
