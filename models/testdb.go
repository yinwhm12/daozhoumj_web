package models

import (
	"daozhoumj/models/mongodb"
	"gopkg.in/mgo.v2/bson"
)

type TestUser struct {
	ID string `bson:"_id"	json:"_id,omitempty"`
	Name string `bson:"name"	json:"name,omitempty"`
	Pwd string	`bson:"pwd" json:"pwd,omitempty"`
}


func Insert(u *TestUser) error {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("users")
	return c.Insert(u)
}

func FindUserById(id string)(testUser *TestUser, err error)  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("users")
	err = c.FindId(id).One(&testUser)
	if err != nil{
		return nil, err
	}
	return testUser, nil
}

func ValidateUser(name, pwd string) error  {
	conn := mongodb.Conn()
	defer conn.Close()

	c := conn.DB("").C("users")
	//testUser := TestUser{}
	var uu interface{}
	return  c.Find(bson.M{"name":name,"pwd":pwd}).One(uu)
}
