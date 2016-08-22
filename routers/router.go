// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/asiainfoLDP/datahub_custom/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/custom",
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.ARequirementController{},
			),
		),
		beego.NSNamespace("/datahub",
			beego.NSInclude(
				&controllers.DRequirementController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

//var auth = func(ctx *context.Context) {
//	loginName := ctx.Request.Header.Get("User")
//	if loginName == ""  {
//		ctx.Abort(401, "no auth")
//	}
//}
