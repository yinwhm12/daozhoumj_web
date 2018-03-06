package mongodb

import (
	"testing"
	"time"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type UserTest struct {
	Id int "_id"
	Name string `json:"name"`
	CreateTime int `json:"create_time"`
}


func TestUserTest(t *testing.T)  {
	conn := Conn()
	defer conn.Close()

	user := UserTest{
		Id:2,
		Name:"yn",
		CreateTime:int(time.Now().Unix()),
	}
	c := conn.DB("ddzhu").C("test")
	c.Insert(&user)

}

func TestFindUser(t *testing.T)  {
	conn := Conn()
	defer conn.Close()

	c := conn.DB("ddzhu").C("test")
	user := UserTest{}
	_ = c.Find(bson.M{"_id":1}).One(&user)
	fmt.Println(user)
}


