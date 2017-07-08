package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type UserData struct {
	ID	string	`bson:"_id" json:"id,omitempty"`
	UserName	string	`bson:"user_name json:"user_name,omitempty""`
	Password	string	`bson:"password" json:"password,omitempty"`
	CreatedTime	int	`bson:"created_time" json:"created_time,omitempty"`
	Token	string	`bson:"token"	json:"token,omitempty"`
}

func ValidateUser(name,pwd string)(user *UserData, err error)  {
	conn := mongodb.Conn()
	defer  conn.Close()

	c := conn.DB("").C("users")
	err = c.Find(bson.M{"user_name": name, "password": pwd}).One(&user)
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
		UserName:name,
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
