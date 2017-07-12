package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["daozhoumj/controllers:ObjectController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:ObjectController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:ObjectController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:ObjectController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:ObjectController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "AddPlayer",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "PlayerCounts",
			Router: `/playerCount`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "IncreaseCount",
			Router: `/increaseCount`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/getAll`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/getOne`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "GetBadPlayers",
			Router: `/badPlayers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "GetOneBadPlayer",
			Router: `/badPlayer`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "GetAPlayer",
			Router: `/getAPlayer`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:PlayerController"],
		beego.ControllerComments{
			Method: "AddBadPlayer",
			Router: `/addBadPlayer`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:ProxyController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:ProxyController"],
		beego.ControllerComments{
			Method: "ChangeClass",
			Router: `/changeClass`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:ProxyController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:ProxyController"],
		beego.ControllerComments{
			Method: "ShowProxy",
			Router: `/showProxy`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:ProxyController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:ProxyController"],
		beego.ControllerComments{
			Method: "SearchId",
			Router: `/searchId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:ProxyController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:ProxyController"],
		beego.ControllerComments{
			Method: "SearchOne",
			Router: `/searchOne`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:ProxyController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:ProxyController"],
		beego.ControllerComments{
			Method: "GetProxyCount",
			Router: `/getProxyCount`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:SoldCardsController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:SoldCardsController"],
		beego.ControllerComments{
			Method: "GetMyTime",
			Router: `/getMyTime`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:SoldCardsController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:SoldCardsController"],
		beego.ControllerComments{
			Method: "GetToTime",
			Router: `/getToTime`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:TestMongo"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:TestMongo"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:TestMongo"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:TestMongo"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:UserController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:UserController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:UserController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:UserController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:UserController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:UserController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:UserController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:UserController"],
		beego.ControllerComments{
			Method: "TokenLogin",
			Router: `/checkToken`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:UserController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
