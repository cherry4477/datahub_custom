package models

import (
	//"database/sql"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/astaxie/beego/config"
	"database/sql"
	"github.com/asiainfoLDP/datahub_custom/conf"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	Requirements  map[string]*Requirement
	mysqlUser     string
	mysqlPassword string
	mysqlDatabase string

	db *sql.DB
)

func AddOne(requirement Requirement) int64 {
	o := orm.NewOrm()


	id, err := o.Insert(&requirement)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(id)

	return id
}

func Get(params map[string]string) []*Requirement {
	o := orm.NewOrm()
	var requirements []*Requirement
	qs := o.QueryTable("requirement")

	if name, _ := params["name"]; name != "" {
		beego.Debug(name)
		qs = qs.Filter("name__iexact", name)
	}
	if phone, _ := params["phone"]; phone != "" {
		qs = qs.Filter("phone__contains", phone)
	}
	if email, _ := params["email"]; email != "" {
		qs = qs.Filter("email__contains", email)
	}
	if company, _ := params["company"]; company != "" {
		qs = qs.Filter("company__iexact", company)
	}
	if content, _ := params["content"]; content != "" {
		qs = qs.Filter("requirement_content__icontains", content)
	}
	qs.All(&requirements)

	beego.Debug(requirements)

	return requirements
}

func GetAll() []*Requirement {
	o := orm.NewOrm()

	var requirements []*Requirement
	_, err := o.QueryTable("requirement").All(&requirements)
	checkErr(err)

	beego.Debug(requirements)

	return requirements

}

func Update(req Requirement) {
	o := orm.NewOrm()
	_, err := o.Update(&req)
	checkErr(err)
}

func Delete(ObjectId string) {
}

func checkErr(err error) {
	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

func init() {
	mysqlUser = conf.GetMysqlUser()
	mysqlPassword = conf.GetMysqlPassword()
	mysqlDatabase = conf.GetMysqlDatabase()

	beego.Debug(mysqlUser, mysqlPassword, mysqlDatabase)

	connstr := mysqlUser + ":" + mysqlPassword + "@tcp(10.1.235.98:3388)/" + mysqlDatabase + "?charset=utf8"

	orm.RegisterDataBase("datahub", "mysql", connstr, 30)

	orm.RunSyncdb("default", false, true)
}
