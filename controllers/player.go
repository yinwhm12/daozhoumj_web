package controllers

import (
	"encoding/json"
	"net/http"
	"daozhoumj/models"
	"time"
	"gopkg.in/mgo.v2/bson"
	"daozhoumj/models/bean"
)

type PlayerController struct {
	BaseController
}

// @Title Login
// @Description Logs user into the system
// @Param	body		body 	models.Player	true		"body for user content"
// @Success 200 {object} ok
// @Failure 403 user not exist
// @router /add [post]
func (c *PlayerController) AddPlayer()  {
	//fmt.Println("in---")
	var v models.Player
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); if err != nil{
		c.RespJSON(http.StatusBadRequest, err.Error())
		return
	}
	v.ID = bson.NewObjectId()
	v.LastGameTime = (int)(time.Now().Unix())
	err = models.AddPlayer(&v)
	if err != nil{
		c.RespJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.RespJSONData("ok")

}

// @Title all players
// @Success 200 {int}
// @Failure 403 empty
// @router /playerCount [get]
func (c *PlayerController)PlayerCounts()  {
	total, err := models.GetPlayerCounts()
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	c.RespJSONData(total)
}

// @Title yesterday increase players
// @Success 200 {int}
// @Failure 403 empty
// @router /increaseCount [get]
func (c *PlayerController)IncreaseCount()  {
	total, err := models.GetYesterdayIncreaseCounts()
	if err != nil {
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	c.RespJSONData(total)
}

// @Title get all players
// @Params	offset	path	int	true	"page offset"
// @Params	limit	path	int	true	"page limit"
// @Success 200	{int,[]object} total	[]models.player	true	"return"
// @Failure	403	empty
// @router /getAll [get]
func (c *PlayerController)GetAll()  {
	offset, err := c.GetInt("offset")
	if err != nil {
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	limit, err := c.GetInt("limit")
	if err != nil {
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	total, ps, err := models.GetPlayersByPage(limit,offset)
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	c.RespJSONDataWithTotal(ps,int64(total))
}

// @Title get one player
// @router /getOne [get]
func (c *PlayerController)GetOne()  {
	idStr := c.GetString("id")
	//fmt.Println("idstr",idStr)
	p, err  := models.GetOnePlayer(idStr)
	if err != nil{
		 c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	c.RespJSONDataWithTotal(p, int64(len(p)))
}

// @Title get bad players
// @router /badPlayers [get]
func (c *PlayerController)GetBadPlayers()  {
	offset, err := c.GetInt("offset")
	if err != nil {
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	limit, err := c.GetInt("limit")
	if err != nil {
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	total, ps, err := models.GetBadPlayersByPage(limit,offset)
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	c.RespJSONDataWithTotal(ps,int64(total))
}

// @Title get a bad player
// @router /badPlayer [get]
func (c *PlayerController)GetOneBadPlayer()  {
	idStr := c.GetString("id")
	//fmt.Println("idstr",idStr)
	p, err  := models.GetOneBadPlayerById(idStr)
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	c.RespJSONDataWithTotal(p, int64(len(p)))
}

// @Title get player by id
// @router /getAPlayer [get]
func (c *PlayerController)GetAPlayer()  {
	idStr := c.GetString("id")
	//fmt.Println("idstr",idStr)
	p, err  := models.GetAPlayer(idStr)
	//fmt.Println("pppp",p)
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	if len(p) == 1{
		c.RespJSONData(p)
	}else{
		c.RespJSON(bean.CODE_Params_Err, "数据异常!")
	}
}

// @Title add a bad player by id
// @router /addBadPlayer [put]
func (c *PlayerController)AddBadPlayer()  {
	idStr := c.GetString("id")
	//fmt.Println("idstr",idStr)
	err  := models.AddBadPlayer(idStr)
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	c.RespJSON(bean.CODE_Success,"ok")
}

