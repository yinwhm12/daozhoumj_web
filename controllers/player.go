package controllers

import (
	"encoding/json"
	"net/http"
	"fmt"
	"daozhoumj/models/client"
)

type PlayerController struct {
	BaseController
}

// @Title Login
// @Description Logs user into the system
// @Param	body		body 	client.CreateSession	true		"body for user content"
// @Success 200 {object} ok
// @Failure 403 user not exist
// @router /add [post]
func (c *PlayerController) AddPlayer()  {
	fmt.Println("in---")
	var v client.CreateSession
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); if err != nil{
		c.RespJSON(http.StatusBadRequest, err.Error())
		return
	}
	//v.LastGameTime = (int)(time.Now().Unix())
	//err = models.AddPlayer(&v)
	//if err != nil{
	//	c.RespJSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	c.RespJSONData("ok")

}
