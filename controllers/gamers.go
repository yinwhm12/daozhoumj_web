package controllers

import (
	"daozhoumj/models"
	"net/http"
	"daozhoumj/models/bean"
)

//玩家信息
type GamersController struct {
	BaseController
}

//显示 玩家的一些信息
type ShowGamerInfo struct {
	Achievements models.Achiment `json:"achievements"`
	Sons int `json:"sons"`
	UserData UserInfo `json:"user_data"`
}

//用户的简化信息
type UserInfo struct {
	NickName string `json:"nick_name"`
	HeadImgUrl string `json:"head_img_url"`
	Gold int `json:"gold"`
	Diamond int `json:"diamond"`
}

// @Title get ten top
// @router /tenTop [get]
func (c *GamersController)GetTenTop()  {
	gs, err := models.GetTenTopData()
	if err != nil{
		c.RespJSON(http.StatusBadRequest, err.Error())
		return
	}
	infos := make([]ShowGamerInfo,len(gs))
	if len(gs) >0 {
		for i, v := range gs{
			infos[i].Achievements = v
			myclient, err := models.GetMyClientByGameId(v.Id)
			if err != nil{
				//todo 又一个读取失败 需处理
				infos[i].Sons = 0
			}else{
				infos[i].Sons = len(myclient.Sons)
			}
			u, err := models.GetUserByGameId(v.Id)
			if err != nil{
				// tod 读取玩家失败
				infos[i].UserData = UserInfo{}
			}else{
				infos[i].UserData.NickName = u.NickName
				infos[i].UserData.HeadImgUrl = u.HeadImgUrl
				infos[i].UserData.Gold = u.Gold
				infos[i].UserData.Diamond = u.Diamond

			}
		}
	}
	c.RespJSON(bean.CODE_Success, infos)
}

// @Title get data by page
// @router /gamers [get]
func (c *GamersController)GetGamers()  {
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
	total, gs, err := models.GetDataByPage(offset, limit)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err, err.Error())
		return
	}
	//获取玩家信息
	infos := make([]ShowGamerInfo,len(gs))
	if len(gs) >0 {
		for i, v := range gs{
			infos[i].Achievements = v
			myclient, err := models.GetMyClientByGameId(v.Id)
			if err != nil{
				//todo 又一个读取失败 需处理
				infos[i].Sons = 0
			}else{
				infos[i].Sons = len(myclient.Sons)
			}
			u, err := models.GetUserByGameId(v.Id)
			if err != nil{
				// tod 读取玩家失败
				infos[i].UserData = UserInfo{}
			}else{
				infos[i].UserData.NickName = u.NickName
				infos[i].UserData.HeadImgUrl = u.HeadImgUrl
				infos[i].UserData.Gold = u.Gold
				infos[i].UserData.Diamond = u.Diamond

			}
		}
	}
	c.RespJSONDataWithTotal(infos, int64(total))
	
}

// 直属人 的数据
type DearGamers struct {
	GameId string `json:"game_id"`
	NickName string `json:"nick_name"`
	HeadImgUrl string `json:"head_img_url"`
	Assert int `json:"assert"` //贡献值
}

//
type OneData struct {
	UserData UserInfo `json:"user_data"`
	Dears []DearGamers `json:"dears"`
	Achievement models.Achiment `json:"achievement"`
}

// @Title get a gamer info
// @router /getOne [get]
func (c *GamersController)GetOne()  {
	id := c.GetString("id")
	u,err := models.GetUserByGameId(id)
	if err != nil{
		c.RespJSON(bean.CODE_Params_Err,err.Error())
		return
	}
	//显示用户的金额等 贡献 直属人的贡献
	if u !=nil {
		oneData := OneData{}
		oneData.UserData.Diamond = u.Diamond
		oneData.UserData.Gold = u.Gold
		oneData.UserData.NickName = u.NickName
		oneData.UserData.HeadImgUrl = u.HeadImgUrl

		//获取其贡献度等数据
		myAchi, err := models.GetAchievemetById(id)
		if err != nil{
			oneData.Achievement = models.Achiment{
				Id:id,
				Achievements:0,
				Commision:0,
				Degree:0,
			}
		}else{
			oneData.Achievement = *myAchi
		}
		//获取其直属 用户
		myclient, err := models.GetMyClientByGameId(id)
		if err!= nil{
			//主推玩家没有
			oneData.Dears = []DearGamers{}
			c.RespJSON(bean.CODE_Success,oneData)
			return
		}
		if myclient != nil{
			if len(myclient.Sons) > 0{
				dears := make([]DearGamers, len(myclient.Sons))
				for i, v := range myclient.Sons{
					av, err := models.GetAchievemetById(v)
					dears[i].GameId = v
					if err != nil{
						//todo 有错误需处理
						dears[i].Assert =0
					}else{
						dears[i].Assert = av.Achievements
					}
					duser, err := models.GetUserByGameId(v)
					if err != nil{
						//todo 有问题 找不到玩家信息
						dears[i].HeadImgUrl =""
						dears[i].NickName = ""
					}else{
						dears[i].HeadImgUrl = duser.HeadImgUrl
						dears[i].NickName = duser.NickName
					}
				}
				oneData.Dears = dears
			}
			c.RespJSON(bean.CODE_Success,oneData)
		}
	}else{
		c.RespJSON(bean.CODE_Params_Err,"没有该玩家的信息,请输入正确的玩家ID")
		return
	}


}

