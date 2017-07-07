package mongodb

import (
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
)

var session *mgo.Session

func Conn()*mgo.Session  {
	return session.Copy()
}

func init() {
	url:= beego.AppConfig.String("mongodb::url")
	//db := beego.AppConfig.String("mongodb::db")

	sess, err := mgo.Dial(url)
	if err != nil{
		panic(err)
	}
	session = sess
	session.SetMode(mgo.Monotonic,true)
	//session.DB(db)
}
