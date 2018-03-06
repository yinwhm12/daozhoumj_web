package filters

import (
	"github.com/astaxie/beego/context"
	"strings"
	"daozhoumj/client_info"
	"daozhoumj/client_info/tokenBean"
	"daozhoumj/models"
	"errors"
	"daozhoumj/models/bean"
)

var AuthLogin = func(ctx *context.Context) {
	if strings.Contains(ctx.Request.RequestURI,"/login"){
		return
	}
	if ctx.Request.Method == "OPTIONS"{
		ctx.Input.SetData("uid",0)
		return
	}
	token := ctx.Request.Header.Get("Authorization")
	//fmt.Println("token===",token)
	if token != ""{
		flag := client_info.ValidateToken(token)
		//fmt.Println("flag==",flag)
		if flag == tokenBean.TOKEN_OK{
			user, err :=models.GetManagerByToken(token)
			//fmt.Println("user",user)
			if err != nil{
				AllowCrows(ctx,err)
				return
			}
			ctx.Input.SetData("uid", user.Id)
			//fmt.Println("user_id",user.ID)
		}else if flag == tokenBean.TOKEN_OVERTIME{
			//errA = errors.New("token 失效!")
			AllowCrows(ctx,errors.New("token out of date"))
			return
		}else {
			//errA = errors.New("token 无法处理!")
			AllowCrows(ctx,errors.New("validate token"))
			return
		}
	}else {
		AllowCrows(ctx,errors.New("validate token"))
		return
	}

}

func AllowCrows(ctx *context.Context,err error)  {
	if ctx.Request.Method == "OPTIONS" {
		ctx.Input.SetData("uid", 0)
		return
	}
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")                               //允许访问源
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST,DELETE, GET, PUT, OPTIONS") //允许post访问
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")     //header的类型
	ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
	ctx.ResponseWriter.ResponseWriter.WriteHeader(bean.CODE_Unauthorized)
	ctx.WriteString(err.Error())
}