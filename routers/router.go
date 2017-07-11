// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"daozhoumj/controllers"
	"daozhoumj/filters"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/*",&controllers.BaseController{},"options:Options"),
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/testUser",
			beego.NSInclude(
				&controllers.TestMongo{},
			),
		),
		beego.NSNamespace("/player",
			beego.NSInclude(
				&controllers.PlayerController{},
			),
		),

	)
	beego.AddNamespace(ns)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, filters.AuthLogin, true) // 验证登陆
}
