package controllers

import (
	"catw/catw/models/bean"
	"daozhoumj/models"
	"strings"
	"strconv"
)

type SoldCardsController struct {
	BaseController
}

// @Title Get my sold-record
// router /getMySoldCards [get]
func (c *SoldCardsController)GetMySoldCards()  {
	idStr := c.GetString("id")
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
	total,s, err := models.GetSoldRecordsByPage(limit, offset,idStr)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	c.RespJSONDataWithTotal(s, int64(total))
}

// @Title Get  sold-record buyer
// router /getToSoldCards [get]
func (c *SoldCardsController)GetToSoldCards()  {
	idStr := c.GetString("id")
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
	total,s, err := models.GetSoldRecordPageByToPlayerId(limit, offset,idStr)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	c.RespJSONDataWithTotal(s, int64(total))
}

// @Title get my sold-cards by time
// @router /getMyTime [get]
func (c *SoldCardsController)GetMyTime()  {
	idStr := c.GetString("id")
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
	timeStr := c.GetString("time")
	times := make([]int, 2)
	if timeStr != ""{
		ss := strings.Split(timeStr, ",")
		if len(ss) != 2{
			c.RespJSON(bean.CODE_Forbidden, "时间有误!")
			return
		}else{
			for i :=0 ;i < len(ss); i++{
				ii, _ :=  strconv.Atoi(ss[i])
				times[i] = ii
			}
		}
	}else{
		c.RespJSON(bean.CODE_Forbidden, "时间有误!")
		return
	}
	total, s, err := models.GetSoldRecordPageByTime(limit, offset,idStr,times)
	c.RespJSONDataWithTotal(s, int64(total))
}

// @Title get buyer sold-cards by time
// @router /getToTime [get]
func (c *SoldCardsController)GetToTime()  {
	idStr := c.GetString("id")
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
	timeStr := c.GetString("time")
	times := make([]int, 2)
	if timeStr != ""{
		ss := strings.Split(timeStr, ",")
		if len(ss) != 2{
			c.RespJSON(bean.CODE_Forbidden, "时间有误!")
			return
		}else{
			for i :=0 ;i < len(ss); i++{
				ii, _ :=  strconv.Atoi(ss[i])
				times[i] = ii
			}
		}
	}else{
		c.RespJSON(bean.CODE_Forbidden, "时间有误!")
		return
	}
	total, s, err := models.GetSoldRecordPageByToPlayerIdTime(limit, offset,idStr,times)
	c.RespJSONDataWithTotal(s, int64(total))
}
