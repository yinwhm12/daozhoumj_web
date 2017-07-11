package controllers

import (
	"github.com/astaxie/beego"
	"daozhoumj/models/bean"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController)RespJSON(code int, data interface{})  {
	c.AllowCross()
	c.Ctx.Output.SetStatus(code)
	var hasIndex = true
	if beego.BConfig.RunMode == beego.PROD {
		hasIndex = false
	}
	c.Ctx.Output.JSON(data,hasIndex,false)
}

func (c *BaseController)RespJSONData(data interface{})  {
	c.AllowCross()
	c.RespJSON(bean.CODE_Success,data)
}

func (c *BaseController)RespJSONDataWithTotal(data interface{},total int64)  {
	c.RespJSON(bean.CODE_Success,map[string]interface{}{
		"data": data,
		"total": total,
	})
}

func (c *BaseController)Uid()string  {
	return c.Ctx.Input.GetData("uid").(string)
}

func (c *BaseController)AllowCross()  {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST,DELETE, GET, PUT, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")     
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}

func (c *BaseController)Options()  {
	c.AllowCross()
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo":"yin"}
	c.ServeJSON()
}