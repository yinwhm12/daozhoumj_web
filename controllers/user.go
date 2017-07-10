package controllers

import (
	"daozhoumj/models"

	"daozhoumj/models/client"
	"encoding/json"
	"net/http"
	"daozhoumj/client_info"
	"daozhoumj/client_info/tokenBean"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	client.CreateSession	true		"body for user content"
// @Success 200 {object} client.LoginSuccessOutPut
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var v client.CreateSession
	err := json.Unmarshal(u.Ctx.Input.RequestBody,&v); if err != nil{
		u.RespJSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := models.ValidateUser(v.Name, v.Password)
	if err != nil{
		u.RespJSON(http.StatusBadRequest, "帐号信息有误!")
		return
	}
	token, err := client_info.CreateToken(v.Name)
	if err != nil{
		u.RespJSON(http.StatusBadRequest,"token error")
		return
	}

	err = models.UpdateToken(token, v.Name,v.Password)
	if err != nil{
		u.RespJSON(http.StatusBadRequest,err.Error())
		return
	}

	user.Password = ""
	v.Password = ""

	u.Ctx.ResponseWriter.Header().Add("Auth",token)
	u.RespJSON(http.StatusOK,client.LoginSuccessOutPut{user.ID,user.Name,user.Token})

}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	//users := models.GetAllUsers()
	//u.Data["json"] = users
	//u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	//uid := u.GetString(":uid")
	//if uid != "" {
	//	user, err := models.GetUser(uid)
	//	if err != nil {
	//		u.Data["json"] = err.Error()
	//	} else {
	//		u.Data["json"] = user
	//	}
	//}
	//u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	//uid := u.GetString(":uid")
	//if uid != "" {
	//	var user models.User
	//	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//	uu, err := models.UpdateUser(uid, &user)
	//	if err != nil {
	//		u.Data["json"] = err.Error()
	//	} else {
	//		u.Data["json"] = uu
	//	}
	//}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	body		body 	client.CreateSession	true		"body for user content"
// @Success 200 {object} client.LoginSuccessOutPut
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var v client.CreateSession
	err := json.Unmarshal(u.Ctx.Input.RequestBody,&v); if err != nil{
		u.RespJSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := models.ValidateUser(v.Name, v.Password)
	if err != nil{
		u.RespJSON(http.StatusBadRequest, "帐号信息有误!")
		return
	}
	token, err := client_info.CreateToken(v.Name)
	if err != nil{
		u.RespJSON(http.StatusBadRequest,"token error")
		return
	}
	err = models.UpdateToken(token, v.Name,v.Password)
	if err != nil{
		u.RespJSON(http.StatusBadRequest,err.Error())
		return
	}
	user.Password = ""
	v.Password = ""
	u.Ctx.ResponseWriter.Header().Add("Auth",token)
	u.RespJSON(http.StatusOK,client.LoginSuccessOutPut{user.ID,user.Name,user.Token})
}

// @Title token Login
// @Param token	query	string	true
// @Success 200	{string}	login	success
// @Failure 403 token out of date
// @router /checkToken [get]
func (u *UserController) TokenLogin()  {
	var tokenLogin	client.TokenLogin
	err := json.Unmarshal(u.Ctx.Input.RequestBody,&tokenLogin)
	if err != nil{
		u.RespJSON(http.StatusBadRequest,"token 无效!")
		return
	}
	tokenFlag := client_info.ValidateToken(tokenLogin.Token)
	if tokenFlag == tokenBean.TOKEN_OK{
		tokenString, err := client_info.CreateToken(tokenLogin.Token)
		if err != nil{
			u.RespJSON(http.StatusBadRequest,"信息异常!")
			return
		}
		tokenLogin.Token = tokenString
		u.RespJSON(http.StatusOK,tokenLogin)
	}else {
		u.RespJSON(http.StatusBadRequest,"信息失效，请重新登录!")

	}
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

