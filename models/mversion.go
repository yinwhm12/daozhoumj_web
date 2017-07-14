package models

import (
	"gopkg.in/mgo.v2/bson"
	"daozhoumj/models/mongodb"
)

type Version struct {
	ID	bson.ObjectId	`bson:"_id" json:"id,omitempty"`
	Content	string	`bson:"content" json:"content,omitempty"`
	Address	string	`bson:"address" json:"address,omitempty"`
	CreatedTime	int	`bson:"created_time" json:"created_time,omitempty"`
}

func AddVersion(v *Version) error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("version")
	return c.Insert(v)
}

func DeletedVersion(id string) error {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("version")
	return c.Remove(bson.M{"_id":bson.ObjectIdHex(id)})
}

func GetAllVersionByPage(offset, limit int)(total int,v []Version, err error)  {
	conn := mongodb.Conn()
	defer  conn.Close()

	c := conn.DB("").C("version")
	total, err = c.Find(nil).Count()
	if err != nil{
		return -1, nil, err
	}
	err = c.Find(nil).Skip(offset).Limit(limit).All(&v)
	return
}

func GetTimeVersionByPage(limit,offset int,t []int)(total int,v []Version,err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("version")
	total, err = c.Find(bson.M{"created_time":bson.M{"$gte":t[0],"$lte":t[1]}}).Count()
	if err != nil{
		return -1, nil, err
	}
	err = c.Find(bson.M{"created_time":bson.M{"$gte":t[0],"$lte":t[1]}}).Skip(offset).Limit(limit).All(&v)
	return
	
}
