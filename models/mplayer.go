package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"time"
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
	JoinProxyTime	int	`bson:"join_proxy_time" json:"join_proxy_time,omitempty"` //加入代理的时间
	HasCardsCount	int	`bson:"has_cards_count" json:"has_cards_count,omitempty"`//持卡数
}

func AddPlayer(p *Player)error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err := c.Insert(p)
	return err
}

func GetPlayerCounts()(total int, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	total, err = c.Find(nil).Count()
	return
}

func GetYesterdayIncreaseCounts()(total int, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	nTime := time.Now()
	yesTime := nTime.AddDate(0,0,-1)
	hadM := yesTime.Hour() * 60*60 + yesTime.Minute() *60 + yesTime.Second()
	time0 := int(yesTime.Unix())-hadM
	time1 := time0 + 24 * 60* 60

	c := conn.DB("").C("player")
	total, err = c.Find(bson.M{"last_game_time":bson.M{"$gte":time0,"$lte":time1}}).Count()
	return
}

func GetAllPlayers() (p []Player,err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err = c.Find(nil).All(&p)
	return
}

func GetAPlayer(id string)(p []Player, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err = c.Find(bson.M{"_id":bson.ObjectIdHex(id),"is_bad_player":0}).All(&p)
	return
}

func GetOnePlayer(id string)(p []Player, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err = c.FindId(bson.ObjectIdHex(id)).All(&p)
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

func GetOneBadPlayerById(id string)(p []Player, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err = c.FindId(bson.ObjectIdHex(id)).All(&p)
	return
}

func AddBadPlayer(id string)error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err := c.Update(bson.ObjectIdHex(id),bson.M{"$set":bson.M{"is_bad_player":1}})
	return  err
}

//
func GetPlayerInfoById(id string)(p []Player, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err = c.Find(bson.M{"_id":bson.ObjectIdHex(id),"is_proxy":bson.M{"$in":[]int{1,2}}}).All(&p)
	return
}

//代理 级别 修改 也可做添加代理
func UpdateProxyClassById(id string,class int)(err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	err = c.Update(bson.ObjectIdHex(id),bson.M{"$set":bson.M{"is_proxy":class}})
	return
}

//获取所有(1 2 ...)的代理
func GetWhichProxy(limit, offset, class int)(total int, p []Player, err error)  {
	conn := mongodb.Conn()
	defer  conn.Close()

	c := conn.DB("").C("player")
	if class == 1||class == 2{//一级代理
		total, err = c.Find(bson.M{"is_proxy":class}).Count()
		if err != nil {
			return -1, nil, err
		}
		err = c.Find(bson.M{"is_proxy":class}).Skip(offset).Limit(limit).All(&p)
	}else{//所有代理
		total, err = c.Find(bson.M{"is_proxy":bson.M{"$in":[]int{1,2}}}).Count()
		if err != nil {
			return -1, nil, err
		}
		err = c.Find(bson.M{"is_proxy":bson.M{"$in":[]int{1,2}}}).Skip(offset).Limit(limit).All(&p)
	}
	return
}

//获取各个代理的数量
func GetAllProxyCount()(all, c1,c2 int, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("player")
	all, err = c.Find(bson.M{"is_proxy":bson.M{"$in":[]int{1,2}}}).Count()
	if err != nil {
		return -1,-1,-1,err
	}
	c1, err = c.Find(bson.M{"is_proxy":1}).Count()
	if err != nil{
		return -1,-1,-1,err
	}
	c2= all - c1
	return all, c1,c2,nil
}
