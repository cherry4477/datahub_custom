package conf

import "github.com/astaxie/beego/config"

var iniConfig config.Configer

func init() {
	iniConfig, _ = config.NewConfig("ini", "conf/app.conf")
}

func GetMysqlUser() string {
	return iniConfig.String("MysqlUser")
}

func GetMysqlPassword() string {
	return iniConfig.String("MysqlPassword")
}

func GetMysqlDatabase() string {
	return iniConfig.String("MysqlDatabase")
}
