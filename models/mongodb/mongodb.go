package mongodb

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func Conn()*mgo.Session  {
	return session.Copy()
}

func init() {
	//url:= beego.AppConfig.String("mongodb::url")
	//db := beego.AppConfig.String("mongodb::db")
	//todo test使用
	url := "39.107.65.67:27777"
	//fmt.Println("----url:",url)
	sess, err := mgo.Dial(url)
	if err != nil{
		panic(err)
	}
	session = sess
	session.SetMode(mgo.Monotonic,true)
	//session.DB(db)
}
