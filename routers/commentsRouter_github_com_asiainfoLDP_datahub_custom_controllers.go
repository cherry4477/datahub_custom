package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:DRequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:DRequirementController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:DRequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:DRequirementController"],
		beego.ControllerComments{
			"Get",
			`/requirement`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ERequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ERequirementController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ERequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ERequirementController"],
		beego.ControllerComments{
			"Get",
			`/requirement`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ERequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ERequirementController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ERequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ERequirementController"],
		beego.ControllerComments{
			"Put",
			`/:reqId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ERequirementController"] = append(beego.GlobalControllerRouter["github.com/asiainfoLDP/datahub_custom/controllers:ERequirementController"],
		beego.ControllerComments{
			"Delete",
			`/:reqId`,
			[]string{"delete"},
			nil})

}
