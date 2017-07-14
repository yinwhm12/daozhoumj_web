package controllers

import (
	"daozhoumj/models/client"
	"encoding/json"
	"daozhoumj/models/bean"
	"daozhoumj/models"
	"gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
	"strings"
	"strconv"
)

type VersionController struct {
	BaseController
}

// @Title create new version
// @router / [post]
func (c *VersionController)Post()  {
	var v client.VersionJSON
	err := json.Unmarshal(c.Ctx.Input.RequestBody,&v)
	if err != nil{
		c.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	mv := models.Version{
		bson.NewObjectId(),
		v.Content,
		v.Address,
		(int)(time.Now().Unix()),
	}
	err = models.AddVersion(&mv)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	c.RespJSON(bean.CODE_Success,"发布成功!")
}

// @Title delete one version
// @router /:id [delete]
func (c *VersionController) Delete()  {
	versionId := c.Ctx.Input.Param(":id")
	fmt.Println("id ===",versionId)
	err := models.DeletedVersion(versionId)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	c.RespJSON(bean.CODE_Success, "删除成功!")
}

// @Title all versions by page
// @router /getAllVersions [get]
func (c *VersionController)GerAllVersion()  {
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
	total, v, err := models.GetAllVersionByPage(offset, limit)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	c.RespJSONDataWithTotal(v, int64(total))
}

// @Title choose version By time
// @router /getVersionsT [get]
func (c *VersionController)GerVersionsT()  {
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
	total, v, err := models.GetTimeVersionByPage(limit,offset,times)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	c.RespJSONDataWithTotal(v, int64(total))
}
