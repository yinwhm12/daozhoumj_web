package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"time"
	"crypto"
	"io"
)

type Manager struct {
	Id int `json:"_id" bson:"_id"`
	NickName string `json:"nick_name" bson:"nick_name"`
	LastTime int `json:"last_time" bson:"last_time"`
	Token string `json:"token" bson:"token"`
	MDPwd string `json:"md_pwd" bson:"md_pwd"`
}

func ValidateAndLogin(nickName, pwd string)(Manager, error)  {
	md5 := crypto.MD5.New()
	io.WriteString(md5,pwd)
	mdpwd := md5.Sum(nil)
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	m := new(Manager)
	err := c.Find(bson.M{"nick_name":nickName,"md_pwd":string(mdpwd)}).One(&m)
	return *m, err

}

func ValidateOkByMDAdnId(id int,pwd string)(Manager,error)  {
	md5 := crypto.MD5.New()
	io.WriteString(md5,pwd)
	mdpwd := md5.Sum(nil)
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	m := new(Manager)
	err := c.Find(bson.M{"_id":id,"md_pwd":string(mdpwd)}).One(&m)
	return  *m,err
}

func GetManagerByMDPwd(md string)(Manager, error)  {
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	m := new(Manager)
	err := c.Find(bson.M{"md_pwd":md}).One(&m)
	return *m, err
}


func GetManagerByToken(token string)(Manager, error)  {
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	m := new(Manager)
	err := c.Find(bson.M{"token":token}).One(&m)
	return *m, err
}

func UpdateManagerToken(token string,pwd string)error  {
	md5 := crypto.MD5.New()
	io.WriteString(md5,pwd)
	mdpwd := md5.Sum(nil)
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	err := c.Update(bson.M{"md_pwd":string(mdpwd)},bson.M{"$set":bson.M{"token":token}})
	return err
}

func UpdateTokenByToken(old_token,new_token string)error  {
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	err := c.Update(bson.M{"token":old_token},bson.M{"$set":bson.M{"token":new_token}})
	return err

}

func UpdateMd5(pwd string,id int)error  {
	md5 := crypto.MD5.New()
	io.WriteString(md5,pwd)
	mdpwd := md5.Sum(nil)
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	err := c.Update(bson.M{"_id":id},bson.M{"$set":bson.M{"md_pwd":string(mdpwd)}})
	return err
}


func GetManagerById(id string)(Manager, error)  {
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	m := new(Manager)
	err := c.Find(bson.M{"_id":id}).One(&m)
	return *m, err
}

func EditNickName(name string,id int)error  {
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	err := c.Update(bson.M{"_id":id},bson.M{"$set":bson.M{"nick_name":name}})
	return err
}

func UpdateLastTimeById(id int)error  {
	lasttime := int(time.Now().Unix())
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	err := c.Update(bson.M{"_id":id},bson.M{"$set":bson.M{"last_time":lasttime}})
	return err
}

func SaveM(m *Manager)error  {
	conn := mongodb.Conn()
	defer conn.Close()
	c := conn.DB("ddzhu").C("manager")
	return c.Insert(&m)

}