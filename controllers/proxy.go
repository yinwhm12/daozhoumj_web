package controllers

import (
	"daozhoumj/models/bean"
	"daozhoumj/models"
)

type ProxyController struct {
	BaseController
}

// @Title change proxy class
// @router /changeClass [put]
func (c *ProxyController)ChangeClass()  {
	idStr := c.GetString("id")
	class, err := c.GetInt("class")
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	err = models.UpdateProxyClassById(idStr, class)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	c.RespJSON(bean.CODE_Success, "ok")
}

// @Title page show which proxy class you want
// @router /showProxy [get]
func (c *ProxyController)ShowProxy()  {
	class, err := c.GetInt("class")
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	offset, err := c.GetInt("offset")
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	limit, err := c.GetInt("limit")
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	total, p, err := models.GetWhichProxy(limit, offset, class)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	c.RespJSONDataWithTotal(p,int64(total))
}

// @Title search which id proxy
// @router /searchId [get]
func (c *ProxyController)SearchId()  {
	idStr := c.GetString("id")
	ps, err := models.GetPlayerInfoById(idStr)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	c.RespJSONDataWithTotal(ps, int64(len(ps)))
}

// @Title search one player(all proxy)
// @router /searchOne [get]
func (c *ProxyController)SearchOne()  {
	idStr := c.GetString("id")
	ps, err := models.GetOnePlayer(idStr)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	c.RespJSONDataWithTotal(ps, int64(len(ps)))
}

type ProxyCount struct {
	All int    `json:"all,omitempty"`
	C1	int	`json:"c_1,omitempty"`
	C2	int	`json:"c_2,omitempty"`
}

// @Title get kinds of proxy counts
// @router /getProxyCount [get]
func (c *ProxyController)GetProxyCount()  {
	all, c1, c2, err := models.GetAllProxyCount()
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	pc := ProxyCount{all,c1,c2}
	c.RespJSON(bean.CODE_Success,pc)
}