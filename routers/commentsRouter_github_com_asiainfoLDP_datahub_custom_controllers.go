package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ARequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ARequirementController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ARequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ARequirementController"],
		beego.ControllerComments{
			"Get",
			`/requirement`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ARequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ARequirementController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ARequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ARequirementController"],
		beego.ControllerComments{
			"Put",
			`/:reqId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ARequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ARequirementController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:DRequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:DRequirementController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

}
