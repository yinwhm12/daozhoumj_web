package controllers

import (
	"encoding/json"
	"daozhoumj/models/bean"
	"daozhoumj/models"
)

type TestMongo struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param body	body	models.TestUser	 true	"user information"
// @Success 200	{string} models.TestUser.ID
// @Failure	403 body is empty
// @router / [post]
func (c *TestMongo)Post()  {
	var user models.TestUser
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	//err := models.ValidateUser(user.Name, user.Pwd)
	//if err !=nil {
	//	c.RespJSON(bean.CODE_Forbidden,err.Error())
	//	return
	//}

	err := models.Insert(&user)
	if err !=nil {
		c.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	c.RespJSONData("ok")
}

// @Title GetUser
// @Description get user by id
// @Param	id	path	string	true
// @Success 200 {object} models.TestUser
// @Failure	403 :id is empty
// @router /:id [get]
func (c *TestMongo)Get()  {
	id := c.GetString(":id")
	if id != ""{
		user, err := models.FindUserById(id)
		if err != nil{
			c.RespJSON(bean.CODE_Forbidden, err.Error())
			return
		}
		c.RespJSONData(user)
	}else {
		c.RespJSON(bean.CODE_Forbidden,nil)
	}
}
