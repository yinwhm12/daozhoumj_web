package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type UserData struct {
	ID	string	`bson:"_id" json:"id,omitempty"`
	Name	string    `bson:"name" json:"name,omitempty"`
	Password	string	`bson:"password" json:"password,omitempty"`
	CreatedTime	int	`bson:"created_time" json:"created_time,omitempty"`
	Token	string	`bson:"token" json:"token,omitempty"`
	ProxyClass	int	`bson:"proxy_class" json:"proxy_class,omitempty"`//代理级别 1为admin总代理级 0为 普通总代代理

	LeftCards	int	`bson:"left_cards" json:"left_cards,omitempty"` //剩下卡的数量
	SoldCards	int	`bson:"sold_cards" json:"sold_cards,omitempty"` //售出卡的数量
	RechargeTimes	int	`bson:"recharge_times" json:"recharge_times,omitempty"` //充卡次数
}



func ValidateUser(name,pwd string)(user *UserData, err error)  {
	conn := mongodb.Conn()
	defer  conn.Close()

	c := conn.DB("").C("users")
	err = c.Find(bson.M{"name": name, "password": pwd}).One(&user)
	return
}

//用户数量
func GetUserCount() (int,error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("users")
	count,err :=  c.Find(bson.M{"_id":bson.M{"$ne":nil}}).Count()
	return count, err
}

//添加用户新用户
func AddUser(name,pwd string) error {
	conn := mongodb.Conn()
	defer conn.Close()

	count, err := GetUserCount()
	if err != nil{
		return err
	}
	user := UserData{
		ID: strconv.Itoa(count+1),
		Name:name,
		Password: pwd,
		CreatedTime: int(time.Now().Unix()),
	}
	c := conn.DB("").C("users")
	return c.Insert(&user)
}

func GetUserById(id string) (user *UserData, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()
	
	c := conn.DB("").C("users")
	err = c.FindId(id).One(&user)
	return
}

//func GetUserIdByToken(token string)(uid string, err error)  {
//	//conn := mongodb.Conn()
//	//defer conn.Close()
//	//
//	//c := conn.DB("").C("users")
//	//var user UserData
//	//c.Find(bson.M{""})
//	user, err :=GetUserByToken(token)
//}

func GetUserByToken(token string)(user *UserData,err error)  {
	conn := mongodb.Conn()
	defer conn.Close()
	
	c := conn.DB("").C("users")
	err = c.Find(bson.M{"token":token}).One(&user)
	return
}

func UpdatePassword(id,oldPwd, newPwd string) error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("users")
	err := c.Update(bson.M{"_id":id,"password":oldPwd},bson.M{"$set":bson.M{"password":newPwd}})
	return err
}

func UpdateToken(token string,name,pwd string) error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("users")
	err := c.Update(bson.M{"name": name, "password": pwd},bson.M{"$set":bson.M{"token":token}})
	return err
}
