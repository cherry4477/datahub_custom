package logs

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.SetLevel(beego.LevelDebug)
}
