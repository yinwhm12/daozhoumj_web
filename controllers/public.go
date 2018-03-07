package controllers

import (
	"daozhoumj/models/client"
	"encoding/json"
	"net/http"
	"daozhoumj/models"
	"time"
	"daozhoumj/models/bean"
)

//公告

type PublicController struct {
	BaseController
}

// @Title Add new Public
// @Success 200 string
// @Failure 403 body is empty
// @router / [post]
func (p *PublicController)Post()  {
	var data client.PublicJSON
	err := json.Unmarshal(p.Ctx.Input.RequestBody,&data); if err != nil{
		p.RespJSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := models.CountPublic()
	if err != nil{
		p.RespJSON(http.StatusBadRequest,err.Error())
		return
	}
	publicData := models.Public{
		Id:id,
		CreateTime:int(time.Now().Unix()),
		Message:data.Message,
	}
	err = models.AddPublic(&publicData)
	if err != nil{
		p.RespJSON(http.StatusForbidden,err.Error())
		return
	}
	p.RespJSON(bean.CODE_Success,"发布成功!")
}

// @Title delete a public
// @router /:id [delete]
func (p *PublicController)Delete()  {
	id := p.Ctx.Input.Param(":id")
	err := models.DeletedPublic(id)
	if err != nil{
		p.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}
	p.RespJSON(bean.CODE_Success,"删除成功")
}

// @Title get all public data by page
// @router /publics [get]
func (p *PublicController)GetAllPublics()  {
	offset, err := p.GetInt("offset")
	if err != nil{
		p.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	limit, err := p.GetInt("limit")
	if err != nil{
		p.RespJSON(bean.CODE_Forbidden, err.Error())
		return
	}
	total, ps , err := models.GetAllPublicByPage(offset,limit)
	if err != nil{
		p.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	p.RespJSONDataWithTotal(ps, int64(total))
}

// @Title get some public datas by time
// @router /publicsByTime [get]
func (p *PublicController)GetPublicsByTime()  {
	var  jsonData client.PageByTime
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &jsonData)
	if err != nil{
		p.RespJSON(bean.CODE_Forbidden,err.Error())
		return
	}
	total, ps, err := models.GetPublicByPageAndTimes(jsonData.Offset,jsonData.Limit,jsonData.Before,jsonData.After)
	if err != nil{
		p.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	p.RespJSONDataWithTotal(ps, int64(total))

}