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
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["daozhoumj/controllers:UserController"] = append(beego.GlobalControllerRouter["daozhoumj/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
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
