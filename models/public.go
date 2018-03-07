package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
)

//公告

type Public struct {
	Id int `json:"_id" bson:"_id"`
	CreateTime int `json:"create_time" bson:"create_time"`
	Message string `json:"message" bson:"message"`
}

func CountPublic()(int,error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("public")
	total, err := c.Find(nil).Count()
	return total, err
}

func AddPublic(p *Public) error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("public")
	return c.Insert(p)
}

func DeletedPublic(id string)error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("public")
	return c.Remove(bson.M{"_id":bson.ObjectIdHex(id)})
}

func GetAllPublicByPage(offset, limit int)(total int,p []Public, err error){

	conn := mongodb.Conn()
	defer  conn.Close()

	c := conn.DB("ddzhu").C("public")
	total, err = c.Find(nil).Count()
	if err != nil{
		return -1, nil, err
	}
	err = c.Find(nil).Skip(offset).Limit(limit).All(&p)
	return
}

func GetPublicByPageAndTimes(limit, offset, t1,t2 int)(total int,p []Public,err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("version")
	total, err = c.Find(bson.M{"created_time":bson.M{"$gte":t1,"$lte":t2}}).Count()
	if err != nil{
		return -1, nil, err
	}
	err = c.Find(bson.M{"created_time":bson.M{"$gte":t1,"$lte":t2}}).Sort("-create_time").Skip(offset).Limit(limit).All(&p)
	return

}
