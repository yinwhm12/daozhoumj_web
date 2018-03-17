package models

import (
	"gopkg.in/mgo.v2/bson"
	"daozhoumj/models/mongodb"
)

//定时 删除这周前的 战绩
func DeleteBeforeThisWeekData(t int)  error{
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("records")
	_,err := c.RemoveAll(bson.M{"create_time":bson.M{"$lt":t}})
	return err
}
